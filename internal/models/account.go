package models

type AccountOption func(account *Account)

type Account struct {
	ID      int    `json:"id"`
	LoginID int    `json:"-"`
	Avatar  string `json:"avatar"`
	Name    string `json:"name" binding:"required"`
	Surname string `json:"surname" binding:"required"`
	Age     int    `json:"age" binding:"required"`
	Sex     rune   `json:"sex" binding:"required,len=1"`
	Hobby   string `json:"hobby" binding:"required"`
	City    string `json:"city" binding:"required"`
}

func newDefaultAccount() *Account {
	return &Account{
		Age: 0,
		Sex: 'm',
	}
}

func NewAccount(options ...AccountOption) *Account {
	user := newDefaultAccount()

	for _, userOption := range options {
		userOption(user)
	}
	return user
}
