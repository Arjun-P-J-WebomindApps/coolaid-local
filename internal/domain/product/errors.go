package product

import "errors"

// Product errors
var (
	// -------------------------------------------------
	// Product
	// -------------------------------------------------
	ErrProductNotFound = errors.New("product not found")
	ErrProductExists   = errors.New("product already exists")

	// -------------------------------------------------
	// Foreign Keys
	// -------------------------------------------------
	ErrModelNotFound    = errors.New("model not found")
	ErrCategoryNotFound = errors.New("category not found")
	ErrBrandNotFound    = errors.New("brand not found")
	ErrCompanyNotFound  = errors.New("company not found")

	// -------------------------------------------------
	// Input Validation
	// -------------------------------------------------
	ErrInvalidInput          = errors.New("invalid input")
	ErrInvalidPartNo         = errors.New("invalid part number")
	ErrInvalidCompany        = errors.New("invalid company")
	ErrInvalidModel          = errors.New("invalid model")
	ErrInvalidBrand          = errors.New("invalid brand")
	ErrInvalidCategory       = errors.New("invalid category")
	ErrInvalidVendor         = errors.New("invalid vendor data")
	ErrInvalidOEM            = errors.New("invalid oem data")
	ErrInvalidInventory      = errors.New("invalid inventory data")
	ErrInvalidOfferData      = errors.New("invalid offer data")
	ErrInvalidOfferDateRange = errors.New("invalid offer date range")
	ErrInvalidPricingData    = errors.New("invalid pricing data")

	// -------------------------------------------------
	// Other Domain Errors
	// -------------------------------------------------
	ErrOfferNotFound     = errors.New("offer not found")
	ErrInventoryNotFound = errors.New("inventory not found")
	ErrPricingNotFound   = errors.New("pricing not found")

	ErrInternal = errors.New("internal error")
)
