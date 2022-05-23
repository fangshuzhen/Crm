package impl

import (
	"fmt"
	"hello/core/entity"
	"hello/core/mongo"

	"github.com/globalsign/mgo/bson"
)

type Logger struct{}
type Admin struct{}

//数据展示
func (this_ *Logger) List(reqData *entity.DtoLoggerList) (interface{}, error) {

	db := mongo.GetDb()
	defer db.Session.Close()
	var logger entity.Logger
	loggerColl := db.C(logger.CollectionName())
	var loggerList []entity.Logger
	query := bson.M{}
	// or 查询
	if reqData.Keyword != "" {
		or := []bson.M{
			{"client_ip": bson.M{"$regex": fmt.Sprintf("^.*%s.*", reqData.Keyword)}},
			{"status_code": bson.M{"$regex": fmt.Sprintf("^.*%s.*", reqData.Keyword)}},
		}
		query["$or"] = or
	}
	var err error
	var count int
	// 查询分页的用户数据
	if err := loggerColl.Find(query).Skip((reqData.Page - 1) * reqData.Size).Limit(reqData.Size).All(&loggerList); err != nil {
		return nil, err
	}
	// 查询总共数据条数
	if count, err = loggerColl.Find(query).Count(); err != nil {
		return nil, err
	}
	// 构造分页数据返回客户端
	pager := entity.Pager{
		Page:    reqData.Page,
		Size:    reqData.Size,
		Total:   count,
		Records: loggerList,
	}

	return pager, nil

}
