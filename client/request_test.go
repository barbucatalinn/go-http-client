//go:build !integration
// +build !integration

package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBaseClient_NewRequest(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	c := &BaseClient{}
	u, _ := url.Parse("https://app.local")

	t.Run("GET request", func(t *testing.T) {
		result, err := c.NewRequest(ctx, http.MethodGet, u.String(), nil)
		assert.Nil(t, err)

		r, _ := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)

		assert.IsType(t, &Request{}, result)
		assert.Nil(t, result.body)
		assert.Equal(t, int64(0), result.contentLength)
		assert.Equal(t, r, result.Request)
	})

	t.Run("POST request", func(t *testing.T) {
		type product struct {
			Code string `json:"code"`
			Name string `json:"name"`
		}

		req := product{
			Code: "pkg1",
			Name: "product 1",
		}

		result, err := c.NewRequest(ctx, http.MethodPost, u.String(), &req)
		assert.Nil(t, err)

		requestByte, err := json.Marshal(&req)
		assert.Nil(t, err)

		b := bytes.NewReader(requestByte)
		r, _ := http.NewRequestWithContext(ctx, http.MethodPost, u.String(), nil)

		assert.Nil(t, err)
		assert.IsType(t, &Request{}, result)
		assert.Equal(t, b, result.body)
		assert.Equal(t, int64(34), result.contentLength)
		assert.Equal(t, r, result.Request)
	})
}

func TestRequest_SetHeader(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	c := &BaseClient{}
	u, _ := url.Parse("https://app.local")

	k := "key"
	v := "value"

	req, err := c.NewRequest(ctx, http.MethodPost, u.String(), nil)
	assert.Nil(t, err)

	assert.Empty(t, req.Header.Get(k))

	req = req.SetHeader(k, v)
	assert.Equal(t, req.Header.Get(k), v)
}

func TestRequest_SetHeaders(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	c := &BaseClient{}
	u, _ := url.Parse("https://app.local")

	k1 := "key1"
	v1 := "value1"

	k2 := "key2"
	v2 := "value2"

	req, err := c.NewRequest(ctx, http.MethodPost, u.String(), nil)
	assert.Nil(t, err)

	assert.Empty(t, req.Header.Get(k1), req.Header.Get(k2))

	req = req.SetHeaders(map[string]string{
		k1: v1,
		k2: v2,
	})
	assert.Equal(t, req.Header.Get(k1), v1)
	assert.Equal(t, req.Header.Get(k2), v2)
}

func TestRequest_setupAuth(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	c := &BaseClient{}
	u, _ := url.Parse("https://app.local")

	scheme := "my-schema"
	token := "secret"

	req, err := c.NewRequest(ctx, http.MethodPost, u.String(), nil)
	assert.Nil(t, err)

	err = req.setupAuth(scheme, token)
	assert.Nil(t, err)

	assert.Equal(t, fmt.Sprintf("%s %s", scheme, token), req.Header.Get(authorizationHeaderKey))
}
