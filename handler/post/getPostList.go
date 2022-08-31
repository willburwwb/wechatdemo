package post

import (
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"
	"wechatdemo/database"
	databasePost "wechatdemo/database/post"
	"wechatdemo/model"
	"wechatdemo/response"
	"wechatdemo/verify"

	"github.com/gin-gonic/gin"
)

func JudgeNow(c *gin.Context) uint {
	tokenString := c.GetHeader("Authorization")
	//验证格式
	if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") { //token为空或者不是以"Bearer "开头
		return 0
	}

	tokenString = tokenString[7:] //丢弃开头部分

	token, claims, err := verify.ParseToken(tokenString)
	if err != nil || !token.Valid { //返回出错或者token无效
		return 0
	}
	userId := claims.UserId
	DB := database.Get()
	var user model.User
	DB.First(&user, userId)
	if user.ID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "用户不存在!"})
		return 0
	}
	return userId
}
func checkPost(c *gin.Context, list *model.ListType) error {
	if list.Mode != "Time" && list.Mode != "Hot" {
		response.Failed(c, 400, "参数Mode有误", "")
		return errors.New("格式错误")
	}
	if list.Limit == 0 {
		list.Limit = 10
	}
	if list.Limit > 50 {
		response.Failed(c, 400, "查询条数太多", "")
		return errors.New("查询条数太多")
	}
	return nil
}
func InitGetPostList(c *gin.Context) (model.ListType, error) {
	var list model.ListType
	if err := c.ShouldBind(&list); err != nil {
		log.Print(err)
		response.Failed(c, 400, "参数错误", err)
		return model.ListType{}, err
	}
	if err := checkPost(c, &list); err != nil {
		return model.ListType{}, err
	}
	return list, nil
}
func ReturnPostList(c *gin.Context, posts []model.Post) {
	len := len(posts)
	userid := JudgeNow(c)
	log.Println("当前正在查询的人:", userid)
	responsePosts := make([]model.ResponsePost, len, 50)
	for i := 0; i < len; i++ {
		if userid != 0 {
			responsePosts[i].IsThumb = databasePost.GetIsThumb(userid, posts[i].ID)
			responsePosts[i].IsFollow = databasePost.GetIsFollow(userid, posts[i].ID)
			responsePosts[i].IsReplied = databasePost.GetIsReply(userid, posts[i].ID)
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
func GetPostList(c *gin.Context) {
	var list model.ListType
	var err error
	if list, err = InitGetPostList(c); err != nil {
		return
	}
	posts := databasePost.GetPostList(c, &list, "", "")
	ReturnPostList(c, posts)
}
func GetPostListByUser(c *gin.Context) {
	offset, ok := strconv.Atoi(c.Query("offset"))
	if ok != nil {
		log.Println("参数错误")
		response.Failed(c, 400, "参数错误", "")
		return
	}
	limit, ok := strconv.Atoi(c.Query("limit"))
	if ok != nil || limit == 0 {
		log.Println("参数错误")
		response.Failed(c, 400, "参数错误", "")
		return
	}
	userid := c.GetUint("user")
	userName, _ := databasePost.GetPostUsername(userid)

	log.Printf("offset = %d limit = %d userName = %s\n", offset, limit, userName)
	if err != nil {
		return
	}
	if userid == 0 {
		log.Println("用户id不存在")
		response.Failed(c, 400, "用户id不存在", "")
		return
	}
	canshu := make(map[string]interface{})
	canshu["offset"] = offset
	canshu["limit"] = limit
	canshu["user_name"] = userName
	databasePost.GetPostListByUSer(c, &canshu)
}
