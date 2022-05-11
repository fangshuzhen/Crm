package service

import (
	"hello/core/entity"
	"hello/crm-center/internal/impl"

	"github.com/gin-gonic/gin"
)

type Contract interface {
	UploadContractFile(c *gin.Context) (interface{}, error)
	List(reqData *entity.DtoContractList) (interface{}, error)

	Create(reqData *entity.Contract) (interface{}, error)
	Update(reqData *entity.Contract) (interface{}, error)
	Delete(reqData *entity.Contract) (interface{}, error)

	AddFile(reqData *entity.Contract) (interface{}, error)
	DeleteFile(reqData *entity.Contract) (interface{}, error)
}

func GetContractServer() Contract {
	return Contract(&impl.Contract{})
}
