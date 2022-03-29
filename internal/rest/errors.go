package rest

import "errors"

var (
	ErrNotAuthenticated    = errors.New("Ошибка авторизации пользователя")
	ErrRequestNotSupported = errors.New("запрос не поддерживается")
)
