package graphql

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/vektah/gqlparser/v2/ast"

	resolver "github.com/webomindapps-dev/coolaid-backend/internal/api/graphql/resolver"
	generated "github.com/webomindapps-dev/coolaid-backend/internal/generated/graphql"
	service "github.com/webomindapps-dev/coolaid-backend/internal/service/container"
)

type Server struct {
	Handler *handler.Server
}

func NewServer(services *service.Container) *Server {
	cfg := generated.Config{
		Resolvers: resolver.NewResolver(services),
	}

	cfg.Directives.Role = resolver.RoleDirective

	h := handler.New(generated.NewExecutableSchema(cfg))

	h.AddTransport(transport.Options{})
	h.AddTransport(transport.GET{})
	h.AddTransport(transport.POST{})

	h.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	h.Use(extension.Introspection{})
	h.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	return &Server{Handler: h}
}
