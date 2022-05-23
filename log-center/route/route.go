package route

import (
	_ "hello/docs"
	"hello/log-center/controller"

	"github.com/gin-gonic/gin"
)

func LoadRoute(router *gin.Engine) {

	//日志相关的接口
	logger := router.Group("/log/logger/")
	{
		logger.POST("/list", controller.LoggerList)

	}

}
