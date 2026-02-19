package product

import "context"

func (s *Service) GetProductPartNos(
	ctx context.Context,
	search string,
) ([]string, error) {

	if search == "" {
		return nil, ErrInvalidInput
	}

	partNos, err := s.DB.Queries().GetProductPartNos(ctx, search)
	if err != nil {
		return nil, ErrInternal
	}

	return partNos, nil
}
