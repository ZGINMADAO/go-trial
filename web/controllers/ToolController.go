package controllers

import (
	"os"
	"io"
	"time"
	"fmt"
	"github.com/kataras/iris"
	"net/smtp"
	"strings"
	"github.com/kataras/iris/mvc"
)

type ToolController struct {
	BaseController
}

func (my *ToolController) PostUpload() {

	file, info, err := my.Ctx.FormFile("file")
	if err != nil {
		my.ReturnJson(true, nil, err.Error())
		return
	}
	defer file.Close()
	fileName := fmt.Sprintf("%v", time.Now().Unix()) + info.Filename
	out, err := os.OpenFile("./web/public/uploads/"+fileName,
		os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		my.ReturnJson(false, nil, err.Error())
		return
	}
	defer out.Close()
	io.Copy(out, file)
	my.ReturnJson(true, iris.Map{"origin": info.Filename, "real": "/public/uploads/" + fileName}, "上传成功")
}

func SendEmail(_ *ToolController) {
	auth := smtp.PlainAuth("", "1287020839@qq.com", "ujpltpaenazsjcae", "smtp.qq.com")
	to := []string{"15555124010@163.com"}
	nickname := "ggc"
	user := "1287020839@qq.com"
	subject := "test mail"
	contentType := "Content-Type: text/plain; charset=UTF-8"
	body := "This is the email body."
	msg := []byte("To: " + strings.Join(to, ",") + "\r\nFrom: " + nickname +
		"<" + user + ">\r\nSubject: " + subject + "\r\n" + contentType + "\r\n\r\n" + body)
	err := smtp.SendMail("smtp.qq.com:25", auth, user, to, msg)
	if err != nil {
		fmt.Printf("send mail error: %v", err)
	}
}

func (my *ToolController) Get() mvc.View {
	return mvc.View{
		Name:   "admin/websockets.html",
		Layout: "",
	}
}
