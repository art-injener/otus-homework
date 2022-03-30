package request

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID                int    `json:"-"`
	Email             string `json:"email"`
	Password          string `json:"password,omitempty"`
	RepeatedPassword  string `json:"repeated_password,omitempty"`
	EncryptedPassword string `json:"-"`
}

const (
	passwordMinLength = 6
	passwordMaxLength = 100
)

func (u *User) Validate() error {
	return validation.ValidateStruct(
		u,
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Password, validation.Required, validation.Length(passwordMinLength, passwordMaxLength)),
		validation.Field(&u.RepeatedPassword, validation.Required, validation.Length(passwordMinLength, passwordMaxLength)))
}

func (u *User) BeforeCreate() error {
	if len(u.Password) > 0 {
		enc, err := encryptString(u.Password)

		if err != nil {
			return err
		}

		u.EncryptedPassword = enc
	}

	return nil
}

func (u *User) ComparePassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.EncryptedPassword), []byte(password)) == nil
}

func encryptString(password string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
