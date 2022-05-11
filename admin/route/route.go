package route

import (
	_ "hello/docs"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func LoadRoute(router *gin.Engine) {

	router.GET("/admin/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

}
