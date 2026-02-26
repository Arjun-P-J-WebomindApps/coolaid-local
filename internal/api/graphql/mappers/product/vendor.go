package product_mapper

import (
	"github.com/webomindapps-dev/coolaid-backend/internal/domain/product"
	"github.com/webomindapps-dev/coolaid-backend/internal/generated/graphql/model"
)

func MapCreateVendorInput(vendors []*model.VendorListData) []product.CreateVendorListingInput {
	vendorNumbers := make([]product.CreateVendorListingInput, 0, len(vendors))

	for _, vendor := range vendors {
		vendorNumbers = append(vendorNumbers, product.CreateVendorListingInput{
			VendorName:   vendor.VendorName,
			VendorPartNo: vendor.VendorPartNo,
			VendorPrice:  vendor.VendorPrice,
		})
	}

	return vendorNumbers
}
