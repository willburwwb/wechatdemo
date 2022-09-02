package jpg

import (
	"encoding/base64"
	"io"
	"wechatdemo/database"
	"wechatdemo/model"
	"wechatdemo/response"

	"github.com/gin-gonic/gin"
)

//获取上传的图片,并保存到数据库
func DownloadJpg(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		response.Failed(c, 400, "打开图片失败!", nil)
		return
	}
	name := file.Filename //名称
	icon, _ := file.Open()
	content, err := io.ReadAll(icon) 
	if err != nil {
		response.Failed(c, 400, "获取图片编码失败!", nil)
		return
	}
	codedata := base64.StdEncoding.EncodeToString(content) //编码结果
	
	f := model.Jpg {
		Jpgname: name,
		Codedata: codedata,
	}
	
	DB := database.Get()
	DB.Create(&f)
}