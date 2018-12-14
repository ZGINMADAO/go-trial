package services

import (
	"github.com/go-xorm/xorm"
	"go-trial/units"
	"go-trial/datamodels"
	"strconv"
)

type ProductService interface {
	ProductWithType(*xorm.Engine, map[string]string) (int64, []result)
}
type productService struct {
}

func NewProduct() ProductService {
	return &productService{}
}

type result struct {
	datamodels.Product
	CreatedAt string
	TypeName  string
}

func (my *productService) ProductWithType(DB *xorm.Engine, requestData map[string]string) (int64, []result) {

	var product []datamodels.Product
	build := DB.Where("")

	if requestData["Title"] != "" {
		build = DB.Where("title like ?", "%"+requestData["Title"]+"%")
	}

	countBuild := *build
	size, _ := strconv.Atoi(requestData["Size"])
	page, _ := strconv.Atoi(requestData["Page"])
	build.Limit(size, (page-1)*size).Desc("Id").Find(&product)
	total, _ := (&countBuild).Count(new(datamodels.Product))

	typeIdList := make([]int, len(product))
	for key, val := range product {
		typeIdList[key] = val.TypeId
	}

	var productType []datamodels.ProductType

	DB.In("id", typeIdList).Find(&productType)

	mapProductType := make(map[int]datamodels.ProductType)
	for _, val := range productType {
		mapProductType[val.Id] = val
	}

	results := make([]result, len(product))
	for inx, val := range product {
		results[inx].CreatedAt = val.CreatedAt.Format(units.CnFormat)
		results[inx].Id = val.Id
		results[inx].Title = val.Title
		results[inx].Image = val.Image
		results[inx].TypeName = mapProductType[val.TypeId].Title
	}

	return total, results
}
