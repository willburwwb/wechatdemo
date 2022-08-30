package comment

import (
	"log"
	databasecomment "wechatdemo/database/comment"
	"wechatdemo/model"
	"wechatdemo/response"

	"github.com/gin-gonic/gin"
)

func GetCommentListByPost(c *gin.Context) {
	var requestComment model.RequestCommentByPost
	if err := c.ShouldBind(&requestComment); err != nil {
		response.Failed(c, 400, "获取评论参数有误", "")
		return
	}
	var comments []model.Comment
	var recomments []model.Comment
	log.Println("获取评论数据")
	comments = databasecomment.GetCommentByPost(requestComment.Postid, 0)
	if comments == nil {
		return
	}
	var replyComments []model.ReplyComments
	for i, comment := range comments {
		log.Println("获取第一个评论", i)
		var replyComment model.ReplyComments
		replyComment.Content = comment.Content
		replyComment.UserName = comment.UserName

		recomments = databasecomment.GetCommentByPost(requestComment.Postid, comment.ID)
		for _, recomment := range recomments {
			var reReplycomment model.ReplyComments
			reReplycomment.Content = recomment.Content
			reReplycomment.UserName = recomment.UserName
			replyComment.ReplyComments = append(replyComment.ReplyComments, reReplycomment)
		}

		replyComments = append(replyComments, replyComment)
	}
	response.Success(c, 200, "成功", replyComments)
}
