package repository

import (
	"errors"
	"fmt"
	"github.com/art-injener/otus/internal/models"
)

type AccountsRepositoryMock struct {
	users map[uint64]models.Account
	id    uint64
}

func NewAccountsRepo() *AccountsRepositoryMock {
	repo := AccountsRepositoryMock{
		users: make(map[uint64]models.Account),
		id:    1,
	}
	return &repo
}

func (r *AccountsRepositoryMock) GetAll() ([]models.Account, error) {
	users := make([]models.Account, 0, len(r.users))
	for _, account := range r.users {
		users = append(users, account)
	}
	return users, nil
}

func (r *AccountsRepositoryMock) GetById(id uint64) (models.Account, error) {
	user, ok := r.users[id]
	if !ok {
		return models.Account{}, errors.New(fmt.Sprintf("Пользователь с id %d не найден", id))
	}
	return user, nil
}

func (r *AccountsRepositoryMock) Add(user models.Account) error {
	r.id++
	r.users[r.id] = user
	return nil
}
