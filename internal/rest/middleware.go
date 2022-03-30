package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/sessions"

	"github.com/art-injener/otus-homework/internal/rest/handlers"
	"github.com/art-injener/otus-homework/internal/service"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
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
			c.AbortWithStatus(http.StatusNoContent)

			return
		}
	}
}

func AuthMiddleware(svc service.SocialNetworkService, store sessions.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		session, err := store.Get(c.Request, handlers.SessionName)
		if err != nil {
			handlers.ErrorResponse(c, http.StatusInternalServerError, errInternalServerError)
			return
		}
		id, ok := session.Values["user_id"]
		if !ok {
			handlers.ErrorResponse(c, http.StatusUnauthorized, errNotAuthenticated)

			return
		}

		userId, ok := id.(int)
		if !ok {
			handlers.ErrorResponse(c, http.StatusUnauthorized, errNotAuthenticated)
			return
		}

		_, err = svc.GetUserByID(c.Request.Context(), userId)
		if err != nil {
			handlers.ErrorResponse(c, http.StatusUnauthorized, errNotAuthenticated)
			return
		}

		account, err := svc.GetAccountByUserID(c.Request.Context(), userId)
		if err != nil {
			handlers.ErrorResponse(c, http.StatusUnauthorized, errNotAuthenticated)
			return
		}

		c.Set("auth_user_id", id)
		c.Set("current_account", account)
	}
}
