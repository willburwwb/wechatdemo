package handler

import (
	"bytes"
	"html/template"
	"log"
	"net/smtp"

	"github.com/jordan-wright/email"
)

// func SendEmail(userEmail string, code string) {
// 	e := email.NewEmail()
// 	e.From = "wwb <709782717@qq.com>"
// 	e.To = []string{userEmail}
// 	e.Subject = "hust二手物品交易平台验证码"
// 	e.Text = []byte(code)
// 	err := e.Send("smtp.qq.com:25", smtp.PlainAuth("", "709782717@qq.com", "jkqlctbwxgvebebc", "smtp.qq.com"))
// 	if err != nil {
// 		log.Println("email send to ", userEmail, "failed && randcode =", code, err)
// 		return
// 	}

// }
func SendEmail(userEmail string, code string) {
	e := email.NewEmail()
	e.From = "不期而喻小程序团队 <husttryanderror@163.com>"
	e.To = []string{userEmail}
	e.Subject = "邮箱验证码"
	//e.Text = []byte(code)
	// file, err := os.Open("D:\\project\\wechatdemo-masternew\\static\\1.tmpl")
	// if err != nil {
	// 	log.Println("读取文件失败", err)
	// 	return
	// }
	tmpl, err := template.ParseFiles("static/1.tmpl")
	if err != nil {
		log.Println(err)
		return
	}
	buf := bytes.NewBuffer([]byte{})
	err = tmpl.Execute(buf, code)
	if err != nil {
		log.Println("读取html失败")
		return
	}
	e.HTML = buf.Bytes()
	err = e.Send("smtp.163.com:25", smtp.PlainAuth("", "husttryanderror@163.com", "VNCJGLQPQCKGDIHM", "smtp.163.com"))
	if err != nil {
		log.Println("email send to ", userEmail, "failed && randcode =", code, err)
		return
	}
}
