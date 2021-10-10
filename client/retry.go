package client

import (
	"context"
	"crypto/x509"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"regexp"
)

var (
	// a regular expression to match the error returned by net/http when the
	// configured number of redirects is exhausted
	redirectsErrorRe = regexp.MustCompile(`stopped after \d+ redirects\z`)

	// a regular expression to match the error returned by net/http when the
	// scheme specified in the URL is invalid
	schemeErrorRe = regexp.MustCompile(`unsupported protocol scheme`)
)

// RetryPolicy specifies a policy for handling retries
type RetryPolicy func(ctx context.Context, resp *http.Response, err error) (bool, error)

// DefaultRetryPolicy provides a default callback for Client.Retry, which
// will retry on connection errors and server errors
func DefaultRetryPolicy(ctx context.Context, resp *http.Response, err error) (bool, error) {
	// no retry for context.Canceled or context.DeadlineExceeded
	if ctx.Err() != nil {
		return false, ctx.Err()
	}
	return baseRetryPolicy(resp, err)
}

// baseRetryPolicy is the func where the logic is keep for the
// default retry policy
func baseRetryPolicy(resp *http.Response, err error) (bool, error) {
	if err != nil {
		if v, ok := err.(*url.Error); ok {
			// to too many redirects - no retry
			if redirectsErrorRe.MatchString(v.Error()) {
				return false, v
			}

			// invalid protocol scheme - no retry
			if schemeErrorRe.MatchString(v.Error()) {
				return false, v
			}

			// TLS cert verification failure - no retry
			if _, ok := v.Err.(x509.UnknownAuthorityError); ok {
				return false, v
			}

			// net package error - retry
			if _, ok := v.Err.(*net.OpError); ok {
				return true, v
			}
		}

		// the error is likely recoverable - retry
		return true, nil
	}

	// 429 Too Many Requests
	if resp.StatusCode == http.StatusTooManyRequests {
		return true, nil
	}

	// check the response code
	if resp.StatusCode == 0 || (resp.StatusCode >= 500 && resp.StatusCode != 501) {
		return true, fmt.Errorf("unexpected HTTP status %s", resp.Status)
	}

	return false, nil
}
