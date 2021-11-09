package _examples

import (
	"context"
	"fmt"
	"time"

	"github.com/barbucatalinn/go-http-client/client"
	"github.com/sirupsen/logrus"
)

func customBackoffExample() {
	// create the logger
	logger := logrus.New()

	// create the client
	c := client.NewClient(logger).WithBackoffStrategy(CustomBackoffStrategy)

	// perform the request
	result, err := c.Get(context.Background(), "https://test.api/products/1")
	if err != nil {
		panic(err)
	}

	// do something with the result
	fmt.Println(result)
}

func CustomBackoffStrategy(i int) time.Duration {
	return time.Duration(i + 2) * time.Second
}
