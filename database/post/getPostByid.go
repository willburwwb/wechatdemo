package post

import (
	"encoding/json"
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

func AnalyzeJson(fileid string) []string {
	var ids []string
	err := json.Unmarshal([]byte(fileid), &ids)
	if err != nil {
		log.Println("解析json出错!")
	}
	return ids
}