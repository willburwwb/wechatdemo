package post

import (
	"log"
	"wechatdemo/database"
	"wechatdemo/model"
)

func GetPostUsername(userid uint) (string, error) {
	var user model.User
	db := database.Get()
	err := db.Where("id=?", userid).Find(&user).Error
	if err != nil {
		log.Println("由id查找本名出错")
		return "", err
	}
	return user.Name, nil
}
