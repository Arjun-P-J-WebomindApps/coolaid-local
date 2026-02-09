package routes

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/webomindapps-dev/coolaid-backend/typesense"
)

// SearchHealthZ checks if search engine is reachable
func SearchHealthZ(c *gin.Context) {
	if typesense.TS == nil || typesense.TS.Client == nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status":  "down",
			"service": "search",
			"error":   "typesense not initialized",
		})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 3*time.Second)
	defer cancel()

	if _, err := typesense.TS.Client.Health(ctx, 3*time.Second); err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status":  "down",
			"service": "search",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"service": "search",
	})
}
