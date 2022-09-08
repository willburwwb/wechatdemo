package handler

import (
	"log"
	"wechatdemo/database"
	"wechatdemo/model"
	"wechatdemo/response"

	"github.com/gin-gonic/gin"
)

func GetUserDetail(c *gin.Context) {
	userid := c.GetUint("user")
	var user model.User
	db := database.Get()
	if err := db.Where("id = ?", userid).First(&user).Error; err != nil {
		log.Println("获取用户信息失败")
		response.Failed(c, 400, "获取用户信息失败", "")
		return
	}
	response.Success(c, 400, "获取用户信息成功", user)
}
func GetUserPostsSum(c *gin.Context) {
	userid := c.GetUint("user")
	db := database.Get()
	var count int64
	if err := db.Model(&model.Post{}).Where("user_id = ?", userid).Count(&count).Error; err != nil {
		log.Println("获取帖子数目失败", err)
		response.Failed(c, 400, "获取帖子数目失败", err)
		return
	}
	log.Println("发现共有", count, "分帖子")
	c.JSON(200, gin.H{
		"code": 200,
		"data": gin.H{
			"count": count,
		},
		"message": "成功返回帖子数目",
	})
}
func GetUserCommentsSum(c *gin.Context) {
	userid := c.GetUint("user")
	db := database.Get()
	var count int64
	if err := db.Model(&model.Comment{}).Where("user_id = ?", userid).Count(&count).Error; err != nil {
		log.Println("获取评论数目失败", err)
		response.Failed(c, 400, "获取评论数目失败", err)
		return
	}
	log.Println("发现共有", count, "分评论")
	c.JSON(200, gin.H{
		"code": 200,
		"data": gin.H{
			"count": count,
		},
		"message": "成功返回评论数目",
	})
}
func GetUserFollowsSum(c *gin.Context) {
	userid := c.GetUint("user")
	db := database.Get()
	var count int64
	if err := db.Model(&model.Follow{}).Where("userid = ?", userid).Distinct("postid").Count(&count).Error; err != nil {
		log.Println("获取收藏数目失败", err)
		response.Failed(c, 400, "获取收藏数目失败", err)
		return
	}
	log.Println("发现共有", count, "分收藏")
	c.JSON(200, gin.H{
		"code": 200,
		"data": gin.H{
			"count": count,
		},
		"message": "成功返回收藏数目",
	})
}
