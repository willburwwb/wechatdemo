package post

import (
	"log"
	"strconv"
	"wechatdemo/database"
	databasepost "wechatdemo/database/post"
	"wechatdemo/model"
	"wechatdemo/response"

	"github.com/gin-gonic/gin"
)

func GetPostByid(c *gin.Context) {
	postid := c.Query("postid")
	if postid, ok := strconv.Atoi(postid); ok != nil || postid == 0 {
		log.Println("参数传递错误", postid)
		response.Failed(c, 400, "参数传递错误", err)
		return
	}
	var post model.Post
	db := database.Get()
	if err := db.Model(&post).Where("id = ?", postid).Find(&post).Error; err != nil || post.ID == 0 {
		log.Println("获取post错误", err)
		response.Failed(c, 400, "获取post错误", err)
		return
	}
	databasepost.ReturnPost(c, &post, c.GetUint("user"))
}
