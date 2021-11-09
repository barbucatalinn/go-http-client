package _examples

import (
	"context"
	"fmt"

	"github.com/barbucatalinn/go-http-client/client"
	"github.com/sirupsen/logrus"
)

func putExample() {
	// create the logger
	logger := logrus.New()

	// create the client
	c := client.NewClient(logger)

	type product struct {
		Code string `json:"code"`
		Name string `json:"name"`
	}

	req := product{
		Code: "pkg8",
		Name: "Product 8",
	}

	// perform the request
	result, err := c.Put(context.Background(), "https://test.api/products/1", "application/json", &req)
	if err != nil {
		panic(err)
	}

	// do something with the result
	fmt.Println(result)
}
