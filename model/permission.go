package model

import "gorm.io/gorm"

type Permission struct {
	gorm.Model
	Code string
	Name string
}

func (Permission) TableName() string {
	return "permission"
}
