package report

import (
	"log"
	"net/smtp"
	"strconv"
	databaseuser "wechatdemo/database/user"
	"wechatdemo/model"
	"wechatdemo/response"

	"github.com/gin-gonic/gin"
	"github.com/jordan-wright/email"
)

func Report(c *gin.Context) {
	var report model.Report
	if err := c.ShouldBind(&report); err != nil {
		log.Println("举报参数错误", err)
		response.Failed(c, 400, "举报参数错误", err)
		return
	}
	userid := c.GetUint("user")
	username, err := databaseuser.GetUserNameByID(userid)
	if err != nil {
		log.Println("查询名字出错", err)
		return
	}
	report.Content = "举报人: " + username + "\n举报帖子: " + strconv.Itoa(int(report.Postid)) + "\n举报内容:" + report.Content
	report.Username = username
	SendReportEmail(&report)
}
func SendReportEmail(report *model.Report) {
	e := email.NewEmail()

	e.From = "用户" + report.Username + " <husttryanderror@163.com>"
	log.Println(e.From)
	e.To = []string{"husttryanderror@163.com"}
	e.Subject = "举报信息"
	e.Text = []byte(report.Content)
	err := e.Send("smtp.163.com:25", smtp.PlainAuth("", "husttryanderror@163.com", "VNCJGLQPQCKGDIHM", "smtp.163.com"))
	if err != nil {
		log.Println("举报信息未成功送出", err)
		return
	}
	log.Println("举报信息成功送出")
}
