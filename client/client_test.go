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

func TestBaseClient_Get(t *testing.T) {
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

		response, err := c.Get(ctx, u)
		assert.Nil(t, err)
		assert.IsType(t, &Response{}, response)

		type product struct {
			Code string `json:"code"`
			Name string `json:"name"`
		}
		var p product

		err = response.UnmarshalJSONResponse(&p)
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

		response, err := c.Get(ctx, u)
		assert.NotNil(t, err)
		assert.IsType(t, &Response{}, response)
	})
}

func TestBaseClient_Head(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	logger := logrus.New()
	c := NewClient(logger).WithRetryMax(0)

	t.Run("success", func(t *testing.T) {
		mux, u, shutdown := setup()
		defer shutdown()

		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Header().Set("ETag", "342602-4f6-4db09b2978ec0")
			w.Header().Set("Last-Modified", "Thu, 25 Apr 2013 13:13:13 GMT")
			w.Header().Set("Content-Length", "1950")
			_, _ = fmt.Fprint(w, ``)
		})

		response, err := c.Head(ctx, u)
		assert.Nil(t, err)
		assert.IsType(t, &Response{}, response)
	})

	t.Run("error", func(t *testing.T) {
		mux, u, shutdown := setup()
		defer shutdown()

		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = fmt.Fprint(w, `Internal server error`)
		})

		response, err := c.Head(ctx, u)
		assert.NotNil(t, err)
		assert.IsType(t, &Response{}, response)
	})
}

func TestBaseClient_Post(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	logger := logrus.New()
	c := NewClient(logger).WithRetryMax(0)

	request := struct {
		Code string `json:"code"`
		Name string `json:"name"`
	}{
		Code: "pkg1",
		Name: "product 1",
	}

	t.Run("success", func(t *testing.T) {
		mux, u, shutdown := setup()
		defer shutdown()

		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusCreated)
			_, _ = fmt.Fprint(w, ``)
		})

		response, err := c.Post(ctx, u, "application/json", &request)
		assert.Nil(t, err)
		assert.IsType(t, &Response{}, response)
	})

	t.Run("error", func(t *testing.T) {
		mux, u, shutdown := setup()
		defer shutdown()

		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = fmt.Fprint(w, `Internal server error`)
		})

		response, err := c.Post(ctx, u, "application/json", &request)
		assert.NotNil(t, err)
		assert.IsType(t, &Response{}, response)
	})
}

func TestBaseClient_Put(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	logger := logrus.New()
	c := NewClient(logger).WithRetryMax(0)

	request := struct {
		Code string `json:"code"`
		Name string `json:"name"`
	}{
		Code: "pkg1",
		Name: "product 1",
	}

	t.Run("success", func(t *testing.T) {
		mux, u, shutdown := setup()
		defer shutdown()

		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, ``)
		})

		response, err := c.Put(ctx, u, "application/json", &request)
		assert.Nil(t, err)
		assert.IsType(t, &Response{}, response)
	})

	t.Run("error", func(t *testing.T) {
		mux, u, shutdown := setup()
		defer shutdown()

		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = fmt.Fprint(w, `Internal server error`)
		})

		response, err := c.Put(ctx, u, "application/json", &request)
		assert.NotNil(t, err)
		assert.IsType(t, &Response{}, response)
	})
}

func TestBaseClient_Patch(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	logger := logrus.New()
	c := NewClient(logger).WithRetryMax(0)

	request := struct {
		Name string `json:"name"`
	}{
		Name: "product 2",
	}

	t.Run("success", func(t *testing.T) {
		mux, u, shutdown := setup()
		defer shutdown()

		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, ``)
		})

		response, err := c.Patch(ctx, u, "application/json", &request)
		assert.Nil(t, err)
		assert.IsType(t, &Response{}, response)
	})

	t.Run("error", func(t *testing.T) {
		mux, u, shutdown := setup()
		defer shutdown()

		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = fmt.Fprint(w, `Internal server error`)
		})

		response, err := c.Patch(ctx, u, "application/json", &request)
		assert.NotNil(t, err)
		assert.IsType(t, &Response{}, response)
	})
}

func TestBaseClient_Delete(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	logger := logrus.New()
	c := NewClient(logger).WithRetryMax(0)

	t.Run("success", func(t *testing.T) {
		mux, u, shutdown := setup()
		defer shutdown()

		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNoContent)
			_, _ = fmt.Fprint(w, ``)
		})

		response, err := c.Delete(ctx, u)
		assert.Nil(t, err)
		assert.IsType(t, &Response{}, response)
	})

	t.Run("error", func(t *testing.T) {
		mux, u, shutdown := setup()
		defer shutdown()

		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = fmt.Fprint(w, `Internal server error`)
		})

		response, err := c.Delete(ctx, u)
		assert.NotNil(t, err)
		assert.IsType(t, &Response{}, response)
	})
}

func TestBaseClient_Options(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	logger := logrus.New()
	c := NewClient(logger).WithRetryMax(0)

	t.Run("success", func(t *testing.T) {
		mux, u, shutdown := setup()
		defer shutdown()

		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Allow", "GET,HEAD,POST,OPTIONS,TRACE")
			_, _ = fmt.Fprint(w, ``)
		})

		response, err := c.Options(ctx, u)
		assert.Nil(t, err)
		assert.IsType(t, &Response{}, response)
	})

	t.Run("error", func(t *testing.T) {
		mux, u, shutdown := setup()
		defer shutdown()

		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = fmt.Fprint(w, `Internal server error`)
		})

		response, err := c.Options(ctx, u)
		assert.NotNil(t, err)
		assert.IsType(t, &Response{}, response)
	})
}
