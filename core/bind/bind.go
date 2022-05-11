package bind

import (
	"fmt"
	"hello/core/entity"
	"hello/core/service"
	"reflect"

	"github.com/gin-gonic/gin"
)

func ShouldBindJSON(c *gin.Context, reqData interface{}) (*entity.UserInfo, error) {

	var err error

	if reqData != nil {
		if err = c.ShouldBind(&reqData); err != nil {
			return nil, err
		}
	}

	// 取Token
	token := c.Request.Header.Get("X-Token")

	// 如果从头信息中取不到token信息，则通过“反射”从请求参数中取token字段
	if token == "" {
		typeof := reflect.TypeOf(reqData)
		valueof := reflect.ValueOf(reqData)
		if typeof.Kind() == reflect.Ptr {
			valueof = valueof.Elem()
			typeof = typeof.Elem()
		}

		var modelValue reflect.Value
		if _, ok := typeof.FieldByName("DtoApiModel"); ok {
			modelValue = valueof.FieldByName("DtoApiModel")
		} else {
			modelValue = valueof
		}

		token = modelValue.FieldByName("Token").String()
	}

	if token == "" {
		return nil, fmt.Errorf("无法读取到Token信息")
	}

	// 登录检查
	if user, err := service.CheckToken(token); err != nil {
		return nil, err
	} else {
		if user.Name == "" {
			return nil, fmt.Errorf("Token已过期或未登录")
		}
		return user, nil
	}
}
