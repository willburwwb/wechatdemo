package post

import (
	"wechatdemo/database"
	"wechatdemo/model"
	"wechatdemo/response"

	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	db := database.Get()
	userId := c.GetUint("user")
	//获取参数
	var post model.Post
	if err := c.ShouldBind(&post); err != nil {
		response.Failed(c, 400, "参数错误", "")
		return
	}
	// avatar := c.PostForm("avatar")
	// title := c.PostForm("title")
	// qq := c.BindJSON("qq")
	// wx := c.PostForm("wx")
	// content := c.PostForm("content")
	// price := c.PostForm("price")
	// location := c.PostForm("location")
	// tag := c.PostForm("tag")
	if post.QQ == "" && post.Wx == "" {
		response.Failed(c, 400, "必须给定至少一个联系方式", nil)
		return
	}
	if post.Content == "" || post.Title == "" {
		response.Failed(c, 400, "content或title未给出", nil)
		return
	}
	post.UserId = userId
	db.Table("post").Create(&post)
	response.Success(c, 200, "创建帖子成功", post)
}
