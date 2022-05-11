package route

import (
	"hello/user-center/controller"

	_ "hello/docs"

	"github.com/gin-gonic/gin"
)

func LoadRoute(router *gin.Engine) {

	//登录相关的接口
	admin := router.Group("/user/admin/")
	{
		admin.POST("/login", controller.Login)
		admin.POST("/logout", controller.Logout)
		admin.POST("/info", controller.Info)

		admin.POST("/list", controller.UserList)
		admin.POST("/create", controller.UserCreate)
		admin.POST("/update", controller.UserUpdate)
		admin.POST("/delete", controller.UserDelete)
	}

}
