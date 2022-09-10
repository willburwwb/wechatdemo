package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email string `gorm:"type:varchar(50);not null"` //邮箱登录
	Name  string `gorm:"type:varchar(50);not null"` //给用户赋予的名字
	QQ    string `gorm:"qq"`
	Wx    string
}

type VerifyCode struct {
	gorm.Model
	Email string `gorm:"type:varchar(50);not null" form:"email" postform:"email" json:"email" binding:"required"` //发送的邮箱
	Code  string `gorm:"type:varchar(10);not null" form:"code" postform:"code" json:"code" binding:"required"`    //验证码
}

type VerifyUser struct {
	Email string `postform:"email" json:"email" form:"email" binding:"required"`
}

type RequestFollow struct {
	gorm.Model
	Limit  int `uri:"limit" json:"limit" form:"limit"`
	Offset int `uri:"offset" json:"offset" form:"offset" bind:"required"`
}
