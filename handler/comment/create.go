package comment

import (
	"log"
	databasecomment "wechatdemo/database/comment"
	databasepost "wechatdemo/database/post"
	"wechatdemo/model"
	"wechatdemo/response"

	"github.com/gin-gonic/gin"
)

func CreateComment(c *gin.Context) {
	var comment model.Comment
	if err := c.ShouldBind(&comment); err != nil {
		response.Failed(c, 400, "创建评论参数有误", "")
		return
	}
	user := c.GetUint("user")
	userName, err := databasepost.GetPostUsername(user)
	log.Println("userName为", userName)
	if err != nil {
		response.Failed(c, 400, "获取名字失败", "")
		return
	}
	comment.UserName = userName
	databasecomment.Create(c, &comment)
}
