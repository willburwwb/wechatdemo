package post

import (
	"wechatdemo/database"
	"wechatdemo/model"
	"wechatdemo/response"

	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) {
	db := database.Get()
	userId := c.GetUint("user")
	var user model.User
	db.Where("id = ?", userId).First(&user)
	//获取参数
	json := make(map[string]interface{})
	var post model.Post
	if err := c.ShouldBindJSON(&json); err != nil {
		response.Failed(c, 400, "参数错误", "")
		return
	}
	db.Where("id = ?", json["postid"]).First(&post)
	if post.UserName != user.Name {
		response.Failed(c, 400, "权限不足", "")
		return
	}
	db.Delete(&post)
	response.Success(c, 200, "删除成功!", post)
}
