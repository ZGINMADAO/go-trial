package admin

import (
	"go-trial/web/controllers"
	"go-trial/datamodels"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris"
)

type SupplierController struct {
	controllers.BaseController
}

func (my *SupplierController) Get() mvc.View {
	return mvc.View{
		Name: "admin/supplier/list.html",
	}
}

func (my *SupplierController) GetList() {

	total, _ := my.DB.Count(new(datamodels.Supplier))
	var supplierList []datamodels.Supplier
	err := my.DB.Find(&supplierList)
	if err != nil {
		my.Ctx.Application().Logger().Warn(err.Error())
		return
	}
	my.Ctx.JSON(iris.Map{"rows": supplierList, "total": total})
}
