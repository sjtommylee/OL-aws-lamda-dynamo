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
	if u.Username == "" {
		return NewInputError("username", "cannot be blank")
	}

	if u.Email == "" {
		return NewInputError("Email", "cannot be blank")
	}

	return nil
}

func ValidatePassword(password string) error {
	if len(password) < MinLength {
		return NewInputError("password", "must be specific length")
	}
	return nil
}
