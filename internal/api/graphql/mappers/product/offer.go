package product_mapper

import (
	"github.com/webomindapps-dev/coolaid-backend/internal/domain/product"
	"github.com/webomindapps-dev/coolaid-backend/internal/generated/graphql/model"
)

func MapCreateOfferInput(input *model.ProductOfferData) product.CreateProductOfferInput {
	return product.CreateProductOfferInput{
		IsOfferActive: input.IsOfferActive,
		StartDate:     input.StartDate,
		EndDate:       input.EndDate,

		AcTrader:   input.AcTrader,
		MultiBrand: input.MultiBrand,
		Autotrader: input.Autotrader,
		AcWorkshop: input.AcWorkshop,
	}
}

func MapUpdateOfferInput(input *model.ProductOfferData) product.UpdateProductOfferInput {
	return product.UpdateProductOfferInput{
		IsOfferActive: &input.IsOfferActive,
		StartDate:     &input.StartDate,
		EndDate:       &input.EndDate,

		AcTrader:   &input.AcTrader,
		MultiBrand: &input.MultiBrand,
		Autotrader: &input.Autotrader,
		AcWorkshop: &input.AcWorkshop,
	}
}

func MapFetchedOffer(input *product.ProductOffer) *model.ProductOffer {
	return &model.ProductOffer{
		IsOfferActive: input.IsOfferActive,
		StartDate:     input.StartDate.Format("01-02-2025"),
		EndDate:       input.EndDate.Format("01-02-2025"),
		AcTrader:      input.AcTrader,
		MultiBrand:    input.MultiBrand,
		Autotrader:    input.Autotrader,
		AcWorkshop:    input.AcWorkshop,
	}
}
