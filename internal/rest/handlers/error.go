package handlers

import (
	"github.com/gin-gonic/gin"
)

var ErrKey = "error"

type response struct {
	Error string `json:"error" example:"message"`
}

func ErrorResponseWithMessage(c *gin.Context, code int, msg string) {
	c.HTML(
		code,
		"error.html",
		gin.H{
			ErrKey: response{
				Error: msg},
		})
}

func ErrorResponse(c *gin.Context, code int, err error) {
	ErrorResponseWithMessage(c, code, err.Error())
}
