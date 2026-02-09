package download

import (
	"context"
	"fmt"
	"io"

	"github.com/webomindapps-dev/coolaid-backend/internal/csv_schema"
	"github.com/webomindapps-dev/coolaid-backend/internal/csv_util"
	repository "github.com/webomindapps-dev/coolaid-backend/internal/repository/csv"
	"github.com/webomindapps-dev/coolaid-backend/internal/repository/csv/category"
)

type Service struct {
	InventoryRepo repository.InventoryRepository
	PricingRepo   repository.PricingRepository
	BasicRepo     repository.BasicRepository
	CategoryRepo  category.CategoryRepository
}

func (s *Service) ExportToCSV(
	ctx context.Context,
	query string,
	categoryName string,
	w io.Writer,
) error {

	//Get Headers
	ts, err := csv_schema.SchemaFor(query, categoryName)
	if err != nil {
		return err
	}

	//2. Fetch Data
	var data []any

	switch query {

	case "inventory":
		items, err := s.InventoryRepo.List(ctx)
		if err != nil {
			return err
		}
		for _, it := range items {
			data = append(data, it)
		}

	case "pricing":
		items, err := s.PricingRepo.List(ctx)
		if err != nil {
			return err
		}
		for _, it := range items {
			data = append(data, it)
		}

	case "basic":
		items, err := s.BasicRepo.List(ctx)
		if err != nil {
			return err
		}
		for _, it := range items {
			data = append(data, it)
		}

	case "category":
		items, err := s.CategoryRepo.List(ctx)
		if err != nil {
			return err
		}
		for _, it := range items {
			data = append(data, it)
		}

	default:
		return fmt.Errorf("unsupported query %q", query)
	}

	// 3️⃣ Convert typed rows → [][]string using schema headers
	matrix, err := csv_util.BuildCSV(ts.Headers, data)
	if err != nil {
		return err
	}

	// 4️⃣ Stream CSV
	return csv_util.WriteCSV(ts.Headers, matrix, w)
}
