package product_mapper

import (
	"context"

	"github.com/google/uuid"
	"github.com/webomindapps-dev/coolaid-backend/internal/domain/product"
	"github.com/webomindapps-dev/coolaid-backend/internal/generated/graphql/model"
)

func MapProduct(ctx context.Context, product *product.Product) *model.Product {
	if product == nil {
		return &model.Product{
			ID:     uuid.New(),
			PartNo: "",
		}
	}

	uid, err := product.ID.ToUUID(ctx)
	if err != nil {
		uid = uuid.New()
	}
	return &model.Product{
		ID:     uid,
		PartNo: product.PartNo,
	}
}

func MapProductList(ctx context.Context, details []*product.ProductDetails) []*model.ProductDetails {
	productsDetails := make([]*model.ProductDetails, 0, len(details))

	if details == nil {
		return []*model.ProductDetails{}
	}

	for _, data := range details {
		productsDetails = append(productsDetails, &model.ProductDetails{
			Company:      data.Product.CompanyName,
			Model:        data.Product.ModelName,
			Category:     data.Product.CategoryName,
			Brand:        data.Product.BrandName,
			BrandImage:   data.Product.BrandImage,
			PartNo:       data.Product.PartNo,
			ModelVariant: MapFetchedVariant(ctx, data.ModelVariant),
			Pricing:      MapFetchedPricing(data.Pricing),
			Inventory:    MapFetchedInventory(data.Inventory),
			Offer:        MapFetchedOffer(data.Offer),
		})
	}

	return productsDetails
}
