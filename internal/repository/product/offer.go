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

func (p *productQueries) GetOfferByPartNo(
	ctx context.Context,
	partNo string,
) (*domain.OfferRow, error) {

	row, err := p.q.GetOfferByPartNo(ctx, partNo)
	if err != nil {
		return nil, err
	}

	return mapOffer(row), nil
}

func (p *productQueries) CreateOffer(
	ctx context.Context,
	params domain.CreateProductOfferParams,
) (*domain.OfferRow, error) {

	id, err := uuid.Parse(params.ID)
	if err != nil {
		return nil, err
	}

	start, err := time.Parse("2006-01-02", params.StartDate)
	if err != nil {
		return nil, err
	}

	end, err := time.Parse("2006-01-02", params.EndDate)
	if err != nil {
		return nil, err
	}

	row, err := p.q.CreateOffer(ctx, sqlc.CreateOfferParams{
		ID:            id,
		PartNo:        params.PartNo,
		IsOfferActive: sqlnull.Bool(&params.IsOfferActive),
		StartDate:     start,
		EndDate:       end,
		AcTrader:      params.AcTrader,
		MultiBrand:    params.MultiBrand,
		Autotrader:    params.Autotrader,
		AcWorkshop:    params.AcWorkshop,
	})
	if err != nil {
		return nil, err
	}

	return mapOffer(row), nil
}

func (p *productQueries) UpdateOffer(
	ctx context.Context,
	params domain.UpdateProductOfferParams,
) (*domain.OfferRow, error) {

	row, err := p.q.UpdateOfferByPartNo(ctx, sqlc.UpdateOfferByPartNoParams{
		PartNo:        params.PartNo,
		IsOfferActive: sqlnull.Bool(params.IsOfferActive),
		StartDate:     params.StartDate,
		EndDate:       params.EndDate,
		AcTrader:      ptr.StringSliceValue(params.AcTrader),
		MultiBrand:    ptr.StringSliceValue(params.MultiBrand),
		Autotrader:    ptr.StringSliceValue(params.Autotrader),
		AcWorkshop:    ptr.StringSliceValue(params.AcWorkshop),
	})
	if err != nil {
		return nil, err
	}

	return mapOffer(row), nil
}

func (p *productQueries) DeleteOffer(
	ctx context.Context,
	partNo string,
) error {
	return p.q.DeleteOfferByPartNo(ctx, partNo)
}

func mapOffer(row sqlc.Offer) *domain.OfferRow {
	return &domain.OfferRow{
		ID:            row.ID.String(),
		PartNo:        row.PartNo,
		IsOfferActive: sqlnull.BoolValue(row.IsOfferActive),
		StartDate:     row.StartDate,
		EndDate:       row.EndDate,
		AcTrader:      row.AcTrader,
		MultiBrand:    row.MultiBrand,
		Autotrader:    row.Autotrader,
		AcWorkshop:    row.AcWorkshop,
	}
}
