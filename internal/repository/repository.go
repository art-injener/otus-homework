package repository

import (
	"context"

	"github.com/art-injener/otus-homework/internal/models"
	"github.com/art-injener/otus-homework/internal/models/request"
)

type AccountsRepository interface {
	GetAllAccounts(ctx context.Context) ([]*models.Account, error)
	GetAccountByID(ctx context.Context, id int) (*models.Account, error)
	AddAccount(context.Context, *models.Account) error
	GetAccountByUserID(context.Context, int) (*models.Account, error)

	GetUserByEmail(ctx context.Context, email string) (*request.User, error)
	GetUserByID(ctx context.Context, id int) (*request.User, error)
	AddNewUser(context.Context, *request.User) error
}
