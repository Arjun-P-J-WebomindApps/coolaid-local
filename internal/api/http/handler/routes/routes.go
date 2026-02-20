package routes

import service "github.com/webomindapps-dev/coolaid-backend/internal/service/container"

type RouteHandler struct {
	Services *service.Container
}

func NewRouteHandler(services *service.Container) *RouteHandler {
	return &RouteHandler{Services: services}
}
