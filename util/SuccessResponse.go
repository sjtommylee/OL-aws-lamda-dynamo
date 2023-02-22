package util

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

func Headers() map[string]string {
	return map[string]string{
		"Access-Control-Allow-Origin":      "*",
		"Access-Control-Allow-Credentials": "true",
	}
}

func NewSuccessresponse(statusCode int, body interface{}) (events.APIGatewayProxyResponse, error) {
	response := events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Headers:    Headers(),
	}

	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return NewErrorResponse(err)
		}

		response.Body = string(jsonBody)
	}

	return response, nil
}
