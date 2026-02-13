package productrepo

import (
	"context"

	"github.com/google/uuid"
	domain "github.com/webomindapps-dev/coolaid-backend/internal/domain/product"
	"github.com/webomindapps-dev/coolaid-backend/internal/generated/sqlc"
	"github.com/webomindapps-dev/coolaid-backend/internal/shared/sqlnull"
)

func (p *productQueries) GetPricingByPartNo(
	ctx context.Context,
	partNo string,
) (*domain.PricingRow, error) {

	row, err := p.q.GetPricingByPartNo(ctx, partNo)
	if err != nil {
		return nil, err
	}

	return mapPricing(row), nil
}

func (p *productQueries) CreatePricing(
	ctx context.Context,
	params domain.CreateProductPricingParams,
) (*domain.PricingRow, error) {

	id, err := uuid.Parse(params.ID)
	if err != nil {
		return nil, err
	}

	row, err := p.q.CreatePricing(ctx, sqlc.CreatePricingParams{
		ID:            id,
		ProductPartID: params.ProductPartID,
		BasicPrice:    params.BasicPrice,
		Freight:       params.Freight,
		Gst:           params.Gst,

		AcWorkshop:    params.AcWorkshop,
		AcWorkshopPer: params.AcWorkshopPer,
		AcWorkshopAmt: params.AcWorkshopAmt,

		MultibrandWorkshop:    params.MultibrandWorkshop,
		MultibrandWorkshopPer: params.MultibrandWorkshopPer,
		MultibrandWorkshopAmt: params.MultibrandWorkshopAmt,

		AutoTrader:    params.AutoTrader,
		AutoTraderPer: params.AutoTraderPer,
		AutoTraderAmt: params.AutoTraderAmt,

		AcTrader:    params.AcTrader,
		AcTraderPer: params.AcTraderPer,
		AcTraderAmt: params.AcTraderAmt,

		OutstationClassA: params.OutstationClassA,
		OutstationNote:   sqlnull.String(params.OutstationNote),

		OemMrp:                  params.OemMrp,
		UnitMeasure:             sqlnull.String(params.UnitMeasure),
		MinimumPurchaseQuantity: params.MinimumPurchaseQuantity,

		CreatedAt: params.CreatedAt,
		UpdatedAt: params.UpdatedAt,
	})
	if err != nil {
		return nil, err
	}

	return mapPricing(row), nil
}

func (p *productQueries) UpdatePricing(
	ctx context.Context,
	params domain.UpdateProductPricingParams,
) (*domain.PricingRow, error) {

	id, err := uuid.Parse(params.ID)
	if err != nil {
		return nil, err
	}

	row, err := p.q.UpdatePricing(ctx, sqlc.UpdatePricingParams{
		ID: id,

		BasicPrice: sqlnull.Float64(params.BasicPrice),
		Freight:    sqlnull.Float64(params.Freight),
		Gst:        sqlnull.Float64(params.Gst),

		AcWorkshop:    sqlnull.Float64(params.AcWorkshop),
		AcWorkshopPer: sqlnull.Float64(params.AcWorkshopPer),
		AcWorkshopAmt: sqlnull.Float64(params.AcWorkshopAmt),

		MultibrandWorkshop:    sqlnull.Float64(params.MultibrandWorkshop),
		MultibrandWorkshopPer: sqlnull.Float64(params.MultibrandWorkshopPer),
		MultibrandWorkshopAmt: sqlnull.Float64(params.MultibrandWorkshopAmt),

		AutoTrader:    sqlnull.Float64(params.AutoTrader),
		AutoTraderPer: sqlnull.Float64(params.AutoTraderPer),
		AutoTraderAmt: sqlnull.Float64(params.AutoTraderAmt),

		AcTrader:    sqlnull.Float64(params.AcTrader),
		AcTraderPer: sqlnull.Float64(params.AcTraderPer),
		AcTraderAmt: sqlnull.Float64(params.AcTraderAmt),

		OutstationClassA: sqlnull.Float64(params.OutstationClassA),
		OutstationNote:   sqlnull.String(params.OutstationNote),

		OemMrp:                  sqlnull.Float64(params.OemMrp),
		UnitMeasure:             sqlnull.String(params.UnitMeasure),
		MinimumPurchaseQuantity: sqlnull.Int32(params.MinimumPurchaseQuantity),

		UpdatedAt: params.UpdatedAt,
	})
	if err != nil {
		return nil, err
	}

	return mapPricing(row), nil
}

func (p *productQueries) DeletePricing(
	ctx context.Context,
	partNo string,
) error {
	return p.q.DeletePricing(ctx, partNo)
}

func mapPricing(row sqlc.ProductPricing) *domain.PricingRow {
	return &domain.PricingRow{
		ID:            row.ID.String(),
		ProductPartID: row.ProductPartID,
		BasicPrice:    row.BasicPrice,
		Freight:       row.Freight,
		Gst:           row.Gst,
		Tax:           row.Tax,

		AcWorkshop:    row.AcWorkshop,
		AcWorkshopPer: row.AcWorkshopPer,
		AcWorkshopAmt: row.AcWorkshopAmt,

		MultibrandWorkshop:    row.MultibrandWorkshop,
		MultibrandWorkshopPer: row.MultibrandWorkshopPer,
		MultibrandWorkshopAmt: row.MultibrandWorkshopAmt,

		AutoTrader:    row.AutoTrader,
		AutoTraderPer: row.AutoTraderPer,
		AutoTraderAmt: row.AutoTraderAmt,

		AcTrader:    row.AcTrader,
		AcTraderPer: row.AcTraderPer,
		AcTraderAmt: row.AcTraderAmt,

		OutstationClassA: row.OutstationClassA,
		OutstationNote:   sqlnull.StringPtr(row.OutstationNote),

		OemMrp:                  row.OemMrp,
		UnitMeasure:             sqlnull.StringPtr(row.UnitMeasure),
		MinimumPurchaseQuantity: row.MinimumPurchaseQuantity,

		CreatedAt: row.CreatedAt,
		UpdatedAt: row.UpdatedAt,
	}
}
