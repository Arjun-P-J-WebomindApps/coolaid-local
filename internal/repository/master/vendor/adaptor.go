package vendorrepo

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"github.com/webomindapps-dev/coolaid-backend/db"
	"github.com/webomindapps-dev/coolaid-backend/internal/domain/master/shared"
	vendor "github.com/webomindapps-dev/coolaid-backend/internal/domain/master/vendors"
	"github.com/webomindapps-dev/coolaid-backend/internal/generated/sqlc"
	"github.com/webomindapps-dev/coolaid-backend/internal/shared/sqlnull"
)

/* ------------------------------------------------------------------ */
/* Repository                                                         */
/* ------------------------------------------------------------------ */

type VendorDB struct {
	db *sql.DB
}

func NewVendorRepository(dbCtx *db.DBContext) *VendorDB {
	return &VendorDB{
		db: dbCtx.SqlDB,
	}
}

type vendorQueries struct {
	q *sqlc.Queries
}

// compile-time safety
var _ vendor.Queries = (*vendorQueries)(nil)

/* ------------------------------------------------------------------ */
/* DB                                                                 */
/* ------------------------------------------------------------------ */

func (v *VendorDB) BeginTx(
	ctx context.Context,
) (*sql.Tx, vendor.Queries, error) {

	tx, err := v.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, nil, err
	}

	return tx, &vendorQueries{
		q: sqlc.New(tx),
	}, nil
}

func (v *VendorDB) Queries() vendor.Queries {
	return &vendorQueries{
		q: sqlc.New(v.db),
	}
}

/* ------------------------------------------------------------------ */
/* Writes                                                             */
/* ------------------------------------------------------------------ */

func (v *vendorQueries) CreateVendor(
	ctx context.Context,
	p vendor.CreateVendorParams,
) (*vendor.VendorRow, error) {

	row, err := v.q.CreateVendor(ctx, sqlc.CreateVendorParams{
		ID:          uuid.MustParse(p.ID),
		CompanyName: p.CompanyName,
	})
	if err != nil {
		return nil, err
	}

	return &vendor.VendorRow{
		ID:          row.ID.String(),
		CompanyName: row.CompanyName,
	}, nil
}

func (v *vendorQueries) DeleteVendor(
	ctx context.Context,
	id shared.ID,
) error {

	u, err := id.ToUUID(ctx)
	if err != nil {
		return err
	}

	return v.q.DeleteVendorByID(ctx, u)
}

func (v *vendorQueries) DeleteVendorContacts(
	ctx context.Context,
	vendorID shared.ID,
) error {

	u, err := vendorID.ToUUID(ctx)
	if err != nil {
		return err
	}

	return v.q.DeleteVendorContactsByID(ctx, u)
}

func (v *vendorQueries) CreateVendorContact(
	ctx context.Context,
	p vendor.CreateVendorContactParams,
) error {

	_, err := v.q.CreateVendorContacts(ctx, sqlc.CreateVendorContactsParams{
		ID:            uuid.MustParse(p.ID),
		VendorID:      uuid.MustParse(p.VendorID),
		ContactPerson: p.ContactPerson,
		MobileNo:      p.MobileNumber,
		EmailID:       p.EmailID,
	})

	return err
}

/* ------------------------------------------------------------------ */
/* Reads                                                              */
/* ------------------------------------------------------------------ */

func (v *vendorQueries) GetVendorByCompanyName(
	ctx context.Context,
	companyName string,
) (*vendor.VendorRow, error) {

	row, err := v.q.GetVendorByCompanyName(ctx, companyName)
	if err != nil {
		return nil, err
	}

	return &vendor.VendorRow{
		ID:          row.ID.String(),
		CompanyName: row.CompanyName,
	}, nil
}

func (v *vendorQueries) GetVendorsWithContacts(
	ctx context.Context,
	companyName string,
) ([]vendor.VendorWithContactRow, error) {

	rows, err := v.q.GetVendorsByCompanyName(ctx, companyName)
	if err != nil {
		return nil, err
	}

	out := make([]vendor.VendorWithContactRow, 0, len(rows))
	for _, r := range rows {
		out = append(out, vendor.VendorWithContactRow{
			VendorID:      r.ID.UUID.String(),
			CompanyName:   sqlnull.StringValueOrEmpty(r.CompanyName),
			ContactPerson: &r.ContactPerson,
			MobileNumber:  &r.MobileNo,
			EmailID:       &r.EmailID,
		})
	}

	return out, nil
}
