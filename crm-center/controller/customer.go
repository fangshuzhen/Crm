package controller

import (
	"fmt"
	"hello/core/bind"
	"hello/core/entity"
	"hello/core/response"
	"hello/crm-center/internal/service"

	"github.com/gin-gonic/gin"
)

// @Tags CRM-customer相关接口
// @Summary 获取客户列表
// @Accept json
// @Produce json
// @Param model body entity.DtoCustomerList true "data"
// @Success 200 {object} response.R
// @Router /crm/customer/list [post]
func CustomerList(c *gin.Context) {

	var reqData entity.DtoCustomerList
	var err error
	var result interface{}
	var user *entity.UserInfo

	if user, err = bind.ShouldBindJSON(c, &reqData); err == nil {

		fmt.Printf("用户[%s]来获取客户列表信息", user.Username)

		result, err = service.GetCustomerServer().List(&reqData)
	}
	response.Json(c, result, err)
}

// @Tags CRM-customer相关接口
// @Summary 添加新用户
// @Description 添加一个新用户
// @Accept json
// @Produce json
// @Param model body entity.Customer true "data"
// @Success 200 {object} response.R
// @Router /crm/customer/create [post]
func CustomerCreate(c *gin.Context) {

	var reqData entity.Customer
	var err error
	var result interface{}

	if err = c.ShouldBindJSON(&reqData); err == nil {
		result, err = service.GetCustomerServer().Create(&reqData)
	}

	response.Json(c, result, err)
}

// @Tags CRM-customer相关接口
// @Summary 修改用户信息
// @Description 修改用户信息
// @Accept json
// @Produce json
// @Param model body entity.Customer true "data"
// @Success 200 {object} response.R
// @Router /crm/customer/update [post]
func CustomerUpdate(c *gin.Context) {

	var reqData entity.Customer
	var err error
	var result interface{}

	if err = c.ShouldBindJSON(&reqData); err == nil {
		result, err = service.GetCustomerServer().Update(&reqData)
	}

	response.Json(c, result, err)
}

// @Tags CRM-customer相关接口
// @Summary 删除用户信息
// @Description 删除用户信息
// @Accept json
// @Produce json
// @Param model body entity.Customer true "data"
// @Success 200 {object} response.R
// @Router /crm/customer/delete [post]
func CustomerDelete(c *gin.Context) {

	var reqData entity.Customer
	var err error
	var result interface{}

	if err = c.ShouldBindJSON(&reqData); err == nil {
		result, err = service.GetCustomerServer().Delete(&reqData)
	}

	response.Json(c, result, err)
}
