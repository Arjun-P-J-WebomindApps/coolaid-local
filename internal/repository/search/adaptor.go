package searchrepo

import (
	"context"
	"database/sql"

	"github.com/webomindapps-dev/coolaid-backend/db"
	domain "github.com/webomindapps-dev/coolaid-backend/internal/domain/search"
	"github.com/webomindapps-dev/coolaid-backend/internal/generated/sqlc"
)

type SearchDB struct {
	db *sql.DB
}

func NewSearchRepository(dbCtx *db.DBContext) *SearchDB {
	return &SearchDB{
		db: dbCtx.SqlDB,
	}
}

type searchQueries struct {
	q *sqlc.Queries
}

// Ensure interface implementation
var _ domain.Queries = (*searchQueries)(nil)

func (s *SearchDB) BeginTx(
	ctx context.Context,
) (*sql.Tx, domain.Queries, error) {

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, nil, err
	}

	return tx, &searchQueries{
		q: sqlc.New(tx),
	}, nil
}

func (s *SearchDB) Queries() domain.Queries {
	return &searchQueries{
		q: sqlc.New(s.db),
	}
}

func (s *SearchDB) NewQueriesFromTx(tx *sql.Tx) domain.Queries {
	return &searchQueries{
		q: sqlc.New(tx),
	}
}
