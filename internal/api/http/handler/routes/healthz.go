package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthZ checks if the API process is running
func HealthZ(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"service": "coolaid-backend",
	})
}
