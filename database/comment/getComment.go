package comment

import (
	"log"
	"wechatdemo/database"
	"wechatdemo/model"
)

func GetCommentByPost(postid uint, responseid uint) []model.Comment {
	var comments []model.Comment
	var db = database.Get()
	log.Println("评论查询:", postid, " responseid:", responseid)
	err := db.Where("postid = ? AND responseid = ?", postid, responseid).Find(&comments).Error
	if err != nil {
		log.Println("失败", err)
		return nil
	}
	log.Printf("查询postid=%d responseid=%d成功\n", postid, responseid)
	return comments
}
func GetCommentByUser(userid uint) {

}
