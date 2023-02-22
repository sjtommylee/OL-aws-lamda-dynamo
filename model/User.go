package model

const MinLength = 0
const MaxLength = 64

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

func ValidatePassword(password string) error {
	return nil
}
