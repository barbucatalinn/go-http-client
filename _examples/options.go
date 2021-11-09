package _examples

import (
	"context"
	"fmt"

	"github.com/barbucatalinn/go-http-client/client"
	"github.com/sirupsen/logrus"
)

func optionsExample() {
	// create the logger
	logger := logrus.New()

	// create the client
	c := client.NewClient(logger)

	// perform the request
	result, err := c.Options(context.Background(), "https://test.api/products")
	if err != nil {
		panic(err)
	}

	// do something with the result
	fmt.Println(result)
}
