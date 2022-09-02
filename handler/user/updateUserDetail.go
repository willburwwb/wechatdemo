package handler

import (
	"log"
	"wechatdemo/database"
	"wechatdemo/model"
	"wechatdemo/response"

	"github.com/gin-gonic/gin"
)

func UpdateUserDetail(c *gin.Context) {
	var json map[string]interface{}
	if err := c.ShouldBindJSON(&json); err != nil {
		log.Println("参数绑定出现问题")
		response.Failed(c, 400, "参数绑定出现问题", err)
		return
	}
	if name, ok := json["name"].(string); !ok || name == "" {
		log.Println("没有传用户名||格式不对")
		response.Failed(c, 400, "没有传用户名||格式不对", "")
		return
	}
	userid := c.GetUint("user")
	db := database.Get()
	if userid != 0 {
		err := db.Model(&model.User{}).Where("id = ?", userid).Updates(json).Error
		if err != nil {
			log.Println("更改失败", err)
			response.Failed(c, 400, "更改失败", err)
			return
		}
		response.Success(c, 200, "更改成功", json)
	}
}
