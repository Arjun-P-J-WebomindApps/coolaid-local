package productrepo

import (
	"context"
	"time"

	"github.com/google/uuid"
	domain "github.com/webomindapps-dev/coolaid-backend/internal/domain/product"
	"github.com/webomindapps-dev/coolaid-backend/internal/generated/sqlc"
	"github.com/webomindapps-dev/coolaid-backend/internal/shared/ptr"
	"github.com/webomindapps-dev/coolaid-backend/internal/shared/sqlnull"
)

func (p *productQueries) GetInventoryByPartNo(
	ctx context.Context,
	partNo string,
) (*domain.InventoryRow, error) {

	row, err := p.q.GetInventoryProductByPartNo(ctx, partNo)
	if err != nil {
		return nil, err
	}

	return mapInventory(row), nil
}

func (p *productQueries) CreateInventory(
	ctx context.Context,
	params domain.CreateProductInventoryParams,
) (*domain.InventoryRow, error) {

	id, err := uuid.Parse(params.ID)
	if err != nil {
		return nil, err
	}

	row, err := p.q.CreateInventoryProduct(ctx, sqlc.CreateInventoryProductParams{
		ID:                   id,
		PartNo:               params.PartNo,
		MinimumOrderLevel:    params.MinimumOrderLevel,
		MaximumOrderLevel:    params.MaximumOrderLevel,
		QtyInStock:           params.QtyInStock,
		Location:             params.Location,
		IsFlash:              sqlnull.Bool(&params.IsFlash),
		IsRequestedForSupply: sqlnull.Bool(&params.IsRequestedForSupply),
		VendorID:             sqlnull.UUID(params.VendorID),
	})
	if err != nil {
		return nil, err
	}

	return mapInventory(row), nil
}

func (p *productQueries) UpdateInventory(
	ctx context.Context,
	params domain.UpdateProductInventoryParams,
) (*domain.InventoryRow, error) {

	row, err := p.q.UpdateInventoryProductByPartNo(ctx, sqlc.UpdateInventoryProductByPartNoParams{
		PartNo:               params.PartNo,
		MinimumOrderLevel:    ptr.Int32Value(params.MinimumOrderLevel),
		MaximumOrderLevel:    ptr.Int32Value(params.MaximumOrderLevel),
		QtyInStock:           ptr.Int32Value(params.QtyInStock),
		Location:             ptr.String(params.Location),
		IsFlash:              sqlnull.Bool(params.IsFlash),
		IsRequestedForSupply: sqlnull.Bool(params.IsRequestedForSupply),
		VendorID:             sqlnull.UUID(params.VendorID),
	})
	if err != nil {
		return nil, err
	}

	return mapInventory(row), nil
}

func (p *productQueries) DeleteInventory(
	ctx context.Context,
	partNo string,
) error {
	return p.q.DeleteInventoryProductByPartNo(ctx, partNo)
}

func mapInventory(row sqlc.Inventory) *domain.InventoryRow {
	return &domain.InventoryRow{
		ID:                   row.ID.String(),
		PartNo:               row.PartNo,
		MinimumOrderLevel:    row.MinimumOrderLevel,
		MaximumOrderLevel:    row.MaximumOrderLevel,
		QtyInStock:           row.QtyInStock,
		Location:             &row.Location,
		IsFlash:              sqlnull.BoolValue(row.IsFlash),
		IsRequestedForSupply: sqlnull.BoolValue(row.IsRequestedForSupply),
		VendorID:             sqlnull.UUIDPtr(row.VendorID),
		UpdatedAt:            time.Now(),
	}
}
