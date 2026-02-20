package routes

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// SearchHealthZ checks if search engine is reachable
func (h RouteHandler) SearchHealthZ(c *gin.Context) {

	typesense := h.Services.TS
	if typesense == nil || typesense.Client == nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status":  "down",
			"service": "search",
			"error":   "typesense not initialized",
		})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 3*time.Second)
	defer cancel()

	if _, err := typesense.Client.Health(ctx, 3*time.Second); err != nil {
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
