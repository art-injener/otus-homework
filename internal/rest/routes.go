package rest

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/art-injener/otus/internal/logger"
	"github.com/art-injener/otus/internal/rest/handlers"
	"github.com/art-injener/otus/internal/service"
)

func newAccountsRoutes(handler *gin.RouterGroup, service service.SocialNetworkService, log *logger.Logger) {
	r := handlers.NewAccountsHandler(service, log)

	h := handler.Group("/accounts")
	{
		h.GET("/all", r.GetAccounts)
		h.GET("/:id", r.GetAccountById)
		h.POST("/new", r.AddAccount)
	}
}

func HealthCheck(c *gin.Context) {
	v, err := json.Marshal(`{"health" : "ok"}`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
	}
	c.JSON(http.StatusOK, v)
}
