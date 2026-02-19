package productrepo

import (
	"context"

	"github.com/google/uuid"
	domain "github.com/webomindapps-dev/coolaid-backend/internal/domain/product"
	"github.com/webomindapps-dev/coolaid-backend/internal/generated/sqlc"
)

func (p *productQueries) GetVendorListingsByIds(
	ctx context.Context,
	ids []string,
) ([]domain.VendorListingRow, error) {

	uIDs := make([]uuid.UUID, 0, len(ids))

	for _, id := range ids {
		uid, err := uuid.Parse(id)
		if err != nil {
			continue // skip invalid UUIDs
		}
		uIDs = append(uIDs, uid)
	}

	// Optional: if no valid UUIDs, return empty slice
	if len(uIDs) == 0 {
		return []domain.VendorListingRow{}, nil
	}

	rows, err := p.q.GetVendorListingsByIds(ctx, uIDs)
	if err != nil {
		return nil, err
	}

	result := make([]domain.VendorListingRow, 0, len(rows))

	for _, r := range rows {
		result = append(result, domain.VendorListingRow{
			ID:            r.ID.String(),
			ProductPartNo: r.ProductPartNo,
			VendorName:    r.VendorName,
			VendorPartNo:  r.VendorPartNo,
			VendorPrice:   r.VendorMrp,
		})
	}

	return result, nil
}

func (p *productQueries) GetVendorsByPartNo(
	ctx context.Context,
	partNo string,
) ([]domain.VendorListingRow, error) {

	rows, err := p.q.GetVendorListingsByProductPartNo(ctx, partNo)
	if err != nil {
		return nil, err
	}

	result := make([]domain.VendorListingRow, 0, len(rows))

	for _, r := range rows {
		result = append(result, domain.VendorListingRow{
			ID:            r.ID.String(),
			ProductPartNo: r.ProductPartNo,
			VendorName:    r.VendorName,
			VendorPartNo:  r.VendorPartNo,
			VendorPrice:   r.VendorMrp,
		})
	}

	return result, nil
}

func (p *productQueries) CreateVendorListingWithID(
	ctx context.Context,
	params domain.CreateVendorListingParams,
) (*domain.VendorListingRow, error) {

	id, err := uuid.Parse(params.ID)
	if err != nil {
		return nil, err
	}

	row, err := p.q.CreateVendorListingWithID(ctx, sqlc.CreateVendorListingWithIDParams{
		ID:            id,
		ProductPartNo: params.ProductPartNo,
		VendorName:    params.VendorName,
		VendorPartNo:  params.VendorPartNo,
		VendorMrp:     params.VendorPrice,
	})
	if err != nil {
		return nil, err
	}

	return &domain.VendorListingRow{
		ID:            row.ID.String(),
		ProductPartNo: row.ProductPartNo,
		VendorName:    row.VendorName,
		VendorPartNo:  row.VendorPartNo,
		VendorPrice:   row.VendorMrp,
	}, nil
}

func (p *productQueries) DeleteVendorsByPartNo(
	ctx context.Context,
	partNo string,
) error {
	return p.q.DeleteVendorListingsByProductPartNo(ctx, partNo)
}
