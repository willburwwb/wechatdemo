package post

import (
	"errors"
	"log"
	"net/http"
	"strings"
	"wechatdemo/database"
	databasePost "wechatdemo/database/post"
	databaseuser "wechatdemo/database/user"
	"wechatdemo/model"
	"wechatdemo/response"
	"wechatdemo/verify"

	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
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

func GetPostList(c *gin.Context) {
	var list model.ListType
	var err error
	if list, err = InitGetPostList(c); err != nil {
		return
	}
	posts := databasePost.GetPostList(c, &list, "", "")
	userid := c.GetUint("user")
	databasePost.ReturnPostList(c, posts, userid)
}
func GetPostListByUser(c *gin.Context) {
	offset := com.StrTo(c.Query("offset")).MustInt64()
	limit := com.StrTo(c.Query("limit")).MustInt()
	userid := c.GetUint("user")
	userName, _ := databaseuser.GetUserNameByID(userid)

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
	canshu["user_id"] = c.GetUint("user")
	databasePost.GetPostListByUSer(c, &canshu)
}
