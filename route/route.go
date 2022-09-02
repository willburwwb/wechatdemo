package route

import (
	comments "wechatdemo/handler/comment"
	posts "wechatdemo/handler/post"
	users "wechatdemo/handler/user"
	"wechatdemo/middle"

	"github.com/gin-gonic/gin"
)

func InitRoute() *gin.Engine {
	route := gin.Default()
	user := route.Group("/user")
	{
		user.POST("/login", users.ReorLo)
		user.POST("/getcode", users.GetVerifyCode)
		user.GET("/getUserDetail", middle.AuthJWT(), users.GetUserDetail)
	}
	post := route.Group("/post")
	{
		post.GET("/getPostsList", posts.GetPostList)
		post.Use(middle.AuthJWT())
		post.POST("/create", posts.Create)
		post.PUT("/update", posts.Update)
		post.DELETE("/delete", posts.Delete)
		post.POST("/thumb", posts.Thumb)
		post.POST("/follow", posts.Follow)
		post.GET("/selectByTitle", posts.GetPostListByTitle)
		//比postlist多一个title字段，get form方式
		post.GET("/selectByTag", posts.GetPostListByTag)
		post.GET("/getPostListByUser", posts.GetPostListByUser)
		//比postlist多一个tag字段，get form方式

	}
	comment := route.Group("/comment", middle.AuthJWT())
	{

		//comment.Use(middle.AuthJWT())
		comment.POST("/create", comments.CreateComment)
		comment.GET("/getCommentListByPost", comments.GetCommentListByPost)
		comment.GET("/getCommentListByUser", comments.GetCommentListByUser)
		comment.GET("/getReCommentListByUser")
		comment.DELETE("/delete", comments.Delete)
	}
	route.GET("/comment/getUserByCommentid", comments.GetUserByCommentID)
	return route
}
