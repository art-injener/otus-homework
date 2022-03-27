package service

import (
	"context"

	"github.com/art-injener/otus/internal/models"
)

type SocialNetworkService interface {
	GetAll(ctx context.Context) ([]models.Account, error)
	GetById(ctx context.Context, id uint64) (models.Account, error)
	Add(ctx context.Context, user models.Account) error
}
