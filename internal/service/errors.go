package service

import "errors"

var (
	errGetAllAccounts = errors.New("Ошибка получения списка профилей")
	errGetAccount     = errors.New("Ошибка получения профиля")
	errAddNewAccount  = errors.New("Ошибка добавления нового профиля")

	errRegistrationNewUser = errors.New("Ошибка регистрации нового пользователя")
	errValidationUser      = errors.New("Ошибка валидации данных нового пользователя при регистрации")
	errGetDataUser         = errors.New("Ошибка получения данных пользователя")
	errPasswordsNotEquals  = errors.New("Введенные пароли не совпадают")
)
