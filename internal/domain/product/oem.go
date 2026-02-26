package product

import (
	"context"

	"github.com/google/uuid"
)

// ============================================================================
// PRODUCT AGGREGATE ( OEM Unique Row )
// ===========================================================================
func collectUniqueOEMIDs(variants []*ModelVariantRow) []string {
	seen := make(map[string]struct{})

	for _, v := range variants {
		for _, id := range v.OemIDs {
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

func (s *Service) buildOEMLookup(
	ctx context.Context,
	Q Queries,
	ids []string,
) (map[string]OEMListing, error) {

	if len(ids) == 0 {
		return map[string]OEMListing{}, nil
	}

	rows, err := Q.GetOemListingsByIds(ctx, ids)
	if err != nil {
		return nil, err
	}

	lookup := make(map[string]OEMListing, len(rows))

	for _, r := range rows {
		lookup[r.ID] = OEMListing{
			ID:        ID(r.ID),
			OemNumber: r.OemNumber,
			Price:     r.Price,
		}
	}

	return lookup, nil
}

func attachOEMToVariants(
	variants []*ModelVariantRow,
	lookup map[string]OEMListing,
) map[string][]OEMListing {

	result := make(map[string][]OEMListing)

	for _, v := range variants {
		partNo := v.PartNo

		for _, id := range v.OemIDs {
			if oem, ok := lookup[id]; ok {
				result[partNo] = append(result[partNo], oem)
			}
		}
	}

	return result
}

//
// ============================================================
// 🔹 OEM LISTING
// ============================================================
//

func (s *Service) getOEMByPartNo(
	ctx context.Context,
	Q Queries,
	partNo string,
) ([]OEMListing, error) {

	if partNo == "" {
		return nil, ErrInvalidInput
	}

	rows, err := Q.GetOEMByPartNo(ctx, partNo)
	if err != nil {
		return nil, err
	}

	result := make([]OEMListing, 0, len(rows))
	for _, r := range rows {
		result = append(result, OEMListing{
			ID:        ID(r.ID),
			OemNumber: r.OemNumber,
			Price:     r.Price,
		})
	}

	return result, nil
}

func (s *Service) createOEMs(
	ctx context.Context,
	Q Queries,
	partNo string,
	oems []CreateOEMInput, // assuming you have an input struct
) ([]string, error) {

	if partNo == "" {
		return nil, ErrInvalidInput
	}

	if len(oems) == 0 {
		return []string{}, nil
	}

	listings := make([]string, 0, len(oems))

	for _, oem := range oems {

		if oem.OemNumber == "" {
			return nil, ErrInvalidInput
		}

		row, err := Q.CreateOemListingWithID(
			ctx,
			CreateOEMParams{
				ID:        uuid.New().String(),
				PartNo:    partNo,
				OemNumber: oem.OemNumber,
				Price:     oem.Price,
			},
		)
		if err != nil {
			return nil, err
		}

		listings = append(listings, row.ID)
	}

	return listings, nil
}

func (s *Service) deleteOEMByPartNo(
	ctx context.Context,
	Q Queries,
	partNo string,
) error {

	if partNo == "" {
		return ErrInvalidInput
	}

	return Q.DeleteOEMByPartNo(ctx, partNo)
}
