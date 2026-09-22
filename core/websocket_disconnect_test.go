package core

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

// fakeGateway は WebSocket の gateway を装う。ログインには応答し、要求は mode に従う。
type fakeGateway struct {
	*httptest.Server
	mode atomic.Value // "respond" | "close" (Close フレームを送って閉じる) | "drop" (TCP を黙って切る)
	got  atomic.Int32
}

func newFakeGateway(t *testing.T) *fakeGateway {
	t.Helper()
	g := &fakeGateway{}
	g.mode.Store("respond")
	up := websocket.Upgrader{}
	g.Server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := up.Upgrade(w, r, nil)
		if err != nil {
			return
		}
		defer c.Close()
		var wmu sync.Mutex
		for {
			_, msg, err := c.ReadMessage()
			if err != nil {
				return
			}
			var req map[string]interface{}
			_ = json.Unmarshal(msg, &req)
			x, _ := req["x_gs2"].(map[string]interface{})
			id, _ := x["requestId"].(string)
			if x["function"] == "login" {
				wmu.Lock()
				_ = c.WriteJSON(map[string]interface{}{"requestId": id, "status": 200, "type": "response",
					"body": map[string]interface{}{"access_token": "pt", "token_type": "Bearer", "expires_in": 3600}})
				wmu.Unlock()
				continue
			}
			g.got.Add(1)
			switch g.mode.Load().(string) {
			case "close":
				// ★応答を返さずに閉じる（gateway の setUserId が自分自身の接続を切る形）。
				_ = c.WriteControl(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "mgmt_delete"), time.Now().Add(time.Second))
				return
			case "drop":
				// ★Close フレームも無しに TCP を切る（ノードの死亡・ネットワーク断）。
				_ = c.NetConn().Close()
				return
			default:
				go func() {
					time.Sleep(time.Duration(1+len(id)%5) * time.Millisecond)
					wmu.Lock()
					_ = c.WriteJSON(map[string]interface{}{"requestId": id, "status": 200, "type": "response",
						"body": map[string]interface{}{"echo": id}})
					wmu.Unlock()
				}()
			}
		}
	}))
	t.Cleanup(g.Close)
	return g
}

func newWsSession(t *testing.T, g *fakeGateway) *Gs2WebSocketSession {
	t.Helper()
	ws := &Gs2WebSocketSession{Credential: BasicGs2Credential{ClientId: "cid", ClientSecret: "sec"},
		Region: "ap-northeast-1", SteadyEndpoint: g.URL}
	if err := ws.Connect(); err != nil {
		t.Fatalf("Connect: %v", err)
	}
	t.Cleanup(ws.Disconnect)
	return ws
}

// request は 1 要求を送り、timeout 以内の結果を返す（来なければ t.Fatal）。
func request(t *testing.T, ws *Gs2WebSocketSession, timeout time.Duration) AsyncResult {
	t.Helper()
	cb := make(chan AsyncResult, 1)
	id := WebSocketRequestId(uuid.New().String())
	job := WebSocketNetworkJob{RequestId: id, Bodies: map[string]interface{}{
		"x_gs2": map[string]interface{}{"service": "gateway", "component": "webSocketSession", "function": "setUserId", "requestId": id}},
		Callback: cb}
	if err := ws.Send(&job, true); err != nil {
		return AsyncResult{Err: err}
	}
	select {
	case r := <-cb:
		return r
	case <-time.After(timeout):
		t.Fatalf("%v 待っても応答も誤りも来ない（以前の永久待ち）", timeout)
		return AsyncResult{}
	}
}

func receiveGoroutines() int {
	buf := make([]byte, 1<<20)
	n := runtime.Stack(buf, true)
	return strings.Count(string(buf[:n]), "Gs2WebSocketSession).receive(")
}

func TestWebSocketRequestResponse(t *testing.T) {
	g := newFakeGateway(t)
	ws := newWsSession(t, g)
	r := request(t, ws, 2*time.Second)
	if r.Err != nil || !strings.Contains(r.Payload, `"echo"`) {
		t.Fatalf("r=%+v", r)
	}
	if n := len(ws.Jobs); n != 0 {
		t.Errorf("応答済みの要求が待ち行列に残っている: %d", n)
	}
}

// ★応答の前にサーバーが Close フレームで閉じたら、待ち中の要求に ConnectionBroken が返る。
func TestWebSocketPendingFailsWhenServerCloses(t *testing.T) {
	g := newFakeGateway(t)
	ws := newWsSession(t, g)
	base := receiveGoroutines() // このセッションの 1 本を含む（他の試験のセッションも数に入る）
	g.mode.Store("close")
	r := request(t, ws, 2*time.Second)
	if !errors.As(r.Err, &ConnectionBroken{}) {
		t.Fatalf("ConnectionBroken のはず: %v", r.Err)
	}
	// 以後の送信は即座に ConnectionBroken。
	r = request(t, ws, time.Second)
	if !errors.As(r.Err, &ConnectionBroken{}) {
		t.Fatalf("切れた後の送信は ConnectionBroken のはず: %v", r.Err)
	}
	// 受信 goroutine は抜けている（以前は panic → 無バッファのチャネルで固まっていた）。
	deadline := time.Now().Add(2 * time.Second)
	for receiveGoroutines() >= base && time.Now().Before(deadline) {
		time.Sleep(20 * time.Millisecond)
	}
	if n := receiveGoroutines(); n != base-1 {
		t.Errorf("受信 goroutine が抜けていない: %d（%d のはず）", n, base-1)
	}
	// 繋ぎ直せる。
	g.mode.Store("respond")
	if err := ws.Connect(); err != nil {
		t.Fatalf("再接続: %v", err)
	}
	if r := request(t, ws, 2*time.Second); r.Err != nil {
		t.Fatalf("再接続後の要求: %v", r.Err)
	}
}

// ★Close フレーム無しの切断（TCP が黙って切れる）でも同じ。
func TestWebSocketPendingFailsWhenConnectionDrops(t *testing.T) {
	g := newFakeGateway(t)
	ws := newWsSession(t, g)
	g.mode.Store("drop")
	r := request(t, ws, 2*time.Second)
	if !errors.As(r.Err, &ConnectionBroken{}) {
		t.Fatalf("ConnectionBroken のはず: %v", r.Err)
	}
}

// ★Disconnect も待ち中の要求に ConnectionBroken を返す。
func TestWebSocketDisconnectFailsPending(t *testing.T) {
	g := newFakeGateway(t)
	ws := newWsSession(t, g)
	// サーバーが応答しない（close も drop もしない）ように、受け取って何もしない mode は無いので
	// 応答を遅らせる代わりに、送ってすぐ Disconnect する。
	cb := make(chan AsyncResult, 1)
	id := WebSocketRequestId(uuid.New().String())
	job := WebSocketNetworkJob{RequestId: id, Bodies: map[string]interface{}{"x_gs2": map[string]interface{}{"requestId": id}}, Callback: cb}
	ws.Disconnect()
	if err := ws.Send(&job, true); !errors.As(err, &ConnectionBroken{}) {
		t.Fatalf("Disconnect 後の送信: %v", err)
	}
}

// ★並列に送っても応答が取り違えられず、全部返る（Jobs の競合）。
func TestWebSocketConcurrentRequests(t *testing.T) {
	g := newFakeGateway(t)
	ws := newWsSession(t, g)
	const n = 50
	var wg sync.WaitGroup
	errs := make(chan error, n)
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			cb := make(chan AsyncResult, 1)
			id := WebSocketRequestId(uuid.New().String())
			job := WebSocketNetworkJob{RequestId: id, Bodies: map[string]interface{}{"x_gs2": map[string]interface{}{"function": "x", "requestId": id}}, Callback: cb}
			if err := ws.Send(&job, true); err != nil {
				errs <- err
				return
			}
			select {
			case r := <-cb:
				if r.Err != nil || !strings.Contains(r.Payload, string(id)) {
					errs <- errors.New("応答が違う: " + r.Payload)
				}
			case <-time.After(3 * time.Second):
				errs <- errors.New("応答が来ない")
			}
		}()
	}
	wg.Wait()
	close(errs)
	for err := range errs {
		t.Error(err)
	}
	if l := len(ws.Jobs); l != 0 {
		t.Errorf("Jobs が残っている: %d", l)
	}
	if got := g.got.Load(); got != n {
		t.Errorf("サーバーが受けた要求 %d != %d", got, n)
	}
}
