package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/javahongxi/golab/gin/cache"
	"github.com/javahongxi/golab/gin/config"
	"github.com/javahongxi/golab/gin/middleware"
	"github.com/javahongxi/golab/gin/model"
	"github.com/javahongxi/golab/gin/routes"
	"go.uber.org/zap"
)

func main() {
	middleware.InitLogger()
	zap.L().Info("starting gin demo server")

	model.InitDB()
	cache.InitRedis()

	r := routes.SetupRouter()

	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", config.Cfg.ServerPort),
		Handler:      r,
		ReadTimeout:  time.Duration(config.Cfg.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(config.Cfg.WriteTimeout) * time.Second,
	}

	zap.L().Info("server listening on", zap.String("port", config.Cfg.ServerPort))
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		zap.L().Panic("failed to start server", zap.Error(err))
	}
}
