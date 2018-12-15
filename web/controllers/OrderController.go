package controllers

import (
	"net/smtp"
	"strings"
	"fmt"
)

type OrderController struct {
	BaseController
}

func (my *OrderController) Get() {
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
