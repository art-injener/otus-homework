package handlers

import "errors"

var (
	errIncorrectDataForNewAccount = errors.New("получены некорректные данные для добавления нового профиля")
	errAccountIsExists            = errors.New("профиль для пользователя уже существует")

	errInternalServerError = errors.New("внутренняя ошибка сервера")
)
