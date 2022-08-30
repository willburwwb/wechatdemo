package post

import (
	"log"
	"reflect"
	"wechatdemo/database"
	"wechatdemo/model"
	"wechatdemo/response"

	"github.com/gin-gonic/gin"
)

func Update(c *gin.Context) {
	db := database.Get()
	userId := c.GetUint("user")
	var user model.User
	log.Println("当前正在更新的人", userId)
	if userId == 0 {
		response.Failed(c, 400, "当前用户不存在token", err)
		return
	}
	db.Where("id = ?", userId).First(&user)
	json := make(map[string]interface{})
	if err := c.BindJSON(&json); err != nil {
		response.Failed(c, 400, "给定更新参数错误!", err)
		return
	}
	var post model.Post
	if json["postid"] == 0 {
		log.Println("postid为0")
		response.Failed(c, 400, "postid不能为0", nil)
		return
	} else {
		log.Println("postid为", json["postid"], " 类型为", reflect.TypeOf(json["postid"]))
	}
	db.Where("id = ?", json["postid"]).First(&post)
	log.Println("post's username :", post.UserName, " 你的名字", user.Name)
	if post.UserName != user.Name {
		response.Failed(c, 400, "权限不足!", err)
		return
	}
	delete(json, "postid")
	//更新
	err := db.Model(&post).Updates(json).Error
	if err != nil {
		response.Failed(c, 400, "更新失败", err)
	}
	response.Success(c, 200, "更新成功!", post)
}
