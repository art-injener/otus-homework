package service

import "errors"

var (
	errGetAllAccounts    = errors.New("ошибка получения списка профилей")
	errGetAccount        = errors.New("ошибка получения профиля")
	errAddNewAccount     = errors.New("ошибка добавления нового профиля")
	errValidationAccount = errors.New("ошибка валидации данных профиля. Имя, фамилия и дата рождения не должны быть пустыми")
	errUpdateAccount     = errors.New("ошибка обновления данных профиля")

	errRegistrationNewUser = errors.New("ошибка регистрации нового пользователя")
	errValidationUser      = errors.New("ошибка валидации данных нового пользователя при регистрации")
	errGetDataUser         = errors.New("ошибка получения данных пользователя")
	errPasswordsNotEquals  = errors.New("введенные пароли не совпадают")

	errMakeFriends = errors.New("ошибка создания дружбы между пользователями")
	errGetFriends  = errors.New("ошибка получения списка друзей")
)
