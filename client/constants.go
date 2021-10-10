package client

import "time"

// Misc.
const (
	defaultRetryMax   int   = 2
	responseReadLimit int64 = 4096
)

// Auth schemes
const (
	basicAuthScheme  string = "Basic"
	bearerAuthScheme string = "Bearer"
)

// Header keys/values used for requests
const (
	authorizationHeaderKey string = "Authorization"
	contentTypeHeaderKey   string = "Content-Type"
	retryAfterHeaderKey    string = "Retry-After"
	userAgentHeaderKey     string = "User-Agent"
	userAgentHeaderValue   string = "go-http-client"
)

// http.Transport constants
const (
	maxIdleConns          int           = 100
	idleConnTimeout       time.Duration = 90 * time.Second
	tlsHandshakeTimeout   time.Duration = 10 * time.Second
	expectContinueTimeout time.Duration = 1 * time.Second

	dialContextTimeout   time.Duration = 30 * time.Second
	dialContextKeepAlive time.Duration = 30 * time.Second
)
