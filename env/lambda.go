//go:build release

package env

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
)

type HandleRequester struct {
	Handler func() (string, error)
}

type Event struct {
	Name string `json:"name"`
}

func Main(handler func() (string, error)) {
	HandleRequest := func(_ context.Context, _ Event) (string, error) { return handler() }
	lambda.Start(HandleRequest)
}
