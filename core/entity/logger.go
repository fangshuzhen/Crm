package entity

import (
	"time"

	"github.com/globalsign/mgo/bson"
)

type Logger struct {
	Id          bson.ObjectId `json:"id" bson:"_id"`
	Time        time.Time     `json:"time" bson:"time"`
	Client_ip   string        `json:"client_ip" bson:"client_ip"`
	Req_uri     string        `json:"req_uri" bson:"req_uri"`
	Status_code int           `json:"status_code" bson:"status_code"`
}

// 用于返回集合名称
func (this_ *Logger) CollectionName() string {
	return "logger"
}

// 日志列表请求数据模板
type DtoLoggerList struct {
	DtoApiModel
	DtoPageList
}
