package _examples

import (
	"context"
	"fmt"

	"github.com/barbucatalinn/go-http-client/client"
	"github.com/sirupsen/logrus"
)

func retryMaxExample() {
	// create the logger
	logger := logrus.New()

	// create the client
	c := client.NewClient(logger).WithRetryMax(10)

	// perform the request
	result, err := c.Get(context.Background(), "https://test.api/products/1")
	if err != nil {
		panic(err)
	}

	// do something with the result
	fmt.Println(result)
}
