package graphql

import (
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
)

// PlaygroundHandler returns a gin handler for GraphQL Playground
func PlaygroundHandler(endpoint string) gin.HandlerFunc {
	h := playground.Handler("GraphQL Playground", endpoint)

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
