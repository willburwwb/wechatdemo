package handler

import (
	"log"
	"net/smtp"

	"github.com/jordan-wright/email"
)

func SendEmail(userEmail string, code string) {
	e := email.NewEmail()
	e.From = "wwb <709782717@qq.com>"
	e.To = []string{userEmail}
	e.Subject = "hust二手物品交易平台验证码"
	e.Text = []byte(code)
	err := e.Send("smtp.qq.com:25", smtp.PlainAuth("", "709782717@qq.com", "jkqlctbwxgvebebc", "smtp.qq.com"))
	if err != nil {
		log.Println("email send to ", userEmail, "failed && randcode =", code, err)
		return
	}
}
