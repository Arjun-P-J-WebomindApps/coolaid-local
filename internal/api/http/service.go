package handler

import (
	routeH "github.com/webomindapps-dev/coolaid-backend/internal/api/http/handler/routes"
	service "github.com/webomindapps-dev/coolaid-backend/internal/service/container"
)

type Server struct {
	Services *service.Container
	Handler  *HttpHandler
}

func NewHttpServer(services *service.Container) *Server {

	routerHandler := routeH.NewRouteHandler(services)

	handler := NewHttpHandler(
		routerHandler,
	)

	return &Server{
		Services: services,
		Handler:  handler,
	}
}
