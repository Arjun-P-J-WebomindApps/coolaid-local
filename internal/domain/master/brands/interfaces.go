package brand

import (
	"context"

	"github.com/webomindapps-dev/coolaid-backend/internal/domain/master/shared"
)

// Domain-level queries (NOT sqlc)
type Queries interface {
	// Reads
	GetBrandByID(ctx context.Context, id shared.ID) (*BrandRow, error)
	GetBrandByName(ctx context.Context, name string) (*BrandRow, error)
	GetBrandListByName(ctx context.Context, name string) ([]BrandRow, error)

	// Writes
	CreateBrand(ctx context.Context, p CreateBrandParams) (*BrandRow, error)
	UpdateBrand(ctx context.Context, p UpdateBrandParams) (*BrandRow, error)
	DeleteBrand(ctx context.Context, id shared.ID) error
}

type DB = shared.DB[Queries]
