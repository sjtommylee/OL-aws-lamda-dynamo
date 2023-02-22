package util

func Headers() map[string]string {
	return map[string]string{
		"Access-Control-Allow-Origin":      "*",
		"Access-Control-Allow-Credentials": "true",
	}
}

func NewSuccessresponse(statusCode int, body interface{}) error {
	return nil
}
