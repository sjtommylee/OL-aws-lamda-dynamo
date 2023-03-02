package service

import (
	"sync"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// init

var once sync.Once
var client *dynamodb.DynamoDB

func initializeSingleton() {
	session := session.Must(session.NewSessionWithOptions(session.Options{}))
	client = dynamodb.New(session)
}

func DynamoClient() *dynamodb.DynamoDB {
	// making sure to initalice client
	once.Do(initializeSingleton)
	return client
}
