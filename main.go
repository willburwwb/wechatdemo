package main

import (
	"log"
	"wechatdemo/database"
	"wechatdemo/model"
	"wechatdemo/route"
)

func main() {
	if err := database.InitLocalDB(); err != nil {
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
	db.AutoMigrate(&model.ResponseComment{})
	engine := route.InitRoute()
	if err := engine.Run(":3000"); err != nil {
		log.Fatal("service failed", err)
	}
}
