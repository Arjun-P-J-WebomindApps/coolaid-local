package product_mapper

import (
	"github.com/google/uuid"
	"github.com/webomindapps-dev/coolaid-backend/internal/domain/product"
	"github.com/webomindapps-dev/coolaid-backend/internal/generated/graphql/model"
)

func MapCreateInventoryInput(input *model.ProductInventoryData) product.CreateProductInventoryInput {
	return product.CreateProductInventoryInput{
		MinimumOrderLevel: input.MinimumOrderLevel,
		MaximumOrderLevel: input.MaximumOrderLevel,
		QtyInStock:        input.QtyInStock,
		Location:          input.Location,
		IsFlash:           input.IsFlash,
		VendorName:        input.VendorName,
	}
}

func MapUpdateInventoryInput(input *model.ProductInventoryData) product.UpdateProductInventoryInput {
	return product.UpdateProductInventoryInput{
		MinimumOrderLevel: &input.MinimumOrderLevel,
		MaximumOrderLevel: &input.MaximumOrderLevel,
		QtyInStock:        &input.QtyInStock,
		Location:          &input.Location,
		IsFlash:           &input.IsFlash,
		VendorName:        &input.VendorName,
	}
}

func MapFetchedInventory(data *product.ProductInventory) *model.ProductInventory {
	//TODO:
	uid := uuid.New()

	return &model.ProductInventory{
		PartNo:               data.PartNo,
		MinimumOrderLevel:    data.MinimumOrderLevel,
		MaximumOrderLevel:    data.MaximumOrderLevel,
		QtyInStock:           data.QtyInStock,
		Location:             data.Location,
		IsFlash:              data.IsFlash,
		IsRequestedForSupply: data.IsRequestedForSupply,
		VendorID:             &uid,
	}
}
