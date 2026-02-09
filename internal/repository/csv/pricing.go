package repository

import (
	"context"

	"github.com/webomindapps-dev/coolaid-backend/db"
)

type PricingRow struct {
	PartNo   string
	Company  string
	Model    string
	Brand    string
	Category string

	BasicPrice float64
	Freight    float64
	GST        float64

	AcWorkshop                float64
	AcWorkshopPercent         float64
	AcWorkshopMargin          float64
	MultibrandWorkshop        float64
	MultibrandWorkshopPercent float64
	MultibrandWorkshopMargin  float64
	AutoTrader                float64
	AutoTraderPercent         float64
	AutoTraderMargin          float64
	AcTrader                  float64
	AcTraderPercent           float64
	AcTraderMargin            float64

	OEMMRP      float64
	UnitMeasure string

	// Keep raw IDs, resolve later
	VendorIDs []string
}

type PricingRepository interface {
	List(ctx context.Context) ([]PricingRow, error)
}

type pricingRepo struct {
	q *db.DBContext
}

func NewPricingRepository(q *db.DBContext) PricingRepository {
	return &pricingRepo{q: q}
}

func (r *pricingRepo) List(ctx context.Context) ([]PricingRow, error) {
	items, err := r.q.Queries.GetProductPricingDownloadDetails(ctx)
	if err != nil {
		return nil, err
	}

	out := make([]PricingRow, 0, len(items))
	for _, it := range items {
		out = append(out, PricingRow{
			PartNo:   it.PartNo.String,
			Company:  it.Company,
			Model:    it.Model,
			Brand:    it.Brand,
			Category: it.Category,

			BasicPrice: it.BasicPrice,
			Freight:    it.Freight,
			GST:        it.GST,

			AcWorkshop:                it.AcWorkshop,
			AcWorkshopPercent:         it.AcWorkshopPer,
			AcWorkshopMargin:          it.AcWorkshopMargin,
			MultibrandWorkshop:        it.MultibrandWorkshop,
			MultibrandWorkshopPercent: it.MultibrandWorkshopPer,
			MultibrandWorkshopMargin:  it.MultibrandWorkshopMargin,
			AutoTrader:                it.AutoTrader,
			AutoTraderPercent:         it.AutoTraderPer,
			AutoTraderMargin:          it.AutoTraderMargin,
			AcTrader:                  it.AcTrader,
			AcTraderPercent:           it.AcTraderPer,
			AcTraderMargin:            it.AcTraderMargin,

			OEMMRP:      it.OEMMRP,
			UnitMeasure: it.UnitMeasure.String,

			VendorIDs: it.VendorID,
		})
	}

	return out, nil
}
