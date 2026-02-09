package customer

import (
	"context"

	"github.com/webomindapps-dev/coolaid-backend/internal/domain/master/shared"
)

func (s *Service) Create(
	ctx context.Context,
	input CreateCustomerInput,
) (*Customer, error) {

	_, err := s.DB.Queries().GetCustomerByUniqueKey(
		ctx,
		input.CustomerCompanyName,
		input.ContactPerson,
		input.Mobile,
		input.Type,
	)
	if err == nil {
		return nil, ErrCustomerExists
	}

	row, err := s.DB.Queries().CreateCustomer(ctx, CreateCustomerParams{
		ID:                  shared.NewID().String(),
		CustomerCompanyName: input.CustomerCompanyName,
		ContactPerson:       input.ContactPerson,
		Mobile:              input.Mobile,
		Type:                input.Type,
		CustomerDesignation: input.CustomerDesignation,
		Address:             input.Address,
		Flat:                input.Flat,
		Street:              input.Street,
		City:                input.City,
		State:               input.State,
		Pincode:             input.Pincode,
		PaymentMode:         input.PaymentMode,
	})
	if err != nil {
		return nil, ErrInternal
	}

	return mapRowToModel(row), nil
}
