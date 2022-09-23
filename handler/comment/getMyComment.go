package comment

import (
	databasecomment "wechatdemo/database/comment"

	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func GetCommentListByMyself(c *gin.Context) {
	userid := c.GetUint("user")
	offset := com.StrTo(c.Query("offset")).MustInt()
	limit := com.StrTo(c.Query("limit")).MustInt()
	databasecomment.GetCommentByMyself(c, userid, limit, offset)
}