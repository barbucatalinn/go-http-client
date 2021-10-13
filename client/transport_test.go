//go:build !integration
// +build !integration

package client

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_createHTTPTransport(t *testing.T) {
	t.Parallel()

	result := createHTTPTransport()
	assert.IsType(t, &http.Transport{}, result)
}
