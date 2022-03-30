package models

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

type AccountOption func(account *Account)

const (
	layoutISO = "2006-01-02"
)

type Account struct {
	ID       int    `json:"id"`
	LoginID  int    `json:"-"`
	Avatar   string `json:"avatar"`
	Name     string `json:"name" binding:"required"`
	Surname  string `json:"surname" binding:"required"`
	Birthday string `json:"birthday" binding:"required"`
	Sex      string `json:"sex" binding:"required,len=1"`
	Hobby    string `json:"hobby" binding:"required"`
	City     string `json:"city" binding:"required"`
}

func NewDefaultAccount() *Account {
	return &Account{
		Birthday: layoutISO,
		Sex:      "m",
	}
}

func NewAccount(options ...AccountOption) *Account {
	user := NewDefaultAccount()

	for _, userOption := range options {
		userOption(user)
	}

	return user
}

func (a *Account) FormattedBirthday(format string) {
	formattedBirthday, err := time.Parse(layoutISO, a.Birthday)
	if err != nil {
		return
	}
	a.Birthday = formattedBirthday.Format(format)
}

func (a *Account) IsEmptyFields() bool {
	return a.Name == "" && a.Surname == "" && a.City == "" && a.Hobby == ""
}

func (a *Account) Validate() error {
	return validation.ValidateStruct(a,
		validation.Field(&a.Name, validation.Required),
		validation.Field(&a.Surname, validation.Required),
		validation.Field(&a.Birthday, validation.Required),
	)
}
