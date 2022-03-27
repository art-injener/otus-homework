package swagger

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag"
)

// @title Сервис социальной сети
// @description API сервиса социальной сети
func RegisterHandler(g *gin.Engine) {

	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	doc, _ := swag.ReadDoc()
	g.GET("/swagger-file", func(c *gin.Context) {
		_, _ = c.Writer.WriteString(doc)
		c.Status(http.StatusOK)
	})
}
