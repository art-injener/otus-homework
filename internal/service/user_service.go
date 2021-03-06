package service

import (
	"context"
	"errors"

	"github.com/art-injener/otus-homework/internal/logger"
	"github.com/art-injener/otus-homework/internal/models"
	"github.com/art-injener/otus-homework/internal/models/request"
	"github.com/art-injener/otus-homework/internal/repository"
)

type user struct {
	repository repository.AccountsRepository
	log        *logger.Logger
}

func NewUserService(repo repository.AccountsRepository, log *logger.Logger) *user {
	return &user{
		repository: repo,
		log:        log,
	}
}

func (s *user) GetAllAccounts(ctx context.Context) ([]*models.Account, error) {
	accounts, err := s.repository.GetAllAccounts(ctx)
	if err != nil {
		logger.LogError(err, s.log)

		return nil, errGetAllAccounts
	}

	return accounts, nil
}

func (s *user) GetAccountByID(ctx context.Context, id int) (*models.Account, error) {
	account, err := s.repository.GetAccountByID(ctx, id)
	if err != nil {
		logger.LogError(err, s.log)

		return nil, errGetAccount
	}

	return account, nil
}

func (s *user) AddNewAccount(ctx context.Context, user *models.Account) error {
	err := s.repository.AddAccount(ctx, user)
	if err != nil {
		logger.LogError(err, s.log)

		return errAddNewAccount
	}

	return nil
}

func (s *user) GetAccountByUserID(ctx context.Context, userID int) (*models.Account, error) {
	account, err := s.repository.GetAccountByUserID(ctx, userID)
	if err != nil {
		logger.LogError(err, s.log)

		return nil, errGetAccount
	}

	return account, nil
}

func (s *user) UpdateAccount(ctx context.Context, acc *models.Account) error {
	if err := acc.Validate(); err != nil {
		logger.LogError(err, s.log)

		return errValidationAccount
	}

	err := s.repository.UpdateAccount(ctx, acc)
	if err != nil {
		logger.LogError(err, s.log)

		return errUpdateAccount
	}

	curAccount, err := s.GetAccountByUserID(ctx, acc.LoginID)
	if err != nil {
		logger.LogError(err, s.log)

		return errUpdateAccount
	}
	acc.ID = curAccount.ID
	return nil
}

func (s *user) GetUserByEmail(ctx context.Context, email string) (*request.User, error) {
	user, err := s.repository.GetUserByEmail(ctx, email)
	if err != nil {
		logger.LogError(err, s.log)

		return nil, errGetDataUser
	}

	return user, nil
}

func (s *user) GetUserByID(ctx context.Context, id int) (*request.User, error) {
	user, err := s.repository.GetUserByID(ctx, id)
	if err != nil {
		logger.LogError(err, s.log)

		return nil, errGetDataUser
	}

	return user, nil
}

func (s *user) AddNewUser(ctx context.Context, user *request.User) error {
	if err := user.Validate(); err != nil {
		logger.LogError(err, s.log)

		return errValidationUser
	}

	if user.Password != user.RepeatedPassword {
		return errPasswordsNotEquals
	}

	if err := user.BeforeCreate(); err != nil {
		logger.LogError(err, s.log)

		return errRegistrationNewUser
	}

	err := s.repository.AddNewUser(ctx, user)
	if err != nil {
		logger.LogError(err, s.log)

		return errRegistrationNewUser
	}

	return nil
}

func (s *user) ExistsUser(ctx context.Context, user *request.User) (bool, error) {
	userByEmail, err := s.GetUserByEmail(ctx, user.Email)
	if err != nil {
		if !errors.Is(err, errGetDataUser) {
			return false, err
		}
	}

	if userByEmail != nil {
		return true, nil
	}

	return false, nil
}

func (s *user) MakeFriends(ctx context.Context, currentAccountID, friendID int) error {
	err := s.repository.MakeFriends(ctx, currentAccountID, friendID)
	if err != nil {
		logger.LogError(err, s.log)

		return errMakeFriends
	}

	return nil
}

func (s *user) GetFriends(ctx context.Context, accountID int) (accounts []*models.Account, err error) {
	friends, err := s.repository.GetFriends(ctx, accountID)
	if err != nil {
		logger.LogError(err, s.log)

		return nil, errGetFriends
	}

	return friends, nil
}
