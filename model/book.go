package model

type Book struct {
	ID     uint   `gorm:"primarykey"`
	Name   string `gorm:"column:name"`
	Title  string `gorm:"column:title"`
	UserID uint   `gorm:"column:user_id"`
	User   *User
}

func (Book) TableName() string {
	return "book"
}
