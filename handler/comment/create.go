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
	if err != nil {
		response.Failed(c, 400, "获取名字失败", "")
		return
	}
	comment.UserName = userName
	databasecomment.Create(c, &comment)
	log.Println("创建成功")
	response.Success(c, 200, "创建成功", "")
}
func CreateReComment(c *gin.Context) {
	var recomment model.ResponseComment
	if err := c.ShouldBind(&recomment); err != nil {
		response.Failed(c, 400, "创建二级评论参数有误", "")
		return
	}
	user := c.GetUint("user")
	userName, err := databasepost.GetPostUsername(user)
	if err != nil {
		response.Failed(c, 400, "获取名字失败", "")
		return
	}
	recomment.UserName = userName
	databasecomment.ReCreate(c, &recomment)
	log.Println("创建成功")
	response.Success(c, 200, "创建成功", "")
}
