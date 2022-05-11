package service

import (
	"hello/core/entity"
	"hello/crm-center/internal/impl"
)

type Customer interface {
	List(reqData *entity.DtoCustomerList) (interface{}, error)

	Create(reqData *entity.Customer) (interface{}, error)
	Update(reqData *entity.Customer) (interface{}, error)
	Delete(reqData *entity.Customer) (interface{}, error)
}

func GetCustomerServer() Customer {
	return Customer(&impl.Customer{})
}
