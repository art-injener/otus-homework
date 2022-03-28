package service

import (
	"context"
	"fmt"

	"github.com/art-injener/otus-homework/internal/logger"
	"github.com/art-injener/otus-homework/internal/models"
	"github.com/art-injener/otus-homework/internal/models/request"
	"github.com/art-injener/otus-homework/internal/repository"
)

type user struct {
	repository repository.AccountsRepository
	log        *logger.Logger
}

func NewUserService(repository repository.AccountsRepository, log *logger.Logger) *user {
	return &user{
		repository: repository,
		log:        log,
	}
}

func (s *user) GetAllAccounts(ctx context.Context) ([]*models.Account, error) {
	accounts, err := s.repository.GetAllAccounts()
	if err != nil {
		logger.LogError(fmt.Errorf(errGetAllAccounts.Error(), err), s.log)
		return nil, errGetAllAccounts
	}
	return accounts, nil
}

func (s *user) GetAccountById(ctx context.Context, id int) (*models.Account, error) {
	account, err := s.repository.GetAccountByID(id)
	if err != nil {
		logger.LogError(fmt.Errorf(errGetAccountByID.Error(), err), s.log)
		return nil, errGetAccountByID
	}
	return account, nil
}

func (s *user) AddNewAccount(ctx context.Context, user *models.Account) error {
	err := s.repository.AddAccount(user)
	if err != nil {
		logger.LogError(fmt.Errorf(errAddNewAccount.Error(), err), s.log)
		return errAddNewAccount
	}
	return nil
}

func (s *user) GetUserByEmail(ctx context.Context, email string) (*request.User, error) {
	return s.repository.GetUserByEmail(email)
}

func (s *user) GetUserByID(ctx context.Context, id int) (*request.User, error) {
	return s.repository.GetUserByID(id)
}

func (s *user) AddNewUser(ctx context.Context, user *request.User) error {
	if err := user.Validate(); err != nil {
		return err
	}

	if err := user.BeforeCreate(); err != nil {
		return err
	}
	return s.repository.AddNewUser(user)
}

func (s *user) ExistsUser(ctx context.Context, user *request.User) (bool, error) {
	userByEmail, err := s.GetUserByEmail(ctx, user.Email)
	if err != nil {
		return false, err
	}

	if userByEmail != nil {
		return true, nil
	}
	return false, nil
}
