package model

import "gorm.io/gorm"

type Token struct {
	gorm.Model
	UserID       uint   `gorm:"column:user_id"`
	RefreshToken string `gorm:"column:refresh_token"`
}

func (Token) TableName() string {
	return "token"
}
