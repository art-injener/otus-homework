package service

import "errors"

var (
	errGetAllAccounts = errors.New("Ошибка получения списка профилей")
	errGetAccountByID = errors.New("Ошибка получения профиля по id")
	errAddNewAccount  = errors.New("Ошибка добавления нового профиля")
)
