package middleware

import (
	"context"

	"github.com/gin-gonic/gin"
)

type ctxKey string

const UserAgentKey ctxKey = "user_agent"

func UserAgentMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ua := c.Request.UserAgent()

		ctx := context.WithValue(
			c.Request.Context(),
			UserAgentKey,
			ua,
		)

		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func GetUserAgent(ctx context.Context) string {
	ua, ok := ctx.Value(UserAgentKey).(string)
	if !ok || ua == "" {
		return "unknown"
	}
	return ua
}
