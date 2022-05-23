package controller

import (
	"fmt"
	"hello/core/bind"
	"hello/core/entity"
	"hello/core/response"
	"hello/log-center/internal/service"

	"github.com/gin-gonic/gin"
)

// @Tags Logger相关的接口
// @Summary 测试Logger日志列表
// @Accept json
// @Produce json
// @Param model body entity.DtoLoggerList true "data"
// @Success 200 {object} response.R
// @Router /log/logger/list [post]
func LoggerList(c *gin.Context) {

	var reqData entity.DtoLoggerList
	var err error
	var result interface{}
	var user *entity.UserInfo

	if user, err = bind.ShouldBindJSON(c, &reqData); err == nil {

		fmt.Printf("用户[%s]来获取客户列表信息", user.Username)
		result, err = service.GetLoggerService().List(&reqData)
	}

	response.Json(c, result, err)
}
