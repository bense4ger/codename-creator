package main

import (
	"bense4ger/codename-creator/foo"
	"context"

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
