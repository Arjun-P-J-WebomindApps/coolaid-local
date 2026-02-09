package modelrepo

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/webomindapps-dev/coolaid-backend/db"
	domainmodel "github.com/webomindapps-dev/coolaid-backend/internal/domain/master/model"
	"github.com/webomindapps-dev/coolaid-backend/internal/domain/master/shared"
	"github.com/webomindapps-dev/coolaid-backend/internal/generated/sqlc"
	"github.com/webomindapps-dev/coolaid-backend/internal/shared/sqlnull"
)

type ModelDB struct {
	db      *sql.DB
	queries *sqlc.Queries
}

func NewModelRepository(dbCtx *db.DBContext) *ModelDB {
	return &ModelDB{
		db:      dbCtx.SqlDB,
		queries: dbCtx.Queries,
	}
}

type modelQueries struct {
	q *sqlc.Queries
}

// compile-time guarantee
var _ domainmodel.Queries = (*modelQueries)(nil)

func (m *ModelDB) BeginTx(
	ctx context.Context,
) (*sql.Tx, domainmodel.Queries, error) {

	tx, err := m.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, nil, err
	}

	return tx, &modelQueries{
		q: sqlc.New(tx),
	}, nil
}

func (m *ModelDB) Queries() domainmodel.Queries {
	return &modelQueries{
		q: sqlc.New(m.db),
	}
}

func (m *modelQueries) GetModelByID(
	ctx context.Context,
	id shared.ID,
) (*domainmodel.ModelRow, error) {

	modelID, err := id.ToUUID(ctx)
	if err != nil {
		return nil, err
	}

	row, err := m.q.GetModelById(ctx, modelID)
	if err != nil {
		return nil, err
	}

	return &domainmodel.ModelRow{
		ID:        row.ID.String(),
		CompanyID: row.CompanyID.String(),
		Name:      row.Name,
		ImageURL:  row.ImageUrl,
	}, nil
}

func (m *modelQueries) GetModelsByCompanyAndModelNames(
	ctx context.Context,
	modelName string,
	companyName string,
) ([]domainmodel.ModelWithCompanyRow, error) {

	rows, err := m.q.GetModelsByCompanyAndModel(ctx, sqlc.GetModelsByCompanyAndModelParams{
		Column1: companyName,
		Column2: modelName,
	})
	if err != nil {
		return nil, err
	}

	out := make([]domainmodel.ModelWithCompanyRow, 0, len(rows))
	for _, r := range rows {
		out = append(out, domainmodel.ModelWithCompanyRow{
			ID:          r.ID.String(),
			CompanyName: r.CompanyName,
			Name:        r.ModelName,
			ImageURL:    r.ImageUrl,
		})
	}

	return out, nil
}

func (m *modelQueries) CreateModel(
	ctx context.Context,
	p domainmodel.CreateModelParams,
) (*domainmodel.ModelRow, error) {

	row, err := m.q.CreateModel(ctx, sqlc.CreateModelParams{
		ID:        uuid.MustParse(p.ID),
		CompanyID: uuid.MustParse(p.CompanyID),
		Name:      p.Name,
		ImageUrl:  p.ImageURL,
	})
	if err != nil {
		return nil, err
	}

	return &domainmodel.ModelRow{
		ID:        row.ID.String(),
		CompanyID: row.CompanyID.String(),
		Name:      row.Name,
		ImageURL:  row.ImageUrl,
	}, nil
}

func (m *modelQueries) UpdateModel(
	ctx context.Context,
	p domainmodel.UpdateModelParams,
) (*domainmodel.ModelRow, error) {

	row, err := m.q.UpdateModelByID(ctx, sqlc.UpdateModelByIDParams{
		ID:        uuid.MustParse(p.ID),
		CompanyID: sqlnull.UUID(p.CompanyID),
		ModelName: sqlnull.String(p.Name),
		ImageUrl:  sqlnull.String(p.ImageURL),
	})
	if err != nil {
		return nil, err
	}

	return &domainmodel.ModelRow{
		ID:        row.ID.String(),
		CompanyID: row.CompanyID.String(),
		Name:      row.Name,
		ImageURL:  row.ImageUrl,
	}, nil
}

func (m *modelQueries) DeleteModel(
	ctx context.Context,
	id shared.ID,
) error {

	modelID, err := id.ToUUID(ctx)
	if err != nil {
		return err
	}

	return m.q.DeleteModelByID(ctx, modelID)
}
