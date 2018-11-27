package main

import "github.com/kataras/iris"

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	app.Run(iris.Addr("0.0.0.0:8080"))
}
