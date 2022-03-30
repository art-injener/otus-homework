package repository

import (
	"context"

	"github.com/art-injener/otus-homework/internal/models"
	"github.com/art-injener/otus-homework/internal/models/request"
)

type AccountsRepository interface {
	GetAllAccounts(context.Context) ([]*models.Account, error)
	GetAccountByID(context.Context, int) (*models.Account, error)
	AddAccount(context.Context, *models.Account) error
	GetAccountByUserID(context.Context, int) (*models.Account, error)
	UpdateAccount(context.Context, *models.Account) error

	GetUserByEmail(context.Context, string) (*request.User, error)
	GetUserByID(context.Context, int) (*request.User, error)
	AddNewUser(context.Context, *request.User) error

	MakeFriends(context.Context, int, int) error
	IsFriends(context.Context, int, int) (bool, error)
	GetFriends(context.Context, int) ([]*models.Account, error)
}
