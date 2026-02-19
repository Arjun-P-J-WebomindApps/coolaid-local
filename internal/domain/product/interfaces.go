package product

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/webomindapps-dev/coolaid-backend/oplog"
)

// ------------------------------------------------------------
// ID
// ------------------------------------------------------------

type ID string

// ToUUID converts product.ID into uuid.UUID with logging
func (id ID) ToUUID(ctx context.Context) (uuid.UUID, error) {
	u, err := uuid.Parse(string(id))
	if err != nil {
		oplog.Error(ctx,
			"failed to convert product.ID to uuid",
			"id=", id,
			"error=", err,
		)
		return uuid.Nil, fmt.Errorf("invalid product id %q: %w", id, err)
	}
	return u, nil
}

// ------------------------------------------------------------
// Queries (Repository must implement this)
// ------------------------------------------------------------

type Queries interface {

	// ============================
	// PRODUCT (main table)
	// ============================

	GetProductByPartNo(ctx context.Context, partNo string) (*ProductRow, error)
	CreateProductPart(ctx context.Context, p CreateProductParams) (*ProductRow, error)
	UpdateProductPart(ctx context.Context, p UpdateProductParams) (*ProductRow, error)
	DeleteProductPart(ctx context.Context, partNo string) error

	// ============================
	// MODEL VARIANT
	// ============================

	GetModelVariantByPartNo(ctx context.Context, partNo string) (*ModelVariantRow, error)
	CreateModelVariant(ctx context.Context, p CreateModelVariantParams) (*ModelVariantRow, error)
	UpdateModelVariant(ctx context.Context, p UpdateModelVariantParams) (*ModelVariantRow, error)
	DeleteModelVariant(ctx context.Context, partNo string) error

	// ============================
	// PRICING
	// ============================

	GetPricingByPartNo(ctx context.Context, partNo string) (*PricingRow, error)
	CreatePricing(ctx context.Context, p CreateProductPricingParams) (*PricingRow, error)
	UpdatePricing(ctx context.Context, p UpdateProductPricingParams) (*PricingRow, error)
	DeletePricing(ctx context.Context, partNo string) error

	// ============================
	// INVENTORY
	// ============================

	GetInventoryByPartNo(ctx context.Context, partNo string) (*InventoryRow, error)
	CreateInventory(ctx context.Context, p CreateProductInventoryParams) (*InventoryRow, error)
	UpdateInventory(ctx context.Context, p UpdateProductInventoryParams) (*InventoryRow, error)
	DeleteInventory(ctx context.Context, partNo string) error

	// ============================
	// OFFER
	// ============================

	GetOfferByPartNo(ctx context.Context, partNo string) (*OfferRow, error)
	CreateOffer(ctx context.Context, p CreateProductOfferParams) (*OfferRow, error)
	UpdateOffer(ctx context.Context, p UpdateProductOfferParams) (*OfferRow, error)
	DeleteOffer(ctx context.Context, partNo string) error

	// ============================
	// OEM LISTING
	// ============================

	GetOEMByPartNo(ctx context.Context, partNo string) ([]OEMListingRow, error)
	GetOemListingsByIds(ctx context.Context, ids []string) ([]OEMListingRow, error)
	CreateOemListingWithID(ctx context.Context, p CreateOEMParams) (*OEMListingRow, error)
	DeleteOEMByPartNo(ctx context.Context, partNo string) error

	// ============================
	// VENDOR LISTING
	// ============================

	GetVendorsByPartNo(ctx context.Context, partNo string) ([]VendorListingRow, error)
	GetVendorListingsByIds(ctx context.Context, ids []string) ([]VendorListingRow, error)
	CreateVendorListingWithID(ctx context.Context, p CreateVendorListingParams) (*VendorListingRow, error)
	DeleteVendorsByPartNo(ctx context.Context, partNo string) error

	// ============================
	// FILTER / READ AGGREGATES
	// ============================
	GetProductDetailsList(ctx context.Context, params ProductsFilterParams) ([]*ProductAggregateRow, error)
	GetFilteredProducts(ctx context.Context, p FilterSelectionParams) ([]FilteredRow, error)
	GetSimilarPricing(ctx context.Context, p SimilarPricingParams) ([]SimilarPricingRow, error)
	GetProductPartNos(ctx context.Context, search string) ([]string, error)
}

// ------------------------------------------------------------
// DB (Transaction Boundary)
// ------------------------------------------------------------

type DB interface {
	BeginTx(ctx context.Context) (*sql.Tx, Queries, error)
	Queries() Queries
	NewQueriesFromTx(tx *sql.Tx) Queries
}
