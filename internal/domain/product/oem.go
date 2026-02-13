package product

import (
	"context"

	"github.com/google/uuid"
)

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
		return nil, nil
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
