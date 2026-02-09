package shared

import (
	"context"
	"database/sql"
)

// DB is shared because only Queries differ per domain.
type DB[Q any] interface {
	BeginTx(ctx context.Context) (*sql.Tx, Q, error)
	Queries() Q
}
