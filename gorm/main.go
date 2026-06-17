package main

import (
	"fmt"
	"log"

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

	log.Println("Database connected successfully")
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
	if err := initDB(); err != nil {
		log.Fatalf("Failed to connect database: %v", err)
	}

	err := db.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatalf("Failed to migrate: %v", err)
	}
	log.Println("Table migrated successfully")

	newUser := &model.User{
		Username: "hongxi",
		Nickname: "洪熙",
		Gender:   1,
		Age:      30,
	}
	if err := createUser(newUser); err != nil {
		log.Printf("Failed to create user: %v", err)
	} else {
		log.Printf("Created user: ID=%d, Username=%s", newUser.ID, newUser.Username)
	}

	user, err := getUserByID(newUser.ID)
	if err != nil {
		log.Printf("Failed to get user: %v", err)
	} else {
		log.Printf("Got user: %+v", user)
	}

	userByUsername, err := getUserByUsername("hongxi")
	if err != nil {
		log.Printf("Failed to get user by username: %v", err)
	} else {
		log.Printf("Got user by username: %+v", userByUsername)
	}

	user.Nickname = "Hongxi"
	user.Age = 31
	if err := updateUser(user); err != nil {
		log.Printf("Failed to update user: %v", err)
	} else {
		log.Printf("Updated user: %+v", user)
	}

	users, err := listUsers()
	if err != nil {
		log.Printf("Failed to list users: %v", err)
	} else {
		log.Printf("User list (%d):", len(users))
		for _, u := range users {
			log.Printf("  - ID=%d, Username=%s, Nickname=%s", u.ID, u.Username, u.Nickname)
		}
	}

	if err := deleteUser(newUser.ID); err != nil {
		log.Printf("Failed to delete user: %v", err)
	} else {
		log.Printf("Deleted user: ID=%d", newUser.ID)
	}

	fmt.Println("\nGORM CRUD example completed!")
}
