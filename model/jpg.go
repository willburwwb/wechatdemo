package model

import "gorm.io/gorm"

type Jpg struct {
	gorm.Model
	Jpgname  string `gorm:"jpgname" json:"jpgname"` //获取图片的途径
	Path string `gorm:"path"`
}