package controller

import (
	"fmt"
	"hello/core/bind"
	"hello/core/entity"
	"hello/core/response"
	"hello/crm-center/internal/service"

	"github.com/gin-gonic/gin"
)

// @Tags CRMcontract相关接口
// @Summary 获取客户列表
// @Accept json
// @Produce json
// @Param model body entity.DtoContractList true "data"
// @Success 200 {object} response.R
// @Router /crm/contract/list [post]
func ContractList(c *gin.Context) {

	var reqData entity.DtoContractList
	var err error
	var result interface{}
	var user *entity.UserInfo

	if user, err = bind.ShouldBindJSON(c, &reqData); err == nil {

		fmt.Printf("用户[%s]来获取客户列表信息", user.Username)

		result, err = service.GetContractServer().List(&reqData)
	}
	response.Json(c, result, err)
}

// @Tags CRMcontract相关接口
// @Summary 上传合同文件
// @Accept json
// @Produce json
// @Success 200 {object} response.R
// @Router /crm/contract/upload [post]
func UploadContractFile(c *gin.Context) {

	var err error
	var result interface{}

	if _, err = bind.ShouldBindJSON(c, nil); err == nil {
		result, err = service.GetContractServer().UploadContractFile(c)
	}
	response.Json(c, result, err)
}

// @Tags CRMcontract相关接口
// @Summary 添加新合同
// @Description 添加一个新合同
// @Accept json
// @Produce json
// @Param model body entity.Contract true "data"
// @Success 200 {object} response.R
// @Router /crm/contract/create [post]
func ContractCreate(c *gin.Context) {

	var reqData entity.Contract
	var err error
	var result interface{}

	if err = c.ShouldBindJSON(&reqData); err == nil {
		result, err = service.GetContractServer().Create(&reqData)
	}

	response.Json(c, result, err)
}

// @Tags CRMcontract相关接口
// @Summary 修改合同信息
// @Description 修改合同信息
// @Accept json
// @Produce json
// @Param model body entity.Contract true "data"
// @Success 200 {object} response.R
// @Router /crm/contract/update [post]
func ContractUpdate(c *gin.Context) {

	var reqData entity.Contract
	var err error
	var result interface{}

	if err = c.ShouldBindJSON(&reqData); err == nil {
		result, err = service.GetContractServer().Update(&reqData)
	}

	response.Json(c, result, err)
}

// @Tags CRMcontract相关接口
// @Summary 删除合同信息
// @Description 删除合同信息
// @Accept json
// @Produce json
// @Param model body entity.Contract true "data"
// @Success 200 {object} response.R
// @Router /crm/contract/delete [post]
func ContractDelete(c *gin.Context) {

	var reqData entity.Contract
	var err error
	var result interface{}

	if err = c.ShouldBindJSON(&reqData); err == nil {
		result, err = service.GetContractServer().Delete(&reqData)
	}

	response.Json(c, result, err)
}

// @Tags CRMcontract相关接口
// @Summary 添加文件
// @Description 添加一个新文件
// @Accept json
// @Produce json
// @Param model body entity.Contract true "data"
// @Success 200 {object} response.R
// @Router /crm/contract/addFile [post]
func ContractFileAdd(c *gin.Context) {

	var reqData entity.Contract
	var err error
	var result interface{}

	if err = c.ShouldBindJSON(&reqData); err == nil {
		result, err = service.GetContractServer().AddFile(&reqData)
	}

	response.Json(c, result, err)
}

// @Tags CRMcontract相关接口
// @Summary 删除文件
// @Description 删除一个文件
// @Accept json
// @Produce json
// @Param model body entity.Contract true "data"
// @Success 200 {object} response.R
// @Router /crm/contract/deleteFile [post]
func ContractFileDelete(c *gin.Context) {

	var reqData entity.Contract
	var err error
	var result interface{}

	if err = c.ShouldBindJSON(&reqData); err == nil {
		result, err = service.GetContractServer().DeleteFile(&reqData)
	}

	response.Json(c, result, err)
}
