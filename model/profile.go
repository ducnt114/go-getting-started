package model

import "gorm.io/gorm"

type Profile struct {
	gorm.Model
	UserName string `gorm:"column:user_name"`
	Bio      string `gorm:"column:bio"`
}

func (Profile) TableName() string {
	return "profile"
}
