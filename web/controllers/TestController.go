package controllers

import (
	"github.com/kataras/iris"
	"fmt"
	"log"
	"go-trial/datamodels"
	"go-trial/datasource"
)

type TestController struct {
	Ctx iris.Context
}

func (my *TestController) Get() string {

	//var err error
	//
	//db := config.Db
	//driveSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
	//	db.User, db.Pass, db.Host, db.Port, db.Name)
	//engine, err := xorm.NewEngine("mysql", driveSource)

	engine := datasource.Instance()

	var admin datamodels.Admin
	admin.Id = 2
	_, errGet := engine.Get(&admin)
	if errGet != nil {
		log.Fatal(errGet)
	}
	fmt.Println("fuck you")
	fmt.Println(admin)

	//var tree datamodels.Tree

	//engine.Sync2(new(datamodels.Tree))
	//engine.Id(1).Delete(&tree)
	//fmt.Println(tree)
	return "fuck you"

}
