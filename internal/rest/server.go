package rest

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/art-injener/otus-homework/internal/service"
	"github.com/gorilla/sessions"

	"github.com/gin-gonic/gin"

	"github.com/art-injener/otus-homework/internal"
	"github.com/art-injener/otus-homework/internal/config"
	"github.com/art-injener/otus-homework/internal/rest/swagger"
)

func NewWebServer(service service.SocialNetworkService, session sessions.Store, cfg *config.Config) (*http.Server, error) {
	if cfg == nil {
		return nil, internal.ErrorEmptyConfig
	}
	mode := gin.ReleaseMode
	if cfg.IsDebug() {
		mode = gin.DebugMode
	}

	gin.SetMode(mode)

	router := gin.Default()
	router.Use(RequestIDMiddleware())
	router.Use(LoggerMiddleware())
	router.Use(CORSMiddleware())
	router.NoRoute(func(c *gin.Context) {
		c.AbortWithStatusJSON(http.StatusNotFound, errors.New("запрос не поддерживается"))
	})

	swagger.RegisterHandler(router)
	registerRoutes(router, service, session, cfg)

	return &http.Server{Addr: fmt.Sprintf(":%d", cfg.ServerPort), Handler: router}, nil
}

func registerRoutes(router *gin.Engine, service service.SocialNetworkService, session sessions.Store, cfg *config.Config) {
	h := router.Group("/v1")
	{
		newAccountsRoutes(h, service, session, cfg.Log)
		newUsersRoutes(h, service, session, cfg)
	}
}
