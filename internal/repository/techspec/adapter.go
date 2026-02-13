package techspecrepo

import (
	"context"
	"database/sql"

	"github.com/webomindapps-dev/coolaid-backend/db"
	"github.com/webomindapps-dev/coolaid-backend/internal/domain/techspec"
	"github.com/webomindapps-dev/coolaid-backend/internal/generated/sqlc"
)

type TechSpecDB struct {
	db *sql.DB
}

func NewTechSpecRepository(dbCtx *db.DBContext) *TechSpecDB {
	return &TechSpecDB{
		db: dbCtx.SqlDB,
	}
}

type techSpecQueries struct {
	q *sqlc.Queries
}

// Compile-time guarantee
var _ techspec.Queries = (*techSpecQueries)(nil)

func (t *TechSpecDB) BeginTx(
	ctx context.Context,
) (*sql.Tx, techspec.Queries, error) {

	tx, err := t.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, nil, err
	}

	return tx, &techSpecQueries{
		q: sqlc.New(tx),
	}, nil
}

func (t *TechSpecDB) Queries() techspec.Queries {
	return &techSpecQueries{
		q: sqlc.New(t.db),
	}
}

func (r *TechSpecDB) NewQueriesFromTx(tx *sql.Tx) techspec.Queries {
	return &techSpecQueries{
		q: sqlc.New(tx),
	}
}
