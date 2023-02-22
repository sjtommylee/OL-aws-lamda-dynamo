package service

// https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_GetItem.html
func GetItemByKey(tableName string, key AWSObject, out interface{}) (bool, error) {
	return true, nil
}

// https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_Query.html
func QueryItems() {}

// https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_BatchGetItem.html
func BatchGetItems() {}

// https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_PutItem.html
func PutItem() {}
