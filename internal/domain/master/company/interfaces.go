package company

import (
	"context"

	"github.com/webomindapps-dev/coolaid-backend/internal/domain/master/shared"
)

// Domain-level queries (NOT sqlc)
type Queries interface {
	// Reads
	GetCompanyByID(ctx context.Context, id shared.ID) (*CompanyRow, error)
	GetCompanyByName(ctx context.Context, name string) (*CompanyRow, error)
	GetCompaniesByName(ctx context.Context, name string) ([]CompanyRow, error)

	// Writes
	CreateCompany(ctx context.Context, p CreateCompanyParams) (*CompanyRow, error)
	UpdateCompany(ctx context.Context, p UpdateCompanyParams) (*CompanyRow, error)
	DeleteCompany(ctx context.Context, id shared.ID) error
}

// Same semantics as auth:
// - Queries(): domain queries
// - BeginTx(): sqlc escape hatch
type DB = shared.DB[Queries]
