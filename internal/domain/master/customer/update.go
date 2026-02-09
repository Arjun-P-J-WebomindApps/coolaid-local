package customer

import "context"

func (s *Service) Update(
	ctx context.Context,
	input UpdateCustomerInput,
) (*Customer, error) {

	row, err := s.DB.Queries().UpdateCustomer(ctx, UpdateCustomerParams{
		ID:                  input.ID,
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
		return nil, ErrCustomerNotFound
	}

	return mapRowToModel(row), nil
}
