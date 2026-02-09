package category

import (
	"context"

	"github.com/webomindapps-dev/coolaid-backend/internal/domain/master/shared"
)

type Queries interface {
	// Reads
	GetCategoriesByName(ctx context.Context, name string) ([]CategoryRow, error)
	GetCategoryByID(ctx context.Context, id shared.ID) (*CategoryRow, error)

	// Writes
	CreateCategory(ctx context.Context, p CreateCategoryParams) (*CategoryRow, error)
	UpdateCategory(ctx context.Context, p UpdateCategoryParams) (*CategoryRow, error)
	DeleteCategory(ctx context.Context, id shared.ID) error
}

type DB = shared.DB[Queries]
