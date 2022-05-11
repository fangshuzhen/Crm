package response

import "github.com/gin-gonic/gin"

type R struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Json(c *gin.Context, resData interface{}, err error) {
	if err != nil {
		Error(c, err, resData)
	} else {
		Success(c, resData)
	}
}

func Success(c *gin.Context, resData interface{}) {
	c.JSON(200, R{Code: 20000, Message: "suceess", Data: resData})
}

func Error(c *gin.Context, err error, resData interface{}) {
	c.JSON(200, R{Code: 50001, Message: err.Error(), Data: resData})
}

func ErrorE(c *gin.Context, code int, message string, resData interface{}) {
	c.JSON(200, R{Code: code, Message: message, Data: resData})
}
