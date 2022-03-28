package models

type AccountOption func(account *Account)

type Account struct {
	ID      int    `json:"-"`
	LoginID int    `json:"-"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Age     int    `json:"age"`
	Sex     string `json:"sex"`
	Hobby   string `json:"hobby"`
	City    string `json:"city"`
}

func newDefaultAccount() *Account {
	return &Account{
		Age: 0,
		Sex: "m",
	}
}

func NewAccount(options ...AccountOption) *Account {
	user := newDefaultAccount()

	for _, userOption := range options {
		userOption(user)
	}
	return user
}
