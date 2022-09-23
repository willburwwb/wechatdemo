package comment

import (
	"log"
	"wechatdemo/database"
	"wechatdemo/model"
)

func GetPostIdByUser(userid uint) []uint { //本人的所有帖子id
	var ids []uint //存储postid
	var posts []model.Post
	DB := database.Get()
	err := DB.Where("user_id = ?", userid).Find(&posts).Error
	if err != nil {
		log.Println("查询本人所有帖子id失败")
	}
	for _, post := range posts {
		ids = append(ids, post.ID)
	}
	return ids
}

func GetCommentIdByUser(userid uint) []uint { //本人所有评论id
	var ids []uint //commentid
	var comments []model.Comment
	DB := database.Get()
	err := DB.Where("user_id = ?", userid).Find(&comments).Error
	if err != nil {
		log.Println("查询本人所有评论id失败")
	}
	for _, comment := range comments {
		ids = append(ids, comment.ID)
	}
	return ids
}