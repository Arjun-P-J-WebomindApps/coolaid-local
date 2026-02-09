package categoryrepo

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/webomindapps-dev/coolaid-backend/db"
	"github.com/webomindapps-dev/coolaid-backend/internal/domain/master/category"
	"github.com/webomindapps-dev/coolaid-backend/internal/domain/master/shared"
	"github.com/webomindapps-dev/coolaid-backend/internal/generated/sqlc"
	"github.com/webomindapps-dev/coolaid-backend/internal/shared/sqlnull"
)

type CategoryDB struct {
	db      *sql.DB
	queries *sqlc.Queries
}

func NewCategoryRepository(dbCtx *db.DBContext) *CategoryDB {
	return &CategoryDB{
		db:      dbCtx.SqlDB,
		queries: dbCtx.Queries,
	}
}

type categoryQueries struct {
	q *sqlc.Queries
}

// compile-time guarantee
var _ category.Queries = (*categoryQueries)(nil)

func (c *CategoryDB) BeginTx(
	ctx context.Context,
) (*sql.Tx, category.Queries, error) {

	tx, err := c.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, nil, err
	}

	return tx, &categoryQueries{
		q: sqlc.New(tx),
	}, nil
}

func (c *CategoryDB) Queries() category.Queries {
	return &categoryQueries{
		q: sqlc.New(c.db),
	}
}

func (c *categoryQueries) GetCategoryByID(
	ctx context.Context,
	id shared.ID,
) (*category.CategoryRow, error) {

	categoryID, err := id.ToUUID(ctx)
	if err != nil {
		return nil, err
	}

	row, err := c.q.GetCategoryById(ctx, categoryID)
	if err != nil {
		return nil, err
	}

	return &category.CategoryRow{
		ID:    row.ID.String(),
		Name:  row.Name,
		Image: sqlnull.StringPtr(row.Image),
	}, nil
}

func (c *categoryQueries) GetCategoriesByName(
	ctx context.Context,
	name string,
) ([]category.CategoryRow, error) {

	rows, err := c.q.GetCategoriesByName(ctx, name)
	if err != nil {
		return nil, err
	}

	out := make([]category.CategoryRow, 0, len(rows))
	for _, r := range rows {
		out = append(out, category.CategoryRow{
			ID:    r.ID.String(),
			Name:  r.Name,
			Image: sqlnull.StringPtr(r.Image),
		})
	}

	return out, nil
}

func (c *categoryQueries) CreateCategory(
	ctx context.Context,
	p category.CreateCategoryParams,
) (*category.CategoryRow, error) {

	row, err := c.q.CreateCategories(ctx, sqlc.CreateCategoriesParams{
		ID:    uuid.MustParse(p.ID),
		Name:  p.Name,
		Image: sqlnull.String(p.Image),
	})
	if err != nil {
		return nil, err
	}

	return &category.CategoryRow{
		ID:    row.ID.String(),
		Name:  row.Name,
		Image: sqlnull.StringPtr(row.Image),
	}, nil
}

func (c *categoryQueries) UpdateCategory(
	ctx context.Context,
	p category.UpdateCategoryParams,
) (*category.CategoryRow, error) {

	row, err := c.q.UpdateCategoryByID(ctx, sqlc.UpdateCategoryByIDParams{
		ID:            uuid.MustParse(p.ID),
		CategoryName:  sqlnull.String(p.Name),
		CategoryImage: sqlnull.String(p.Image),
	})
	if err != nil {
		return nil, err
	}

	return &category.CategoryRow{
		ID:    row.ID.String(),
		Name:  row.Name,
		Image: sqlnull.StringPtr(row.Image),
	}, nil
}

func (c *categoryQueries) DeleteCategory(
	ctx context.Context,
	id shared.ID,
) error {

	categoryID, err := id.ToUUID(ctx)
	if err != nil {
		return err
	}

	return c.q.DeleteCategoryByID(ctx, categoryID)
}
