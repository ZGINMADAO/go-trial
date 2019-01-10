package admin

import (
	"github.com/kataras/iris/mvc"
	adminService "go-trial/services/admin"
	"github.com/kataras/iris"
	"go-trial/web/controllers"
)

type OrderController struct {
	controllers.BaseController
}

func (my *OrderController) Get() mvc.View {
	return mvc.View{
		Name: "admin/order/lists.html",
	}
}

func (my *OrderController) GetList() {

	service := adminService.NewOrder()
	page := my.Ctx.URLParam("page")
	size := my.Ctx.URLParam("size")
	keyWord := my.Ctx.URLParam("keyWord")
	var requestData = make(map[string]string)
	requestData["KeyWord"] = keyWord
	requestData["Size"] = size
	requestData["Page"] = page
	total, results := service.OrderList(my.DB, requestData)
	my.Ctx.JSON(iris.Map{"rows": results, "total": total})
}
