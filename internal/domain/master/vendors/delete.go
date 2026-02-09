package vendor

import (
	"context"

	"github.com/webomindapps-dev/coolaid-backend/internal/domain/master/shared"
)

func (s *Service) Delete(ctx context.Context, id string) error {

	if err := s.DB.Queries().DeleteVendor(ctx, shared.ID(id)); err != nil {
		return ErrVendorNotFound
	}

	return nil
}
