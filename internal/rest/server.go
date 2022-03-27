package rest

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/art-injener/otus/internal/service"

	"github.com/gin-gonic/gin"

	"github.com/art-injener/otus/internal"
	"github.com/art-injener/otus/internal/config"
	"github.com/art-injener/otus/internal/rest/swagger"
)

func CreateWebServer(service service.SocialNetworkService, cfg *config.Config) (*http.Server, error) {
	if cfg == nil {
		return nil, internal.ErrorEmptyConfig
	}
	mode := gin.ReleaseMode
	if cfg.IsDebug() {
		mode = gin.DebugMode
	}

	gin.SetMode(mode)
	//TODO: при использовании gin.Default() внутри вызывается gin.New() и устанавливаются handlers gin.Logger() и gin.Recovery()
	router := gin.Default()

	router.Use(LoggerMiddleware())
	router.Use(AuthMiddleware())
	router.Use(CORSMiddleware())
	router.NoRoute(func(c *gin.Context) {
		c.AbortWithStatusJSON(http.StatusNotFound, errors.New("запрос не поддерживается"))
	})

	swagger.RegisterHandler(router)
	registerAccountsRoutes(router, service, cfg)

	return &http.Server{Addr: fmt.Sprintf(":%d", cfg.ServerPort), Handler: router}, nil
}

func registerAccountsRoutes(router *gin.Engine, service service.SocialNetworkService, cfg *config.Config) {
	h := router.Group("/v1")
	{
		newAccountsRoutes(h, service, cfg.Log)
	}
}
