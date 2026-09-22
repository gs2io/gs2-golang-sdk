package core

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/gorilla/websocket"
)

// ---------------------------------------------------------------- 偽サーバー

type steadyServer struct {
	*httptest.Server
	mu     sync.Mutex
	paths  []string
	bodies []string
	status int
}

func (r *steadyServer) hits(path string) int {
	r.mu.Lock()
	defer r.mu.Unlock()
	n := 0
	for _, p := range r.paths {
		if p == path {
			n++
		}
	}
	return n
}

// newSteadyServer はフリートの基点を装う: プロジェクトトークンのログインは 200、その他は status で {}。
func newSteadyServer(t *testing.T, status int) *steadyServer {
	t.Helper()
	rs := &steadyServer{status: status}
	rs.Server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		rs.mu.Lock()
		rs.paths = append(rs.paths, r.URL.Path)
		rs.bodies = append(rs.bodies, string(body))
		rs.mu.Unlock()
		if r.URL.Path == "/identifier/projectToken/login" {
			_, _ = io.WriteString(w, `{"access_token":"pt","token_type":"Bearer","expires_in":3600}`)
			return
		}
		w.WriteHeader(rs.status)
		_, _ = io.WriteString(w, `{"message":"[]"}`)
	}))
	t.Cleanup(rs.Close)
	return rs
}

func newRestSession(t *testing.T, steady string) *Gs2RestSession {
	t.Helper()
	s := NewGs2RestSession(BasicGs2Credential{ClientId: "cid", ClientSecret: "secret"}, Region("ap-northeast-1"))
	s.SteadyEndpoint = steady
	if err := s.Connect(); err != nil {
		t.Fatalf("Connect に失敗: %v", err)
	}
	return s
}

func sendJob(t *testing.T, s *Gs2RestSession, method HttpMethod, target Url, bodies map[string]interface{}) error {
	t.Helper()
	callback := make(chan AsyncResult, 1)
	job := NetworkJob{Url: target, Method: method, Headers: s.CreateAuthorizationHeader(), Bodies: bodies, Callback: callback}
	_ = s.Send(&job, true)
	return (<-callback).Err
}

// failFirstDials は最初の n 回の dial を「接続拒否」で落とし、以後は本物の dial に通す Transport。
func failFirstDials(s *Gs2RestSession, n int32) *int32 {
	var dials int32
	real := (&net.Dialer{Timeout: SteadyConnectTimeout}).DialContext
	s.connection.Client().Transport = &http.Transport{
		DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
			if atomic.AddInt32(&dials, 1) <= n {
				return nil, &net.OpError{Op: "dial", Net: network, Err: errors.New("connection refused")}
			}
			return real(ctx, network, addr)
		},
	}
	return &dials
}

// ---------------------------------------------------------------- URL

// ★SteadyEndpoint が空なら、従来の EndpointHost と byte 単位で同じ URL（override 2 種も従来どおり）。
func TestEndpointHostWithoutSteadyIsUnchanged(t *testing.T) {
	s := NewGs2RestSession(BasicGs2Credential{ClientId: "cid", ClientSecret: "secret"}, Region("ap-northeast-1"))
	if got, want := s.EndpointHost("account", nil), Url("https://account.ap-northeast-1.gen2.gs2io.com"); got != want {
		t.Errorf("got %s want %s", got, want)
	}
	saved := EndpointHost
	EndpointHost = "https://{service}.{region}.dev.gen2.gs2io.com"
	defer func() { EndpointHost = saved }()
	if got := s.EndpointHost("account", nil); got != "https://account.ap-northeast-1.dev.gen2.gs2io.com" {
		t.Errorf("global override: got %s", got)
	}
	override := "https://custom.example.test/{service}"
	if got := s.EndpointHost("account", &override); got != "https://custom.example.test/account" {
		t.Errorf("per-service override: got %s", got)
	}
}

func TestEndpointHostWithSteady(t *testing.T) {
	s := NewGs2RestSession(BasicGs2Credential{ClientId: "cid", ClientSecret: "secret"}, Region("ap-northeast-1"))
	s.SteadyEndpoint = "https://bs-dev.example.test/ "
	if got := s.EndpointHost("account", nil); got != "https://bs-dev.example.test/account" {
		t.Errorf("steady: got %s", got)
	}
	// サービスごとの override は Steady より強い。
	override := "https://custom.example.test/{service}"
	if got := s.EndpointHost("account", &override); got != "https://custom.example.test/account" {
		t.Errorf("per-service override: got %s", got)
	}
}

// ★プロジェクトトークンのログインも Steady の基点へ向く（共有クラウドの identifier へ行かない）。
func TestConnectLogsInAtSteady(t *testing.T) {
	srv := newSteadyServer(t, http.StatusOK)
	s := newRestSession(t, srv.URL)
	if got := srv.hits("/identifier/projectToken/login"); got != 1 {
		t.Fatalf("login hits = %d", got)
	}
	if s.connection.Client().Transport == nil {
		t.Fatalf("Steady では接続段階に上限のある Transport が要る")
	}
}

func TestSharedCloudClientIsPlain(t *testing.T) {
	if c := newHTTPClient(""); c.Transport != nil {
		t.Fatalf("共有クラウドは従来どおり new(http.Client)")
	}
}

// ---------------------------------------------------------------- 再送

// ★接続段階の失敗 → 同じ要求をもう 1 回だけ（POST の本文も保たれる）。
func TestConnectFailureIsRetriedOnce(t *testing.T) {
	srv := newSteadyServer(t, http.StatusOK)
	s := newRestSession(t, srv.URL)
	dials := failFirstDials(s, 1)
	err := sendJob(t, s, Post, s.EndpointHost("account", nil).AppendPath("/status", nil), map[string]interface{}{"k": "v"})
	if err != nil {
		t.Fatalf("再送で成功するはず: %v", err)
	}
	if got := srv.hits("/account/status"); got != 1 {
		t.Errorf("server hits = %d", got)
	}
	if atomic.LoadInt32(dials) != 2 {
		t.Errorf("dials = %d (1 回落ちて 1 回通る)", *dials)
	}
	// 本文は圧縮されて届く（Content-Encoding: gzip）ので中身は見ず、空でないことだけ見る。
	srv.mu.Lock()
	last := srv.bodies[len(srv.bodies)-1]
	srv.mu.Unlock()
	if last == "" {
		t.Errorf("再送の本文が空")
	}
}

// ★2 回続けて接続段階で落ちたら諦める（3 回目は無い）。
func TestConnectFailureTwiceGivesUp(t *testing.T) {
	srv := newSteadyServer(t, http.StatusOK)
	s := newRestSession(t, srv.URL)
	dials := failFirstDials(s, 100)
	err := sendJob(t, s, Get, s.EndpointHost("account", nil).AppendPath("/status", nil), nil)
	if err == nil {
		t.Fatalf("失敗が返るはず")
	}
	if atomic.LoadInt32(dials) != 2 {
		t.Errorf("dials = %d (再送は 1 回だけ)", *dials)
	}
	if got := srv.hits("/account/status"); got != 0 {
		t.Errorf("server hits = %d", got)
	}
}

// ★送信後の失敗（5xx）は再送しない。
func TestServerErrorIsNotRetried(t *testing.T) {
	srv := newSteadyServer(t, http.StatusInternalServerError)
	s := newRestSession(t, srv.URL)
	if err := sendJob(t, s, Post, s.EndpointHost("account", nil).AppendPath("/status", nil), map[string]interface{}{}); err == nil {
		t.Fatalf("5xx は誤りとして返るはず")
	}
	if got := srv.hits("/account/status"); got != 1 {
		t.Errorf("server hits = %d (再送してはいけない)", got)
	}
}

// ★Steady 未設定なら接続失敗でも再送しない（共有クラウドの挙動は不変）。
func TestNoRetryWithoutSteady(t *testing.T) {
	srv := newSteadyServer(t, http.StatusOK)
	s := NewGs2RestSession(BasicGs2Credential{ClientId: "cid", ClientSecret: "secret"}, Region("ap-northeast-1"))
	override := srv.URL + "/{service}"
	// ログイン先だけ偽サーバーへ（Steady は使わない）。
	saved := EndpointHost
	EndpointHost = override
	defer func() { EndpointHost = saved }()
	if err := s.Connect(); err != nil {
		t.Fatalf("Connect: %v", err)
	}
	dials := failFirstDials(s, 1)
	if err := sendJob(t, s, Get, s.EndpointHost("account", nil).AppendPath("/status", nil), nil); err == nil {
		t.Fatalf("再送しないので失敗が返るはず")
	}
	if atomic.LoadInt32(dials) != 1 {
		t.Errorf("dials = %d", *dials)
	}
}

// ---------------------------------------------------------------- 判定

func TestIsConnectFailure(t *testing.T) {
	cases := []struct {
		name string
		err  error
		want bool
	}{
		{"nil", nil, false},
		{"dns", &net.DNSError{Err: "no such host", Name: "x"}, true},
		{"dial refused", &net.OpError{Op: "dial", Err: errors.New("connection refused")}, true},
		{"dial timeout", &net.OpError{Op: "dial", Err: &timeoutError{}}, true},
		{"read (sent)", &net.OpError{Op: "read", Err: errors.New("connection reset by peer")}, false},
		{"write (sent)", &net.OpError{Op: "write", Err: errors.New("broken pipe")}, false},
		{"tls verify", &tls.CertificateVerificationError{Err: x509.UnknownAuthorityError{}}, true},
		{"x509 hostname", x509.HostnameError{Host: "x"}, true},
		{"tls handshake timeout", &timeoutError{msg: "net/http: TLS handshake timeout"}, true},
		{"client timeout", &timeoutError{msg: "Client.Timeout exceeded while awaiting headers"}, false},
		{"eof", io.EOF, false},
		{"other", errors.New("boom"), false},
	}
	for _, c := range cases {
		if got := isConnectFailure(c.err); got != c.want {
			t.Errorf("%s: got %v want %v", c.name, got, c.want)
		}
	}
}

type timeoutError struct{ msg string }

func (e *timeoutError) Error() string {
	if e.msg == "" {
		return "i/o timeout"
	}
	return e.msg
}
func (e *timeoutError) Timeout() bool   { return true }
func (e *timeoutError) Temporary() bool { return true }

func TestIsSteadyUrl(t *testing.T) {
	if !isSteadyUrl("https://bs.example.test/", "https://bs.example.test/account/status") {
		t.Error("配下は真")
	}
	if isSteadyUrl("https://bs.example.test", "https://bs.example.test.evil/x") {
		t.Error("前方一致だけでは駄目")
	}
	if isSteadyUrl("", "https://bs.example.test/account") {
		t.Error("Steady 未設定は偽")
	}
}

// ---------------------------------------------------------------- WebSocket

func TestWebSocketUrlAndDialer(t *testing.T) {
	ws := &Gs2WebSocketSession{Credential: BasicGs2Credential{ClientId: "cid", ClientSecret: "secret"}, Region: Region("ap-northeast-1")}
	if got := ws.webSocketUrl(); got != strings.ReplaceAll(WsEndpointHost, "{region}", "ap-northeast-1") {
		t.Errorf("未設定: got %s", got)
	}
	if ws.dialer() != websocket.DefaultDialer {
		t.Errorf("未設定は DefaultDialer")
	}
	ws.SteadyEndpoint = "https://bs-dev.example.test/"
	if got := ws.webSocketUrl(); got != "wss://bs-dev.example.test/" {
		t.Errorf("https: got %s", got)
	}
	if d := ws.dialer(); d == websocket.DefaultDialer || d.HandshakeTimeout != SteadyConnectTimeout {
		t.Errorf("Steady は handshake に上限")
	}
	ws.SteadyEndpoint = "http://127.0.0.1:8080"
	if got := ws.webSocketUrl(); got != "ws://127.0.0.1:8080/" {
		t.Errorf("http: got %s", got)
	}
	if steadyWebSocketUrl("not a url") != "" {
		t.Errorf("壊れた基点は空")
	}
	_ = time.Second
}
