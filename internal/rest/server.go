package rest

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"

	"github.com/art-injener/otus-homework/internal"
	"github.com/art-injener/otus-homework/internal/config"
	"github.com/art-injener/otus-homework/internal/service"
)

func NewWebServer(svc service.SocialNetworkService, session sessions.Store, cfg *config.Config) (*http.Server, error) {
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
		c.AbortWithStatusJSON(http.StatusNotFound, errRequestNotSupported)
	})

	registerRoutes(router, svc, session, cfg)

	router.Static("/css", "./static/css")
	router.Static("/js", "./static/js")
	router.StaticFile("/favicon.ico", "./img/favicon.ico")

	router.LoadHTMLGlob("templates/*")

	return &http.Server{Addr: fmt.Sprintf(":%d", cfg.ServerPort), Handler: router}, nil
}

func registerRoutes(router *gin.Engine, svc service.SocialNetworkService, session sessions.Store, cfg *config.Config) {
	routerGroup := router.Group("/v1")
	{
		newAccountsRoutes(routerGroup, svc, session, cfg.Log)
		newUsersRoutes(routerGroup, svc, session, cfg)
	}

	router.GET("/", func(ctx *gin.Context) {
		currentAccount, _ := ctx.Get("current_account")

		ctx.HTML(http.StatusOK, "index.html",
			gin.H{
				"currentAccount": currentAccount,
			})
	})
}
