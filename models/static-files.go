package models

import "gorm.io/gorm"

type StaticFiles struct {
	gorm.Model
	Title string	
	UserRefer int 
	Owner User `gorm:"foreignKey:UserRefer"`
}