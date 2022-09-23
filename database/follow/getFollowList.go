package follow

import (
	"log"
	"wechatdemo/database"
	"wechatdemo/model"
	"wechatdemo/response"

	"github.com/gin-gonic/gin"
)

func GetFollowList(c *gin.Context, userId uint, requestFollow *model.RequestFollow) *[]model.Follow {
	db := database.Get()
	log.Println(requestFollow.Offset, requestFollow.Limit)
	var follows []model.Follow
	err := db.Offset(requestFollow.Offset).Limit(requestFollow.Limit).Where("userid = ?", userId).Order("id desc").Find(&follows).Error
	if err != nil {
		response.Failed(c, 400, "查询followlist错误", nil)
		log.Println("查询followlist错误", err)
		return nil
	}
	log.Println("获取followlist成功")
	return &follows
}
func GetFollowsSumByPost(postid uint) int {
	var follows []model.Follow
	db := database.Get()
	db.Where("postid = ?", postid).Find(&follows)
	return len(follows)
}
func GetThumbsSumByPost(postid uint) int {
	var thumbs []model.Thumb
	db := database.Get()
	db.Where("postid = ?", postid).Find(&thumbs)
	return len(thumbs)
}
