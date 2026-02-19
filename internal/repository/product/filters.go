package productrepo

import (
	"context"

	domain "github.com/webomindapps-dev/coolaid-backend/internal/domain/product"
	"github.com/webomindapps-dev/coolaid-backend/internal/generated/sqlc"
	"github.com/webomindapps-dev/coolaid-backend/internal/shared/ptr"
	"github.com/webomindapps-dev/coolaid-backend/internal/shared/sqlnull"
)

func (p *productQueries) GetFilteredProducts(
	ctx context.Context,
	params domain.FilterSelectionParams,
) ([]domain.FilteredRow, error) {

	rows, err := p.q.GetFilteredDetails(ctx, sqlc.GetFilteredDetailsParams{
		Companyname:  ptr.String(params.Company),
		Modelname:    ptr.String(params.Model),
		Brandname:    ptr.String(params.Brand),
		Categoryname: ptr.String(params.Category),
		Unicode:      ptr.String(params.Unicode),
	})
	if err != nil {
		return nil, err
	}

	result := make([]domain.FilteredRow, 0, len(rows))

	for _, r := range rows {
		result = append(result, domain.FilteredRow{
			PartNo:        r.PartNo,
			CompanyName:   r.CompanyName,
			CompanyImage:  r.CompanyImage,
			ModelName:     r.ModelName,
			ModelImage:    r.ModelImage,
			BrandName:     r.BrandName,
			BrandImage:    r.BrandImage,
			CategoryName:  r.CategoryName,
			CategoryImage: sqlnull.StringValueOrEmpty(r.CategoryImage),
		})
	}

	return result, nil
}
