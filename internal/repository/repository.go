package repository

import (
	"github.com/art-injener/otus/internal/models"
)

type AccountsRepository interface {
	GetAll() ([]models.Account, error)
	GetById(id uint64) (models.Account, error)
	Add(user models.Account) error
}
