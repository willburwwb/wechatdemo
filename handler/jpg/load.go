package jpg

import (
	"fmt"
	"os"
	"wechatdemo/database"
	"wechatdemo/model"
	"wechatdemo/response"

	"github.com/gin-gonic/gin"
)

//获取上传的图片,保存到本地服务器
func DownloadJpg(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		response.Failed(c, 400, "打开图片失败!", nil)
		return
	}
	pwd, _ := os.Getwd()   //获取当前项目的根目录
	path := fmt.Sprintf("%s/static/%s", pwd, file.Filename)  //要在当前根目录下创建一个static文件夹
	err  = c.SaveUploadedFile(file, path) //保存进去
	if err != nil {
		response.Failed(c, 400, "保存失败!", nil)
	}

	DB := database.Get()
	f := model.Jpg{
		Jpgname: file.Filename,
		Path: path,
	}
	DB.Create(&f)

	response.Success(c, 200, "保存成功!", "ip和端口//域名" + "/static/" + file.Filename) //返回url,前面需要 ip地址 + 端口号 或者 域名
}