package rest

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/sessions"

	"github.com/art-injener/otus-homework/internal/rest/handlers"
	"github.com/art-injener/otus-homework/internal/service"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("!!")
	}
}

func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := uuid.New().String()
		c.Writer.Header().Set("X-Request-Id", id)
		c.Set("request_id", id)
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers",
			"Content-Type, Content-Length, Accept-Encoding,"+
				" X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)

			return
		}
	}
}

func AuthMiddleware(service service.SocialNetworkService, sessions sessions.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		session, err := sessions.Get(c.Request, handlers.SessionName)
		if err != nil {
			handlers.ErrorResponse(c, http.StatusInternalServerError, err)
			return
		}
		id, ok := session.Values["user_id"]
		if !ok {
			handlers.ErrorResponse(c, http.StatusUnauthorized, ErrNotAuthenticated)
			return
		}

		user, err := service.GetUserByID(c.Request.Context(), id.(int))
		if err != nil {
			handlers.ErrorResponse(c, http.StatusUnauthorized, ErrNotAuthenticated)
			return
		}
		c.Set("auth_user", user)
	}
}
