package admin

import (
	"github.com/go-xorm/xorm"
	"go-trial/datamodels"
	"strconv"
	"go-trial/units"
	"fmt"
)

type OrderService interface {
	OrderList(DB *xorm.Engine, requestData map[string]string) (int64, []map[string]interface{})
}

type orderService struct {
}

func NewOrder() OrderService {
	return &orderService{}
}

func (my *orderService) OrderList(DB *xorm.Engine, requestData map[string]string) (int64, []map[string]interface{}) {
	var orders []datamodels.Order
	build := DB.Where("")

	if requestData["KeyWord"] != "" {
		build = DB.Where("title like ?", "%"+requestData["KeyWord"]+"%")
	}

	countBuild := *build
	size, _ := strconv.Atoi(requestData["Size"])
	page, _ := strconv.Atoi(requestData["Page"])

	build.Limit(size, (page-1)*size).Desc("Id").Find(&orders)
	total, _ := (&countBuild).Count(new(datamodels.Order))

	results := make([]map[string]interface{}, len(orders))
	for key := range results {
		results[key] = make(map[string]interface{})
	}
	fmt.Println(orders)

	for inx, val := range orders {
		results[inx]["Id"] = val.Id
		results[inx]["TradeNo"] = val.TradeNo
		results[inx]["ProductTitle"] = val.ProductTitle
		results[inx]["ProductTypeTitle"] = val.ProductTypeTitle
		results[inx]["OriginAmount"] = val.OriginAmount
		results[inx]["RealAmount"] = val.RealAmount
		results[inx]["Stock"] = val.Stock
		results[inx]["StatusTitle"] = val.StatusTitle()
		results[inx]["CreatedAt"] = val.CreatedAt.Format(units.CnFormat)
	}
	return total, results
}
