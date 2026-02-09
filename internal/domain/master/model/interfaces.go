package models

import (
	"context"

	"github.com/webomindapps-dev/coolaid-backend/internal/domain/master/shared"
)

// Domain-level queries (NOT sqlc)
type Queries interface {
	// Reads
	GetModelByID(ctx context.Context, id shared.ID) (*ModelRow, error)
	GetModelsByCompanyAndModelNames(ctx context.Context, modelName string, companyName string) ([]ModelWithCompanyRow, error)

	// Writes
	CreateModel(ctx context.Context, p CreateModelParams) (*ModelRow, error)
	UpdateModel(ctx context.Context, p UpdateModelParams) (*ModelRow, error)
	DeleteModel(ctx context.Context, id shared.ID) error
}

type DB = shared.DB[Queries]
