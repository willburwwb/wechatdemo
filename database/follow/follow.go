package follow

import (
	"log"
	"wechatdemo/database"
	"wechatdemo/model"
	"wechatdemo/response"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InsertFollow(c *gin.Context, follow *model.Follow) {
	db := database.Get()
	var post model.Post
	db.Model(&model.Post{}).Where("id = ?", follow.Postid).Find(&post)
	if post.ID != 0 {
		result := db.Model(&model.Follow{}).Where("userid = ? AND postid = ?", follow.Userid, follow.Postid).Find(&model.Follow{}).RowsAffected
		if result != 0 {
			log.Println("已创建收藏", result)
			return
		}
		if err := db.Model(&model.Follow{}).Create(follow).Error; err != nil {
			log.Println("创建收藏出现错误", err)
			response.Failed(c, 400, "创建收藏出现错误", err)
			return
		}
		log.Println("创建收藏成功")
		if err := db.Model(&model.Post{}).Where("id = ?", follow.Postid).UpdateColumn("follow", gorm.Expr("follow + ?", 1)).Error; err != nil {
			log.Println("增加帖子收藏数失败")
		}
		response.Success(c, 200, "创建收藏成功", *follow)
	}
}
func DeleteFollow(follow *model.Follow) (msg interface{}, err error) {
	db := database.Get()
	log.Println("删除收藏", follow.Userid, " ", follow.Postid)
	if err = db.Model(&model.Follow{}).Where("userid = ? AND postid = ?", follow.Userid, follow.Postid).Delete(&model.Post{}).Error; err != nil {
		log.Println("删除收藏出现问题", err)
		//response.Failed(c, 400, "取消收藏出现错误", err)
		msg = "取消收藏出现错误"
		return
	}
	log.Println("删除收藏成功")
	if err := db.Model(&model.Post{}).Where("id = ?", follow.Postid).UpdateColumn("follow", gorm.Expr("follow - ?", 1)).Error; err != nil {
		log.Println("减少帖子收藏数失败")
	}
	msg = &follow
	return msg, nil
}
