package post

import (
	"log"
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
func GetPostListByUSer(c *gin.Context, params *map[string]interface{}) {
	db := database.Get()
	var posts []model.Post
	if (*params)["limit"] != nil {
		//log.Println(reflect.TypeOf((*params)["limit"]))
		if value, ok := (*params)["limit"].(int); ok {
			log.Println("value limit=", int(value))
			db = db.Limit(int(value))
		}
	}
	if (*params)["offset"] != nil {
		if value, ok := (*params)["offset"].(int); ok {
			db = db.Offset(int(value))
		}
	}
	if (*params)["user_name"] != nil {
		if value, ok := (*params)["user_name"].(string); ok {
			db = db.Where("user_name = ?", value)
		}
	}
	err := db.Find(&posts).Error
	if err != nil {
		log.Println("根据用户搜寻帖子出现错误", err)
		response.Failed(c, 400, "根据用户搜寻帖子出现错误", nil)
		return
	}
	response.Success(c, 200, "根据用户搜寻帖子成功", posts)
	log.Println("根据用户搜寻帖子成功")
}
