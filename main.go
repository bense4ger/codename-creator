package main

import (
	"context"

	"github.com/bense4ger/codename-creator/foo"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/bense4ger/codename-creator/responder"
)

func handleRequest(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		Body:       responder.GetMessage(),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handleRequest)
}
