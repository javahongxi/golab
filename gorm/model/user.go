package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID         uint           `gorm:"column:id;primaryKey;autoIncrement"`
	Username   string         `gorm:"column:username;size:20;not null"`
	Nickname   string         `gorm:"column:nickname;size:20"`
	Gender     int8           `gorm:"column:gender;type:tinyint"`
	Age        int            `gorm:"column:age"`
	CreateDate time.Time      `gorm:"column:create_date;autoCreateTime"`
	UpdateDate time.Time      `gorm:"column:update_date;autoUpdateTime"`
	DeletedAt  gorm.DeletedAt `gorm:"column:deleted_at"`
}

func (u User) TableName() string {
	return "go_user"
}
