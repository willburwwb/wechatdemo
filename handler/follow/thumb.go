package post

import (
	"log"
	"wechatdemo/database"
	"wechatdemo/model"
	"wechatdemo/response"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Thumb(c *gin.Context) {
	db := database.Get()
	user := c.GetUint("user")
	var thumb model.Thumb
	if err := c.ShouldBind(&thumb); err != nil {
		response.Failed(c, 400, "点赞参数错误", err)
		return
	}
	var thumbnew model.Thumb
	var code = 0
	db.Where("userid =?  AND postid = ?", user, thumb.Postid).Find(&thumbnew)
	if thumbnew.ID == 0 {
		thumbnew.Postid = thumb.Postid
		thumbnew.Userid = user
		db.Create(&thumbnew)
		code = 1
		log.Println("增加点赞")
	} else {
		db.Where("userid =?  AND postid = ?", user, thumb.Postid).Delete(&model.Thumb{})
		code = -1
		log.Println("删除点赞")
	}
	err2 := db.Model(&model.Post{}).Where("id=?", thumb.Postid).Update("thumb", gorm.Expr("thumb + ?", code)).Error

	log.Println("thumb", thumb.Postid)

	if err2 != nil {
		response.Failed(c, 400, "点赞失败", "")
		return
	}
	response.Success(c, 200, "更新点赞数成功", "")
}
