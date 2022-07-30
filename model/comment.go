package model

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Postid   uint   `json:"postid" postform:"postid" form:"postid" uri:"postid" binding:"required"`
	Content  string `json:"content" postform:"content" form:"content" uri:"content" binding:"required"`
	UserName string
}
type ResponseComment struct {
	gorm.Model
	Postid     uint   `json:"postid" postform:"postid" form:"postid" uri:"postid" binding:"required"`
	Responseid uint   `json:"responseid" postform:"responseid" form:"responseid" uri:"responseid"`
	Content    string `json:"content" postform:"content" form:"content" uri:"content" binding:"required"`
	UserName   string
}

type ReplyComment struct {
	UserName         string
	Commentid          uint `json:"postid" postform:"postid" form:"postid" uri:"postid" binding:"required"`
	Content          string
	ResponseComments []ResponseComment
}
