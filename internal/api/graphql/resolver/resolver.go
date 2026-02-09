package resolvers

import service "github.com/webomindapps-dev/coolaid-backend/internal/service/container"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Services *service.Container
}

func NewResolver(services *service.Container) *Resolver {
	return &Resolver{
		Services: services,
	}
}
