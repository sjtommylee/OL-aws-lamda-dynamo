package service

import (
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// value vs key: values are mapping the attirubte values to the dynamodb table

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

func IntegerKey(name string, value int) AWSObject {
	return AWSObject{
		name: IntegerValue(value),
	}
}

// takes in a int value as input and returns a pointer to dynamo struct
// we are converting the input value into a string representation before passing it into AWS
func IntegerValue(value int) *dynamodb.AttributeValue {
	return &dynamodb.AttributeValue{
		N: aws.String(strconv.Itoa((value))),
	}
}

func Integer64Key(name string, value int64) AWSObject {
	return AWSObject{
		name: Integer64Value(value),
	}
}

func Integer64Value(value int64) *dynamodb.AttributeValue {
	return &dynamodb.AttributeValue{
		N: aws.String(strconv.FormatInt(value, 10)),
	}
}

// TODO: Blob value?
func Blobvalue(value []byte) *dynamodb.AttributeValue {
	return &dynamodb.AttributeValue{
		B: value,
	}
}

//TODO: The ReverseIndexInt64 function takes a slice of int64 values as input and returns a map that maps each value in the input slice to its corresponding index in the slice.
