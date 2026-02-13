package product

import (
	"context"

	"github.com/google/uuid"
)

func (s *Service) createVendors(
	ctx context.Context,
	Q Queries,
	partNo string,
	vendors []CreateVendorListingInput, // your domain input struct
) ([]string, error) {

	if partNo == "" {
		return nil, ErrInvalidInput
	}

	if len(vendors) == 0 {
		return nil, nil
	}

	listings := make([]string, 0, len(vendors))

	for _, vendor := range vendors {

		if vendor.VendorName == "" {
			return nil, ErrInvalidInput
		}

		row, err := Q.CreateVendorListingWithID(
			ctx,
			CreateVendorListingParams{
				ID:            uuid.New().String(),
				ProductPartNo: partNo,
				VendorName:    vendor.VendorName,
				VendorPartNo:  vendor.VendorPartNo,
				VendorPrice:   vendor.VendorPrice,
			},
		)
		if err != nil {
			return nil, err
		}

		listings = append(listings, row.ID)
	}

	return listings, nil
}
