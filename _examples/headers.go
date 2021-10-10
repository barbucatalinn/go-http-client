package _examples

import (
	"context"
	"fmt"
	"net/http"

	"github.com/barbucatalinn/go-http-client/client"
	"github.com/sirupsen/logrus"
)

func headersExample() {
	// create the logger
	logger := logrus.New()

	// create the client
	c := client.NewClient(logger)

	// create a new request
	req, err := c.NewRequest(context.Background(), http.MethodGet, "https://test.api/products/1", nil)
	if err != nil {
		panic(err)
	}

	// set the headers
	req.SetHeader("key", "value")

	req.SetHeaders(map[string]string{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
	})

	// perform the request
	result, err := c.Do(req)
	if err != nil {
		panic(err)
	}

	// do something with the result
	fmt.Println(result)
}
