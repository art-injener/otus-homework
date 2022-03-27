package handlers

import (
	"fmt"
	"github.com/art-injener/otus/internal/logger"
	"github.com/art-injener/otus/internal/models"
	"github.com/art-injener/otus/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type AccountsHandler struct {
	service service.SocialNetworkService
	log     *logger.Logger
}

func NewAccountsHandler(s service.SocialNetworkService, log *logger.Logger) *AccountsHandler {
	return &AccountsHandler{
		service: s,
		log:     log,
	}
}

func (a *AccountsHandler) GetAccounts(ctx *gin.Context) {
	users, err := a.service.GetAll(ctx.Request.Context())
	if err != nil {
		logger.LogError(err, a.log)
		ErrorResponse(ctx, http.StatusBadRequest, "Ошибка при получения списка пользователей")
	}
	ctx.IndentedJSON(http.StatusOK, users)
}

func (a *AccountsHandler) GetAccountById(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logger.LogError(err, a.log)
		ErrorResponse(ctx, http.StatusBadRequest, "Неверный параметр id")
	}
	user, err := a.service.GetById(ctx.Request.Context(), uint64(id))
	if err != nil {
		logger.LogError(err, a.log)
		ErrorResponse(ctx, http.StatusBadRequest, fmt.Sprintf("Пользователь с id = %d не найден", id))
	}
	ctx.IndentedJSON(http.StatusOK, user)
}

func (a *AccountsHandler) AddAccount(ctx *gin.Context) {
	var user models.Account
	if err := ctx.BindJSON(&user); err != nil {
		logger.LogError(err, a.log)
		ErrorResponse(ctx, http.StatusBadRequest, "Получено некорректное сообщение")
	}
	if err := a.service.Add(ctx.Request.Context(), user); err != nil {
		logger.LogError(err, a.log)
		ErrorResponse(ctx, http.StatusInternalServerError, "Ошибка при добавлении нового пользователя")
	}
}
