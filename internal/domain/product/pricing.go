package product

import (
	"context"
	"time"

	"github.com/google/uuid"
)

//
// ============================================================
// 🔹 PRICING (Single Table Only)
// ============================================================
//

func (s *Service) getPricingByPartNo(
	ctx context.Context,
	Q Queries,
	partNo string,
) (*ProductPricing, error) {

	if partNo == "" {
		return nil, ErrInvalidInput
	}

	row, err := Q.GetPricingByPartNo(ctx, partNo)
	if err != nil {
		return nil, ErrPricingNotFound
	}

	return mapPricingRowToModel(row), nil
}

func (s *Service) createPricing(
	ctx context.Context,
	Q Queries,
	productPartID string,
	input CreateProductPricingInput,
) (*ProductPricing, error) {

	if productPartID == "" {
		return nil, ErrInvalidInput
	}

	now := time.Now()

	row, err := Q.CreatePricing(ctx, CreateProductPricingParams{
		ID:                      uuid.NewString(),
		ProductPartID:           productPartID,
		BasicPrice:              input.BasicPrice,
		Freight:                 input.Freight,
		Gst:                     input.Gst,
		AcWorkshop:              input.AcWorkshop,
		AcWorkshopPer:           input.AcWorkshopPer,
		AcWorkshopAmt:           input.AcWorkshopAmt,
		MultibrandWorkshop:      input.MultibrandWorkshop,
		MultibrandWorkshopPer:   input.MultibrandWorkshopPer,
		MultibrandWorkshopAmt:   input.MultibrandWorkshopAmt,
		AutoTrader:              input.AutoTrader,
		AutoTraderPer:           input.AutoTraderPer,
		AutoTraderAmt:           input.AutoTraderAmt,
		AcTrader:                input.AcTrader,
		AcTraderPer:             input.AcTraderPer,
		AcTraderAmt:             input.AcTraderAmt,
		OutstationClassA:        input.OutstationClassA,
		OutstationNote:          &input.OutstationNote,
		OemMrp:                  input.OemMrp,
		UnitMeasure:             &input.UnitMeasure,
		MinimumPurchaseQuantity: input.MinimumPurchaseQuantity,
		CreatedAt:               now,
		UpdatedAt:               now,
	})
	if err != nil {
		return nil, err
	}

	return mapPricingRowToModel(row), nil
}

func (s *Service) updatePricing(
	ctx context.Context,
	Q Queries,
	partNo string,
	input UpdateProductPricingInput,
) (*ProductPricing, error) {

	if partNo == "" {
		return nil, ErrInvalidInput
	}

	now := time.Now()

	row, err := Q.UpdatePricing(ctx, UpdateProductPricingParams{
		partNo: partNo,

		BasicPrice: input.BasicPrice,
		Freight:    input.Freight,
		Gst:        input.Gst,

		AcWorkshop:    input.AcWorkshop,
		AcWorkshopPer: input.AcWorkshopPer,
		AcWorkshopAmt: input.AcWorkshopAmt,

		MultibrandWorkshop:    input.MultibrandWorkshop,
		MultibrandWorkshopPer: input.MultibrandWorkshopPer,
		MultibrandWorkshopAmt: input.MultibrandWorkshopAmt,

		AutoTrader:    input.AutoTrader,
		AutoTraderPer: input.AutoTraderPer,
		AutoTraderAmt: input.AutoTraderAmt,

		AcTrader:    input.AcTrader,
		AcTraderPer: input.AcTraderPer,
		AcTraderAmt: input.AcTraderAmt,

		OutstationClassA: input.OutstationClassA,
		OutstationNote:   input.OutstationNote,

		OemMrp:      input.OemMrp,
		UnitMeasure: input.UnitMeasure,

		MinimumPurchaseQuantity: input.MinimumPurchaseQuantity,

		UpdatedAt: &now,
	})
	if err != nil {
		return nil, err
	}

	return mapPricingRowToModel(row), nil
}

//
// ============================================================
// 🔹 DELETE PRICING
// ============================================================
//

func (s *Service) deletePricing(
	ctx context.Context,
	Q Queries,
	id string,
) error {

	if id == "" {
		return ErrInvalidInput
	}

	// Optional: check exists first
	_, err := Q.GetPricingByPartNo(ctx, id)
	if err != nil {
		return ErrPricingNotFound
	}

	return Q.DeletePricing(ctx, id)
}
