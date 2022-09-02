package comment

import (
	"log"
	"wechatdemo/database"
	databaseuser "wechatdemo/database/user"
	"wechatdemo/model"
	"wechatdemo/response"

	"github.com/gin-gonic/gin"
)

func GetUserByCommentID(c *gin.Context) {
	commentid := c.Query("commentid")
	db := database.Get()
	var comment model.Comment
	if err := db.Model(&model.Comment{}).Where("id = ?", commentid).Find(&comment).Error; err != nil {
		log.Println("查询不到", commentid, "的user")
		response.Failed(c, 400, "查询不到user", err)
		return
	}
	name, _ := databaseuser.GetUserNameByID(comment.UserId)
	response.Success(c, 200, "成功", map[string]interface{}{"name": name})
}
