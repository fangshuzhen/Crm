package service

import (
	"hello/core/entity"
	"hello/log-center/internal/impl"
)

type Logger interface {
	List(reqData *entity.DtoLoggerList) (interface{}, error)
}

func GetLoggerService() Logger {

	return Logger(&impl.Logger{})
}
