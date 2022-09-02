package jpg

import (
	"wechatdemo/database"
	"wechatdemo/model"
	"wechatdemo/response"

	"github.com/gin-gonic/gin"
)

func GetJpg(c *gin.Context) { 
	name := c.Param("filename")
	DB := database.Get()
	var f model.Jpg
	DB.Where("jpgname = ?", name).First(&f)
	if f.ID == 0 {
		response.Failed(c, 400, "不存在该图片!", nil)
	} else {
		response.Success(c, 200, "获取成功!", f)
	}
}