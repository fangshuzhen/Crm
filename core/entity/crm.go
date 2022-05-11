package entity

import (
	"time"

	"github.com/globalsign/mgo/bson"
)

type Customer struct {
	Id            bson.ObjectId `json:"id" bson:"_id"`
	Title         string        `json:"title" bson:"title"`
	Name          string        `json:"name" bson:"name"`
	Url           string        `json:"url" bson:"url"`
	Communication string        `json:"communication" bson:"communication"`
	Type          []string      `json:"type" bson:"type"`
}

// 用于返回集合名称
func (this_ *Customer) CollectionName() string {
	return "customer"
}

// 客户列表请求数据模板
type DtoCustomerList struct {
	DtoApiModel
	DtoPageList
}

// --------------------------------------

// 合同
type Contract struct {
	Id         bson.ObjectId `json:"id" bson:"_id"`
	Title      string        `json:"title" bson:"title"`
	State      string        `json:"state" bson:"state"`
	StartDate  time.Time     `json:"startDate" bson:"startDate"`
	EndDate    time.Time     `json:"endDate" bson:"endDate"`
	Desc       string        `json:"desc" bson:"desc"`
	CustomerId bson.ObjectId `json:"customerId" bson:"customerId"`
	Files      []string      `json:"files" bson:"files"`
	Name       string
}

// state： 0未签约，1签约中，2已失效，3已终止，4已过期

// 用于返回集合名称
func (this_ *Contract) CollectionName() string {
	return "contract"
}

// 合同列表请求数据模板
type DtoContractList struct {
	DtoApiModel
	DtoPageList
}

//合同文件名
type ContractAddFile struct {
	Id    bson.ObjectId `json:"id" bson:"_id"`
	Files string        `json:"files" bson:"files"`
}

func (this_ *ContractAddFile) CollectionName() string {
	return "contract"
}
