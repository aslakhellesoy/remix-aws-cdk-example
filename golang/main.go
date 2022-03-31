package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	body := event.Body
	headers := map[string]string{
		"Content-Type": event.Headers["content-type"],
	}
	return events.APIGatewayProxyResponse{
		StatusCode:      200,
		Headers:         headers,
		Body:            body,
		IsBase64Encoded: false,
	}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
