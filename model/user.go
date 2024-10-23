package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name    string   `gorm:"column:name"`
	Age     int      `gorm:"column:age"`
	Pass    string   `gorm:"column:pass"`
	Profile *Profile `gorm:"foreignKey:UserName;references:Name"`
	Books   []*Book
}

func (User) TableName() string {
	return "user"
}
