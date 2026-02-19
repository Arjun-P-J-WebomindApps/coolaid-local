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

	row, err := p.q.GetProductPricingFromPartNo(ctx, partNo)
	if err != nil {
		return nil, err
	}

	return mapPricing(row), nil
}

func (p *productQueries) GetSimilarPricing(
	ctx context.Context,
	params domain.SimilarPricingParams,
) ([]domain.SimilarPricingRow, error) {

	rows, err := p.q.GetSimilarPricingByPartNo(ctx, sqlc.GetSimilarPricingByPartNoParams{
		Name:   params.Company,
		Name_2: params.Model,
		Name_3: params.Category,
		Type:   params.Type,
	})
	if err != nil {
		return nil, err
	}

	result := make([]domain.SimilarPricingRow, 0, len(rows))

	for _, r := range rows {
		result = append(result, domain.SimilarPricingRow{
			BrandName: r.BrandName,
			Pricing: domain.Pricing{
				BasicPrice:              r.BasicPrice,
				Freight:                 r.Freight,
				Gst:                     r.Gst,
				Tax:                     r.Tax,
				AcWorkshop:              r.AcWorkshop,
				AcWorkshopPer:           r.AcWorkshopPer,
				AcWorkshopAmt:           r.AcWorkshopAmt,
				MultibrandWorkshop:      r.MultibrandWorkshop,
				MultibrandWorkshopPer:   r.MultibrandWorkshopPer,
				MultibrandWorkshopAmt:   r.MultibrandWorkshopAmt,
				AutoTrader:              r.AutoTrader,
				AutoTraderPer:           r.AutoTraderPer,
				AutoTraderAmt:           r.AutoTraderAmt,
				AcTrader:                r.AcTrader,
				AcTraderPer:             r.AcTraderPer,
				AcTraderAmt:             r.AcTraderAmt,
				OutstationClassA:        0,
				OutstationNote:          "",
				OemMrp:                  r.OemMrp,
				UnitMeasure:             r.UnitMeasure.String,
				MinimumPurchaseQuantity: r.MinimumPurchaseQuantity,
			},
		})
	}

	return result, nil
}

func (p *productQueries) CreatePricing(
	ctx context.Context,
	params domain.CreateProductPricingParams,
) (*domain.PricingRow, error) {

	id, err := uuid.Parse(params.ID)
	if err != nil {
		return nil, err
	}

	pID, err := uuid.Parse(params.ProductPartID)
	if err != nil {
		return nil, err
	}

	row, err := p.q.CreateProductPrice(ctx, sqlc.CreateProductPriceParams{
		ID:            id,
		ProductPartID: pID,
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

	row, err := p.q.UpdateProductPriceByPartNo(ctx, sqlc.UpdateProductPriceByPartNoParams{
		PartNo: params.PartNo,

		NewBasicPrice: sqlnull.Float64(params.BasicPrice),
		NewFreight:    sqlnull.Float64(params.Freight),
		NewGst:        sqlnull.Float64(params.Gst),

		NewAcWorkshop:    sqlnull.Float64(params.AcWorkshop),
		NewAcWorkshopPer: sqlnull.Float64(params.AcWorkshopPer),
		NewAcWorkshopAmt: sqlnull.Float64(params.AcWorkshopAmt),

		NewMultibrandWorkshop:    sqlnull.Float64(params.MultibrandWorkshop),
		NewMultibrandWorkshopPer: sqlnull.Float64(params.MultibrandWorkshopPer),
		NewMultibrandWorkshopAmt: sqlnull.Float64(params.MultibrandWorkshopAmt),

		NewAutoTrader:    sqlnull.Float64(params.AutoTrader),
		NewAutoTraderPer: sqlnull.Float64(params.AutoTraderPer),
		NewAutoTraderAmt: sqlnull.Float64(params.AutoTraderAmt),

		NewAcTrader:    sqlnull.Float64(params.AcTrader),
		NewAcTraderPer: sqlnull.Float64(params.AcTraderPer),
		NewAcTraderAmt: sqlnull.Float64(params.AcTraderAmt),

		NewOutstationClassA: sqlnull.Float64(params.OutstationClassA),
		NewOutstationNote:   sqlnull.String(params.OutstationNote),

		NewOemMrp:                  sqlnull.Float64(params.OemMrp),
		NewUnitMeasure:             sqlnull.String(params.UnitMeasure),
		NewMinimumPurchaseQuantity: sqlnull.Int32(params.MinimumPurchaseQuantity),
	})
	if err != nil {
		return nil, err
	}

	return mapPricing(sqlc.ProductPartPricing{
		ID:            row.ID,
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
		OutstationNote:   row.OutstationNote,

		MinimumPurchaseQuantity: row.MinimumPurchaseQuantity,
		MrpTemp:                 row.MrpTemp,
		OemMrp:                  row.OemMrp,

		UnitMeasure: row.UnitMeasure,

		CreatedAt: row.CreatedAt,
		UpdatedAt: row.UpdatedAt,
	}), nil

}

func (p *productQueries) DeletePricing(
	ctx context.Context,
	partNo string,
) error {
	return p.q.DeleteProductPriceByPartNo(ctx, partNo)
}

func mapPricing(row sqlc.ProductPartPricing) *domain.PricingRow {
	return &domain.PricingRow{
		ID:         row.ID.String(),
		BasicPrice: row.BasicPrice,
		Freight:    row.Freight,
		Gst:        row.Gst,
		Tax:        row.Tax,

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
