package repository

import (
	"context"

	"github.com/webomindapps-dev/coolaid-backend/db"
)

type BasicRow struct {
	PartNo   string
	Company  string
	Model    string
	Brand    string
	Category string

	Type       string
	YearStart  int
	YearEnd    int
	Generation string

	FuelTypes     []string
	Transmission  []string
	EngineCC      float64
	PlatformCodes []string

	HSNCode   string
	Make      string
	Placement string

	AdditionalInfo string
	Unicode        []string

	VendorIDs []string
	OEMIDs    []string
}

type BasicRepository interface {
	List(ctx context.Context) ([]BasicRow, error)
}

type basicRepo struct {
	q *db.DBContext
}

func NewBasicRepository(q *db.DBContext) BasicRepository {
	return &basicRepo{q: q}
}

func (r *basicRepo) List(ctx context.Context) ([]BasicRow, error) {
	items, err := r.q.Queries.GetModelVariantDownloadDetails(ctx)
	if err != nil {
		return nil, err
	}

	out := make([]BasicRow, 0, len(items))
	for _, it := range items {
		out = append(out, BasicRow{
			PartNo:   it.PartNo,
			Company:  it.Company,
			Model:    it.Model,
			Brand:    it.Brand,
			Category: it.Category,

			Type:       it.Type,
			YearStart:  int(it.YearStart.Int32),
			YearEnd:    int(it.YearEnd.Int32),
			Generation: it.Generation.String,

			FuelTypes:     it.FuelTypes,
			Transmission:  it.Transmission,
			EngineCC:      it.EngineCC.Float64,
			PlatformCodes: it.PlatformCodes,

			HSNCode:        it.HSNCode.String,
			Make:           it.Make,
			Placement:      it.Placement,
			AdditionalInfo: it.AdditionalInfo.String,
			Unicode:        it.Unicode,

			VendorIDs: it.VendorIDs,
			OEMIDs:    it.OEMIDs,
		})
	}

	return out, nil
}
