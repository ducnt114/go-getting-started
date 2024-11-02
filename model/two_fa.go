package model

import "gorm.io/gorm"

type TwoFactor struct {
	gorm.Model
	UserID uint   `gorm:"column:user_id"`
	Secret string `gorm:"column:secret"`
}

func (TwoFactor) TableName() string {
	return "two_fa"
}
