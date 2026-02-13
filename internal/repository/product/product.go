package productrepo

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	domain "github.com/webomindapps-dev/coolaid-backend/internal/domain/product"
	"github.com/webomindapps-dev/coolaid-backend/internal/generated/sqlc"
)

func (p *productQueries) GetProductByPartNo(
	ctx context.Context,
	partNo string,
) (*domain.ProductRow, error) {

	row, err := p.q.GetProductPartsByPartNo(ctx, partNo)
	if err != nil {
		return nil, err
	}

	return mapProduct(row), nil
}

func (p *productQueries) CreateProduct(
	ctx context.Context,
	params domain.CreateProductParams,
) (*domain.ProductRow, error) {

	row, err := p.q.CreateProductParts(ctx, sqlc.CreateProductPartsParams{
		ID:         uuid.MustParse(params.ID),
		PartNo:     params.PartNo,
		CompanyID:  uuid.MustParse(params.CompanyID),
		ModelID:    uuid.MustParse(params.ModelID),
		BrandID:    uuid.MustParse(params.BrandID),
		CategoryID: uuid.MustParse(params.CategoryID),
		IsActive:   sql.NullBool{Valid: true, Bool: true},
		CreatedAt:  params.CreatedAt,
		UpdatedAt:  params.UpdatedAt,
	})
	if err != nil {
		return nil, err
	}

	return mapProduct(row), nil
}

func (p *productQueries) DeleteProduct(
	ctx context.Context,
	partNo string,
) error {
	return p.q.DeleteProductPartByPartNo(ctx, partNo)
}

func mapProduct(row sqlc.ProductPart) *domain.ProductRow {
	return &domain.ProductRow{
		ID:         row.ID.String(),
		PartNo:     row.PartNo,
		CompanyID:  row.CompanyID.String(),
		ModelID:    row.ModelID.String(),
		BrandID:    row.BrandID.String(),
		CategoryID: row.CategoryID.String(),
		IsActive:   true,
		CreatedAt:  row.CreatedAt,
		UpdatedAt:  row.UpdatedAt,
	}
}
