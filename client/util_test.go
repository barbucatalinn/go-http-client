//go:build !integration
// +build !integration

package client

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_drainBody(t *testing.T) {
	t.Parallel()

	body := io.NopCloser(strings.NewReader("test"))
	err := drainBody(body)
	assert.Nil(t, err)

	b, err := ioutil.ReadAll(body)
	assert.Nil(t, err)
	assert.Empty(t, b)
}

func Test_getBodyReader(t *testing.T) {
	t.Parallel()

	type product struct {
		Code string `json:"code"`
		Name string `json:"name"`
	}

	body := product{
		Code: "pkg1",
		Name: "Product 1",
	}

	br, err := getBodyReader(&body)
	assert.Nil(t, err)

	b, err := ioutil.ReadAll(br)
	assert.Nil(t, err)

	var r product

	err = json.Unmarshal(b, &r)
	assert.Nil(t, err)
	assert.Equal(t, r, body)
}

func Test_getContentLength(t *testing.T) {
	t.Parallel()

	type product struct {
		Code string `json:"code"`
		Name string `json:"name"`
	}

	body := product{
		Code: "pkg1",
		Name: "Product 1",
	}

	br, err := getBodyReader(&body)
	assert.Nil(t, err)

	result, err := getContentLength(br)
	assert.Nil(t, err)
	assert.Equal(t, int64(34), result)
}

func Test_basicAuth(t *testing.T) {
	t.Parallel()

	result := basicAuth("u", "p")
	assert.Equal(t, "dTpw", result)
}
