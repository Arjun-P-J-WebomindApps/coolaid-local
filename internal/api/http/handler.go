package handler

import (
	"github.com/gin-gonic/gin"
	routeHandler "github.com/webomindapps-dev/coolaid-backend/internal/api/http/handler/routes"
)

type HttpHandler struct {
	RouteHandler *routeHandler.RouteHandler
}

func NewHttpHandler(
	routeHandler *routeHandler.RouteHandler,
) *HttpHandler {
	return &HttpHandler{
		RouteHandler: routeHandler,
	}
}

func (h *HttpHandler) HandleRequest(ctx *gin.Context) {
	// Example: Access a service from the container
	// authService := h.Server.Services.Auth

	// // Example usage of the service
	// _, err := authService.Login(ctx.Request.Context(), "username", "password")
	// if err != nil {
	// 	ctx.JSON(500, gin.H{"error": "Login failed"})
	// 	return
	// }

	// ctx.JSON(200, gin.H{"message": "Login successful"})
}
