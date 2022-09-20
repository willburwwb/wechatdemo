package comment

import (
	"log"
	"wechatdemo/database"
	"wechatdemo/model"
)

func Delete(comment *model.Comment) (msg interface{}, err error) {
	var db = database.Get()

	tx := db.Begin() // 开启事务
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	err = tx.Delete(comment).Error
	if err != nil {
		log.Println("删除出现问题")
		msg = "删除出现问题"
		tx.Rollback()
		return
	}
	log.Println("删除评论成功")
	msg = *comment
	if comment.Responseid == 0 {
		err = tx.Model(&model.Comment{}).Where("responseid = ?", comment.ID).Delete(&model.Comment{}).Error
		if err != nil {
			log.Println("删除", comment.ID, "的二级评论失败", err)
			msg = "删除二级评论失败"
			log.Println("回滚删除事务")
			tx.Rollback()
			return
		} else {
			log.Println("删除", comment.ID, "的二级评论成功", err)
		}
	}
	tx.Commit()
	// err = db.Model(&model.Post{}).Where("id=?", comment.Postid).Update("reply", gorm.Expr("reply - ?", 1)).Error
	// if err != nil {
	// 	log.Println("更新评论数失败")
	// }

	return
}
