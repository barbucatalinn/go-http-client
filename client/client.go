package client

import (
	"context"
	"net/http"
)

// Get provides the functionality to send "GET" requests
func (c *BaseClient) Get(ctx context.Context, url string) (*Response, error) {
	// create a new request
	req, err := c.NewRequest(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	// set the headers
	req.SetHeaders(map[string]string{
		userAgentHeaderKey: userAgentHeaderValue,
	})

	return c.Do(req)
}

// Head provides the functionality to send "HEAD" requests
func (c *BaseClient) Head(ctx context.Context, url string) (*Response, error) {
	// create a new request
	req, err := c.NewRequest(ctx, http.MethodHead, url, nil)
	if err != nil {
		return nil, err
	}

	// set the headers
	req.SetHeaders(map[string]string{
		userAgentHeaderKey: userAgentHeaderValue,
	})

	return c.Do(req)
}

// Post provides the functionality to send "POST" requests
func (c *BaseClient) Post(ctx context.Context, url, contentType string, body interface{}) (*Response, error) {
	// create a new request
	req, err := c.NewRequest(ctx, http.MethodPost, url, body)
	if err != nil {
		return nil, err
	}

	// set the headers
	req.SetHeaders(map[string]string{
		contentTypeHeaderKey: contentType,
		userAgentHeaderKey:   userAgentHeaderValue,
	})

	return c.Do(req)
}

// Put provides the functionality to send "PUT" requests
func (c *BaseClient) Put(ctx context.Context, url, contentType string, body interface{}) (*Response, error) {
	// create a new request
	req, err := c.NewRequest(ctx, http.MethodPut, url, body)
	if err != nil {
		return nil, err
	}

	// set the headers
	req.SetHeaders(map[string]string{
		contentTypeHeaderKey: contentType,
		userAgentHeaderKey:   userAgentHeaderValue,
	})

	return c.Do(req)
}

// Patch provides the functionality to send "PATCH" requests
func (c *BaseClient) Patch(ctx context.Context, url, contentType string, body interface{}) (*Response, error) {
	// create a new request
	req, err := c.NewRequest(ctx, http.MethodPatch, url, body)
	if err != nil {
		return nil, err
	}

	// set the headers
	req.SetHeaders(map[string]string{
		contentTypeHeaderKey: contentType,
		userAgentHeaderKey:   userAgentHeaderValue,
	})

	return c.Do(req)
}

// Delete provides the functionality to send "DELETE" requests
func (c *BaseClient) Delete(ctx context.Context, url string) (*Response, error) {
	// create a new request
	req, err := c.NewRequest(ctx, http.MethodDelete, url, nil)
	if err != nil {
		return nil, err
	}

	// set the headers
	req.SetHeaders(map[string]string{
		userAgentHeaderKey: userAgentHeaderValue,
	})

	return c.Do(req)
}

// Options provides the functionality to send "OPTIONS" requests
func (c *BaseClient) Options(ctx context.Context, url string) (*Response, error) {
	// create a new request
	req, err := c.NewRequest(ctx, http.MethodOptions, url, nil)
	if err != nil {
		return nil, err
	}

	// set the headers
	req.SetHeaders(map[string]string{
		userAgentHeaderKey: userAgentHeaderValue,
	})

	return c.Do(req)
}
