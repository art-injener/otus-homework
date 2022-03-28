package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/art-injener/otus-homework/internal/logger"
	"github.com/art-injener/otus-homework/internal/models"
	"github.com/art-injener/otus-homework/internal/service"
	"github.com/gin-gonic/gin"
)

type accounts struct {
	service service.SocialNetworkService
	log     *logger.Logger
}

func NewAccountsHandler(s service.SocialNetworkService, log *logger.Logger) *accounts {
	return &accounts{
		service: s,
		log:     log,
	}
}

func (a *accounts) GetAccounts(ctx *gin.Context) {
	users, err := a.service.GetAllAccounts(ctx.Request.Context())
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	ctx.IndentedJSON(http.StatusOK, users)
}

func (a *accounts) GetAccountById(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logger.LogError(err, a.log)
		ErrorResponse(ctx, http.StatusBadRequest, err)
	}
	user, err := a.service.GetAccountById(ctx.Request.Context(), id)
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	ctx.IndentedJSON(http.StatusOK, user)
}

func (a *accounts) AddAccount(ctx *gin.Context) {
	var user *models.Account
	if err := ctx.BindJSON(&user); err != nil {
		logger.LogError(err, a.log)
		ErrorResponse(ctx, http.StatusBadRequest, fmt.Errorf("Получено некорректное сообщение", err))
		return
	}
	if err := a.service.AddNewAccount(ctx.Request.Context(), user); err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	//TODO: сделать redirects
}
