package rest

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/art-injener/otus/internal"
	"github.com/art-injener/otus/internal/config"
)

func CreateWebServer(cfg *config.Config) (*http.Server, error) {
	if cfg == nil {
		return nil, internal.ErrorEmptyConfig
	}
	mode := gin.ReleaseMode
	if cfg.IsDebug() {
		mode = gin.DebugMode
	}

	gin.SetMode(mode)
	router := gin.New()

	router.Use(LoggerMiddleware())
	router.Use(AuthMiddleware())
	router.Use(CORSMiddleware())
	router.NoRoute(func(c *gin.Context) {
		c.AbortWithStatusJSON(http.StatusNotFound, errors.New("запрос не поддерживается"))
	})

	registerRoutes(router)

	return &http.Server{Addr: fmt.Sprintf(":%d", cfg.ServerPort), Handler: router}, nil
}

func registerRoutes(g *gin.Engine) {
	g.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Hello world")
	})
}
