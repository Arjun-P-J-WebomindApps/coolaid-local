package vendor

import (
	"context"

	"github.com/webomindapps-dev/coolaid-backend/internal/domain/master/shared"
)

func (s *Service) Create(
	ctx context.Context,
	input CreateVendorInput,
) (*Vendor, error) {

	tx, qtx, err := s.DB.BeginTx(ctx)
	if err != nil {
		return nil, ErrInternal
	}
	defer tx.Rollback()

	id := shared.NewID().String()

	row, err := qtx.CreateVendor(ctx, CreateVendorParams{
		ID:          id,
		CompanyName: input.CompanyName,
	})
	if err != nil {
		return nil, ErrVendorExists
	}

	for _, c := range input.VendorContacts {
		contactID := shared.NewID().String()

		if err := qtx.CreateVendorContact(ctx, CreateVendorContactParams{
			ID:            contactID,
			VendorID:      id,
			ContactPerson: c.VendorContactPerson,
			MobileNumber:  c.MobileNumber,
			EmailID:       c.EmailID,
		}); err != nil {
			return nil, ErrInternal
		}
	}

	if err := tx.Commit(); err != nil {
		return nil, ErrInternal
	}

	v := mapRowToModel(row)
	for _, c := range input.VendorContacts {
		v.VendorContacts = append(v.VendorContacts, VendorContact{
			ContactPerson: c.VendorContactPerson,
			MobileNumber:  c.MobileNumber,
			EmailID:       c.EmailID,
		})
	}

	return v, nil
}
