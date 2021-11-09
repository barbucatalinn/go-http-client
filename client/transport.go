package client

import (
	"crypto/tls"
	"net"
	"net/http"
	"runtime"
)

// createHTTPTransport returns a new http.Transport with custom values
func createHTTPTransport() *http.Transport {
	return &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   dialContextTimeout,
			KeepAlive: dialContextKeepAlive,
		}).DialContext,
		MaxIdleConns:          maxIdleConns,
		IdleConnTimeout:       idleConnTimeout,
		TLSHandshakeTimeout:   tlsHandshakeTimeout,
		ExpectContinueTimeout: expectContinueTimeout,
		ForceAttemptHTTP2:     true,
		MaxIdleConnsPerHost:   runtime.GOMAXPROCS(0) + 1,

		// in the case we are receiving from the API a 'tls: no renegotiation' error
		// in order to fix this problem we need to run the http client with a TLS config
		// where we need to specify the Renegotiation type
		// see: https://stackoverflow.com/questions/57420833/tls-no-renegotiation-error-on-http-request
		TLSClientConfig: &tls.Config{
			MinVersion:         tls.VersionTLS12,
			Renegotiation:      tls.RenegotiateOnceAsClient,
			InsecureSkipVerify: false,
		},
	}
}
