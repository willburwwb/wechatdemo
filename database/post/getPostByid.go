package post

import (
	"log"
	"wechatdemo/database"
	"wechatdemo/model"
)

func GetPostByID(postid uint) *model.Post {
	var post model.Post
	db := database.Get()
	if err := db.Model(&post).Where("id = ?", postid).Find(&post).Error; err != nil || post.ID == 0 {
		log.Println("获取post错误", err)
		return nil
	}
	return &post
}
