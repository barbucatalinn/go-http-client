package _examples

import (
	"context"
	"fmt"

	"github.com/barbucatalinn/go-http-client/client"
	"github.com/sirupsen/logrus"
)

func basicAuthExample() {
	// create the logger
	logger := logrus.New()

	// create the client
	c := client.NewClient(logger).WithBasicAuth("username", "password")

	// perform the request
	result, err := c.Get(context.Background(), "https://test.api/products/1")
	if err != nil {
		panic(err)
	}

	// do something with the result
	fmt.Println(result)
}

func bearerAuthExample() {
	// create the logger
	logger := logrus.New()

	// create the client
	c := client.NewClient(logger).WithBearerAuth("secret")

	// perform the request
	result, err := c.Get(context.Background(), "https://test.api/products/1")
	if err != nil {
		panic(err)
	}

	// do something with the result
	fmt.Println(result)
}

func customAuthExample() {
	// create the logger
	logger := logrus.New()

	// create the client
	c := client.NewClient(logger).WithCustomAuth("my-auth-scheme", "secret")

	// perform the request
	result, err := c.Get(context.Background(), "https://test.api/products/1")
	if err != nil {
		panic(err)
	}

	// do something with the result
	fmt.Println(result)
}
