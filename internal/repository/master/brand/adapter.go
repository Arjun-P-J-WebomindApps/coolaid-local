package brand_repo

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/webomindapps-dev/coolaid-backend/db"
	brand "github.com/webomindapps-dev/coolaid-backend/internal/domain/master/brands"
	"github.com/webomindapps-dev/coolaid-backend/internal/domain/master/shared"
	"github.com/webomindapps-dev/coolaid-backend/internal/generated/sqlc"
	"github.com/webomindapps-dev/coolaid-backend/internal/shared/sqlnull"
)

type BrandDB struct {
	db      *sql.DB
	queries *sqlc.Queries
}

func NewBrandRepository(dbCtx *db.DBContext) *BrandDB {
	return &BrandDB{
		db:      dbCtx.SqlDB,
		queries: dbCtx.Queries,
	}
}

type brandQueries struct {
	q *sqlc.Queries
}

// compile-time guarantee
var _ brand.Queries = (*brandQueries)(nil)

func (b *BrandDB) BeginTx(
	ctx context.Context,
) (*sql.Tx, brand.Queries, error) {

	tx, err := b.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, nil, err
	}

	return tx, &brandQueries{
		q: sqlc.New(tx),
	}, nil
}

func (b *BrandDB) Queries() brand.Queries {
	return &brandQueries{
		q: sqlc.New(b.db),
	}
}

func (b *brandQueries) GetBrandByID(
	ctx context.Context,
	id shared.ID,
) (*brand.BrandRow, error) {

	brandID, err := id.ToUUID(ctx)
	if err != nil {
		return nil, err
	}

	row, err := b.q.GetBrandById(ctx, brandID)
	if err != nil {
		return nil, err
	}

	return &brand.BrandRow{
		ID:    row.ID.String(),
		Name:  row.Name,
		Image: row.Image,
	}, nil
}

func (b *brandQueries) GetBrandByName(
	ctx context.Context,
	name string,
) (*brand.BrandRow, error) {

	row, err := b.q.GetBrandByName(ctx, name)
	if err != nil {
		return nil, err
	}

	return &brand.BrandRow{
		ID:    row.ID.String(),
		Name:  row.Name,
		Image: row.Image,
	}, nil
}

func (b *brandQueries) GetBrandListByName(
	ctx context.Context,
	name string,
) ([]brand.BrandRow, error) {

	rows, err := b.q.GetBrandListByName(ctx, name)
	if err != nil {
		return nil, err
	}

	out := make([]brand.BrandRow, 0, len(rows))
	for _, r := range rows {
		out = append(out, brand.BrandRow{
			ID:    r.ID.String(),
			Name:  r.Name,
			Image: r.Image,
		})
	}

	return out, nil
}

func (b *brandQueries) CreateBrand(
	ctx context.Context,
	p brand.CreateBrandParams,
) (*brand.BrandRow, error) {

	row, err := b.q.CreateBrand(ctx, sqlc.CreateBrandParams{
		ID:    uuid.MustParse(p.ID),
		Name:  p.Name,
		Image: p.Image,
	})
	if err != nil {
		return nil, err
	}

	return &brand.BrandRow{
		ID:    row.ID.String(),
		Name:  row.Name,
		Image: row.Image,
	}, nil
}

func (b *brandQueries) UpdateBrand(
	ctx context.Context,
	p brand.UpdateBrandParams,
) (*brand.BrandRow, error) {

	row, err := b.q.UpdateBrandByID(ctx, sqlc.UpdateBrandByIDParams{
		ID:        uuid.MustParse(p.ID),
		BrandName: sqlnull.String(p.Name),
		ImageUrl:  sqlnull.String(p.Image),
	})
	if err != nil {
		return nil, err
	}

	return &brand.BrandRow{
		ID:    row.ID.String(),
		Name:  row.Name,
		Image: row.Image,
	}, nil
}

func (b *brandQueries) DeleteBrand(
	ctx context.Context,
	id shared.ID,
) error {

	brandID, err := id.ToUUID(ctx)
	if err != nil {
		return err
	}

	return b.q.DeleteBrandByID(ctx, brandID)
}
