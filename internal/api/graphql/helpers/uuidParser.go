package helpers

import (
	"github.com/google/uuid"
	"github.com/webomindapps-dev/coolaid-backend/internal/domain/techspec"
)

func ParseUUIDOrInternal(id string) (uuid.UUID, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return uuid.Nil, techspec.ErrInternal
	}
	return uid, nil
}
