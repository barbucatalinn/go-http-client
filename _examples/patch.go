package _examples

import (
	"context"
	"fmt"

	"github.com/barbucatalinn/go-http-client/client"
	"github.com/sirupsen/logrus"
)

func patchExample() {
	// create the logger
	logger := logrus.New()

	// create the client
	c := client.NewClient(logger)

	type Product struct {
		Code string `json:"code,omitempty"`
		Name string `json:"name,omitempty"`
	}

	req := Product{
		Code: "pkg8",
	}

	// perform the request
	result, err := c.Patch(context.Background(), "https://test.api/products/1", "application/json", &req)
	if err != nil {
		panic(err)
	}

	// do something with the result
	fmt.Println(result)
}