package model

type Book struct {
	ID     uint   `gorm:"primarykey"`
	Name   string `gorm:"column:name"`
	UserID uint   `gorm:"column:user_id"`
	User   *User
}
