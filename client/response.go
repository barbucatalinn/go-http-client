package client

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// DataDump is a struct containing the request and the response
type DataDump struct {
	RequestDump  []byte
	ResponseDump []byte
}

// Response is a wrapper for the response
type Response struct {
	RawResponse *http.Response

	DataDump *DataDump
}

// GetStatus returns the status string of the response
func (r *Response) GetStatus() string {
	return r.RawResponse.Status
}

// GetStatusCode returns the status code of the response
func (r *Response) GetStatusCode() int {
	return r.RawResponse.StatusCode
}

// GetHeaders returns the header map of the response
func (r *Response) GetHeaders() http.Header {
	return r.RawResponse.Header
}

// GetBody returns the body as []byte array
func (r *Response) GetBody() ([]byte, error) {
	if r.RawResponse == nil {
		return []byte{}, nil
	}
	defer r.RawResponse.Body.Close()
	return ioutil.ReadAll(r.RawResponse.Body)
}

// GetStringBody returns the body as string
func (r *Response) GetStringBody() (string, error) {
	b, err := r.GetBody()
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// UnmarshalJSON unmarshalls the response body into the provided target object
func (r *Response) UnmarshalJSON(target interface{}) error {
	b, err := r.GetBody()
	if err != nil {
		return err
	}
	return json.Unmarshal(b, &target)
}
