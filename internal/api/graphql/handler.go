package graphql

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/webomindapps-dev/coolaid-backend/config"
	resolvers "github.com/webomindapps-dev/coolaid-backend/internal/api/graphql/resolver"
)

func GinHandler(server *Server) gin.HandlerFunc {
	return func(c *gin.Context) {

		// 1️⃣ Base context
		ctx := context.WithValue(
			c.Request.Context(),
			config.GinContextKey,
			c,
		)

		// 2️⃣ Attach CookieWriter
		writer := resolvers.NewGinCookieWriter(c)
		ctx = context.WithValue(ctx, resolvers.CookieWriterKey{}, writer)

		server.Handler.ServeHTTP(
			c.Writer,
			c.Request.WithContext(ctx),
		)
	}
}
