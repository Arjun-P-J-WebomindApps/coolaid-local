package product_mapper

import (
	"github.com/webomindapps-dev/coolaid-backend/internal/domain/product"
	"github.com/webomindapps-dev/coolaid-backend/internal/generated/graphql/model"
)

func ToFilterOrigin(item product.FilterItem) *model.FilterOrigin {
	return &model.FilterOrigin{
		CompanyName:   item.CompanyName,
		CompanyImage:  item.CompanyImage,
		ModelName:     item.ModelName,
		ModelImage:    item.ModelImage,
		CategoryName:  item.CategoryName,
		CategoryImage: item.CategoryImage,
		BrandName:     item.BrandName,
		BrandImage:    item.BrandImage,
		PartNo:        item.PartNo,
	}
}
