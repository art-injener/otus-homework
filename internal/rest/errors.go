package rest

import "errors"

var (
	errNotAuthenticated    = errors.New("пользователь не авторизован")
	errRequestNotSupported = errors.New("запрос не поддерживается")
	errInternalServerError = errors.New("внутренняя ошибка сервера")
)
