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

func newAccountsRoutes(handler *gin.RouterGroup, service service.SocialNetworkService, session sessions.Store, log *logger.Logger) {
	r := handlers.NewAccountsHandler(service, log)

	h := handler.Group("/accounts")
	h.Use(AuthMiddleware(service, session))
	{
		h.GET("/all", r.GetAccounts)
		h.GET("/:id", r.GetAccountById)
		h.POST("/new", r.AddAccount)
	}
}

func newUsersRoutes(handler *gin.RouterGroup, service service.SocialNetworkService, session sessions.Store, cfg *config.Config) {
	usersHandler := handlers.NewUsersHandler(service, cfg.Log, session)

	handler.POST("/registration", usersHandler.RegisterNewUser)
	handler.POST("/login", usersHandler.LoginUser)
}

func HealthCheck(c *gin.Context) {
	v, err := json.Marshal(`{"health" : "ok"}`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
	}
	c.JSON(http.StatusOK, v)
}
