package model

import "gorm.io/gorm"

type Thumb struct {
	gorm.Model
	Userid uint `json:"userid" postform:"userid"`
	Postid uint `json:"postid" postform:"postid" form:"postid" bind:"required"`
}

//收藏
type Follow struct {
	ID     uint `gorm:"primary_key"`
	Userid uint `json:"userid" postform:"userid"`
	Postid uint `json:"postid" postform:"postid" bind:"required"`
}
