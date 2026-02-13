package search

import "context"

type Service interface {
	Index(ctx context.Context, req IndexRequest) error
	Update(ctx context.Context, req IndexRequest) error
	Delete(ctx context.Context, collection string, id string) error
}
