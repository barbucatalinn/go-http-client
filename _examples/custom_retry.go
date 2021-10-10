package _examples

import (
	"context"
	"fmt"
	"net/http"

	"github.com/barbucatalinn/go-http-client/client"
	"github.com/sirupsen/logrus"
)

func customRetryExample() {
	// create the logger
	logger := logrus.New()

	// create the client
	c := client.NewClient(logger).WithRetryPolicy(CustomRetryPolicy)

	// perform the request
	result, err := c.Get(context.Background(), "https://test.api/products/1")
	if err != nil {
		panic(err)
	}

	// do something with the result
	fmt.Println(result)
}

func CustomRetryPolicy(ctx context.Context, resp *http.Response, err error) (bool, error) {
	// do something here
	return true, err
}