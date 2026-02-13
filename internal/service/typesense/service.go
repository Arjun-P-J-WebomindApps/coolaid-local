package typesense

import (
	"github.com/webomindapps-dev/coolaid-backend/internal/domain/search"
)

type Service struct {
	ctx *Context
}

func NewService(ctx *Context) search.Service {
	return &Service{ctx: ctx}
}
