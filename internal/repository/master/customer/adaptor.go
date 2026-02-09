package customerrepo

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"github.com/webomindapps-dev/coolaid-backend/db"
	"github.com/webomindapps-dev/coolaid-backend/internal/domain/master/customer"
	"github.com/webomindapps-dev/coolaid-backend/internal/domain/master/shared"
	"github.com/webomindapps-dev/coolaid-backend/internal/generated/sqlc"
	"github.com/webomindapps-dev/coolaid-backend/internal/shared/sqlnull"
)

type CustomerDB struct {
	db      *sql.DB
	queries *sqlc.Queries
}

func NewCustomerRepository(dbCtx *db.DBContext) *CustomerDB {
	return &CustomerDB{
		db:      dbCtx.SqlDB,
		queries: dbCtx.Queries,
	}
}

type customerQueries struct {
	q *sqlc.Queries
}

// 🔒 compile-time guarantee
var _ customer.Queries = (*customerQueries)(nil)

/* ------------------------------------------------------------------ */
/* DB                                                                 */
/* ------------------------------------------------------------------ */

func (c *CustomerDB) BeginTx(
	ctx context.Context,
) (*sql.Tx, customer.Queries, error) {

	tx, err := c.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, nil, err
	}

	return tx, &customerQueries{
		q: sqlc.New(tx),
	}, nil
}

func (c *CustomerDB) Queries() customer.Queries {
	return &customerQueries{
		q: sqlc.New(c.db),
	}
}

/* ------------------------------------------------------------------ */
/* Writes                                                             */
/* ------------------------------------------------------------------ */

func (c *customerQueries) CreateCustomer(
	ctx context.Context,
	p customer.CreateCustomerParams,
) (*customer.CustomerRow, error) {

	row, err := c.q.CreateCustomer(ctx, sqlc.CreateCustomerParams{
		ID:                  uuid.MustParse(p.ID),
		CustomerCompanyName: p.CustomerCompanyName,
		ContactPerson:       p.ContactPerson,
		Mobile:              p.Mobile,
		Type:                p.Type,
		CustomerDesignation: sqlnull.String(p.CustomerDesignation),
		Address:             sqlnull.String(p.Address),
		Flat:                sqlnull.String(p.Flat),
		Street:              sqlnull.String(p.Street),
		City:                sqlnull.String(p.City),
		State:               sqlnull.String(p.State),
		Pincode:             sqlnull.String(p.Pincode),
		PaymentMode:         sqlnull.String(p.PaymentMode),
	})
	if err != nil {
		return nil, err
	}

	return mapCustomerRow(row), nil
}

func (c *customerQueries) UpdateCustomer(
	ctx context.Context,
	p customer.UpdateCustomerParams,
) (*customer.CustomerRow, error) {

	row, err := c.q.UpdateCustomerByID(ctx, sqlc.UpdateCustomerByIDParams{
		ID:                  uuid.MustParse(p.ID),
		CustomerCompanyName: sqlnull.String(p.CustomerCompanyName),
		ContactPerson:       sqlnull.String(p.ContactPerson),
		Mobile:              sqlnull.String(p.Mobile),
		Type:                sqlnull.String(p.Type),
		CustomerDesignation: sqlnull.String(p.CustomerDesignation),
		Address:             sqlnull.String(p.Address),
		Flat:                sqlnull.String(p.Flat),
		Street:              sqlnull.String(p.Street),
		City:                sqlnull.String(p.City),
		State:               sqlnull.String(p.State),
		Pincode:             sqlnull.String(p.Pincode),
		PaymentMode:         sqlnull.String(p.PaymentMode),
	})
	if err != nil {
		return nil, err
	}

	return mapCustomerRow(row), nil
}

func (c *customerQueries) DeleteCustomer(
	ctx context.Context,
	id shared.ID,
) error {

	u, err := id.ToUUID(ctx)
	if err != nil {
		return err
	}

	_, errCustomer := c.q.DeleteCustomerByID(ctx, u)

	return errCustomer
}

/* ------------------------------------------------------------------ */
/* Reads                                                              */
/* ------------------------------------------------------------------ */

func (c *customerQueries) GetCustomerByID(
	ctx context.Context,
	id shared.ID,
) (*customer.CustomerRow, error) {

	u, err := id.ToUUID(ctx)
	if err != nil {
		return nil, err
	}

	row, err := c.q.GetCustomerFromID(ctx, u)
	if err != nil {
		return nil, err
	}

	return mapCustomerRow(row), nil
}

func (c *customerQueries) GetCustomerByUniqueKey(
	ctx context.Context,
	company string,
	contactPerson string,
	mobile string,
	customerType string,
) (*customer.CustomerRow, error) {

	row, err := c.q.GetCustomerByMobileCompanyContactPersonAndCustomerType(ctx, sqlc.GetCustomerByMobileCompanyContactPersonAndCustomerTypeParams{
		Customercompanyname: company,
		Mobile:              mobile,
		Contactperson:       contactPerson,
		Customertype:        customerType,
	})
	if err != nil {
		return nil, err
	}

	return mapCustomerRow(row), nil
}

func (c *customerQueries) SearchCustomers(
	ctx context.Context,
	company string,
	contactPerson string,
	mobile string,
) ([]customer.CustomerRow, error) {

	rows, err := c.q.SearchCustomersByAttributes(ctx, sqlc.SearchCustomersByAttributesParams{
		CustomerCompanyName: company,
		ContactPerson:       contactPerson,
		Mobile:              mobile,
	})
	if err != nil {
		return nil, err
	}

	out := make([]customer.CustomerRow, 0, len(rows))
	for _, r := range rows {
		out = append(out, *mapCustomerRow(r))
	}

	return out, nil
}

/* ------------------------------------------------------------------ */
/* Mapper                                                             */
/* ------------------------------------------------------------------ */
func mapCustomerRow(r sqlc.Customer) *customer.CustomerRow {
	return &customer.CustomerRow{
		ID:                  r.ID.String(),
		CustomerCompanyName: r.CustomerCompanyName,
		ContactPerson:       r.ContactPerson,
		Mobile:              r.Mobile,
		Type:                r.Type,
		CustomerDesignation: sqlnull.StringValueOrEmpty(r.CustomerDesignation),
		Address:             sqlnull.StringPtr(r.Address),
		Flat:                sqlnull.StringPtr(r.Flat),
		Street:              sqlnull.StringPtr(r.Street),
		City:                sqlnull.StringValueOrEmpty(r.City),
		State:               sqlnull.StringValueOrEmpty(r.State),
		Pincode:             sqlnull.StringPtr(r.Pincode),
		PaymentMode:         sqlnull.StringValueOrEmpty(r.PaymentMode),
	}
}
