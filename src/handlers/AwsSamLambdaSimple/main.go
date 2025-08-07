package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
)

var (
	lambdaStart = lambda.Start
)

func handler(ctx context.Context) (string, error) {
	return "Hello, World!", nil
}

func main() {
	lambdaStart(handler)
}
