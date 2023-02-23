package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/sjtommylee/go-aws-lamda-dynamo/util"
)

type Response struct {
	Body string `json:"body"`
}

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Received body: ", request.Body)

	response := Response{
		Body: "ping",
	}
	// return events.APIGatewayProxyResponse{Body: request.Body, StatusCode: 200}, nil
	return util.NewSuccessresponse(200, response)
}

func main() {
	lambda.Start(Handler)
}
