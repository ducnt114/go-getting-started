package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name  string `gorm:"column:name"`
	Age   int    `gorm:"column:age"`
	Books []*Book
}

func (User) TableName() string {
	return "user"
}
