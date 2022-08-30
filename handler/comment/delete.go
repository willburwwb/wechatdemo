package comment

import (
	"log"
	databasecomment "wechatdemo/database/comment"
	"wechatdemo/response"

	"github.com/gin-gonic/gin"
)

//传递commentid参数
func Delete(c *gin.Context) {
	json := make(map[string]interface{})
	userId := c.GetUint("user")
	c.BindJSON(&json)
	commentid, ok := json["commentid"].(float64)
	log.Println("useid为", userId, " commentid为", commentid)
	if !ok || commentid == 0 {
		log.Println("删除评论参数不全")
		response.Failed(c, 400, "删除评论参数不全", "")
		return
	}
	databasecomment.Delete(c, userId, uint(commentid))
}
