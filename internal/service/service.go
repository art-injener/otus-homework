package service

import (
	"context"

	"github.com/art-injener/otus-homework/internal/models"
	"github.com/art-injener/otus-homework/internal/models/request"
)

type SocialNetworkService interface {
	GetAllAccounts(context.Context) ([]*models.Account, error)
	GetAccountByID(context.Context, int) (*models.Account, error)
	AddNewAccount(context.Context, *models.Account) error
	GetAccountByUserID(context.Context, int) (*models.Account, error)
	UpdateAccount(context.Context, *models.Account) error

	GetUserByEmail(context.Context, string) (*request.User, error)
	GetUserByID(context.Context, int) (*request.User, error)
	AddNewUser(context.Context, *request.User) error
	ExistsUser(context.Context, *request.User) (bool, error)

	MakeFriends(context.Context, int, int) error
	GetFriends(context.Context, int) ([]*models.Account, error)
}
