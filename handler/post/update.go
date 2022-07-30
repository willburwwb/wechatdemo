package post

import (
	"wechatdemo/database"
	"wechatdemo/model"
	"wechatdemo/response"

	"github.com/gin-gonic/gin"
)

func Update(c *gin.Context) {
	db := database.Get()
	userId := c.GetUint("user")
	var user model.User
	db.Where("id = ?", userId).First(&user)
	json := make(map[string]interface{})
	if err := c.BindJSON(&json); err != nil {
		response.Failed(c, 400, "给定更新参数错误!", err)
		return
	}
	var post model.Post
	db.Where("id = ?", json["postid"]).First(&post)
	if post.UserName != user.Name {
		response.Failed(c, 400, "权限不足!", err)
		return
	}
	delete(json, "postid")
	//更新
	db.Model(&post).Updates(json)
	response.Success(c, 200, "更新成功!", post)
}
