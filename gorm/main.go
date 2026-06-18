package main

import (
	"fmt"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/javahongxi/golab/gorm/model"
)

var db *gorm.DB

func initDB() error {
	dsn := "root:root1234@tcp(localhost:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	zap.L().Info("Database connected successfully")
	return nil
}

func createUser(user *model.User) error {
	result := db.Create(user)
	return result.Error
}

func getUserByID(id uint) (*model.User, error) {
	var user model.User
	result := db.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func getUserByUsername(username string) (*model.User, error) {
	var user model.User
	result := db.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func updateUser(user *model.User) error {
	result := db.Save(user)
	return result.Error
}

func deleteUser(id uint) error {
	result := db.Delete(&model.User{}, id)
	return result.Error
}

func listUsers() ([]model.User, error) {
	var users []model.User
	result := db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func main() {
	logger, _ := zap.NewProduction()
	undo := zap.ReplaceGlobals(logger)
	defer undo()

	if err := initDB(); err != nil {
		zap.L().Fatal("Failed to connect database", zap.Error(err))
	}

	err := db.AutoMigrate(&model.User{})
	if err != nil {
		zap.L().Fatal("Failed to migrate", zap.Error(err))
	}
	zap.L().Info("Table migrated successfully")

	newUser := &model.User{
		Username: "hongxi",
		Nickname: "洪熙",
		Gender:   1,
		Age:      30,
	}
	if err := createUser(newUser); err != nil {
		zap.L().Error("Failed to create user", zap.Error(err))
	} else {
		zap.L().Info("Created user", zap.Uint("id", newUser.ID), zap.String("username", newUser.Username))
	}

	user, err := getUserByID(newUser.ID)
	if err != nil {
		zap.L().Error("Failed to get user", zap.Error(err))
	} else {
		zap.L().Info("Got user", zap.Any("user", user))
	}

	userByUsername, err := getUserByUsername("hongxi")
	if err != nil {
		zap.L().Error("Failed to get user by username", zap.Error(err))
	} else {
		zap.L().Info("Got user by username", zap.Any("user", userByUsername))
	}

	user.Nickname = "Hongxi"
	user.Age = 31
	if err := updateUser(user); err != nil {
		zap.L().Error("Failed to update user", zap.Error(err))
	} else {
		zap.L().Info("Updated user", zap.Any("user", user))
	}

	users, err := listUsers()
	if err != nil {
		zap.L().Error("Failed to list users", zap.Error(err))
	} else {
		zap.L().Info("User list", zap.Int("count", len(users)))
		for _, u := range users {
			zap.L().Info("  - user", zap.Uint("id", u.ID), zap.String("username", u.Username), zap.String("nickname", u.Nickname))
		}
	}

	if err := deleteUser(newUser.ID); err != nil {
		zap.L().Error("Failed to delete user", zap.Error(err))
	} else {
		zap.L().Info("Deleted user", zap.Uint("id", newUser.ID))
	}

	fmt.Println("\nGORM CRUD example completed!")
}
