package search

import "context"

type SearchEngine interface {
	Search(ctx context.Context, req SearchRequest) (*SearchResponse, error)
}

type Indexer interface {
	Index(ctx context.Context, req IndexRequest) error
	Update(ctx context.Context, req IndexRequest) error
	Delete(ctx context.Context, collection string, id string) error
}

type Queries interface {
	GetSimilarModels(ctx context.Context, token string) ([]string, error)
	GetPartSuggestions(ctx context.Context, token string) ([]PartSuggestionResponse, error)
	GetOemSuggestions(ctx context.Context, token string) ([]PartSuggestionResponse, error)
	GetVendorSuggestions(ctx context.Context, token string) ([]PartSuggestionResponse, error)
}

type DB interface {
	Queries() Queries
}
