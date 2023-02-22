package model

type User struct {
	Fullname     string
	Username     string
	Email        string
	PasswordHash []byte
	Bio          string
}

func (u *User) Validate() error {
	return nil
}

func ValidatePassword() error {
	return nil
}
