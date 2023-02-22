package service

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// define type alias for a map that maps strings to pointers of dynao db attributes
type AWSObject = map[string]*dynamodb.AttributeValue

// takes in a name, value returns an AWS Object with the provided value
func StringKey(name, value string) AWSObject {
	return AWSObject{
		name: StringValue(value),
	}
}

// takes a string input, returns a pointer to the dynamo struct
func StringValue(value string) *dynamodb.AttributeValue {
	return &dynamodb.AttributeValue{
		S: aws.String(value),
	}
}
