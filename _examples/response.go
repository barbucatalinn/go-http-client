package _examples

import (
	"context"
	"fmt"

	"github.com/barbucatalinn/go-http-client/client"
	"github.com/sirupsen/logrus"
)

func responseExample() {
	// create the logger
	logger := logrus.New()

	// create the client
	c := client.NewClient(logger)

	// perform the request
	result, err := c.Get(context.Background(), "https://test.api/products/1")
	if err != nil {
		panic(err)
	}

	// do something with the result

	// get the status
	status := result.GetStatus()
	fmt.Println(status)

	// get the status code
	statusCode := result.GetStatusCode()
	fmt.Println(statusCode)

	// get the headers
	headers := result.GetHeaders()
	fmt.Println(headers)

	// get the body
	body, err := result.GetBody()
	if err != nil {
		panic(err)
	}
	fmt.Println(body)

	// get the string body
	sBody, err := result.GetStringBody()
	if err != nil {
		panic(err)
	}
	fmt.Println(sBody)

	// unmarshal
	type apiResponse struct {
		Field1 string `json:"field1"`
		Field2 string `json:"field2"`
	}

	var rsp apiResponse

	err = result.UnmarshalJson(rsp)
	if err != nil {
		panic(rsp)
	}
	fmt.Println(rsp.Field1)
	fmt.Println(rsp.Field2)
}
