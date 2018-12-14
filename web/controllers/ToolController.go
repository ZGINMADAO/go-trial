package controllers

import (
	"os"
	"io"
	"time"
	"fmt"
	"github.com/kataras/iris"
)

type ToolController struct {
	BaseController
}

func (my *ToolController) PostUpload() {

	file, info, err := my.Ctx.FormFile("file")
	if err != nil {
		my.Ctx.JSON(apiResource(true, nil, err.Error()))
		return
	}
	defer file.Close()
	fileName := fmt.Sprintf("%v", time.Now().Unix()) + info.Filename
	out, err := os.OpenFile("./web/public/uploads/"+fileName,
		os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		my.Ctx.JSON(apiResource(false, nil, err.Error()))
		return
	}
	defer out.Close()
	io.Copy(out, file)
	my.Ctx.JSON(apiResource(true, iris.Map{"origin": info.Filename, "real": "/public/uploads/" + fileName}, "上传成功"))
}
