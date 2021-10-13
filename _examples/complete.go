package _examples

import (
	"context"
	"fmt"
	"net/http"

	"github.com/barbucatalinn/go-http-client/client"
	"github.com/sirupsen/logrus"
)

func completeExample() {
	// create the logger
	logger := logrus.New()

	// create the client
	c := client.NewClient(logger).
		// set the basic auth
		WithBasicAuth("username", "password").
		// set the maximum number of retries
		WithRetryMax(10).
		// set the retry policy
		WithRetryPolicy(client.DefaultRetryPolicy).
		// set the backoff strategy
		WithBackoffStrategy(client.LinearBackoffStrategy)

	// create a new request
	req, err := c.NewRequest(context.Background(), http.MethodGet, "https://test.api/products/1", nil)
	if err != nil {
		panic(err)
	}

	// set the headers
	req.SetHeader("kev", "value")

	// perform the request
	result, err := c.Do(req)
	if err != nil {
		panic(err)
	}

	// unmarshal
	type apiResponse struct {
		Field1 string `json:"field1"`
		Field2 string `json:"field2"`
	}

	var rsp apiResponse

	err = result.UnmarshalJSON(rsp)
	if err != nil {
		panic(rsp)
	}
	fmt.Println(rsp.Field1)
	fmt.Println(rsp.Field2)
}
