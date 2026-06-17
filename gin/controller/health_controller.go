package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/javahongxi/golab/gin/middleware"
	"github.com/javahongxi/golab/gin/response"
)

type HealthController struct{}

func NewHealthController() *HealthController {
	return &HealthController{}
}

func (c *HealthController) Ping(ctx *gin.Context) {
	response.Success(ctx, gin.H{
		"message":    "pong",
		"request_id": middleware.GetRequestID(ctx),
	})
}

func (c *HealthController) Health(ctx *gin.Context) {
	response.Success(ctx, gin.H{
		"status":  "ok",
		"service": "gin-demo",
	})
}