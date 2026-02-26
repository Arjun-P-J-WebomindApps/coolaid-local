package product_mapper

import (
	"github.com/webomindapps-dev/coolaid-backend/internal/domain/product"
	"github.com/webomindapps-dev/coolaid-backend/internal/generated/graphql/model"
)

func MapCreateOemNumberInput(oems []*model.OEMListingInput) []product.CreateOEMInput {

	oemNumbers := make([]product.CreateOEMInput, 0, len(oems))

	for _, oem := range oems {
		oemNumbers = append(oemNumbers, product.CreateOEMInput{
			OemNumber: oem.OemNumber,
			Price:     oem.Price,
		})
	}

	return oemNumbers
}
