package comment

import (
	"log"
	"wechatdemo/database"
	"wechatdemo/model"
	"wechatdemo/response"

	"github.com/gin-gonic/gin"
)

func GetCommentByPost(postid uint, responseid uint) []model.Comment {
	var comments []model.Comment
	var db = database.Get()
	log.Println("查询postid=", postid, " responseid=", responseid)
	err := db.Where("postid = ? AND responseid = ?", postid, responseid).Order("id desc").Find(&comments).Error
	if err != nil {
		log.Println("失败", err)
		return nil
	}
	log.Printf("查询postid=%d responseid=%d成功\n", postid, responseid)
	return comments
}
func GetCommentByUser(c *gin.Context, userid uint, limit int, offset int) {
	var comments []model.Comment
	var db = database.Get()
	log.Println("评论查询 userid ", userid)
	if userid == 0 {
		log.Println("未查询到该用户/token不存在")
		response.Failed(c, 400, "未查询到该用户", "")
		return
	}
	err := db.Offset(offset).Limit(limit).Where("user_id = ?", userid).Order("id desc").Find(&comments).Error
	if err != nil {
		log.Println("查询数据失败")
		response.Failed(c, 400, "查询数据失败", err)
		return
	}
	log.Println("查询成功")
	response.Success(c, 200, "查询数据成功", comments)
}
func GetCommentsSumByPost(postid uint) int {
	var comments []model.Comment
	db := database.Get()
	db.Where("postid = ?", postid).Find(&comments)
	return len(comments)
}

func GetCommentByMyself(c *gin.Context, userid uint, limit int, offset int) {
	postids := GetCommentIdByUser(userid)
	commentids := GetCommentIdByUser(userid)
	var comments []model.Comment
	DB := database.Get()
	err := DB.Offset(offset).Limit(limit).Where("postid IN ?", postids).Or("responseid IN ?", commentids).Order("id desc").Find(&comments).Error
	if err != nil {
		response.Failed(c, 400, "查询回复失败!", nil)
		return
	}
	response.Success(c, 200, "查询回复成功!", comments)
}