package admin

import (
	"go-trial/web/controllers"
	"github.com/kataras/iris/mvc"
	"go-trial/datamodels"
	"github.com/kataras/iris"
)

type MemberController struct {
	controllers.BaseController
}

func (my *MemberController) Get() mvc.View {
	return mvc.View{
		Name: "admin/member/list.html",
	}
}
func (my *MemberController) GetList() {
	var members []datamodels.Member
	err := my.DB.Find(&members)
	if err != nil {
		my.Ctx.Application().Logger().Warn(err.Error())
		my.ReturnJson(false, "", err.Error())
		return
	}
	total, _ := my.DB.Count(new(datamodels.Member))
	my.Ctx.JSON(iris.Map{"rows": members, "total": total})
}