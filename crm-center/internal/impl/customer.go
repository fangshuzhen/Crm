package impl

import (
	"fmt"
	"hello/core/entity"
	"hello/core/mongo"

	"github.com/globalsign/mgo/bson"
)

type Customer struct{}

func (this_ *Customer) List(reqData *entity.DtoCustomerList) (interface{}, error) {
	db := mongo.GetDb()
	defer db.Session.Close()
	var customer entity.Customer
	customerColl := db.C(customer.CollectionName())
	var customerList []entity.Customer
	query := bson.M{}
	// or 查询
	if reqData.Keyword != "" {
		or := []bson.M{
			{"title": bson.M{"$regex": fmt.Sprintf("^.*%s.*", reqData.Keyword)}},
			{"name": bson.M{"$regex": fmt.Sprintf("^.*%s.*", reqData.Keyword)}},
			{"type": bson.M{"$regex": fmt.Sprintf("^.*%s.*", reqData.Keyword)}},
		}
		query["$or"] = or
	}
	var err error
	var count int
	// 查询分页的用户数据
	if err := customerColl.Find(query).Skip((reqData.Page - 1) * reqData.Size).Limit(reqData.Size).All(&customerList); err != nil {
		return nil, err
	}
	// 查询总共数据条数
	if count, err = customerColl.Find(query).Count(); err != nil {
		return nil, err
	}
	// 构造分页数据返回客户端
	pager := entity.Pager{
		Page:    reqData.Page,
		Size:    reqData.Size,
		Total:   count,
		Records: customerList,
	}
	return pager, nil
}

func (this_ *Customer) Create(reqData *entity.Customer) (interface{}, error) {

	db := mongo.GetDb()
	defer db.Session.Close()

	var customer entity.Customer
	customerColl := db.C(customer.CollectionName())

	// 将客户端提交用户信息保存
	reqData.Id = bson.NewObjectId()
	if err := customerColl.Insert(reqData); err != nil {
		return nil, err
	}

	return nil, nil
}

// 修改数据
func (this_ *Customer) Update(reqData *entity.Customer) (interface{}, error) {

	var err error

	db := mongo.GetDb()
	defer db.Session.Close()

	var customer entity.Customer
	customerColl := db.C(customer.CollectionName())

	// 对用户信息进行修改
	if err = customerColl.Update(bson.M{"_id": reqData.Id}, reqData); err != nil {
		return nil, err
	}

	return nil, err
}

// 删除数据
func (this_ *Customer) Delete(reqData *entity.Customer) (interface{}, error) {

	var err error

	db := mongo.GetDb()
	defer db.Session.Close()

	var customer entity.Customer
	customerColl := db.C(customer.CollectionName())

	if err = customerColl.RemoveId(reqData.Id); err != nil {
		return nil, err
	}

	return nil, err
}
