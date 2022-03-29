package useroption

import (
	"log"

	"github.com/art-injener/otus-homework-homework/internal/models"
)

func WithName(name string) models.UserOption {
	return func(a *models.Account) {
		a.Name = name
	}
}

func WithSurname(surname string) models.UserOption {
	return func(a *models.Account) {
		a.Surname = surname
	}
}

func WithAge(age int) models.UserOption {
	return func(a *models.Account) {
		if age < 0 {
			log.Fatal("Возраст должен быть больше 0!")
		}
		a.Age = age
	}
}

func WithSex(sex string) models.UserOption {
	return func(a *models.Account) {
		if sex != "male" || sex != "female" {
			log.Fatal("Значение должно быть либо \"male\", либо \"female\"")
		}
		a.Sex = sex
	}
}

func WithHobby(hobby string) models.UserOption {
	return func(a *models.Account) {
		a.Hobby = hobby
	}
}

func WithCity(city string) models.UserOption {
	return func(a *models.Account) {
		a.City = city
	}
}
