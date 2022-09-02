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
	log.Println("评论查询:", postid, " responseid:", responseid)
	err := db.Where("postid = ? AND responseid = ?", postid, responseid).Find(&comments).Error
	if err != nil {
		log.Println("失败", err)
		return nil
	}
	log.Printf("查询postid=%d responseid=%d成功\n", postid, responseid)
	return comments
}
func GetCommentByUser(c *gin.Context, userid uint) {
	var comments []model.Comment
	var db = database.Get()
	log.Println("评论查询 userid ", userid)

	if userid == 0 {
		log.Println("未查询到该用户/token不存在")
		response.Failed(c, 400, "未查询到该用户", "")
		return
	}
	err := db.Where("user_id = ?", userid).Find(&comments).Error
	if err != nil {
		log.Println("查询数据失败")
		response.Failed(c, 400, "查询数据失败", err)
		return
	}
	log.Println("查询成功")
	response.Success(c, 200, "查询数据成功", comments)
}
