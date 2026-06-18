package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/javahongxi/golab/config"
	"github.com/javahongxi/golab/gin/cache"
	ginconfig "github.com/javahongxi/golab/gin/config"
	"github.com/javahongxi/golab/gin/middleware"
	"github.com/javahongxi/golab/gin/model"
	"github.com/javahongxi/golab/gin/routes"
	"go.uber.org/zap"
)

func main() {
	// 初始化 Viper 配置管理
	config.Init("config/config.yaml")

	// 从 Viper 加载 gin 配置
	ginconfig.Init()

	// 监听配置文件变化，支持热更新
	config.WatchConfig(func() {
		ginconfig.Reload()
		zap.L().Info("config reloaded successfully")
	})

	middleware.InitLogger()
	zap.L().Info("starting gin demo server")

	model.InitDB()
	cache.InitRedis()

	r := routes.SetupRouter()

	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", ginconfig.Cfg.ServerPort),
		Handler:      r,
		ReadTimeout:  time.Duration(ginconfig.Cfg.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(ginconfig.Cfg.WriteTimeout) * time.Second,
	}

	zap.L().Info("server listening on", zap.String("port", ginconfig.Cfg.ServerPort))
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		zap.L().Panic("failed to start server", zap.Error(err))
	}
}
