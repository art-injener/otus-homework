package handlers

import "errors"

var (
	errIncorrectEmailOrPassword   = errors.New("Введенные логин или пароль некорректны")
	errIncorrectDataForNewAccount = errors.New("Получены некорректные данные для добавления нового профиля")
	errIncorrectDataLogin         = errors.New("Получены некорректные данные для аутентификации")

	errIncorrectDataForNewUser = errors.New("Получены некорректные данные для добавления нового пользователя")
)
