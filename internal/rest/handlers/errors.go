package handlers

import "errors"

var (
	ErrIncorrectEmailOrPassword = errors.New("Введенные логин или пароль некорректны")
)
