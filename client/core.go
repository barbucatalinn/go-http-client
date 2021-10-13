package client

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

// auth struct holds the auth data
type auth struct {
	Scheme string
	Token  string
}

// BaseClient wraps the http.Client and exposes all the functionality of the http.Client
// but with additional functionality
type BaseClient struct {
	// http client
	hc *http.Client

	// maximum number of retries
	retryMax int

	// retry policy
	retryPolicy RetryPolicy

	// backoff strategy
	backoffStrategy BackoffStrategy

	// auth
	auth auth

	// logger
	logger *logrus.Logger

	clientInit sync.Once
}

// NewClient creates a new BaseClient with default values
func NewClient(l *logrus.Logger) *BaseClient {
	return &BaseClient{
		hc:              getHTTPClient(),
		retryMax:        defaultRetryMax,
		retryPolicy:     DefaultRetryPolicy,
		backoffStrategy: DefaultBackoffStrategy,
		logger:          l,
	}
}

// WithRetryMax sets the RetryMax value and returns the BaseClient
func (c *BaseClient) WithRetryMax(retryMax int) *BaseClient {
	if retryMax >= 0 {
		c.retryMax = retryMax
	}
	return c
}

// WithBackoffStrategy sets the backoff value and returns the BaseClient
func (c *BaseClient) WithBackoffStrategy(backoffStrategy BackoffStrategy) *BaseClient {
	c.backoffStrategy = backoffStrategy
	return c
}

// WithRetryPolicy sets the retry value and returns the BaseClient
func (c *BaseClient) WithRetryPolicy(retryPolicy RetryPolicy) *BaseClient {
	c.retryPolicy = retryPolicy
	return c
}

// WithBasicAuth sets the auth object values (basic auth) and returns the BaseClient
func (c *BaseClient) WithBasicAuth(username, password string) *BaseClient {
	c.auth = auth{basicAuthScheme, basicAuth(username, password)}
	return c
}

// WithBearerAuth sets the auth object values (bearer auth) and returns the BaseClient
func (c *BaseClient) WithBearerAuth(token string) *BaseClient {
	c.auth = auth{bearerAuthScheme, token}
	return c
}

// WithCustomAuth sets the auth object values (custom auth) and returns the BaseClient
func (c *BaseClient) WithCustomAuth(scheme, token string) *BaseClient {
	c.auth = auth{scheme, token}
	return c
}

// Do wraps calling an HTTP method with retries
func (c *BaseClient) Do(req *Request) (*Response, error) {
	// get the logger
	logger := c.getLogger()

	// log the action
	logger.Debugf("%s %s", req.Method, req.URL)

	// setup auth
	_ = req.setupAuth(c.auth.Scheme, c.auth.Token)

	// re-create the http client
	c.clientInit.Do(func() {
		if c.hc == nil {
			c.hc = getHTTPClient()
		}
	})

	var respObj Response
	var resp *http.Response
	var dataDump DataDump
	var attempt int
	var shouldRetry bool
	var doErr, retryErr error

	// set the request dump
	dataDump.RequestDump, _ = httputil.DumpRequestOut(req.Request, req.body != nil)

	for i := 0; ; i++ {
		attempt++

		var code int // HTTP response code

		// always rewind the request body when non-nil
		if req.body != nil {
			req.Body = ioutil.NopCloser(req.body)
		}

		// attempt the request
		resp, doErr = c.hc.Do(req.Request)
		if resp != nil {
			code = resp.StatusCode
		}

		// check the retry
		shouldRetry, retryErr = c.retryPolicy(req.Context(), resp, doErr)

		if doErr != nil {
			logger.WithError(doErr).Errorf("%s %s request failed", req.Method, req.URL)
		}

		if !shouldRetry {
			break
		}

		// check the remaining number of retries
		remain := c.retryMax - i
		if remain == 0 {
			break
		}

		// consume any response to reuse the connection
		if doErr == nil {
			drainBodyErr := drainBody(resp.Body)
			if drainBodyErr != nil {
				logger.WithError(drainBodyErr).Error("error reading response body")
			}
		}

		var wait time.Duration

		bsc := backoffStatusCheck(resp)
		if bsc != nil {
			wait = *bsc
		} else {
			wait = c.backoffStrategy(i)
		}

		desc := fmt.Sprintf("%s %s", req.Method, req.URL)
		if code > 0 {
			desc = fmt.Sprintf("%s (status: %d)", desc, code)
		}
		logger.Debugf("%s: retrying in %s (%d left)", desc, wait, remain)

		select {
		case <-req.Context().Done():
			c.hc.CloseIdleConnections()
			return nil, req.Context().Err()
		case <-time.After(wait):
		}

		// make shallow copy of http.Request so that we can modify its body
		// without racing against the closeBody call in persistConn.writeLoop
		httpReq := *req.Request
		req.Request = &httpReq
	}

	// set the raw response
	respObj.RawResponse = resp

	// return successful response
	if doErr == nil && retryErr == nil && !shouldRetry {

		// set the response dump
		dataDump.ResponseDump, _ = httputil.DumpResponse(resp, true)

		// set data dump
		respObj.DataDump = &dataDump

		// return the response
		return &respObj, nil
	}

	defer c.hc.CloseIdleConnections()

	err := doErr
	if retryErr != nil {
		err = retryErr
	}

	// consume the response
	if resp != nil {
		drainBodyErr := drainBody(resp.Body)
		if drainBodyErr != nil {
			logger.WithError(drainBodyErr).Error("error reading response body")
		}
	}

	logger.Debugf("%s %s giving up after %d attempt(s)", req.Method, req.URL, attempt)

	// return the error
	return &respObj, err
}

// getHTTPClient returns a new http.Client with similar default
// values to http.Client but with a custom http.Transport
func getHTTPClient() *http.Client {
	return &http.Client{
		Transport: createHTTPTransport(),
	}
}

// getLogger returns the logger
func (c *BaseClient) getLogger() *logrus.Logger {
	return c.logger
}
