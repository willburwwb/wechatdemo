package comment

import (
	"log"
	"wechatdemo/database"
	"wechatdemo/model"
	"wechatdemo/response"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Delete(c *gin.Context, userid uint, commentid uint) {
	var db = database.Get()
	var user model.User
	var comment model.Comment
	db.Where("id = ?", userid).First(&user)
	log.Println("username为", user.Name)
	db.Where("id = ?", commentid).First(&comment)
	if userid != comment.UserId {
		log.Println("该用户不具备删除权限或评论不存在")
		response.Failed(c, 400, "该用户不具备删除权限或评论不存在", "")
		return
	}
	err := db.Delete(&comment).Error
	if err != nil {
		log.Println("删除出现问题")
		response.Failed(c, 400, "删除出现问题", "")
		return
	}
	log.Println("删除成功")
	response.Success(c, 200, "成功删除", comment)
	err = db.Model(&model.Post{}).Where("id=?", comment.Postid).Update("reply", gorm.Expr("reply - ?", 1)).Error
	if err != nil {
		log.Println("更新评论数失败")
	}
	log.Println("删除评论成功")
}
