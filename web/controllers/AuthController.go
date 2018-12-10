package controllers

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"fmt"
	"go-trial/datamodels"
	"go-trial/services"
)

type AuthController struct {
	ApiJson
	BaseController
}

func (_ *AuthController) GetLogin() mvc.View {
	return mvc.View{
		Name: "admin/login.html",
		Data: iris.Map{"url": "/admin/login"},
	}
}

func (my *AuthController) BeforeActivation(b mvc.BeforeActivation) {
	b.Dependencies().Add(services.NewAuth())
}

func (my *AuthController) BeginRequest(ctx iris.Context) {
	fmt.Println("BeginRequest被调用了啊哈啊哈哈")
}

func (my *AuthController) EndRequest(ctx iris.Context) {
	fmt.Println("EndRequest被调用了啊哈啊哈哈")
}

func (my *AuthController) PostLogin() {

	username := my.Ctx.PostValue("username")
	password := my.Ctx.PostValue("password")
	fmt.Printf("类型为%T,值为%v", username, username)
	if len(username) < 5 || len(password) < 5 {
		my.Ctx.JSON(apiResource(false, nil, "用户或密码不能少于5位"))
	} else {
		//apiResource(true, nil, "登录成功")
		var admin datamodels.Admin
		admin.Username = username
		admin.Password = password
		ok, _ := my.DB.Get(&admin)
		if ok == true {
			my.Ctx.JSON(iris.Map{"status": true, "data": nil, "message": "登录成功"})
		} else {
			my.Ctx.JSON(apiResource(false, nil, "登录失败"))
		}
	}
}

type List struct {
	Title    string
	ParentId int
	Id       int
	Child    []List
}

func Recursive(Trees []datamodels.Tree, parentId int, all []*List) {


	for _, val := range Trees {

		if parentId == val.ParentId {
			child := List{Title: val.Title, ParentId: val.ParentId, Id: val.Id, Child: node}
			node = append(all, child)
			Recursive(Trees, val.ParentId, node)

		}
	}
}

func (my *AuthController) Get() mvc.View {

	var result []datamodels.Tree
	my.DB.Find(&result)

	var all []*List

	Recursive(result, 0, all)

	//m := make(map[string]string)
	//m["a"] = ""
	//
	//for _, v := range m {
	//
	//}
	//make()

	return mvc.View{
		Name: "admin/index.html",
		Data: iris.Map{"Tree": result},
	}
}
