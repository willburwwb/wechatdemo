package post

import (
	"log"
	"wechatdemo/database"
	"wechatdemo/model"
	"wechatdemo/response"

	"github.com/gin-gonic/gin"
)

func GetIsFollow(user uint, postid uint) bool {
	db := database.Get()
	var follow model.Follow
	db.Where("userid = ? AND postid = ?", user, postid).Find(&follow)
	return follow.ID != 0
}
func GetIsThumb(user uint, postid uint) bool {
	db := database.Get()
	var thumb model.Thumb
	db.Where("userid = ? AND postid = ?", user, postid).Find(&thumb)
	return thumb.ID != 0
}
func GetIsReply(userid uint, postid uint) bool {
	db := database.Get()
	var comment model.Comment
	var user model.User
	db.Where("id = ?", userid).Find(&user)
	db.Where("postid = ? AND user_name = ?", postid, user.Name).Find(&comment)
	return comment.ID != 0
}
func ReturnPostList(c *gin.Context, posts []model.Post) {
	len := len(posts)
	userid := c.GetUint("user")
	log.Println("当前正在查询的人:", userid)
	responsePosts := make([]model.ResponsePost, len, 50)
	for i := 0; i < len; i++ {
		if userid != 0 {
			responsePosts[i].IsThumb = GetIsThumb(userid, posts[i].ID)
			responsePosts[i].IsFollow = GetIsFollow(userid, posts[i].ID)
			responsePosts[i].IsReplied = GetIsReply(userid, posts[i].ID)
		}
		responsePosts[i].UserName = posts[i].UserName
		responsePosts[i].ID = posts[i].ID
		responsePosts[i].Avatar = posts[i].Avatar
		responsePosts[i].Title = posts[i].Title
		responsePosts[i].QQ = posts[i].QQ
		responsePosts[i].Wx = posts[i].Wx
		responsePosts[i].Content = posts[i].Content
		responsePosts[i].Price = posts[i].Price
		responsePosts[i].Location = posts[i].Location
		responsePosts[i].Thumb = posts[i].Thumb
		responsePosts[i].Reply = posts[i].Reply
		responsePosts[i].Follow = posts[i].Follow
		responsePosts[i].CreatedAt = posts[i].CreatedAt
		responsePosts[i].Tag = posts[i].Tag
	}
	response.Success(c, 200, "成功返回列表", responsePosts)
}
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
	//response.Success(c, 200, "根据用户搜寻帖子成功", posts)
	log.Println("根据用户搜寻帖子成功")
	ReturnPostList(c, posts)
}
