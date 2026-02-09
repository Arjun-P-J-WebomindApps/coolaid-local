package customer

import (
	"context"

	"github.com/webomindapps-dev/coolaid-backend/internal/domain/master/shared"
)

type Queries interface {
	// Reads
	GetCustomerByID(ctx context.Context, id shared.ID) (*CustomerRow, error)
	GetCustomerByUniqueKey(
		ctx context.Context,
		company string,
		contactPerson string,
		mobile string,
		customerType string,
	) (*CustomerRow, error)

	SearchCustomers(
		ctx context.Context,
		company string,
		contactPerson string,
		mobile string,
	) ([]CustomerRow, error)

	// Writes
	CreateCustomer(ctx context.Context, p CreateCustomerParams) (*CustomerRow, error)
	UpdateCustomer(ctx context.Context, p UpdateCustomerParams) (*CustomerRow, error)
	DeleteCustomer(ctx context.Context, id shared.ID) error
}

type DB = shared.DB[Queries]
