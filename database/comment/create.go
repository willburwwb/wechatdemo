package comment

import (
	"log"
	"wechatdemo/database"
	"wechatdemo/model"
	"wechatdemo/response"

	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context, comment *model.Comment) {
	db := database.Get()
	err := db.Create(comment).Error
	if err != nil {
		log.Println("创建失败")
		response.Failed(c, 400, "创建一级评论失败", "")
		return
	}
}
func ReCreate(c *gin.Context, comment *model.ResponseComment) {
	db := database.Get()
	err := db.Create(comment).Error
	if err != nil {
		log.Println("创建失败")
		response.Failed(c, 400, "创建一级评论失败", "")
		return
	}
}
