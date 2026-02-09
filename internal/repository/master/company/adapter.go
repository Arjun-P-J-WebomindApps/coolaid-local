package companyrepo

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/webomindapps-dev/coolaid-backend/db"
	"github.com/webomindapps-dev/coolaid-backend/internal/domain/master/company"
	"github.com/webomindapps-dev/coolaid-backend/internal/domain/master/shared"
	"github.com/webomindapps-dev/coolaid-backend/internal/generated/sqlc"
	"github.com/webomindapps-dev/coolaid-backend/internal/shared/sqlnull"
)

type CompanyDB struct {
	db      *sql.DB
	queries *sqlc.Queries
}

func NewCompanyRepository(dbCtx *db.DBContext) *CompanyDB {
	return &CompanyDB{
		db:      dbCtx.SqlDB,
		queries: dbCtx.Queries,
	}
}

type companyQueries struct {
	q *sqlc.Queries
}

// compile-time guarantee (VERY IMPORTANT)
var _ company.Queries = (*companyQueries)(nil)

func (c *CompanyDB) BeginTx(
	ctx context.Context,
) (*sql.Tx, company.Queries, error) {

	tx, err := c.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, nil, err
	}

	return tx, &companyQueries{
		q: sqlc.New(tx),
	}, nil
}

func (c *CompanyDB) Queries() company.Queries {
	return &companyQueries{
		q: sqlc.New(c.db),
	}
}

func (c *companyQueries) CreateCompany(
	ctx context.Context,
	p company.CreateCompanyParams,
) (*company.CompanyRow, error) {

	row, err := c.q.CreateCompanies(ctx, sqlc.CreateCompaniesParams{
		ID:       uuid.New(),
		Name:     p.Name,
		Status:   p.Status,
		ImageUrl: p.ImageURL,
	})
	if err != nil {
		return nil, err
	}

	return &company.CompanyRow{
		ID:       row.ID.String(),
		Name:     row.Name,
		Status:   row.Status,
		ImageURL: row.ImageUrl,
	}, nil
}

func (c *companyQueries) GetCompanyByID(
	ctx context.Context,
	id shared.ID,
) (*company.CompanyRow, error) {

	uuidCompany, err := id.ToUUID(ctx)

	if err != nil {
		return nil, err
	}

	row, err := c.q.GetCompanyById(ctx, uuidCompany)
	if err != nil {
		return nil, err
	}

	return &company.CompanyRow{
		ID:       row.ID.String(),
		Name:     row.Name,
		Status:   row.Status,
		ImageURL: row.ImageUrl,
	}, nil
}

func (c *companyQueries) GetCompanyByName(
	ctx context.Context,
	name string,
) (*company.CompanyRow, error) {

	row, err := c.q.GetCompanyByName(ctx, name)
	if err != nil {
		return nil, err
	}

	return &company.CompanyRow{
		ID:       row.ID.String(),
		Name:     row.Name,
		Status:   row.Status,
		ImageURL: row.ImageUrl,
	}, nil
}

func (c *companyQueries) GetCompaniesByName(
	ctx context.Context,
	name string,
) ([]company.CompanyRow, error) {

	rows, err := c.q.GetCompaniesByName(ctx, name)
	if err != nil {
		return nil, err
	}

	out := make([]company.CompanyRow, 0, len(rows))
	for _, r := range rows {
		out = append(out, company.CompanyRow{
			ID:       r.ID.String(),
			Name:     r.Name,
			Status:   r.Status,
			ImageURL: r.ImageUrl,
		})
	}

	return out, nil
}

func (c *companyQueries) UpdateCompany(
	ctx context.Context,
	p company.UpdateCompanyParams,
) (*company.CompanyRow, error) {

	row, err := c.q.UpdateCompanyByID(ctx, sqlc.UpdateCompanyByIDParams{
		CompanyName:   sqlnull.String(p.Name),
		CompanyStatus: sqlnull.Bool(p.Status),
		ImageUrl:      sqlnull.String(p.ImageURL),
	})
	if err != nil {
		return nil, err
	}

	return &company.CompanyRow{
		ID:       row.ID.String(),
		Name:     row.Name,
		Status:   row.Status,
		ImageURL: row.ImageUrl,
	}, nil
}

func (c *companyQueries) DeleteCompany(
	ctx context.Context,
	id shared.ID,
) error {
	return nil
}
