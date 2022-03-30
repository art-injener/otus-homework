package mock

import (
	"context"

	"github.com/art-injener/otus-homework/internal/models"
	"github.com/art-injener/otus-homework/internal/models/request"
	"github.com/art-injener/otus-homework/internal/repository"
)

type accountsRepositoryMock struct {
	accounts map[int]*models.Account
	users    map[int]*request.User
}

var _ repository.AccountsRepository = &accountsRepositoryMock{}

// NewAccountsRepo ...
func NewAccountsRepo() *accountsRepositoryMock {
	repo := accountsRepositoryMock{
		accounts: make(map[int]*models.Account),
		users:    make(map[int]*request.User),
	}

	return &repo
}

func (r *accountsRepositoryMock) GetAllAccounts(_ context.Context) ([]*models.Account, error) {
	users := make([]*models.Account, 0, len(r.accounts))
	for _, account := range r.accounts {
		users = append(users, account)
	}

	return users, nil
}

func (r *accountsRepositoryMock) GetAccountByID(_ context.Context, id int) (*models.Account, error) {
	user, ok := r.accounts[id]
	if !ok {
		return &models.Account{}, repository.ErrAccountNotFound
	}

	return user, nil
}

func (r *accountsRepositoryMock) AddAccount(_ context.Context, user *models.Account) error {
	user.ID = len(r.accounts)
	r.accounts[len(r.accounts)] = user

	return nil
}

func (r *accountsRepositoryMock) GetAccountByUserID(_ context.Context, userID int) (*models.Account, error) {
	for _, v := range r.accounts {
		if v.LoginID == userID {
			return v, nil
		}
	}

	return nil, repository.ErrAccountNotFound
}

func (r *accountsRepositoryMock) UpdateAccount(context.Context, *models.Account) error {
	return nil
}

func (r *accountsRepositoryMock) GetUserByEmail(_ context.Context, email string) (*request.User, error) {
	for _, v := range r.users {
		if v.Email == email {
			return v, nil
		}
	}

	return nil, repository.ErrAccountNotFound
}

func (r *accountsRepositoryMock) GetUserByID(_ context.Context, id int) (*request.User, error) {
	if user, ok := r.users[id]; ok {
		return user, nil
	}

	return nil, repository.ErrUserNotFound
}

func (r *accountsRepositoryMock) AddNewUser(_ context.Context, user *request.User) error {
	user.ID = len(r.users)
	r.users[len(r.users)] = user

	return nil
}

func (r *accountsRepositoryMock) MakeFriends(_ context.Context, _, _ int) error {
	return nil
}

func (r *accountsRepositoryMock) GetFriends(_ context.Context, _ int) ([]*models.Account, error) {
	return nil, nil
}

func (r *accountsRepositoryMock) IsFriends(ctx context.Context, currentUserID, friendID int) (bool, error) {
	return false, nil
}
