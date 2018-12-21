package controllers

import (
	"github.com/kataras/iris"
	"fmt"
	//"log"
	"go-trial/datamodels"
	"go-trial/datasource"
)

type TestController struct {
	Ctx iris.Context
}

func (my *TestController) Get() {
	my.Ctx.View("user/login.html")
	return

	var s1 = struct {
		Id   int
		Name string
	}{
		Id:   2,
		Name: "tom",
	}
	my.Ctx.JSON(s1)
	//var err error
	//
	//db := config.Db
	//driveSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
	//	db.User, db.Pass, db.Host, db.Port, db.Name)
	//engine, err := xorm.NewEngine("mysql", driveSource)

	engine := datasource.Instance()

	var tree []datamodels.Tree

	//tree := make([]datamodels.Tree, 0)
	engine.Find(&tree)

	fmt.Println(tree)

	//var tree datamodels.Tree

	//engine.Sync2(new(datamodels.Tree))
	//engine.Id(1).Delete(&tree)
	//fmt.Println(tree)
	//return "fuck you"

}

func (my *TestController) GetOne() {

}
