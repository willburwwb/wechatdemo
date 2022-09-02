package comment

import (
	"log"
	databasecomment "wechatdemo/database/comment"
	databaseuser "wechatdemo/database/user"
	"wechatdemo/model"
	"wechatdemo/response"

	"github.com/gin-gonic/gin"
)

func CreateComment(c *gin.Context) {
	var comment model.Comment
	if err := c.ShouldBind(&comment); err != nil {
		log.Println(comment, "err:", err)
		log.Println(comment.Postid, " ", comment.Content, " ", comment.Responseid)
		response.Failed(c, 400, "创建评论参数有误", err)
		return
	}
	user := c.GetUint("user")
	userName, err := databaseuser.GetUserNameByID(user)
	log.Println("userName为", userName)
	if err != nil {
		response.Failed(c, 400, "获取名字失败", "")
		return
	}
	//comment.UserName = userName
	comment.UserId = user
	databasecomment.Create(c, &comment)
}
