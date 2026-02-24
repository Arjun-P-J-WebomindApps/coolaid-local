package typesense

import (
	"context"
	"fmt"
)

func (s *Service) Delete(
	ctx context.Context,
	collection string,
	id string,
) error {

	_, err := s.client.
		Collection(collection).
		Document(id).
		Delete(ctx)

	if err != nil {
		return fmt.Errorf("typesense delete failed: %w", err)
	}

	return nil
}
