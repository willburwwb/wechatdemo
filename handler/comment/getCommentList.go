package comment

import (
	"wechatdemo/database"
	"wechatdemo/model"
	"wechatdemo/response"

	"github.com/gin-gonic/gin"
)

func GetRePostList(commentid uint) []model.ResponseComment { //获取二级评论
	db := database.Get()
	var ResponseComments []model.ResponseComment
	db.Where("responseid = ?", commentid).Find(&ResponseComments)
	return ResponseComments
}
func GetCommentByPost(c *gin.Context) {
	var replyComment model.ReplyComment
	if err := c.ShouldBind(&replyComment); err != nil {
		response.Failed(c, 400, "获取回复评论的参数有误", "")
		return
	}
	db := database.Get()
	var replyComments []model.ReplyComment //所有的一级评论,每一条评论包含其二级评论的数组
	var comments []model.Comment           //获取所有的一级评论
	db.Where("postid = ?", replyComment.Commentid).Order("id desc").Find(&comments)
	for i := range comments {
		replyComment.UserName = comments[i].UserName
		replyComment.Content = comments[i].Content
		replyComment.ResponseComments = GetRePostList(comments[i].ID) //获取二级评论
		replyComments = append(replyComments, replyComment)
	}
	response.Success(c, 200, "成功返回评论列表!", replyComments)
}

func GetCommetListByPost(c *gin.Context) {
}
