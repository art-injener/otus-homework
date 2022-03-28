package repository

import "errors"

var (
	ErrAccountNotFound = errors.New("аккаунт не найден")
	ErrUserNotFound    = errors.New("пользователь не найден")
)
