package auth

import (
	"context"
	"database/sql"

	"github.com/webomindapps-dev/coolaid-backend/db"
	"github.com/webomindapps-dev/coolaid-backend/internal/domain/auth"
	"github.com/webomindapps-dev/coolaid-backend/internal/generated/sqlc"
)

// AuthDB is the database adapter for the auth domain.
// It implements auth.DB and returns auth.Queries.
type AuthDB struct {
	db      *sql.DB
	queries *sqlc.Queries
}

func NewAuthRepository(dbCtx *db.DBContext) *AuthDB {
	return &AuthDB{
		db:      dbCtx.SqlDB,
		queries: dbCtx.Queries,
	}
}

type authQueries struct {
	q *sqlc.Queries
}

// Compile-time guarantee that authQueries implements auth.Queries
var _ auth.Queries = (*authQueries)(nil)

func (a *AuthDB) BeginTx(
	ctx context.Context,
) (*sql.Tx, auth.Queries, error) {

	tx, err := a.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, nil, err
	}

	return tx, &authQueries{
		q: sqlc.New(tx),
	}, nil
}

func (a *AuthDB) Queries() auth.Queries {
	return &authQueries{
		q: sqlc.New(a.db),
	}
}
