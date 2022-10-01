package post

import (
	"log"
	databasefollow "wechatdemo/database/follow"
	databasepost "wechatdemo/database/post"
	"wechatdemo/model"
	"wechatdemo/response"

	"github.com/gin-gonic/gin"
)

func Follow(c *gin.Context) {
	user := c.GetUint("user")
	var follow model.Follow
	if err := c.ShouldBind(&follow); err != nil {
		response.Failed(c, 400, "收藏参数错误", err)
		return
	}
	follow.Userid = user
	databasefollow.InsertFollow(c, &follow)
}
func DeleteFollow(c *gin.Context) {
	user := c.GetUint("user")
	var follow model.Follow
	if err := c.ShouldBind(&follow); err != nil {
		response.Failed(c, 400, "取消收藏参数错误", err)
		return
	}
	follow.Userid = user
	msg, err := databasefollow.DeleteFollow(&follow)
	if err != nil {
		response.Failed(c, 400, err.Error(), msg)
		return
	}
	response.Success(c, 200, "成功删除", msg)
}
func GetFollowList(c *gin.Context) {
	user := c.GetUint("user")
	var requestFollow model.RequestFollow
	if err := c.ShouldBind(&requestFollow); err != nil {
		response.Failed(c, 400, "获取收藏列表参数错误", err)
		return
	}
	var follows *[]model.Follow
	var posts []model.Post
	follows = databasefollow.GetFollowList(c, user, &requestFollow)
	for _, values := range *follows {
		log.Println("此时正在获取", values.Postid)
		post := databasepost.GetPostByID(values.Postid)
		if post != nil {
			posts = append(posts, *post)
		}
	}
	databasepost.ReturnPostList(c, &posts, user)
}
