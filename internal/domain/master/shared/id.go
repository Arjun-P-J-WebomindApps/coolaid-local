package shared

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

type ID string

func NewID() ID {
	return ID(uuid.NewString())
}

func (id ID) String() string {
	return string(id)
}

func (id ID) ToUUID(ctx context.Context) (uuid.UUID, error) {
	u, err := uuid.Parse(string(id))
	if err != nil {
		return uuid.Nil, fmt.Errorf("invalid id %q: %w", id, err)
	}
	return u, nil
}
