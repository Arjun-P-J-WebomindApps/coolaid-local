package vendor

import (
	"context"

	"github.com/webomindapps-dev/coolaid-backend/internal/domain/master/shared"
)

func (s *Service) Update(
	ctx context.Context,
	input UpdateVendorInput,
) (*Vendor, error) {

	tx, qtx, err := s.DB.BeginTx(ctx)
	if err != nil {
		return nil, ErrInternal
	}
	defer tx.Rollback()

	row, err := qtx.GetVendorByCompanyName(ctx, input.CompanyName)
	if err != nil {
		return nil, ErrVendorNotFound
	}

	if err := qtx.DeleteVendorContacts(ctx, shared.ID(row.ID)); err != nil {
		return nil, ErrInternal
	}

	for _, c := range input.VendorContacts {
		id := shared.NewID().String()

		if err := qtx.CreateVendorContact(ctx, CreateVendorContactParams{
			ID:            id,
			VendorID:      row.ID,
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
