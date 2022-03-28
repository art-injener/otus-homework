package rest

import "errors"

var (
	ErrNotAuthenticated = errors.New("Ошибка авторизации пользователя")
)
