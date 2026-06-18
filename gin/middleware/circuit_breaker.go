package middleware

import (
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/javahongxi/golab/gin/response"
	"go.uber.org/zap"
)

type circuitBreaker struct {
	mu            sync.Mutex
	state         string
	failCount     int
	lastFail      time.Time
	resetAfter    time.Duration
	failThreshold int
}

const (
	stateClosed   = "closed"
	stateOpen     = "open"
	stateHalfOpen = "half-open"
)

var breaker *circuitBreaker

func init() {
	breaker = &circuitBreaker{
		state:         stateClosed,
		failCount:     0,
		resetAfter:    30 * time.Second,
		failThreshold: 5,
	}
}

func CircuitBreaker() gin.HandlerFunc {
	return func(c *gin.Context) {
		breaker.mu.Lock()
		state := breaker.state

		if state == stateOpen {
			if time.Since(breaker.lastFail) > breaker.resetAfter {
				breaker.state = stateHalfOpen
				state = stateHalfOpen
				zap.L().Info("circuit breaker transitioning to half-open")
			} else {
				breaker.mu.Unlock()
				response.ErrorWithCode(c, 503, "service unavailable")
				c.Abort()
				return
			}
		}
		breaker.mu.Unlock()

		c.Next()

		statusCode := c.Writer.Status()
		if statusCode >= 500 {
			breaker.mu.Lock()
			breaker.failCount++
			breaker.lastFail = time.Now()

			if breaker.failCount >= breaker.failThreshold {
				breaker.state = stateOpen
				zap.L().Warn("circuit breaker tripped, state changed to open")
			}
			breaker.mu.Unlock()
		} else {
			breaker.mu.Lock()
			if breaker.state == stateHalfOpen && statusCode < 500 {
				breaker.state = stateClosed
				breaker.failCount = 0
				zap.L().Info("circuit breaker reset, state changed to closed")
			}
			breaker.mu.Unlock()
		}
	}
}
