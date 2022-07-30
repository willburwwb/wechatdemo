package post

import (
	"log"
	"wechatdemo/database"
	"wechatdemo/model"
	"wechatdemo/response"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Follow(c *gin.Context) {
	db := database.Get()
	user := c.GetUint("user")
	var follow model.Follow
	if err := c.ShouldBind(&follow); err != nil {
		response.Failed(c, 400, "收藏参数错误", err)
		return
	}
	var follownew model.Follow
	var code = 0
	db.Where("userid =?  AND postid = ?", user, follow.Postid).Find(&follownew)
	if follownew.ID == 0 {
		follownew.Postid = follow.Postid
		follownew.Userid = user
		db.Table("follow").Create(&follownew)
		code = 1
		log.Println("收藏")
	} else {
		db.Where("userid =?  AND postid = ?", user, follow.Postid).Delete(&model.Follow{})
		code = -1
		log.Println("取消收藏")
	}
	err2 := db.Model(&model.Post{}).Where("id=?", follow.Postid).Update("follow", gorm.Expr("follow + ?", code)).Error
	if err2 != nil {
		response.Failed(c, 400, "收藏失败", "")
		return
	}
	response.Success(c, 200, "收藏/取消收藏成功", "")
}
func GetIsFollow(user uint, postid uint) bool {
	db := database.Get()
	var follow model.Follow
	db.Where("userid = ? AND postid = ?", user, postid).Find(&follow)
	return follow.ID != 0
}
