package sqlnull

import (
	"github.com/google/uuid"
)

func UUID(id *string) uuid.NullUUID {
	if id == nil {
		return uuid.NullUUID{}
	}

	u, err := uuid.Parse(*id)
	if err != nil {
		// invalid UUID → treat as NULL
		return uuid.NullUUID{}
	}

	return uuid.NullUUID{
		UUID:  u,
		Valid: true,
	}
}
