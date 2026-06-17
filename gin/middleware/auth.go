package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/javahongxi/golab/gin/response"
	"github.com/javahongxi/golab/gin/util"
)

const userIDKey = "user_id"

func GetUserID(c *gin.Context) uint64 {
	if uid, exists := c.Get(userIDKey); exists {
		return uid.(uint64)
	}
	return 0
}

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response.Unauthorized(c)
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			response.Unauthorized(c)
			return
		}

		token := parts[1]
		claims, err := util.ParseToken(token)
		if err != nil {
			response.Unauthorized(c)
			return
		}

		c.Set(userIDKey, claims.UserID)
		c.Next()
	}
}