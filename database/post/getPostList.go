package post

import (
	"wechatdemo/database"
	"wechatdemo/model"
	"wechatdemo/response"

	"github.com/gin-gonic/gin"
)

func GetPostList(c *gin.Context, list *model.ListType, methodname string, method string) []model.Post {
	db := database.Get()
	var posts []model.Post
	var err error
	if list.Mode == "Time" {
		if method == "" {
			err = db.Limit(int(list.Limit)).Offset(int(list.Offset)).Order("id desc").Find(&(posts)).Error
		} else if method == "title" {
			err = db.Limit(int(list.Limit)).Offset(int(list.Offset)).Order("id desc").Where("title Like ?", "%"+methodname+"%").Find(&(posts)).Error
		} else {

			err = db.Where("tag = ?", methodname).Limit(int(list.Limit)).Offset(int(list.Offset)).Order("id desc").Find(&(posts)).Error
		}
	}
	if list.Mode == "Hot" {
		if method == "" {
			err = db.Limit(int(list.Limit)).Offset(int(list.Offset)).Order("thumb desc").Find(&(posts)).Error
		} else if method == "title" {
			err = db.Limit(int(list.Limit)).Offset(int(list.Offset)).Order("thumb desc").Where("title like ?", "%"+methodname+"%").Find(&(posts)).Error
		} else {
			err = db.Where("tag = ?", methodname).Limit(int(list.Limit)).Offset(int(list.Offset)).Order("thumb desc").Find(&(posts)).Error
		}
	}
	if err != nil {
		response.Failed(c, 400, "未成功查询", err)
		return nil
	}
	return posts
}
