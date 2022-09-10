package comment

import (
	"log"
	"wechatdemo/database"
	"wechatdemo/model"

	"gorm.io/gorm"
)

func Delete(comment *model.Comment) (msg interface{}, err error) {
	var db = database.Get()

	err = db.Delete(comment).Error
	if err != nil {
		log.Println("删除出现问题")
		msg = "删除出现问题"
		return
	}
	log.Println("删除成功")
	msg = *comment
	err = db.Model(&model.Post{}).Where("id=?", comment.Postid).Update("reply", gorm.Expr("reply - ?", 1)).Error
	if err != nil {
		log.Println("更新评论数失败")
	}
	log.Println("删除评论成功")
	return
}
