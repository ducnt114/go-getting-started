package model

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Code        string
	Name        string
	Permissions []*Permission `gorm:"many2many:role_permission;"`
}

func (Role) TableName() string {
	return "role"
}
