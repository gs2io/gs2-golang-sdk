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

const SteadyConnectTimeout = 5 * time.Second

func normalizeSteadyEndpoint(v string) string {
	return strings.TrimRight(strings.TrimSpace(v), "/")
}

func steadyRestTemplate(steady string) string {
	steady = normalizeSteadyEndpoint(steady)
	if steady == "" {
		return ""
	}
	return steady + "/{service}"
}

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

func isSteadyUrl(steady string, requestUrl string) bool {
	base := normalizeSteadyEndpoint(steady)
	if base == "" {
		return false
	}
	return requestUrl == base || strings.HasPrefix(requestUrl, base+"/")
}

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
	var netErr net.Error
	if errors.As(err, &netErr) && netErr.Timeout() && strings.Contains(err.Error(), "TLS handshake timeout") {
		return true
	}
	return false
}

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
