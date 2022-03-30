package accountoption

import (
	"log"

	"github.com/art-injener/otus-homework/internal/models"
)

const (
	femaleTypeName = "f"
	maleTypeName   = "m"
)

func WithName(name string) models.AccountOption {
	return func(a *models.Account) {
		a.Name = name
	}
}

func WithSurname(surname string) models.AccountOption {
	return func(a *models.Account) {
		a.Surname = surname
	}
}

func WithSex(sex string) models.AccountOption {
	return func(a *models.Account) {
		if sex != maleTypeName && sex != femaleTypeName {
			log.Println("Значение должно быть либо \"male\", либо \"female\"")

			return
		}
		a.Sex = sex
	}
}

func WithHobby(hobby string) models.AccountOption {
	return func(a *models.Account) {
		a.Hobby = hobby
	}
}

func WithCity(city string) models.AccountOption {
	return func(a *models.Account) {
		a.City = city
	}
}
