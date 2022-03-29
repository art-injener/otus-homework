package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/art-injener/otus-homework/internal/logger"
	"github.com/art-injener/otus-homework/internal/models"
	"github.com/art-injener/otus-homework/internal/service"
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

	ctx.HTML(http.StatusOK, "users_list.html",
		gin.H{
			"users": users,
		})
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
	account := models.Account{}
	if err := ctx.BindJSON(&account); err != nil {
		logger.LogError(err, a.log)
		ErrorResponse(ctx, http.StatusBadRequest, errIncorrectDataForNewAccount)
		return
	}
	userId, _ := ctx.Get("auth_user_id")
	account.LoginID = userId.(int)

	if err := a.service.AddNewAccount(ctx.Request.Context(), &account); err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	//TODO: сделать redirects
}
