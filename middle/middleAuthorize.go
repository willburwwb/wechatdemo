package middle

import (
	"net/http"
	"strings"
	"wechatdemo/database"
	"wechatdemo/model"
	"wechatdemo/verify"

	"github.com/gin-gonic/gin"
)

func AuthJWT() gin.HandlerFunc { //中间件
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		//验证格式
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") { //token为空或者不是以"Bearer "开头
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			c.Abort() //抛弃
			return
		}

		tokenString = tokenString[7:] //丢弃开头部分

		token, claims, err := verify.ParseToken(tokenString)
		if err != nil || !token.Valid { //返回出错或者token无效
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			c.Abort() //抛弃
			return
		}

		userId := claims.UserId
		DB := database.Get()
		var user model.User
		DB.First(&user, userId)

		if user.ID == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "用户不存在!"})
			c.Abort() //抛弃
			return
		}

		//否则
		c.Set("user", userId) //写入上下文
		c.Next()
	}
}
