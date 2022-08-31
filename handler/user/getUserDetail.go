package handler

import (
	"log"
	"wechatdemo/database"
	"wechatdemo/model"
	"wechatdemo/response"

	"github.com/gin-gonic/gin"
)

func GetUserDetail(c *gin.Context) {
	userid := c.GetUint("user")
	var user model.User
	db := database.Get()
	if err := db.Where("id = ?", userid).First(&user).Error; err != nil {
		log.Println("获取用户信息失败")
		response.Failed(c, 400, "获取用户信息失败", "")
		return
	}
	response.Success(c, 400, "获取用户信息成功", user)
}
