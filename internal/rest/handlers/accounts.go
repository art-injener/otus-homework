package handlers

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"image/jpeg"
	"image/png"
	"io"
	"mime/multipart"
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

	account, _ := ctx.Get("current_account")

	ctx.HTML(http.StatusOK, "users_list.html",
		gin.H{
			"users":          users,
			"currentAccount": account,
		})
}

func (a *accounts) GetAccountByID(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logger.LogError(err, a.log)
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	account, err := a.service.GetAccountByID(ctx.Request.Context(), id)
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	if account.Avatar == "" {
		account.Avatar = defaultAvatar
	}

	friends, err := a.service.GetFriends(ctx.Request.Context(), account.ID)
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	currentAccount, _ := ctx.Get("current_account")

	ctx.HTML(http.StatusOK, "user_page.html",
		gin.H{
			"user":           account,
			"friends":        friends,
			"currentAccount": currentAccount,
		})
}

func (a *accounts) AddAccount(ctx *gin.Context) {
	context := ctx.Request.Context()
	userID, _ := ctx.Get("auth_user_id")

	acc, err := a.service.GetAccountByUserID(context, userID.(int))
	if err != nil {
		logger.LogError(err, a.log)
		a.NewAccountWithError(ctx, errAccountIsExists.Error())
		return
	}

	if acc != nil && !acc.IsEmptyFields() {
		logger.LogInfo(errAccountIsExists.Error(), a.log)
		a.NewAccountWithError(ctx, errAccountIsExists.Error())
		return
	}

	form, err := ctx.MultipartForm()
	if err != nil {
		panic(err)
	}

	accountMap := make(map[string]interface{})
	for i := range form.Value {
		accountMap[i] = form.Value[i][0]
	}

	if accountMap["sex"] == "U" {
		a.NewAccountWithError(ctx, "Половая принадлежность профиля не выбрана")
		return
	}

	bytes, err := json.Marshal(accountMap)
	if err != nil {
		logger.LogError(err, a.log)
		//ErrorResponse(ctx, http.StatusBadRequest, errIncorrectDataForNewAccount)
		a.NewAccountWithError(ctx, errIncorrectDataForNewAccount.Error())
		return
	}

	account := models.Account{}
	if err = json.Unmarshal(bytes, &account); err != nil {
		logger.LogError(err, a.log)
		a.NewAccountWithError(ctx, errIncorrectDataForNewAccount.Error())
		//	ErrorResponse(ctx, http.StatusBadRequest, errIncorrectDataForNewAccount)
		return
	}

	account.LoginID = userID.(int)

	saveAvatar(ctx, form, &account)

	if acc.IsEmptyFields() {
		err = a.service.UpdateAccount(context, &account)
	} else {
		err = a.service.AddNewAccount(context, &account)
	}

	if err != nil {
		//		ErrorResponse(ctx, http.StatusInternalServerError, err)
		a.NewAccountWithError(ctx, err.Error())
		return
	}
	ctx.Redirect(http.StatusFound, fmt.Sprintf("/v1/accounts/%d", account.ID))
}

func (a *accounts) MakeFriends(ctx *gin.Context) {
	friendIDStr := ctx.Param("friend-id")
	friendID, err := strconv.Atoi(friendIDStr)
	if err != nil {
		logger.LogError(err, a.log)
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	userId := ctx.Value("auth_user_id").(int)
	account, err := a.service.GetAccountByUserID(ctx.Request.Context(), userId)
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	err = a.service.MakeFriends(ctx.Request.Context(), account.ID, friendID)
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	ctx.Redirect(http.StatusFound, "/v1/accounts/all")
}

func saveAvatar(ctx *gin.Context, form *multipart.Form, account *models.Account) {
	fileHeaders := form.File["avatar"]

	if len(fileHeaders) == 0 {
		account.Avatar = defaultAvatar
		return
	}

	file, err := fileHeaders[0].Open()
	if err != nil {
		ErrorResponseWithMessage(ctx, http.StatusInternalServerError, "Внутренняя ошибка сервера")
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		ErrorResponseWithMessage(ctx, http.StatusInternalServerError, "Внутренняя ошибка сервера")
		return
	}
	contentType := http.DetectContentType(data)

	switch contentType {
	case "image/png":
		fmt.Println("Image type is already PNG.")
	case "image/jpeg":
		img, err := jpeg.Decode(bytes.NewReader(data))
		if err != nil {
			ErrorResponseWithMessage(ctx, http.StatusInternalServerError, "Внутренняя ошибка сервера: unable to decode jpeg")
			return
		}

		var buf bytes.Buffer
		if err := png.Encode(&buf, img); err != nil {
			ErrorResponseWithMessage(ctx, http.StatusInternalServerError, "Внутренняя ошибка сервера : unable to encode png")
			return
		}
		data = buf.Bytes()
	default:
		ErrorResponseWithMessage(ctx, http.StatusInternalServerError, "Внутренняя ошибка сервера : unsupported content type")
	}
	if len(data) == 0 {
		account.Avatar = defaultAvatar
	} else {
		account.Avatar = base64.StdEncoding.EncodeToString(data)
	}
}

func (a *accounts) NewAccountWithError(ctx *gin.Context, errText string) {
	ctx.HTML(http.StatusOK, "new_account.html",
		gin.H{ErrKey: errText})
}
