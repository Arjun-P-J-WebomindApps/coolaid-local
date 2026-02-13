package service

import (
	"github.com/webomindapps-dev/coolaid-backend/db"
	"github.com/webomindapps-dev/coolaid-backend/internal/domain/auth"
	brand "github.com/webomindapps-dev/coolaid-backend/internal/domain/master/brands"
	"github.com/webomindapps-dev/coolaid-backend/internal/domain/master/category"
	"github.com/webomindapps-dev/coolaid-backend/internal/domain/master/company"
	"github.com/webomindapps-dev/coolaid-backend/internal/domain/master/customer"
	models "github.com/webomindapps-dev/coolaid-backend/internal/domain/master/model"
	vendor "github.com/webomindapps-dev/coolaid-backend/internal/domain/master/vendors"
	"github.com/webomindapps-dev/coolaid-backend/internal/domain/techspec"
	repository "github.com/webomindapps-dev/coolaid-backend/internal/repository/auth"
	brand_repo "github.com/webomindapps-dev/coolaid-backend/internal/repository/master/brand"
	categoryrepo "github.com/webomindapps-dev/coolaid-backend/internal/repository/master/category"
	companyrepo "github.com/webomindapps-dev/coolaid-backend/internal/repository/master/company"
	customerrepo "github.com/webomindapps-dev/coolaid-backend/internal/repository/master/customer"
	modelrepo "github.com/webomindapps-dev/coolaid-backend/internal/repository/master/model"
	vendorrepo "github.com/webomindapps-dev/coolaid-backend/internal/repository/master/vendor"
	techspecrepo "github.com/webomindapps-dev/coolaid-backend/internal/repository/techspec"
	"github.com/webomindapps-dev/coolaid-backend/internal/service/crypto"
	"github.com/webomindapps-dev/coolaid-backend/internal/service/mailer"
	"github.com/webomindapps-dev/coolaid-backend/typesense"
)

type Container struct {
	// Infra
	DB        *db.DBContext
	Typesense *typesense.TypesenseContext

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
}

func NewContainer(
	dbCtx *db.DBContext,
	ts *typesense.TypesenseContext,
) *Container {

	// 1️⃣ Build infra adapters
	cryptoSvc := crypto.NewService() // implements auth.Crypto
	mailerSvc := mailer.NewService() // implements auth.Mailer

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

	return &Container{
		DB:        dbCtx,
		Typesense: ts,

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
	}
}
