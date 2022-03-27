package repository

import (
	"github.com/art-injener/otus/internal/models"
)

type AccountsRepositoryImpl struct {
}

func NewAccountsRepo() *AccountsRepositoryImpl {
	return &AccountsRepositoryImpl{}
}

func (r *AccountsRepositoryImpl) GetAll() ([]models.Account, error) {
	return nil, nil
}

func (r *AccountsRepositoryImpl) GetById(id uint64) (models.Account, error) {
	return models.Account{}, nil
}

func (r *AccountsRepositoryImpl) Add(user models.Account) error {
	return nil
}
