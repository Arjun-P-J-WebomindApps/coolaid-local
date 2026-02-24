package search

import "context"

// Business search service
type Service interface {
	GetSuggestions(ctx context.Context, query string, queryBy string) (*Results, error)
}

// Typesense adapter
type SearchEngine interface {
	Search(ctx context.Context, req SearchRequest) (*SearchResponse, error)
}

// Index adapter
type Indexer interface {
	Index(ctx context.Context, req IndexRequest) error
	Update(ctx context.Context, req IndexRequest) error
	Delete(ctx context.Context, collection string, id string) error
}

// DB-based priority suggestions
type PriorityRepository interface {

	//Even though number, check if its a model name
	GetModelSimilar(ctx context.Context, token string) (bool, error)

	// Base Data + Primary Lookup ID (Part No)
	GetPartSuggestions(ctx context.Context, token string, limit int) ([]PartData, error)
	GetOemSuggestions(ctx context.Context, token string, limit int) ([]PartData, error)
	GetVendorSuggestions(ctx context.Context, token string, limit int) ([]PartData, error)
}
