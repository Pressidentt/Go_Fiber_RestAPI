package models

import (
	"time"
)

type User struct{
	ID		uint		`gorm:"primaryKey"`
	CreatedAt time.Time 
	Name 	string	
	Email string 
	//Relations
	StaticFiles []StaticFiles `gorm:"foreignKey:UserRefer"`
}



