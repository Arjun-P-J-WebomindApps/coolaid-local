package productrepo

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	domain "github.com/webomindapps-dev/coolaid-backend/internal/domain/product"
	"github.com/webomindapps-dev/coolaid-backend/internal/generated/sqlc"
	"github.com/webomindapps-dev/coolaid-backend/internal/shared/sqlnull"
)

func (p *productQueries) GetProductPartNos(
	ctx context.Context,
	search string,
) ([]string, error) {

	rows, err := p.q.GetProductPartNos(ctx, sqlnull.String(&search))
	if err != nil {
		return nil, err
	}

	result := make([]string, 0, len(rows))

	for _, r := range rows {
		result = append(result, r)
	}

	return result, nil
}

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

func (p *productQueries) CreateProductPart(
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

func (p *productQueries) UpdateProductPart(
	ctx context.Context,
	params domain.UpdateProductParams,
) (*domain.ProductRow, error) {

	id, err := uuid.Parse(params.ID)
	if err != nil {
		return nil, err
	}

	row, err := p.q.UpdateProductPartByID(ctx, sqlc.UpdateProductPartByIDParams{
		ID:            id,
		NewCompanyID:  sqlnull.UUID(params.CompanyID),
		NewModelID:    sqlnull.UUID(params.ModelID),
		NewBrandID:    sqlnull.UUID(params.BrandID),
		NewCategoryID: sqlnull.UUID(params.CategoryID),
		NewPartNo:     sqlnull.String(params.PartNo),
	})
	if err != nil {
		return nil, err
	}

	return mapProduct(row), nil
}

func (p *productQueries) DeleteProductPart(
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
