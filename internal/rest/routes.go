package rest

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"

	"github.com/art-injener/otus-homework/internal/config"
	"github.com/art-injener/otus-homework/internal/logger"
	"github.com/art-injener/otus-homework/internal/rest/handlers"
	"github.com/art-injener/otus-homework/internal/service"
)

func newAccountsRoutes(handler *gin.RouterGroup, srv service.SocialNetworkService, session sessions.Store, log *logger.Logger) {
	r := handlers.NewAccountsHandler(srv, log)

	h := handler.Group("/accounts")
	h.Use(AuthMiddleware(srv, session))
	{
		h.GET("/all", r.GetAccounts)
		h.GET("/:id", r.GetAccountByID)
		h.POST("/new", r.AddAccount)
		h.GET("/new", func(context *gin.Context) {
			context.HTML(http.StatusOK, "new_account.html", nil)
		})
		h.GET("make-friend/:friend-id", r.MakeFriends)
	}
}

func newUsersRoutes(handler *gin.RouterGroup, srv service.SocialNetworkService, session sessions.Store, cfg *config.Config) {
	usersHandler := handlers.NewUsersHandler(srv, cfg.Log, session)

	handler.POST("/registration", usersHandler.RegisterNewUser)
	handler.GET("/registration", func(context *gin.Context) {
		context.HTML(http.StatusOK, "new_user.html", nil)
	})
	handler.POST("/login", usersHandler.LoginUser)
	handler.GET("/login", usersHandler.LoginForm)
	handler.GET("/logout", usersHandler.Logout)
}

func HealthCheck(c *gin.Context) {
	v, err := json.Marshal(`{"health" : "ok"}`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
	}
	c.JSON(http.StatusOK, v)
}
