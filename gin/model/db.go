package model

import (
	"fmt"

	"github.com/javahongxi/golab/gin/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"go.uber.org/zap"
)

var DB *gorm.DB

func InitDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Cfg.DBUser,
		config.Cfg.DBPassword,
		config.Cfg.DBHost,
		config.Cfg.DBPort,
		config.Cfg.DBName,
	)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		zap.L().Panic("failed to connect database", zap.Error(err))
	}

	err = DB.AutoMigrate(&User{})
	if err != nil {
		zap.L().Error("failed to migrate database", zap.Error(err))
	}

	zap.L().Info("database connected successfully")
}