package handlers

import (
	"github.com/gin-gonic/gin"
)

type response struct {
	Error string `json:"error" example:"message"`
}

func ErrorResponseWithMessage(c *gin.Context, code int, msg string) {
	c.AbortWithStatusJSON(code, response{msg})
}

func ErrorResponse(c *gin.Context, code int, err error) {
	ErrorResponseWithMessage(c, code, err.Error())
}
