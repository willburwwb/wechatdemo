package comment

import (
	"log"
	"wechatdemo/database"
	databasecomment "wechatdemo/database/comment"
	"wechatdemo/model"
	"wechatdemo/response"

	"github.com/gin-gonic/gin"
)

//传递commentid参数
func Delete(c *gin.Context) {
	json := make(map[string]interface{})
	userId := c.GetUint("user")
	c.BindJSON(&json)
	commentid, ok := json["commentid"].(float64)
	log.Println("useid为", userId, " commentid为", commentid)
	if !ok || commentid == 0 {
		log.Println("删除评论参数不全")
		response.Failed(c, 400, "删除评论参数不全", "")
		return
	}
	db := database.Get()
	var user model.User
	var comment model.Comment
	db.Where("id = ?", userId).First(&user)
	log.Println("username为", user.Name)
	db.Where("id = ?", commentid).First(&comment)
	if userId != comment.UserId {
		log.Println("该用户不具备删除权限或评论不存在")
		response.Failed(c, 400, "该用户不具备删除权限或评论不存在", nil)
	}
	msg, err := databasecomment.Delete(&comment)
	if err != nil {
		response.Failed(c, 400, err.Error(), msg)
		return
	}
	response.Success(c, 400, "成功删除", msg)
}
