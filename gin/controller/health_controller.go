package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/javahongxi/golab/gin/middleware"
)

type HealthController struct{}

func NewHealthController() *HealthController {
	return &HealthController{}
}

func (c *HealthController) Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message":    "pong",
		"request_id": middleware.GetRequestID(ctx),
	})
}

func (c *HealthController) Health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"service": "gin-demo",
	})
}