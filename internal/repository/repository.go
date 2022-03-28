package repository

import (
	"github.com/art-injener/otus-homework/internal/models"
	"github.com/art-injener/otus-homework/internal/models/request"
)

type AccountsRepository interface {
	GetAllAccounts() ([]*models.Account, error)
	GetAccountByID(id int) (*models.Account, error)
	AddAccount(*models.Account) error

	GetUserByEmail(email string) (*request.User, error)
	GetUserByID(id int) (*request.User, error)
	AddNewUser(*request.User) error
}
