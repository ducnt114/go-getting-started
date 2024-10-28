package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name    string   `gorm:"column:name"`
	Age     int      `gorm:"column:age"`
	Pass    string   `gorm:"column:pass"`
	Salt    string   `gorm:"column:salt"`
	TwoFA   string   `gorm:"column:2fa"`
	Profile *Profile `gorm:"foreignKey:UserName;references:Name"`
	Books   []*Book
	Tags    Tags `gorm:"column:tags"`
}

func (*User) TableName() string {
	return "user"
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.Age < 50 {
		u.Age = u.Age * 2
	}
	if u.Tags == nil || len(u.Tags) <= 0 {
		u.Tags = []Tag{
			{
				Key: "key-1",
				Val: "val-1",
			},
			{
				Key: "key-2",
				Val: "val-2",
			},
		}
	}
	return
}

func (u *User) AfterCreate(tx *gorm.DB) (err error) {
	if u.Pass == "" {
		tx.Model(u).Update("pass", fmt.Sprintf("random-pass-%v", u.ID))
	}
	return
}

type Tag struct {
	Key string `json:"key"`
	Val string `json:"value"`
}

type Tags []Tag

func (t *Tags) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("failed to convert database value to []byte")
	}
	return json.Unmarshal(b, t)
}

func (t Tags) Value() (driver.Value, error) {
	b, err := json.Marshal(t)
	return b, err
}
