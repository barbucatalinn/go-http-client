package _examples

import (
	"context"
	"fmt"

	"github.com/barbucatalinn/go-http-client/client"
	"github.com/sirupsen/logrus"
)

func postExample() {
	// create the logger
	logger := logrus.New()

	// create the client
	c := client.NewClient(logger)

	type product struct {
		Code string `json:"code"`
		Name string `json:"name"`
	}

	req := product{
		Code: "pkg1",
		Name: "Product 1",
	}

	// perform the request
	result, err := c.Post(context.Background(), "https://test.api/products", "application/json", &req)
	if err != nil {
		panic(err)
	}

	// do something with the result
	fmt.Println(result)
}
