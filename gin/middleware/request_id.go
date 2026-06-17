package middleware

import (
	"crypto/rand"
	"encoding/hex"

	"github.com/gin-gonic/gin"
)

const requestIDKey = "X-Request-ID"

func GetRequestID(c *gin.Context) string {
	if rid, exists := c.Get(requestIDKey); exists {
		return rid.(string)
	}
	return ""
}

func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		rid := generateRequestID()
		c.Set(requestIDKey, rid)
		c.Header(requestIDKey, rid)
		c.Next()
	}
}

func generateRequestID() string {
	b := make([]byte, 16)
	rand.Read(b)
	return hex.EncodeToString(b)
}