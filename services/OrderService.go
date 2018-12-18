package services

import (
	"github.com/go-xorm/xorm"
	"go-trial/units"
	"go-trial/datamodels"
	"strconv"
	"fmt"
)

type OrderService interface {
	OrderList(DB *xorm.Engine, requestData map[string]string) (int64)
}

type orderService struct {
}

func NewOrder() OrderService {
	return &orderService{}
}

func (my *orderService) OrderList(DB *xorm.Engine, requestData map[string]string) (int64) {
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

	var results []map[string]string

	for inx, val := range orders {

		fmt.Println(inx,val,results)
		//results[inx].CreatedAt = val.CreatedAt.Format(units.CnFormat)
		//results[inx].Id = val.Id
		//results[inx].Title = val.Title
		//results[inx].Image = val.Image
		//results[inx].TypeName = mapProductType[val.TypeId].Title
	}
	return  total
}
