package service

import "github.com/aws/aws-sdk-go/service/dynamodb"

type AWSObject = map[string]*dynamodb.AttributeValue
