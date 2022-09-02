package post

import (
	"log"
	"wechatdemo/database"
	"wechatdemo/model"
)

func GetUserNameByID(userid uint) (string, error) {
	var user model.User
	db := database.Get()
	err := db.Where("id=?", userid).Find(&user).Error
	if err != nil {
		log.Println("由id查找本名出错")
		return "", err
	}
	log.Println("userid ", userid, " 对应 ", user.Name)
	return user.Name, nil
}
