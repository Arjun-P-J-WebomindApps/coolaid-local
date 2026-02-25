package handler

import (
	"github.com/gin-gonic/gin"
	routeHandler "github.com/webomindapps-dev/coolaid-backend/internal/api/http/handler/routes"
	search_handler "github.com/webomindapps-dev/coolaid-backend/internal/api/http/handler/search"
)

type HttpHandler struct {
	RouteHandler  *routeHandler.RouteHandler
	SearchHandler *search_handler.SearchHandler
}

func NewHttpHandler(
	routeHandler *routeHandler.RouteHandler,
	searchHandler *search_handler.SearchHandler,
) *HttpHandler {
	return &HttpHandler{
		RouteHandler:  routeHandler,
		SearchHandler: searchHandler,
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
