//go:build !integration
// +build !integration

package client

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	t.Parallel()

	logger := logrus.New()

	result := NewClient(logger)
	assert.IsType(t, &BaseClient{}, result)
	assert.IsType(t, &http.Client{}, result.hc)
	assert.Equal(t, defaultRetryMax, result.retryMax)
	assert.IsType(t, new(RetryPolicy), &result.retryPolicy)
	assert.IsType(t, new(BackoffStrategy), &result.backoffStrategy)
	assert.Equal(t, auth{}, result.auth)
	assert.IsType(t, &logrus.Logger{}, result.logger)
}

func TestBaseClient_WithRetryMax(t *testing.T) {
	t.Parallel()

	logger := logrus.New()

	c := NewClient(logger)
	assert.Equal(t, defaultRetryMax, c.retryMax)

	c = c.WithRetryMax(defaultRetryMax + 1)
	assert.Equal(t, defaultRetryMax+1, c.retryMax)
}

func TestBaseClient_WithBackoffStrategy(t *testing.T) {
	t.Parallel()

	logger := logrus.New()

	c := NewClient(logger)
	assert.IsType(t, new(BackoffStrategy), &c.backoffStrategy)

	c = c.WithBackoffStrategy(ExponentialBackoffStrategy)
	assert.IsType(t, new(BackoffStrategy), &c.backoffStrategy)
}

func TestBaseClient_WithRetryPolicy(t *testing.T) {
	t.Parallel()

	logger := logrus.New()

	c := NewClient(logger)
	assert.IsType(t, new(RetryPolicy), &c.retryPolicy)

	c = c.WithRetryPolicy(DefaultRetryPolicy)
	assert.IsType(t, new(RetryPolicy), &c.retryPolicy)
}

func TestBaseClient_WithBasicAuth(t *testing.T) {
	t.Parallel()

	logger := logrus.New()

	c := NewClient(logger)
	c.WithBasicAuth("u", "p")

	assert.Equal(t, auth{basicAuthScheme, "dTpw"}, c.auth)
}

func TestBaseClient_WithBearerAuth(t *testing.T) {
	t.Parallel()

	logger := logrus.New()

	c := NewClient(logger)
	c.WithBearerAuth("token")

	assert.Equal(t, auth{bearerAuthScheme, "token"}, c.auth)
}

func TestBaseClient_WithCustomAuth(t *testing.T) {
	t.Parallel()

	logger := logrus.New()

	c := NewClient(logger)
	c.WithCustomAuth("scheme", "token")

	assert.Equal(t, auth{"scheme", "token"}, c.auth)
}

func TestBaseClient_Do(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	logger := logrus.New()

	c := NewClient(logger).WithRetryMax(0)

	t.Run("success", func(t *testing.T) {
		mux, u, shutdown := setup()
		defer shutdown()

		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)

			_, _ = fmt.Fprint(w, `
						{
						  "code": "code 1",
						  "name": "name 1"
						}`)
		})

		req, err := c.NewRequest(ctx, http.MethodGet, u, nil)
		assert.Nil(t, err)

		response, err := c.Do(req)
		assert.Nil(t, err)
		assert.IsType(t, &Response{}, response)

		type product struct {
			Code string `json:"code"`
			Name string `json:"name"`
		}
		var p product

		err = response.UnmarshalJSON(&p)
		assert.Nil(t, err)
		assert.Equal(t, product{"code 1", "name 1"}, p)
	})

	t.Run("error", func(t *testing.T) {
		mux, u, shutdown := setup()
		defer shutdown()

		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = fmt.Fprint(w, `Internal server error`)
		})

		req, err := c.NewRequest(ctx, http.MethodGet, u, nil)
		assert.Nil(t, err)

		response, err := c.Do(req)
		assert.NotNil(t, err)
		assert.IsType(t, &Response{}, response)
	})

}

func Test_getHTTPClient(t *testing.T) {
	t.Parallel()

	result := getHTTPClient()
	assert.IsType(t, &http.Client{}, result)
}

func TestBaseClient_getLogger(t *testing.T) {
	t.Parallel()

	logger := logrus.New()
	c := NewClient(logger)

	result := c.getLogger()
	assert.IsType(t, &logrus.Logger{}, result)
	assert.Equal(t, logger, result)
}
