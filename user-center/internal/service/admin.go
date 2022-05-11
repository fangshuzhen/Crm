package service

import (
	"hello/core/entity"
	"hello/user-center/internal/service/impl"

	"github.com/gin-gonic/gin"
)

type Admin interface {
	Login(reqData *entity.DtoAdminLogin) (interface{}, error)
	Logout(c *gin.Context) (interface{}, error)

	Info(reqData *entity.DtoAdminInfo) (interface{}, error)
	Check(tokenStr string) (*entity.UserInfo, error)
	List(reqData *entity.DtoAdminUserList) (interface{}, error)

	Create(reqData *entity.UserInfo) (interface{}, error)
	Update(reqData *entity.UserInfo) (interface{}, error)
	Delete(reqData *entity.UserInfo) (interface{}, error)
}

func GetAdminService() Admin {
	return Admin(&impl.Admin{})
}
