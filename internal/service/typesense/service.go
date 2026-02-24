package typesense

import (
	"github.com/typesense/typesense-go/typesense"
	"github.com/webomindapps-dev/coolaid-backend/internal/domain/search"
)

type Service struct {
	client *typesense.Client
}

func NewService(client *typesense.Client) *Service {
	return &Service{client: client}
}

// Ensure interface implementation
var _ search.SearchEngine = (*Service)(nil)
var _ search.Indexer = (*Service)(nil)
