package service

import (
	"context"

	"github.com/art-injener/otus/internal/logger"
	"github.com/art-injener/otus/internal/models"
	"github.com/art-injener/otus/internal/repository"
)

type UserService struct {
	repository repository.AccountsRepository
	log        *logger.Logger
}

func NewUserService(repository repository.AccountsRepository, log *logger.Logger) *UserService {
	return &UserService{
		repository: repository,
		log:        log,
	}
}

func (s *UserService) GetAll(ctx context.Context) ([]models.Account, error) {
	return s.repository.GetAll()
}

func (s *UserService) GetById(ctx context.Context, id uint64) (models.Account, error) {
	return s.repository.GetById(uint64(id))
}

func (s *UserService) Add(ctx context.Context, user models.Account) error {
	return s.repository.Add(user)
}
