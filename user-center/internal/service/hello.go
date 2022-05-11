package service

import (
	"hello/core/entity"
	"hello/user-center/internal/service/impl"
)

type Hello interface {
	UserList(reqData *entity.DtoHello) (interface{}, error)
	InsertUser(reqData *entity.DtoUser) (interface{}, error)
	DeleteUser(reqData *entity.DtoFindUser) (interface{}, error)
	FindUser(reqData *entity.DtoFindUser) (interface{}, error)
	EditUser(reqData *entity.DtoUser) (interface{}, error)
}

func GetHelloService() Hello {
	return Hello(&impl.Hello{})
}
