package models

type UserOption func(account *Account)

type Account struct {
	ID      uint64 `json:"-"`
	LoginID uint64 `json:"-"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Age     int    `json:"age"`
	Sex     string `json:"sex"`
	Hobby   string `json:"hobby"`
	City    string `json:"city"`
}

func newDefaultAccount() *Account {
	return &Account{
		Name:    "",
		Surname: "",
		Age:     18,
		Sex:     "male",
		Hobby:   "",
		City:    "",
	}
}

func NewAccount(options ...UserOption) *Account {
	user := newDefaultAccount()

	for _, userOption := range options {
		userOption(user)
	}
	return user
}
