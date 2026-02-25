package service

import (
	"context"

	"github.com/webomindapps-dev/coolaid-backend/config"
	"github.com/webomindapps-dev/coolaid-backend/db"
	"github.com/webomindapps-dev/coolaid-backend/internal/domain/auth"
	brand "github.com/webomindapps-dev/coolaid-backend/internal/domain/master/brands"
	"github.com/webomindapps-dev/coolaid-backend/internal/domain/master/category"
	"github.com/webomindapps-dev/coolaid-backend/internal/domain/master/company"
	"github.com/webomindapps-dev/coolaid-backend/internal/domain/master/customer"
	models "github.com/webomindapps-dev/coolaid-backend/internal/domain/master/model"
	vendor "github.com/webomindapps-dev/coolaid-backend/internal/domain/master/vendors"
	"github.com/webomindapps-dev/coolaid-backend/internal/domain/product"
	"github.com/webomindapps-dev/coolaid-backend/internal/domain/search"
	"github.com/webomindapps-dev/coolaid-backend/internal/domain/techspec"
	repository "github.com/webomindapps-dev/coolaid-backend/internal/repository/auth"
	brand_repo "github.com/webomindapps-dev/coolaid-backend/internal/repository/master/brand"
	categoryrepo "github.com/webomindapps-dev/coolaid-backend/internal/repository/master/category"
	companyrepo "github.com/webomindapps-dev/coolaid-backend/internal/repository/master/company"
	customerrepo "github.com/webomindapps-dev/coolaid-backend/internal/repository/master/customer"
	modelrepo "github.com/webomindapps-dev/coolaid-backend/internal/repository/master/model"
	vendorrepo "github.com/webomindapps-dev/coolaid-backend/internal/repository/master/vendor"
	productrepo "github.com/webomindapps-dev/coolaid-backend/internal/repository/product"
	searchrepo "github.com/webomindapps-dev/coolaid-backend/internal/repository/search"
	techspecrepo "github.com/webomindapps-dev/coolaid-backend/internal/repository/techspec"
	"github.com/webomindapps-dev/coolaid-backend/internal/service/crypto"
	"github.com/webomindapps-dev/coolaid-backend/internal/service/mailer"
	typesense "github.com/webomindapps-dev/coolaid-backend/internal/service/typesense"
	"github.com/webomindapps-dev/coolaid-backend/oplog"
)

type Container struct {
	// Infra
	DB  *db.DBContext
	ctx context.Context

	// Domain services
	Auth *auth.Service

	// Master services
	Company  *company.Service
	Model    *models.Service
	Category *category.Service
	Brand    *brand.Service
	Customer *customer.Service
	Vendor   *vendor.Service

	//TechSpecs
	TechSpec *techspec.Service

	//Product
	Product *product.Service

	//External Services
	Search *search.Service
	TS     *typesense.Context
}

func NewContainer(
	dbCtx *db.DBContext,
	ctx context.Context,
) *Container {

	// 1️⃣ Build infra adapters
	cryptoSvc := crypto.NewService() // implements auth.Crypto
	mailerSvc := mailer.NewService() // implements auth.Mailer
	tsCtx, err := typesense.Connect(ctx, config.SearchEngine.TypesenseAPIEndpoint, config.SearchEngine.TypesenseAPIKey)

	if err != nil {
		oplog.Error(ctx, "Failed to connect to typesense")
	}

	//Auth
	authRepo := repository.NewAuthRepository(dbCtx) // implements auth.DB + auth.Queries

	//Master
	companyRepo := companyrepo.NewCompanyRepository(dbCtx)
	modelRepo := modelrepo.NewModelRepository(dbCtx)
	categoryRepo := categoryrepo.NewCategoryRepository(dbCtx)
	brandRepo := brand_repo.NewBrandRepository(dbCtx)
	customerRepo := customerrepo.NewCustomerRepository(dbCtx)
	vendorRepo := vendorrepo.NewVendorRepository(dbCtx)

	//TechRepo
	techRepo := techspecrepo.NewTechSpecRepository(dbCtx)

	//SearchRepo
	searchRepo := searchrepo.NewSearchRepository(dbCtx)

	//ProductRepo
	productRepo := productrepo.NewProductRepository(dbCtx)

	// Auth Service
	authSvc := auth.NewService(
		authRepo,
		cryptoSvc,
		mailerSvc,
	)

	//Master Service
	companySvc := company.NewService(companyRepo)
	modelSvc := models.NewService(modelRepo, companySvc)
	categorySvc := category.NewService(categoryRepo)
	brandSvc := brand.NewService(brandRepo)
	customerSvc := customer.NewService(customerRepo)
	vendorSvc := vendor.NewService(vendorRepo)

	//TechRepo
	techSvc := techspec.NewService(techRepo)

	//Search
	tsSvc := typesense.NewService(tsCtx.Client)
	search := search.NewService(tsSvc, searchRepo)

	//Product
	productSvc := product.NewService(productRepo, companySvc, modelSvc, brandSvc, categorySvc, vendorSvc, tsSvc, techSvc)

	return &Container{
		DB: dbCtx,

		//Auth
		Auth: authSvc,

		//Master
		Company:  companySvc,
		Model:    modelSvc,
		Category: categorySvc,
		Brand:    brandSvc,
		Customer: customerSvc,
		Vendor:   vendorSvc,

		//TechSpec
		TechSpec: techSvc,

		//Product
		Product: productSvc,

		Search: search,
		TS:     tsCtx,
	}
}
