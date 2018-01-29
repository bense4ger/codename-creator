package main

import (
	"context"

	"github.com/bense4ger/codename-creator/foo"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handleRequest(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		Body:       foo.GetMessage(),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handleRequest)
}
