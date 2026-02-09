package customer

import (
	"context"

	"github.com/webomindapps-dev/coolaid-backend/internal/domain/master/shared"
)

func (s *Service) GetByID(ctx context.Context, id string) (*Customer, error) {
	row, err := s.DB.Queries().GetCustomerByID(ctx, shared.ID(id))
	if err != nil {
		return nil, ErrCustomerNotFound
	}
	return mapRowToModel(row), nil
}

func (s *Service) Search(
	ctx context.Context,
	company *string,
	contactPerson *string,
	mobile *string,
) ([]Customer, error) {

	rows, err := s.DB.Queries().SearchCustomers(
		ctx,
		shared.ValueOrEmpty(company),
		shared.ValueOrEmpty(contactPerson),
		shared.ValueOrEmpty(mobile),
	)
	if err != nil {
		return nil, ErrInternal
	}

	out := make([]Customer, 0, len(rows))
	for _, r := range rows {
		out = append(out, *mapRowToModel(&r))
	}

	return out, nil
}
