package product

import (
	"context"

	"github.com/google/uuid"
)

// ==============================================================
//
//	PRODUCT AGGREGATE (Get Vendors data)
//
// ===============================================================

// Assemble a list of unique ids i.e deduplicate from the given list
func collectUniqueVendorIDs(variants []*ModelVariantRow) []string {
	seen := make(map[string]struct{})

	for _, v := range variants {
		for _, id := range v.VendorIDs {
			if id != "" {
				seen[id] = struct{}{}
			}
		}
	}

	ids := make([]string, 0, len(seen))
	for id := range seen {
		ids = append(ids, id)
	}

	return ids
}

// Create a lookUp structure to looup vendor details from id
func (s *Service) buildVendorLookup(
	ctx context.Context,
	Q Queries,
	ids []string,
) (map[string]VendorListing, error) {

	if len(ids) == 0 {
		return map[string]VendorListing{}, nil
	}

	rows, err := Q.GetVendorListingsByIds(ctx, ids)
	if err != nil {
		return nil, err
	}

	lookup := make(map[string]VendorListing, len(rows))

	for _, r := range rows {
		lookup[r.ID] = VendorListing{
			ID:           ID(r.ID),
			VendorName:   r.VendorName,
			VendorPartNo: r.VendorPartNo,
			VendorPrice:  r.VendorPrice,
		}
	}

	return lookup, nil
}

// Lookup function to get data from the look up structure
func attachVendorsToVariants(
	variants []*ModelVariantRow,
	lookup map[string]VendorListing,
) map[string][]VendorListing {

	result := make(map[string][]VendorListing)

	for _, v := range variants {
		partNo := v.PartNo

		for _, id := range v.VendorIDs {
			if vendor, ok := lookup[id]; ok {
				result[partNo] = append(result[partNo], vendor)
			}
		}
	}

	return result
}

// =================================================================
// MUTATIONS
// =================================================================

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
		return []string{}, nil
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

func (s *Service) deleteVendors(
	ctx context.Context,
	Q Queries,
	partNo string,
) error {

	if partNo == "" {
		return ErrInvalidInput
	}

	if err := Q.DeleteVendorsByPartNo(ctx, partNo); err != nil {
		return ErrInternal
	}

	return nil
}
