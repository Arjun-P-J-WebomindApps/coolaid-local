package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
)

type CORSConfig struct {
	AllowedOrigins []string
	AllowedMethods []string
	AllowedHeaders []string
}

func CORSMiddleware(cfg CORSConfig) gin.HandlerFunc {
	allowedOrigins := make(map[string]struct{})
	for _, o := range cfg.AllowedOrigins {
		allowedOrigins[o] = struct{}{}
	}

	return func(c *gin.Context) {
		origin := c.GetHeader("Origin")

		if _, ok := allowedOrigins[origin]; ok {
			h := c.Writer.Header()
			h.Set("Access-Control-Allow-Origin", origin)
			h.Set("Access-Control-Allow-Headers", strings.Join(cfg.AllowedHeaders, ","))
			h.Set("Access-Control-Allow-Methods", strings.Join(cfg.AllowedMethods, ","))
			h.Set("Access-Control-Allow-Credentials", "true")
		}

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
