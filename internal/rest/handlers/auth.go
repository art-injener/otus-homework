package handlers

import (
	"fmt"
	"net/http"

	"github.com/art-injener/otus-homework/internal/models"

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
	user.Email = ctx.PostForm("email")
	user.Password = ctx.PostForm("password")
	user.RepeatedPassword = ctx.PostForm("repeated_password")

	exists, err := u.service.ExistsUser(context, &user)
	if err != nil {
		u.RegistrationFormWithError(ctx, err.Error())
		return
	}

	if exists {
		u.RegistrationFormWithError(ctx, "Пользователь с указанным email уже существует")
		return
	}

	if err = u.service.AddNewUser(context, &user); err != nil {
		u.RegistrationFormWithError(ctx, err.Error())

		return
	}

	account := models.NewDefaultAccount()
	account.LoginID = user.ID
	if err = u.service.AddNewAccount(context, account); err != nil {
		u.RegistrationFormWithError(ctx, err.Error())

		return
	}

	ctx.Redirect(http.StatusFound, "/v1/login")
}

func (u *users) LoginUser(ctx *gin.Context) {
	usr := request.User{
		Email:    ctx.Request.FormValue("email"),
		Password: ctx.Request.FormValue("password"),
	}

	user, err := u.service.GetUserByEmail(ctx.Request.Context(), usr.Email)
	if err != nil {
		logger.LogError(err, u.log)
		u.LoginFormWithError(ctx, err.Error())
	}

	if user != nil && !user.ComparePassword(usr.Password) {
		str := "Не удаётся войти. Пожалуйста, проверьте правильность написания логина и пароля."
		logger.LogInfo(str, u.log)
		u.LoginFormWithError(ctx, str)
	}

	session, err := u.sessionStore.Get(ctx.Request, SessionName)
	if err != nil {
		logger.LogError(err, u.log)
		ErrorResponse(ctx, http.StatusInternalServerError, errInternalServerError)
		return
	}

	session.Values["user_id"] = user.ID
	if err = u.sessionStore.Save(ctx.Request, ctx.Writer, session); err != nil {
		logger.LogError(err, u.log)
		ErrorResponse(ctx, http.StatusInternalServerError, err)

		return
	}

	account, err := u.service.GetAccountByUserID(ctx.Request.Context(), user.ID)
	if err != nil || account.IsEmptyFields() {
		ctx.Redirect(http.StatusMovedPermanently, "/v1/accounts/new")

		return
	}

	ctx.Redirect(http.StatusMovedPermanently, fmt.Sprintf("/v1/accounts/%d", account.ID))
}

func (u *users) LoginForm(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login.html",
		gin.H{})
}

func (u *users) LoginFormWithError(ctx *gin.Context, errText string) {

	ctx.HTML(http.StatusOK, "login.html",
		gin.H{ErrKey: errText})
}

func (u *users) RegistrationFormWithError(ctx *gin.Context, errText string) {
	ctx.HTML(http.StatusOK, "new_user.html",
		gin.H{ErrKey: errText})
}

func (u *users) Logout(ctx *gin.Context) {
	session, _ := u.sessionStore.Get(ctx.Request, SessionName)
	session.Values["user_id"] = nil
	if err := u.sessionStore.Save(ctx.Request, ctx.Writer, session); err != nil {
		logger.LogError(err, u.log)
		ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.Redirect(http.StatusFound, "/v1/login")
}
