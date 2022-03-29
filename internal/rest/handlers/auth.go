package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"

	"github.com/art-injener/otus-homework/internal/logger"
	"github.com/art-injener/otus-homework/internal/models/request"
	"github.com/art-injener/otus-homework/internal/service"
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

	user := request.User{}
	if err := ctx.BindJSON(&user); err != nil {
		logger.LogError(fmt.Errorf(errIncorrectDataForNewUser.Error(), err), u.log)
		ErrorResponse(ctx, http.StatusBadRequest, errIncorrectDataForNewUser)
		return
	}

	exists, err := u.service.ExistsUser(context, &user)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	if exists {
		ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Пользователь с указанным email уже существует"})
		return
	}

	if err = u.service.AddNewUser(context, &user); err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Пользователь успешно создан"})
}

func (u *users) LoginUser(ctx *gin.Context) {
	usr := request.User{}
	if err := ctx.BindJSON(&usr); err != nil {
		logger.LogError(fmt.Errorf(errIncorrectDataLogin.Error(), err), u.log)
		ErrorResponse(ctx, http.StatusBadRequest, errIncorrectDataLogin)
		return
	}

	user, err := u.service.GetUserByEmail(ctx.Request.Context(), usr.Email)
	if err != nil {
		if user != nil && !user.ComparePassword(usr.Password) {
			logger.LogInfo(fmt.Sprintf("Аутентификация пользователя %s не пройдена, пароли не совпадают", user.Email), u.log)
			ErrorResponse(ctx, http.StatusUnauthorized, errIncorrectEmailOrPassword)
		} else {
			ErrorResponse(ctx, http.StatusUnauthorized, err)
		}
		return
	}

	session, err := u.sessionStore.Get(ctx.Request, SessionName)
	if err != nil {
		logger.LogError(fmt.Errorf("Внутренняя ошибка сервера", err), u.log)
		ErrorResponse(ctx, http.StatusInternalServerError, errors.New("Внутренняя ошибка сервера"))
		return
	}
	session.Values["user_id"] = user.ID
	if err := u.sessionStore.Save(ctx.Request, ctx.Writer, session); err != nil {
		logger.LogError(fmt.Errorf("Внутренняя ошибка сервера", err), u.log)
		ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	account, err := u.service.GetAccountByUserID(ctx.Request.Context(), user.ID)
	if err != nil {
		// TODO: если у пользователя еще нет профиля, нужен редирект на страницу создания
		return
	}
	ctx.Redirect(http.StatusMovedPermanently, fmt.Sprintf("/v1/accounts/%d", account.ID))
}
