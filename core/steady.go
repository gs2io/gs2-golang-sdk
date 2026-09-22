// Steady（専用フリート）の基点。
//
// フリートは 1 つの名前（SteadyEndpoint、例 https://bs-dev.ap-northeast-1.dev.gen2.gs2io.com）で受け、
// REST は <steady>/<service>/...、WebSocket は wss://<host>/ を使う。名前はフリートのノードへ直接
// 解決される（間に ALB は無い）ので、フリートが手放した公開 IP に当たると SYN が落ちる。
// そのため Steady のときだけ接続段階に上限（SteadyConnectTimeout）を置き、接続段階の失敗
// （1 バイトも送っていない）だけは同じ要求をもう 1 回だけ送る。送信後の失敗は届いたかもしれない
// ので再送しない（非冪等要求の二重実行を作らない）。
package core

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// SteadyConnectTimeout は Steady の基点への dial / TLS handshake の上限。
// フリートが手放した公開 IP は SYN を落とすので、OS 既定（数十秒〜数分）に任せない。
const SteadyConnectTimeout = 5 * time.Second

// normalizeSteadyEndpoint は末尾の / と空白を落とす。
func normalizeSteadyEndpoint(v string) string {
	return strings.TrimRight(strings.TrimSpace(v), "/")
}

// steadyRestTemplate は SteadyEndpoint から REST の template（{service} 付き）を作る。空なら ""。
func steadyRestTemplate(steady string) string {
	steady = normalizeSteadyEndpoint(steady)
	if steady == "" {
		return ""
	}
	return steady + "/{service}"
}

// steadyWebSocketUrl は SteadyEndpoint から WebSocket の接続先を作る。空なら ""。
// http:// の基点（ローカルの試験・開発）は ws:// に、https:// は wss:// に。
func steadyWebSocketUrl(steady string) string {
	base := normalizeSteadyEndpoint(steady)
	if base == "" {
		return ""
	}
	u, err := url.Parse(base)
	if err != nil || u.Host == "" {
		return ""
	}
	scheme := "wss"
	if u.Scheme == "http" {
		scheme = "ws"
	}
	return scheme + "://" + u.Host + "/"
}

// isSteadyUrl は要求 URL が Steady の基点宛か。
func isSteadyUrl(steady string, requestUrl string) bool {
	base := normalizeSteadyEndpoint(steady)
	if base == "" {
		return false
	}
	return requestUrl == base || strings.HasPrefix(requestUrl, base+"/")
}

// isConnectFailure は「1 バイトも送っていない」失敗か（DNS / TCP の dial / TLS）。
// これだけが再送の対象。
func isConnectFailure(err error) bool {
	if err == nil {
		return false
	}
	var dnsErr *net.DNSError
	if errors.As(err, &dnsErr) {
		return true
	}
	var opErr *net.OpError
	if errors.As(err, &opErr) && opErr.Op == "dial" {
		return true
	}
	var certErr *tls.CertificateVerificationError
	if errors.As(err, &certErr) {
		return true
	}
	var unknownAuthority x509.UnknownAuthorityError
	if errors.As(err, &unknownAuthority) {
		return true
	}
	var hostnameErr x509.HostnameError
	if errors.As(err, &hostnameErr) {
		return true
	}
	var invalidErr x509.CertificateInvalidError
	if errors.As(err, &invalidErr) {
		return true
	}
	// net/http の TLS handshake timeout は非公開型（net.Error で Timeout() が真、文言で見分ける）。
	var netErr net.Error
	if errors.As(err, &netErr) && netErr.Timeout() && strings.Contains(err.Error(), "TLS handshake timeout") {
		return true
	}
	return false
}

// newHTTPClient は既定の HTTP クライアント。Steady のときだけ接続段階に上限（SteadyConnectTimeout）を置く。
// 共有クラウド（SteadyEndpoint が空）は従来どおり new(http.Client)。
func newHTTPClient(steadyEndpoint string) *http.Client {
	if normalizeSteadyEndpoint(steadyEndpoint) == "" {
		return new(http.Client)
	}
	return &http.Client{
		Transport: &http.Transport{
			Proxy:               http.ProxyFromEnvironment,
			DialContext:         (&net.Dialer{Timeout: SteadyConnectTimeout, KeepAlive: 30 * time.Second}).DialContext,
			TLSHandshakeTimeout: SteadyConnectTimeout,
			ForceAttemptHTTP2:   true,
			MaxIdleConns:        100,
			IdleConnTimeout:     90 * time.Second,
		},
	}
}
