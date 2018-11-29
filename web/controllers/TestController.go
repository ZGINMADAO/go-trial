package controllers

import (
	"github.com/kataras/iris"
	"fmt"
	"github.com/go-xorm/xorm"
	"go-trial/config"
	"log"
	"go-trial/datamodels"
)

type TestController struct {
	Ctx iris.Context
}

func (my *TestController) Get() string{

	var err error

	db := config.Db
	driveSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
		db.User, db.Pass, db.Host, db.Port, db.Name)
	engine, err := xorm.NewEngine("mysql", driveSource)

	if err != nil {
		log.Fatal("fail")
	}

	var ret = &datamodels.Admin{Id: 2}
	_, errGet := engine.Get(ret)
	if errGet != nil {
		log.Fatal(errGet)
	}
	fmt.Println(ret)

	fmt.Println(engine.DBMetas())
	return "fuck you"

}
