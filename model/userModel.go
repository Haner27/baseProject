package model

import (
	"gorm.io/gorm"
	"time"
)

type (
	UserModel struct {
		gorm.Model
		Username string     `gorm:"column:username;size:32;unique_index;not:null"`
		Password string     `gorm:"column:password;size:32;not:null"`
		Nickname string     `gorm:"column:username;size:32;unique_index;not:null"`
		Gender   int 		`gorm:"column:gender;not:null"`
		Birthday time.Time  `gorm:"column:birthday"`
	}
)

func (UserModel) TableName() string {
	return "user"
}
