package model

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Postid  uint   `gorm:"postid" json:"postid" postform:"postid" form:"postid" uri:"postid" binding:"required"`
	Content string `json:"content" postform:"content" form:"content" uri:"content" binding:"required"`
	//UserName   string
	UserId     uint
	Responseid uint `gorm:"responseid" json:"responseid" postform:"responseid" form:"responseid"`
}

type RequestCommentByPost struct {
	Postid uint `json:"postid" postform:"postid" form:"postid" uri:"postid" binding:"required"`
}
type ReplyComments struct {
	UserName      string
	Content       string
	ReplyComments []ReplyComments
}
