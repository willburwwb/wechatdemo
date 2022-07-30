package post

import (
	databasePost "wechatdemo/database/post"
	"wechatdemo/model"
	"wechatdemo/response"

	"github.com/gin-gonic/gin"
)

var list model.ListType
var err error

func GetPostListByTitle(c *gin.Context) {
	list, err = InitGetPostList(c)
	if list, err = InitGetPostList(c); err != nil {
		return
	}
	posts := databasePost.GetPostList(c, &list, c.Query("title"), "title")
	ReturnPostList(c, posts)
}
func GetPostListByTag(c *gin.Context) {
	list, err = InitGetPostList(c)
	if list, err = InitGetPostList(c); err != nil {
		return
	}
	tag := c.Query("tag")
	if tag == "" {
		response.Failed(c, 400, "tag参数错误", "")
		return
	}
	posts := databasePost.GetPostList(c, &list, tag, "tag")
	ReturnPostList(c, posts)
}
