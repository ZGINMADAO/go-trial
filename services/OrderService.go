package services

import "github.com/go-xorm/xorm"

type OrderService interface {
	OrderList(DB *xorm.Engine, requestData map[string]string) (int64)
}

type orderService struct {
}

func NewOrder() OrderService {
	return &orderService{}
}

func (my *orderService) OrderList(DB *xorm.Engine, requestData map[string]string) (int64) {
	return 122
}
