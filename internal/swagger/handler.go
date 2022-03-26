package swagger

import "net/http"

// @title Сервис социальной сети
// @description API сервиса социальной сети
func registerHandler( //nolint
	r *http.Server,
) {
	//api.SwaggerInfo.Version = info.Version
	//api.SwaggerInfo.BasePath = "/api/v1"
	//
	//r.Gin.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//
	//doc, _ := swag.ReadDoc()
	//r.Gin.GET("/swagger-file", func(c *gin.Context) {
	//	_, _ = c.Writer.WriteString(doc)
	//	c.Status(http.StatusOK)
	//})
	//
	//logging.Tracef("Swagger - %s/swagger/index.html", r.Server)
}
