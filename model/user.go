package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email string `gorm:"type:varchar(50);not null"` //邮箱登录
	Name  string `gorm:"type:varchar(50);not null"`
}

type VerifyCode struct {
	gorm.Model
	Email string `gorm:"type:varchar(50);not null" form:"email" postform:"email" json:"email" binding:"required"` //发送的邮箱
	Code  string `gorm:"type:varchar(10);not null" form:"code" postform:"code" json:"code" binding:"required"`    //验证码
}

type VerifyUser struct {
	Email string `postform:"email" json:"email" form:"email" binding:"required"`
}
