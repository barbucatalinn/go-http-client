package client

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io"
	"io/ioutil"
)

// drainBody reads the body
func drainBody(body io.ReadCloser) error {
	defer body.Close()
	_, err := io.Copy(ioutil.Discard, io.LimitReader(body, responseReadLimit))
	return err
}

// getBodyReader encodes the payload into the body reader
// and returns it
func getBodyReader(rawBody interface{}) (io.ReadWriter, error) {
	var bodyReader io.ReadWriter

	if rawBody != nil {
		bodyReader = new(bytes.Buffer)
		enc := json.NewEncoder(bodyReader)
		err := enc.Encode(rawBody)
		if err != nil {
			return nil, err
		}
	}
	return bodyReader, nil
}

// getContentLength returns the length of the payload encoded
// into the provided reader
func getContentLength(bodyReader io.Reader) (int64, error) {
	buf := new(bytes.Buffer)
	n, err := buf.ReadFrom(bodyReader)
	if err != nil {
		return 0, err
	}
	return n, nil
}

// basicAuth returns the basic auth token based on the username
// and password provided
func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}
