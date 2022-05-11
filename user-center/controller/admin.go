package controller

import (
	"hello/core/entity"
	"hello/core/response"
	"hello/user-center/internal/service"

	"github.com/gin-gonic/gin"
)

// @Tags Admin相关接口
// @Summary 登录
// @Accept json
// @Produce json
// @Param model body entity.DtoAdminLogin true "data"
// @Success 200 {object} response.R
// @Router /user/admin/login [post]
func Login(c *gin.Context) {

	var reqData entity.DtoAdminLogin
	var err error
	var result interface{}

	if err = c.ShouldBindJSON(&reqData); err == nil {
		result, err = service.GetAdminService().Login(&reqData)
	}

	response.Json(c, result, err)
}

// @Tags Admin相关接口
// @Summary 登出
// @Accept json
// @Produce json
// @Success 200 {object} response.R
// @Router /user/admin/logout [post]
func Logout(c *gin.Context) {

	result, err := service.GetAdminService().Logout(c)

	response.Json(c, result, err)
}

// @Tags Admin相关接口
// @Summary 获取用户信息
// @Accept json
// @Produce json
// @Success 200 {object} response.R
// @Router /user/admin/info [post]
func Info(c *gin.Context) {

	var reqData entity.DtoAdminInfo
	var err error
	var result interface{}

	if err = c.ShouldBindJSON(&reqData); err == nil {
		result, err = service.GetAdminService().Info(&reqData)
	}

	response.Json(c, result, err)
}

// @Tags Admin相关的接口
// @Summary 测试userlist
// @Accept json
// @Produce json
// @Param model body entity.DtoAdminUserList true "data"
// @Success 200 {object} response.R
// @Router /user/admin/list [post]
func UserList(c *gin.Context) {

	var reqData entity.DtoAdminUserList
	var err error
	var result interface{}

	if err = c.ShouldBindJSON(&reqData); err == nil {
		result, err = service.GetAdminService().List(&reqData)
	}

	response.Json(c, result, err)
}

// @Tags Admin相关的接口
// @Summary 添加新用户
// @Description 添加一个新用户
// @Accept json
// @Produce json
// @Param model body entity.UserInfo true "data"
// @Success 200 {object} response.R
// @Router /user/admin/create [post]
func UserCreate(c *gin.Context) {

	var reqData entity.UserInfo
	var err error
	var result interface{}

	if err = c.ShouldBindJSON(&reqData); err == nil {
		result, err = service.GetAdminService().Create(&reqData)
	}

	response.Json(c, result, err)
}

// @Tags Admin相关的接口
// @Summary 修改用户信息
// @Description 修改用户信息
// @Accept json
// @Produce json
// @Param model body entity.UserInfo true "data"
// @Success 200 {object} response.R
// @Router /user/admin/update [post]
func UserUpdate(c *gin.Context) {

	var reqData entity.UserInfo
	var err error
	var result interface{}

	if err = c.ShouldBindJSON(&reqData); err == nil {
		result, err = service.GetAdminService().Update(&reqData)
	}

	response.Json(c, result, err)
}

// @Tags Admin相关的接口
// @Summary 删除用户信息
// @Description 删除用户信息
// @Accept json
// @Produce json
// @Param model body entity.UserInfo true "data"
// @Success 200 {object} response.R
// @Router /user/admin/delete [post]
func UserDelete(c *gin.Context) {

	var reqData entity.UserInfo
	var err error
	var result interface{}

	if err = c.ShouldBindJSON(&reqData); err == nil {
		result, err = service.GetAdminService().Delete(&reqData)
	}

	response.Json(c, result, err)
}
