package controller

import (
	"hello/core/entity"
	"hello/core/response"
	"hello/user-center/internal/service"

	"github.com/gin-gonic/gin"
)

// @Tags 测试相关的接口
// @Summary 测试Info
// @Description 接口说明1
// @Description 接口说明2
// @Description 接口说明3
// @Description 接口说明4
// @Description 接口说明5
// @Description 接口说明6
// @Accept json
// @Produce json
// @Param model body entity.DtoHello true "data"
// @Success 200 {object} response.R
// @Router /user/hello/userList [post]
// func UserList(c *gin.Context) {

// 	var reqData entity.DtoHello
// 	var err error
// 	var result interface{}

// 	if err = c.ShouldBindJSON(&reqData); err == nil {
// 		result, err = service.GetHelloService().UserList(&reqData)
// 	}

// 	response.Json(c, result, err)
// }

// @Tags 测试相关的接口
// @Summary 测试Insert
// @Accept json
// @Produce json
// @Param model body entity.DtoUser true "data"
// @Success 200 {object} response.R
// @Router /user/hello/insert [post]
// 新增用户记录
func InsertUser(c *gin.Context) {

	var reqData entity.DtoUser
	var err error
	var result interface{}

	if err = c.ShouldBindJSON(&reqData); err == nil {
		result, err = service.GetHelloService().InsertUser(&reqData)
	}

	response.Json(c, result, err)
}

// @Tags 测试相关的接口
// @Summary 测试Delete
// @Accept json
// @Produce json
// @Param model body entity.DtoFindUser true "data"
// @Success 200 {object} response.R
// @Router /user/hello/delete [post]
// 删除用户记录
func DeleteUser(c *gin.Context) {

	var reqData entity.DtoFindUser
	var err error
	var result interface{}

	if err = c.ShouldBindJSON(&reqData); err == nil {
		result, err = service.GetHelloService().DeleteUser(&reqData)
	}

	response.Json(c, result, err)
}

// @Tags 测试相关的接口
// @Summary 测试Find
// @Accept json
// @Produce json
// @Param model body entity.DtoFindUser true "data"
// @Success 200 {object} response.R
// @Router /user/hello/find [post]
// 查询用户记录
func FindUser(c *gin.Context) {

	var reqData entity.DtoFindUser
	var err error
	var result interface{}

	if err = c.ShouldBindJSON(&reqData); err == nil {
		result, err = service.GetHelloService().FindUser(&reqData)
	}

	response.Json(c, result, err)
}

// @Tags 测试相关的接口
// @Summary 测试Edit
// @Accept json
// @Produce json
// @Param model body entity.DtoUser true "data"
// @Success 200 {object} response.R
// @Router /user/hello/edit [post]
// 新增用户记录
func EditUser(c *gin.Context) {

	var reqData entity.DtoUser
	var err error
	var result interface{}

	if err = c.ShouldBindJSON(&reqData); err == nil {
		result, err = service.GetHelloService().EditUser(&reqData)
	}

	response.Json(c, result, err)
}
