package search_handler

import service "github.com/webomindapps-dev/coolaid-backend/internal/service/container"

type SearchHandler struct {
	Services *service.Container
}

func NewSearchHandler(services *service.Container) *SearchHandler {
	return &SearchHandler{
		Services: services,
	}
}
