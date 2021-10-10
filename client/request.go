package client

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

// Request wraps the metadata needed to create HTTP requests
type Request struct {
	body io.ReadWriter

	*http.Request
}

// NewRequest creates a new wrapped request with the provided context
func (c *BaseClient) NewRequest(ctx context.Context, method, url string, rawBody interface{}) (*Request, error) {
	// get the body reader
	bodyReader, err := getBodyReader(rawBody)
	if err != nil {
		return nil, err
	}

	// create a http request with context
	req, err := http.NewRequestWithContext(ctx, method, url, nil)
	if err != nil {
		return nil, err
	}

	// set the content length
	if method != http.MethodGet && method != http.MethodDelete && rawBody != nil {
		cl, err := getContentLength(bodyReader)
		if err != nil {
			return nil, err
		}
		req.ContentLength = cl
	}

	return &Request{bodyReader, req}, nil
}

// SetHeader method is to set a single header key/value pair
func (r *Request) SetHeader(key, value string) *Request {
	r.Header.Set(key, value)
	return r
}

// SetHeaders method sets multiple headers key/value pairs
func (r *Request) SetHeaders(headers map[string]string) *Request {
	for k, v := range headers {
		r.SetHeader(k, v)
	}
	return r
}

// setupAuth setups the auth for the request
func (r *Request) setupAuth(scheme, token string) error {
	if scheme != "" && token != "" {
		r.SetHeader(authorizationHeaderKey, fmt.Sprintf("%s %s", scheme, token))
	}
	return nil
}
