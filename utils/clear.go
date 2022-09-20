package utils

import (
	"log"
	"wechatdemo/database"
	"wechatdemo/model"
)

func Clear() {
	db := database.Get()
	var posts []model.Post

	db.Find(&posts)
	for _, post := range posts {
		if post.UserId == 0 {
			db.Delete(&post)
			log.Println("删除帖子", post.ID, post.UserId)
		}
	}
	var comments []model.Comment
	db.Model(&model.Comment{}).Find(&comments)
	for _, comment := range comments {
		err := db.Where("id = ?", comment.Postid).Find(&model.Post{}).RowsAffected
		if err == 0 {
			db.Delete(&comment)
			log.Println("删除评论", comment.ID, comment.UserId)
		}
		if comment.Responseid != 0 {
			err := db.Where("id = ?", comment.Responseid).Find(&model.Comment{}).RowsAffected
			if err == 0 {
				db.Delete(&comment)
				log.Println("删除二级评论", comment.ID, comment.UserId, comment.Responseid)
			}
		}
	}

	var follows []model.Follow
	db.Model(&model.Follow{}).Find(&follows)
	for _, follow := range follows {
		err := db.Where("id = ?", follow.Postid).Find(&model.Post{}).RowsAffected
		if err == 0 {
			db.Delete(&follow)
			log.Println("删除收藏", follow.ID, follow.Postid, follow.Userid)
		}

	}

	var thumbs []model.Thumb
	db.Model(&model.Thumb{}).Find(&thumbs)
	for _, thumb := range thumbs {
		err := db.Where("id = ?", thumb.Postid).Find(&model.Post{}).RowsAffected
		if err == 0 {
			db.Delete(&thumb)
			log.Println("删除点赞", thumb.ID, thumb.Postid, thumb.Userid)
		}

	}
}
