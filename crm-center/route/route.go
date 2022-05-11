package route

import (
	"hello/crm-center/controller"
	_ "hello/docs"

	"github.com/gin-gonic/gin"
)

func LoadRoute(router *gin.Engine) {

	//客户相关的接口
	customer := router.Group("/crm/customer/")
	{
		customer.POST("/list", controller.CustomerList)

		customer.POST("/create", controller.CustomerCreate)
		customer.POST("/update", controller.CustomerUpdate)
		customer.POST("/delete", controller.CustomerDelete)
	}

	// 合同相关接口
	contract := router.Group("/crm/contract/")
	{
		contract.POST("/list", controller.ContractList)
		contract.POST("/upload", controller.UploadContractFile)

		contract.POST("/create", controller.ContractCreate)
		contract.POST("/update", controller.ContractUpdate)
		contract.POST("/delete", controller.ContractDelete)

		contract.POST("/addFile", controller.ContractFileAdd)
		contract.POST("/deleteFile", controller.ContractFileDelete)
	}

}
