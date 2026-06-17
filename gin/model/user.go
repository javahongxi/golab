package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID         uint64    `json:"id" gorm:"primaryKey;autoIncrement"`
	Username   string    `json:"username" gorm:"type:varchar(20);unique;not null"`
	Nickname   string    `json:"nickname" gorm:"type:varchar(20)"`
	Gender     int32     `json:"gender" gorm:"type:tinyint"`
	Age        int32     `json:"age" gorm:"type:int"`
	CreateDate time.Time `json:"create_date" gorm:"autoCreateTime"`
	UpdateDate time.Time `json:"update_date" gorm:"autoUpdateTime"`
}

func (User) TableName() string {
	return "go_user"
}

type CreateUserRequest struct {
	Username string `json:"username" binding:"required,min=3,max=20"`
	Nickname string `json:"nickname" binding:"max=20"`
	Gender   int32  `json:"gender" binding:"min=0,max=2"`
	Age      int32  `json:"age" binding:"min=0,max=150"`
}

type UpdateUserRequest struct {
	Nickname string `json:"nickname" binding:"max=20"`
	Age      int32  `json:"age" binding:"min=0,max=150"`
}

type UserRepository interface {
	Create(user *User) error
	FindByID(id uint64) (*User, error)
	FindByUsername(username string) (*User, error)
	Update(user *User) error
	Delete(id uint64) error
	List(page, limit int) ([]*User, int64, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) FindByID(id uint64) (*User, error) {
	var user User
	err := r.db.First(&user, id).Error
	return &user, err
}

func (r *userRepository) FindByUsername(username string) (*User, error) {
	var user User
	err := r.db.Where("username = ?", username).First(&user).Error
	return &user, err
}

func (r *userRepository) Update(user *User) error {
	return r.db.Save(user).Error
}

func (r *userRepository) Delete(id uint64) error {
	return r.db.Delete(&User{}, id).Error
}

func (r *userRepository) List(page, limit int) ([]*User, int64, error) {
	var users []*User
	var total int64
	err := r.db.Model(&User{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	err = r.db.Offset((page - 1) * limit).Limit(limit).Order("id DESC").Find(&users).Error
	return users, total, err
}