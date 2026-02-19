package productrepo

import (
	"context"

	"github.com/google/uuid"
	domain "github.com/webomindapps-dev/coolaid-backend/internal/domain/product"
	"github.com/webomindapps-dev/coolaid-backend/internal/generated/sqlc"
)

func (p *productQueries) GetOemListingsByIds(
	ctx context.Context,
	ids []string,
) ([]domain.OEMListingRow, error) {

	uIDs := make([]uuid.UUID, 0, len(ids))

	for _, id := range ids {
		uid, err := uuid.Parse(id)
		if err != nil {
			continue // skip invalid UUIDs
		}
		uIDs = append(uIDs, uid)
	}

	// Optional: early return if nothing valid
	if len(uIDs) == 0 {
		return []domain.OEMListingRow{}, nil
	}

	rows, err := p.q.GetOemListingsByIds(ctx, uIDs)
	if err != nil {
		return nil, err
	}

	result := make([]domain.OEMListingRow, 0, len(rows))

	for _, r := range rows {
		result = append(result, domain.OEMListingRow{
			ID:        r.ID.String(),
			PartNo:    r.PartNo,
			OemNumber: r.OemNumber,
			Price:     r.Price,
		})
	}

	return result, nil
}

func (p *productQueries) GetOEMByPartNo(
	ctx context.Context,
	partNo string,
) ([]domain.OEMListingRow, error) {

	rows, err := p.q.GetOemListingsByPartNo(ctx, partNo)
	if err != nil {
		return nil, err
	}

	result := make([]domain.OEMListingRow, 0, len(rows))

	for _, r := range rows {
		result = append(result, domain.OEMListingRow{
			ID:        r.ID.String(),
			PartNo:    r.PartNo,
			OemNumber: r.OemNumber,
			Price:     r.Price,
		})
	}

	return result, nil
}

func (p *productQueries) CreateOemListingWithID(
	ctx context.Context,
	params domain.CreateOEMParams,
) (*domain.OEMListingRow, error) {

	row, err := p.q.CreateOemListing(ctx, sqlc.CreateOemListingParams{
		PartNo:    params.PartNo,
		OemNumber: params.OemNumber,
		Price:     params.Price,
	})
	if err != nil {
		return nil, err
	}

	return &domain.OEMListingRow{
		ID:        row.ID.String(),
		PartNo:    row.PartNo,
		OemNumber: row.OemNumber,
		Price:     row.Price,
	}, nil
}

func (p *productQueries) DeleteOEMByPartNo(
	ctx context.Context,
	partNo string,
) error {
	return p.q.DeleteOemListingsByPartNo(ctx, partNo)
}
