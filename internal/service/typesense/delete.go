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

	if id == "" {
		return fmt.Errorf("typesense: missing document ID")
	}

	_, err := s.ctx.Client.
		Collection(collection).
		Document(id).
		Delete(ctx)

	if err != nil {
		return fmt.Errorf("typesense delete failed: %w", err)
	}

	return nil
}
