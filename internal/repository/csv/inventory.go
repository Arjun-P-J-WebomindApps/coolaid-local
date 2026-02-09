package repository

import (
	"context"

	"github.com/webomindapps-dev/coolaid-backend/db"
)

type InventoryRow struct {
	PartNo        string
	Company       string
	Model         string
	Brand         string
	Category      string
	QtyInStock    int
	MaxOrderLevel int
	MinOrderLevel int
	Location      string
	IsFlash       bool
}

type InventoryRepository interface {
	List(ctx context.Context) ([]InventoryRow, error)
}

type inventoryRepo struct {
	q *db.DBContext
}

func NewInventoryRepository(q *db.DBContext) InventoryRepository {
	return &inventoryRepo{q: q}
}

func (r *inventoryRepo) List(ctx context.Context) ([]InventoryRow, error) {
	items, err := r.q.Queries.GetInventoryDetailsForBulkDownload(ctx)
	if err != nil {
		return nil, err
	}

	out := make([]InventoryRow, 0, len(items))
	for _, it := range items {
		out = append(out, InventoryRow{
			PartNo:        it.PartNo,
			Company:       it.Company,
			Model:         it.Model,
			Brand:         it.Brand,
			Category:      it.Category,
			QtyInStock:    int(it.QtyInStock),
			MaxOrderLevel: int(it.MaximumOrderLevel),
			MinOrderLevel: int(it.MinimumOrderLevel),
			Location:      it.Location,
			IsFlash:       it.IsFlash.Valid && it.IsFlash.Bool,
		})
	}

	return out, nil
}
