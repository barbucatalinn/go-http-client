package _examples

import (
	"context"
	"fmt"
	"github.com/barbucatalinn/go-http-client/client"
	"github.com/sirupsen/logrus"
)

func defaultRetryExample() {
	// create the logger
	logger := logrus.New()

	// create the client
	c := client.NewClient(logger).

		/*
			Available retry policies:

			client.DefaultRetryPolicy
		*/
		WithRetryPolicy(client.DefaultRetryPolicy)

	// perform the request
	result, err := c.Get(context.Background(), "https://test.api/products/1")
	if err != nil {
		panic(err)
	}

	// do something with the result
	fmt.Println(result)
}
