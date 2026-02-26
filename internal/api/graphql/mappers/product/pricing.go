package product_mapper

import (
	"github.com/webomindapps-dev/coolaid-backend/internal/domain/product"
	"github.com/webomindapps-dev/coolaid-backend/internal/generated/graphql/model"
)

func MapCreatePricing(pricingData *model.ProductPricingData) product.CreateProductPricingInput {

	return product.CreateProductPricingInput{
		BasicPrice: pricingData.BasicPrice,
		Freight:    pricingData.Freight,
		Gst:        pricingData.Gst,

		AcWorkshop:    pricingData.AcWorkshop,
		AcWorkshopPer: pricingData.AcWorkshopPer,
		AcWorkshopAmt: pricingData.AcWorkshopAmt,

		MultibrandWorkshop:    pricingData.MultibrandWorkshop,
		MultibrandWorkshopPer: pricingData.MultibrandWorkshopPer,
		MultibrandWorkshopAmt: pricingData.MultibrandWorkshopAmt,

		AutoTrader:    pricingData.AutoTrader,
		AutoTraderPer: pricingData.AutoTraderPer,
		AutoTraderAmt: pricingData.AutoTraderAmt,

		AcTrader:    pricingData.AcTrader,
		AcTraderPer: pricingData.AcTraderPer,
		AcTraderAmt: pricingData.AcTraderAmt,

		OutstationClassA: pricingData.OutstationClassA,
		OutstationNote:   pricingData.OutstationNote,

		OemMrp:                  pricingData.OemMrp,
		UnitMeasure:             pricingData.UnitMeasure,
		MinimumPurchaseQuantity: pricingData.MinimumPurchaseQuantity,
	}
}

func MapUpdatePricing(pricingData *model.ProductPricingData) product.UpdateProductPricingInput {

	return product.UpdateProductPricingInput{
		BasicPrice: &pricingData.BasicPrice,
		Freight:    &pricingData.Freight,
		Gst:        &pricingData.Gst,

		AcWorkshop:    &pricingData.AcWorkshop,
		AcWorkshopPer: &pricingData.AcWorkshopPer,
		AcWorkshopAmt: &pricingData.AcWorkshopAmt,

		MultibrandWorkshop:    &pricingData.MultibrandWorkshop,
		MultibrandWorkshopPer: &pricingData.MultibrandWorkshopPer,
		MultibrandWorkshopAmt: &pricingData.MultibrandWorkshopAmt,

		AutoTrader:    &pricingData.AutoTrader,
		AutoTraderPer: &pricingData.AutoTraderPer,
		AutoTraderAmt: &pricingData.AutoTraderAmt,

		AcTrader:    &pricingData.AcTrader,
		AcTraderPer: &pricingData.AcTraderPer,
		AcTraderAmt: &pricingData.AcTraderAmt,

		OutstationClassA: &pricingData.OutstationClassA,
		OutstationNote:   &pricingData.OutstationNote,

		OemMrp:                  &pricingData.OemMrp,
		UnitMeasure:             &pricingData.UnitMeasure,
		MinimumPurchaseQuantity: &pricingData.MinimumPurchaseQuantity,
	}
}

func MapFetchedPricing(priceData *product.ProductPricing) *model.ProductPricing {

	return &model.ProductPricing{
		BasicPrice: priceData.BasicPrice,
		Freight:    priceData.Freight,
		Gst:        priceData.Gst,

		AcWorkshop:    priceData.AcWorkshop,
		AcWorkshopPer: priceData.AcWorkshopPer,
		AcWorkshopAmt: priceData.AcWorkshopAmt,

		MultibrandWorkshop:    priceData.MultibrandWorkshop,
		MultibrandWorkshopPer: priceData.MultibrandWorkshopPer,
		MultibrandWorkshopAmt: priceData.MultibrandWorkshopAmt,

		AutoTrader:    priceData.AutoTrader,
		AutoTraderPer: priceData.AutoTraderPer,
		AutoTraderAmt: priceData.AutoTraderAmt,

		AcTrader:    priceData.AcTrader,
		AcTraderPer: priceData.AcTraderPer,
		AcTraderAmt: priceData.AcTraderAmt,

		OutstationClassA: priceData.OutstationClassA,
		OutstationNote:   priceData.OutstationNote,

		OemMrp:                  priceData.OemMrp,
		UnitMeasure:             priceData.UnitMeasure,
		MinimumPurchaseQuantity: priceData.MinimumPurchaseQuantity,
	}
}
