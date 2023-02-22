package model

import "encoding/json"

type InputError map[string][]string

func (e InputError) Error() string {
	js, err := json.Marshal(e)
	if err != nil {
		return err.Error()
	}
	return string(js)
}

// constructor that returns new InputError instance
func NewInputError(name, message string) InputError {
	return InputError{
		name: {message},
	}
}
