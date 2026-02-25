package handler

import (
	routeH "github.com/webomindapps-dev/coolaid-backend/internal/api/http/handler/routes"
	search_handler "github.com/webomindapps-dev/coolaid-backend/internal/api/http/handler/search"
	service "github.com/webomindapps-dev/coolaid-backend/internal/service/container"
)

type Server struct {
	Services *service.Container
	Handler  *HttpHandler
}

func NewHttpServer(services *service.Container) *Server {

	routerHandler := routeH.NewRouteHandler(services)
	searchHandler := search_handler.NewSearchHandler(services)

	handler := NewHttpHandler(
		routerHandler,
		searchHandler,
	)

	return &Server{
		Services: services,
		Handler:  handler,
	}
}
