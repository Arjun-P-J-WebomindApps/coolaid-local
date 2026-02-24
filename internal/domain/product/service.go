package product

import (
	brand "github.com/webomindapps-dev/coolaid-backend/internal/domain/master/brands"
	"github.com/webomindapps-dev/coolaid-backend/internal/domain/master/category"
	"github.com/webomindapps-dev/coolaid-backend/internal/domain/master/company"
	models "github.com/webomindapps-dev/coolaid-backend/internal/domain/master/model"
	vendor "github.com/webomindapps-dev/coolaid-backend/internal/domain/master/vendors"
	"github.com/webomindapps-dev/coolaid-backend/internal/domain/search"
	"github.com/webomindapps-dev/coolaid-backend/internal/domain/techspec"
)

type Service struct {
	DB              DB
	CompanyService  *company.Service
	ModelService    *models.Service
	BrandService    *brand.Service
	CategoryService *category.Service
	VendorService   *vendor.Service

	//Typesense
	SearchEngine search.SearchEngine
	Indexer      search.Indexer

	//Technical Specs
	TechnicalService *techspec.Service
}

func NewService(
	db DB,
	company *company.Service,
	model *models.Service,
	brand *brand.Service,
	category *category.Service,
	vendor *vendor.Service,

	searchPort search.Port,

	technicalSpecs *techspec.Service,

) *Service {
	return &Service{
		DB:              db,
		CompanyService:  company,
		ModelService:    model,
		BrandService:    brand,
		CategoryService: category,
		VendorService:   vendor,

		SearchEngine: searchPort,
		Indexer:      searchPort,

		TechnicalService: technicalSpecs,
	}

}
