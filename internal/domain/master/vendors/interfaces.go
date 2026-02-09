package vendor

import (
	"context"

	"github.com/webomindapps-dev/coolaid-backend/internal/domain/master/shared"
)

type Queries interface {
	// Reads
	GetVendorByCompanyName(ctx context.Context, companyName string) (*VendorRow, error)
	GetVendorsWithContacts(ctx context.Context, companyName string) ([]VendorWithContactRow, error)

	// Writes
	CreateVendor(ctx context.Context, p CreateVendorParams) (*VendorRow, error)
	DeleteVendor(ctx context.Context, id shared.ID) error

	DeleteVendorContacts(ctx context.Context, vendorID shared.ID) error
	CreateVendorContact(ctx context.Context, p CreateVendorContactParams) error
}

type DB = shared.DB[Queries]
