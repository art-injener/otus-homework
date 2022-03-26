package rest

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthCheck(c *gin.Context) {
	v, err := json.Marshal(`{"health" : "ok"}`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
	}
	c.JSON(http.StatusOK, v)
}
