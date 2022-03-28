package accounts

import (
	"github.com/art-injener/otus-homework/internal/models"
	"github.com/art-injener/otus-homework/internal/models/request"
)

type accountsRepositoryImpl struct {
}

func NewAccountsRepo() *accountsRepositoryImpl {
	return &accountsRepositoryImpl{}
}

func (r *accountsRepositoryImpl) GetAllAccounts() ([]*models.Account, error) {
	return nil, nil
}

func (r *accountsRepositoryImpl) GetAccountByID(id int) (*models.Account, error) {
	return &models.Account{}, nil
}

func (r *accountsRepositoryImpl) AddAccount(user *models.Account) error {
	return nil
}

func (r *accountsRepositoryImpl) GetUserByEmail(email string) (*request.User, error) {
	return nil, nil
}

func (r *accountsRepositoryImpl) GetUserByID(id int) (*request.User, error) {
	return nil, nil
}

func (r *accountsRepositoryImpl) AddNewUser(*request.User) error {
	return nil
}
