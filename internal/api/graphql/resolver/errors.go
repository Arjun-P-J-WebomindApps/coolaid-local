package resolvers

import (
	"github.com/vektah/gqlparser/v2/gqlerror"
	"github.com/webomindapps-dev/coolaid-backend/internal/domain/auth"
	"github.com/webomindapps-dev/coolaid-backend/internal/domain/master/company"
	models "github.com/webomindapps-dev/coolaid-backend/internal/domain/master/model"
	"github.com/webomindapps-dev/coolaid-backend/internal/domain/product"
)

func mapAuthError(err error) error {
	switch err {
	case auth.ErrInvalidCredentials:
		return gqlerror.Errorf("invalid credentials")
	case auth.ErrInvalidOTP:
		return gqlerror.Errorf("invalid otp")
	case auth.ErrExpiredOTP:
		return gqlerror.Errorf("otp expired")
	case auth.ErrUnauthorized:
		return gqlerror.Errorf("unauthorized")
	default:
		return gqlerror.Errorf("internal error")
	}
}

func mapCompanyError(err error) error {
	switch err {
	case company.ErrCompanyNotFound:
		return gqlerror.Errorf("company not found")

	case company.ErrCompanyExists:
		return gqlerror.Errorf("company already exists")

	case company.ErrInvalidInput:
		return gqlerror.Errorf("invalid input")

	default:
		return gqlerror.Errorf("internal error")
	}
}

func mapModelError(err error) error {
	switch err {

	case models.ErrModelNotFound:
		return gqlerror.Errorf("model not found")

	case models.ErrModelExists:
		return gqlerror.Errorf("model already exists")

	case models.ErrCompanyNotFound:
		return gqlerror.Errorf("company not found")

	case models.ErrInvalidInput:
		return gqlerror.Errorf("invalid input")

	default:
		return gqlerror.Errorf("internal error")
	}
}

func mapProductError(err error) error {
	switch err {

	// Product
	case product.ErrProductNotFound:
		return gqlerror.Errorf("product not found")

	case product.ErrProductExists:
		return gqlerror.Errorf("product already exists")

	// Foreign Keys
	case product.ErrModelNotFound:
		return gqlerror.Errorf("model not found")

	case product.ErrCategoryNotFound:
		return gqlerror.Errorf("category not found")

	case product.ErrBrandNotFound:
		return gqlerror.Errorf("brand not found")

	case product.ErrCompanyNotFound:
		return gqlerror.Errorf("company not found")

	// Validation
	case product.ErrInvalidInput:
		return gqlerror.Errorf("invalid input")

	case product.ErrInvalidPartNo:
		return gqlerror.Errorf("invalid part number")

	case product.ErrInvalidCompany:
		return gqlerror.Errorf("invalid company")

	case product.ErrInvalidModel:
		return gqlerror.Errorf("invalid model")

	case product.ErrInvalidBrand:
		return gqlerror.Errorf("invalid brand")

	case product.ErrInvalidCategory:
		return gqlerror.Errorf("invalid category")

	case product.ErrInvalidVendor:
		return gqlerror.Errorf("invalid vendor data")

	case product.ErrInvalidOEM:
		return gqlerror.Errorf("invalid oem data")

	case product.ErrInvalidInventory:
		return gqlerror.Errorf("invalid inventory data")

	case product.ErrInvalidOfferData:
		return gqlerror.Errorf("invalid offer data")

	case product.ErrInvalidOfferDateRange:
		return gqlerror.Errorf("invalid offer date range")

	case product.ErrInvalidPricingData:
		return gqlerror.Errorf("invalid pricing data")

	// Other
	case product.ErrOfferNotFound:
		return gqlerror.Errorf("offer not found")

	case product.ErrInventoryNotFound:
		return gqlerror.Errorf("inventory not found")

	case product.ErrPricingNotFound:
		return gqlerror.Errorf("pricing not found")

	case product.ErrInternal:
		return gqlerror.Errorf("internal error")

	default:
		return gqlerror.Errorf("internal error")
	}
}
