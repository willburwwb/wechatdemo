package comment

import (
	"log"
	"wechatdemo/database"
	"wechatdemo/model"
	"wechatdemo/response"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Create(c *gin.Context, comment *model.Comment) {
	db := database.Get()
	var post model.Post
	if err := db.Model(&post).Where("id = ?", comment.Postid).Find(&post).Error; err != nil {
		log.Println("未成功找到post", err)
		response.Failed(c, 400, "未成功找到post", err)
		return
	}
	if comment.Responseid != 0 {
		if err := db.Where("id = ?", comment.Responseid).Find(&model.Comment{}).RowsAffected; err == 0 {
			response.Failed(c, 400, "未成功找到原评论", err)
			return
		}
	}
	err := db.Create(comment).Error
	if err != nil {
		log.Println("创建失败")
		response.Failed(c, 400, "创建一级评论失败", "")
	}
	log.Println("创建评论成功")
	response.Success(c, 200, "创建成功", comment)
	err = db.Model(&model.Post{}).Where("id = ?", comment.Postid).UpdateColumn("reply", gorm.Expr("reply + ?", 1)).Error
	if err != nil {
		log.Println("更新评论数失败")
		return
	}
	log.Println("更新评论数成功")
}
