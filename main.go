package main

import (
	"log"
	"wechatdemo/database"
	"wechatdemo/model"
	"wechatdemo/route"
	"wechatdemo/utils"
)

func main() {
	if err := database.InitDB(); err != nil {
		log.Println("数据库连接失败")
		return
	}
	db := database.Get()
	db.AutoMigrate(&model.Post{})
	db.AutoMigrate(&model.VerifyCode{})
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Thumb{})
	db.AutoMigrate(&model.Follow{})
	db.AutoMigrate(&model.Comment{})
	db.AutoMigrate(&model.Jpg{})
	db.AutoMigrate(&model.File{})
	engine := route.InitRoute()
	utils.Clear()
	if err := engine.Run(":3000"); err != nil {
		log.Fatal("service failed", err)
	}
}
