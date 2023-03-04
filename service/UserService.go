package service

import "github.com/sjtommylee/go-aws-lamda-dynamo/model"

func PutUser(user model.User) error {
	err := user.Validate()
	if err != nil {
		return err
	}

	return nil
}

func GetCurrentuser(auth string) (*model.User, string, error) {
	// username, token, err := model.VerifyAuth(auth)
	// return &user, token, nil
	return nil
}
