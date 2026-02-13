package productrepo

import (
	"context"
	"database/sql"

	"github.com/webomindapps-dev/coolaid-backend/db"
	domain "github.com/webomindapps-dev/coolaid-backend/internal/domain/product"
	"github.com/webomindapps-dev/coolaid-backend/internal/generated/sqlc"
)

type ProductDB struct {
	db      *sql.DB
	queries *sqlc.Queries
}

func NewProductRepository(dbCtx *db.DBContext) *ProductDB {
	return &ProductDB{
		db:      dbCtx.SqlDB,
		queries: dbCtx.Queries,
	}
}

type productQueries struct {
	q *sqlc.Queries
}

var _ domain.DB = (*ProductDB)(nil)
var _ domain.Queries = (*productQueries)(nil)

func (p *ProductDB) BeginTx(
	ctx context.Context,
) (*sql.Tx, domain.Queries, error) {

	tx, err := p.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, nil, err
	}

	return tx, &productQueries{
		q: sqlc.New(tx),
	}, nil
}

func (p *ProductDB) Queries() domain.Queries {
	return &productQueries{
		q: sqlc.New(p.db),
	}
}

func (p *ProductDB) NewQueriesFromTx(tx *sql.Tx) domain.Queries {
	return &productQueries{
		q: sqlc.New(tx),
	}
}
