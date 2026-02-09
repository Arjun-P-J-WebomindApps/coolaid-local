package category

import "context"

func (s *Service) ListByName(
	ctx context.Context,
	name string,
) ([]Category, error) {

	rows, err := s.DB.Queries().GetCategoriesByName(ctx, name)
	if err != nil {
		return nil, ErrInternal
	}

	out := make([]Category, 0, len(rows))
	for _, r := range rows {
		out = append(out, *mapRowToModel(&r))
	}

	return out, nil
}
