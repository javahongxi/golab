package model

import (
	"time"

	"github.com/javahongxi/golab/gin/cache"
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

const userCachePrefix = "user:"
const userListCacheKey = "user:list"

func getUserCacheKey(id uint64) string {
	return userCachePrefix + string(rune(id+'0'))
}

func getUsernameCacheKey(username string) string {
	return userCachePrefix + "username:" + username
}

func (r *userRepository) Create(user *User) error {
	if err := r.db.Create(user).Error; err != nil {
		return err
	}
	go func() {
		cache.Set(getUserCacheKey(user.ID), user, 5*time.Minute)
		cache.Set(getUsernameCacheKey(user.Username), user, 5*time.Minute)
		cache.Del(userListCacheKey)
	}()
	return nil
}

func (r *userRepository) FindByID(id uint64) (*User, error) {
	var user User
	cacheKey := getUserCacheKey(id)
	if err := cache.Get(cacheKey, &user); err == nil {
		return &user, nil
	}

	err := r.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}

	go cache.Set(cacheKey, user, 5*time.Minute)
	return &user, nil
}

func (r *userRepository) FindByUsername(username string) (*User, error) {
	var user User
	cacheKey := getUsernameCacheKey(username)
	if err := cache.Get(cacheKey, &user); err == nil {
		return &user, nil
	}

	err := r.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}

	go cache.Set(cacheKey, user, 5*time.Minute)
	return &user, nil
}

func (r *userRepository) Update(user *User) error {
	if err := r.db.Save(user).Error; err != nil {
		return err
	}
	go func() {
		cache.Set(getUserCacheKey(user.ID), user, 5*time.Minute)
		cache.Set(getUsernameCacheKey(user.Username), user, 5*time.Minute)
		cache.Del(userListCacheKey)
	}()
	return nil
}

func (r *userRepository) Delete(id uint64) error {
	user, err := r.FindByID(id)
	if err != nil {
		return err
	}
	if err := r.db.Delete(&User{}, id).Error; err != nil {
		return err
	}
	go func() {
		cache.Del(getUserCacheKey(id))
		cache.Del(getUsernameCacheKey(user.Username))
		cache.Del(userListCacheKey)
	}()
	return nil
}

func (r *userRepository) List(page, limit int) ([]*User, int64, error) {
	var users []*User
	var total int64

	err := r.db.Model(&User{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	err = r.db.Offset((page - 1) * limit).Limit(limit).Order("id DESC").Find(&users).Error
	if err != nil {
		return nil, 0, err
	}

	return users, total, nil
}
