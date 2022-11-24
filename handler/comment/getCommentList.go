package comment

import (
	"log"
	databasecomment "wechatdemo/database/comment"
	databaseuser "wechatdemo/database/user"
	"wechatdemo/model"
	"wechatdemo/response"

	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func GetCommentListByPost(c *gin.Context) {
	var requestComment model.RequestCommentByPost
	if err := c.ShouldBind(&requestComment); err != nil {
		response.Failed(c, 400, "获取评论参数有误", "")
		return
	}
	var comments []model.Comment
	var recomments []model.Comment
	log.Println("根据帖子获取评论数据")
	comments = databasecomment.GetCommentByPost(requestComment.Postid, 0)
	if comments == nil {
		return
	}
	var replyComments []model.ReplyComments
	for i, comment := range comments {
		log.Println("获取第", i+1, "个评论")
		var replyComment model.ReplyComments
		replyComment.ID = comment.ID
		replyComment.Content = comment.Content
		userName, err := databaseuser.GetUserNameByID(comment.UserId)
		if err == nil {
			replyComment.UserName = userName
		} else {
			continue
		}
		replyComment.Fileid, replyComment.QQ, replyComment.Wx = databaseuser.GetUserDetailById(comment.UserId)
		recomments = databasecomment.GetCommentByPost(requestComment.Postid, comment.ID)
		for _, recomment := range recomments {
			var reReplycomment model.ReplyComments
			reReplycomment.ID = recomment.ID
			log.Println("回复的评论", reReplycomment.ID)
			reReplycomment.Content = recomment.Content
			userName, err := databaseuser.GetUserNameByID(recomment.UserId)
			if err == nil {
				reReplycomment.UserName = userName
			}
			reReplycomment.Fileid, reReplycomment.QQ, reReplycomment.Wx = databaseuser.GetUserDetailById(recomment.UserId)
			replyComment.ReplyComments = append(replyComment.ReplyComments, reReplycomment)
		}

		replyComments = append(replyComments, replyComment)
	}
	response.Success(c, 200, "成功", replyComments)
}
func GetCommentListByUser(c *gin.Context) {
	userid := c.GetUint("user")
	offset := com.StrTo(c.Query("offset")).MustInt()
	limit := com.StrTo(c.Query("limit")).MustInt()
	databasecomment.GetCommentByUser(c, userid, limit, offset)
}
