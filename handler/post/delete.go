package post

import (
	"log"
	"reflect"
	"wechatdemo/database"
	"wechatdemo/model"
	"wechatdemo/response"

	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) {
	db := database.Get()
	userId := c.GetUint("user")
	var user model.User
	log.Println("当前正在删除的人", userId)
	db.Where("id = ?", userId).First(&user)
	//获取参数
	if userId == 0 {
		response.Failed(c, 400, "当前用户不存在token", err)
		return
	}
	json := make(map[string]interface{})
	var post model.Post
	if err := c.ShouldBindJSON(&json); err != nil {
		response.Failed(c, 400, "参数错误", "")
		return
	}	
	if json["postid"] == nil {
		log.Println("postid为0")
		response.Failed(c, 400, "postid不能为0", nil)
		return
	} else {
		log.Println("postid为", json["postid"], " 类型为", reflect.TypeOf(json["postid"]))
	}
	db.Where("id = ?", json["postid"]).First(&post)
	log.Println("post's username :", post.UserName, " 你的名字", user.Name)
	if post.UserName != user.Name {
		response.Failed(c, 400, "权限不足", "")
		return
	}
	err := db.Delete(&post).Error
	if err != nil {
		log.Println("删除失败", err)
	}
	response.Success(c, 200, "删除成功!", post)
}
