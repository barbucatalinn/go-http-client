package client

import (
	"net/http"
	"net/http/httptest"
)

func setup() (*http.ServeMux, string, func()) {
	mux := http.NewServeMux()

	h := http.NewServeMux()
	h.Handle("/", mux)

	server := httptest.NewServer(h)

	return mux, server.URL, server.Close
}
