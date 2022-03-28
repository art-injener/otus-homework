package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/art-injener/otus-homework/internal/logger"
	"github.com/art-injener/otus-homework/internal/models/request"
	"github.com/art-injener/otus-homework/internal/repository"
	"github.com/art-injener/otus-homework/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

const (
	SessionName = "socialNetwork"
)

type users struct {
	service      service.SocialNetworkService
	log          *logger.Logger
	sessionStore sessions.Store
}

func NewUsersHandler(s service.SocialNetworkService, log *logger.Logger, sessionStore sessions.Store) *users {
	return &users{
		service:      s,
		log:          log,
		sessionStore: sessionStore,
	}
}

func (u *users) RegisterNewUser(ctx *gin.Context) {
	context := ctx.Request.Context()

	user := &request.User{}

	if err := ctx.BindJSON(user); err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, fmt.Errorf("Некорректные данные нового пользователя", err))
	}

	exists, err := u.service.ExistsUser(context, user)
	if err != nil {
		if !errors.Is(err, repository.ErrAccountNotFound) {
			ErrorResponse(ctx, http.StatusInternalServerError, fmt.Errorf("Ошибка поиска пользователя", err))
			return
		}
	}
	if exists {
		ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Пользователь с указанным email уже существует"})
		return
	}

	if err = u.service.AddNewUser(context, user); err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, fmt.Errorf("Ошибка добавления нового пользователя", err))
		return
	}
	user.Sanitize()
	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Пользователь успешно создан"})
}

func (u *users) LoginUser(ctx *gin.Context) {
	usr := &request.User{}
	if err := ctx.BindJSON(usr); err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, fmt.Errorf("Некорректные данные нового пользователя", err))
		return
	}

	user, err := u.service.GetUserByEmail(ctx.Request.Context(), usr.Email)
	if err != nil || !user.ComparePassword(usr.Password) {
		ErrorResponse(ctx, http.StatusUnauthorized, ErrIncorrectEmailOrPassword)
		return
	}

	session, err := u.sessionStore.Get(ctx.Request, SessionName)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	session.Values["user_id"] = user.ID
	if err := u.sessionStore.Save(ctx.Request, ctx.Writer, session); err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
}
