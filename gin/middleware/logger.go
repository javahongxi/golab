package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var logger *zap.Logger

func InitLogger() {
	var err error
	logger, err = zap.NewProduction()
	if err != nil {
		panic(err)
	}
	zap.ReplaceGlobals(logger)
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		method := c.Request.Method

		c.Next()

		latency := time.Now().Sub(start)
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()

		logger.Info("incoming request",
			zap.String("method", method),
			zap.String("path", path),
			zap.String("ip", clientIP),
			zap.Int("status", statusCode),
			zap.Duration("latency", latency),
			zap.String("request_id", GetRequestID(c)),
		)
	}
}
