package middleware

import (
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/javahongxi/golab/gin/config"
	"github.com/javahongxi/golab/gin/response"
)

type rateLimiter struct {
	mu      sync.Mutex
	clients map[string]*clientRate
}

type clientRate struct {
	requests int
	lastTime time.Time
}

var limiter *rateLimiter

func init() {
	limiter = &rateLimiter{
		clients: make(map[string]*clientRate),
	}
}

func RateLimit() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientIP := c.ClientIP()
		limiter.mu.Lock()

		if _, exists := limiter.clients[clientIP]; !exists {
			limiter.clients[clientIP] = &clientRate{
				requests: 0,
				lastTime: time.Now(),
			}
		}

		client := limiter.clients[clientIP]
		now := time.Now()

		if now.Sub(client.lastTime) > time.Minute {
			client.requests = 0
			client.lastTime = now
		}

		client.requests++
		limiter.mu.Unlock()

		if client.requests > config.Cfg.RateLimit {
			response.ErrorWithCode(c, 429, "too many requests")
			c.Abort()
			return
		}

		c.Next()
	}
}