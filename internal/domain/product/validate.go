package product

import (
	"context"
	"errors"

	"github.com/webomindapps-dev/coolaid-backend/oplog"
)

func (s *Service) validateCreateProduct(input CreateProductInput) error {

	// -------------------------------------------------
	// 1️⃣ Main Required Fields
	// -------------------------------------------------

	if input.Main.PartNo == "" {
		return ErrInvalidPartNo
	}

	if input.Main.CompanyName == "" {
		return ErrInvalidCompany
	}

	if input.Main.ModelName == "" {
		return ErrInvalidModel
	}

	if input.Main.BrandName == "" {
		return ErrInvalidBrand
	}

	if input.Main.CategoryName == "" {
		return ErrInvalidCategory
	}

	// -------------------------------------------------
	// 2️⃣ Offer Validation
	// -------------------------------------------------

	if input.Offer.IsOfferActive {

		if input.Offer.StartDate == "" || input.Offer.EndDate == "" {
			return ErrInvalidOfferData
		}

		start, err := parseDate(input.Offer.StartDate)
		if err != nil {
			return ErrInvalidOfferData
		}

		end, err := parseDate(input.Offer.EndDate)
		if err != nil {
			return ErrInvalidOfferData
		}

		if end.Before(start) {
			return ErrInvalidOfferDateRange
		}
	}

	// -------------------------------------------------
	// 3️⃣ Pricing Validation
	// -------------------------------------------------

	if input.Pricing.BasicPrice < 0 ||
		input.Pricing.Gst < 0 ||
		input.Pricing.MinimumPurchaseQuantity < 0 ||
		input.Pricing.OemMrp < 0 {
		return ErrInvalidPricingData
	}

	// -------------------------------------------------
	// 4️⃣ Inventory Validation
	// -------------------------------------------------

	if input.Inventory.QtyInStock < 0 ||
		input.Inventory.MinimumOrderLevel < 0 ||
		input.Inventory.MaximumOrderLevel < input.Inventory.MinimumOrderLevel {
		return ErrInvalidInventory
	}

	// -------------------------------------------------
	// 5️⃣ Vendor Validation
	// -------------------------------------------------

	for _, vendor := range input.Vendors {

		if vendor.VendorName == "" {
			return ErrInvalidVendor
		}

		if vendor.VendorPartNo == "" {
			return ErrInvalidVendor
		}

		if vendor.VendorPrice < 0 {
			return ErrInvalidVendor
		}
	}

	// -------------------------------------------------
	// 6️⃣ OEM Validation
	// -------------------------------------------------

	for _, oem := range input.OemNumbers {

		if oem.OemNumber == "" {
			return ErrInvalidOEM
		}

		if oem.Price < 0 {
			return ErrInvalidOEM
		}
	}

	return nil
}

// validateUpdateProduct validates update input safely.
// - Only acts on fields explicitly set (pointer != nil)
// - Always validates Vendors/OEMs contents (slices never nil)
// - Blocks immutable changes (partNo + core FKs)
// - Direct access to struct fields, safe deref on pointers
func (s *Service) validateUpdateProduct(
	ctx context.Context,
	input UpdateProductInput,
	current *ProductRow,
) error {

	partNo := input.Main.PartNo

	// PartNo cannot change
	if partNo != current.PartNo {
		oplog.Warn(ctx, "attempt to change partNo", "current", current.PartNo, "attempted", partNo)
		return errors.New("cannot change part number")
	}

	// Company — only if explicitly set
	if input.Main.CompanyName != nil {
		company, err := s.CompanyService.DB.Queries().GetCompanyByName(ctx, *input.Main.CompanyName)
		if err != nil {
			oplog.Error(ctx, "company resolution failed", "name", *input.Main.CompanyName, "error", err)
			return err
		}
		if company.ID != current.CompanyID {
			return errors.New("cannot change company")
		}
	}

	// Model — only if set
	if input.Main.ModelName != nil {
		model, err := s.ModelService.DB.Queries().GetModelByName(ctx, *input.Main.ModelName)
		if err != nil {
			oplog.Error(ctx, "model resolution failed", "name", *input.Main.ModelName, "error", err)
			return err
		}
		if model.ID != current.ModelID {
			return errors.New("cannot change model")
		}
	}

	// Brand — only if set
	if input.Main.BrandName != nil {
		brand, err := s.BrandService.DB.Queries().GetBrandByName(ctx, *input.Main.BrandName)
		if err != nil {
			oplog.Error(ctx, "brand resolution failed", "name", *input.Main.BrandName, "error", err)
			return err
		}
		if brand.ID != current.BrandID {
			return errors.New("cannot change brand")
		}
	}

	// Category — only if set
	if input.Main.CategoryName != nil {
		category, err := s.CategoryService.DB.Queries().GetCategoryByName(ctx, *input.Main.CategoryName)
		if err != nil {
			oplog.Error(ctx, "category resolution failed", "name", *input.Main.CategoryName, "error", err)
			return err
		}
		if category.ID != current.CategoryID {
			return errors.New("cannot change category")
		}
	}

	// Pricing — only if fields are set (pointers inside)
	p := input.Pricing
	if p.BasicPrice != nil && *p.BasicPrice < 0 ||
		p.Gst != nil && *p.Gst < 0 ||
		p.OemMrp != nil && *p.OemMrp < 0 ||
		p.MinimumPurchaseQuantity != nil && *p.MinimumPurchaseQuantity < 0 {
		oplog.Warn(ctx, "invalid pricing values in update", "partNo", partNo)
		return ErrInvalidPricingData
	}

	// Inventory — same pattern
	i := input.Inventory
	if i.QtyInStock != nil && *i.QtyInStock < 0 ||
		i.MinimumOrderLevel != nil && *i.MinimumOrderLevel < 0 ||
		(i.MaximumOrderLevel != nil && *i.MaximumOrderLevel > 0 &&
			i.MinimumOrderLevel != nil && *i.MaximumOrderLevel < *i.MinimumOrderLevel) {
		oplog.Warn(ctx, "invalid inventory values", "partNo", partNo)
		return ErrInvalidInventory
	}

	// Offer — only if IsOfferActive is set and true
	o := input.Offer
	if o.IsOfferActive != nil && *o.IsOfferActive {
		// Require dates only when activating
		if o.StartDate == nil || *o.StartDate == "" ||
			o.EndDate == nil || *o.EndDate == "" {
			oplog.Warn(ctx, "offer active but missing dates", "partNo", partNo)
			return ErrInvalidOfferData
		}

		start, err := parseDate(*o.StartDate)
		if err != nil {
			oplog.Warn(ctx, "invalid offer start date", "partNo", partNo, "error", err)
			return ErrInvalidOfferData
		}

		end, err := parseDate(*o.EndDate)
		if err != nil {
			oplog.Warn(ctx, "invalid offer end date", "partNo", partNo, "error", err)
			return ErrInvalidOfferData
		}

		if end.Before(start) {
			oplog.Warn(ctx, "offer date range invalid (end before start)", "partNo", partNo)
			return ErrInvalidOfferDateRange
		}
	}

	// Vendors — always validate contents
	for _, v := range input.Vendors {
		if v.VendorName == "" || v.VendorPartNo == "" || v.VendorPrice < 0 {
			oplog.Warn(ctx, "invalid vendor",
				"partNo", partNo,
				"name", v.VendorName,
				"part_no", v.VendorPartNo,
				"price", v.VendorPrice)
			return ErrInvalidVendor
		}
	}

	// OEMs — always validate contents
	for _, o := range input.OemNumbers {
		if o.OemNumber == "" || o.Price < 0 {
			oplog.Warn(ctx, "invalid OEM",
				"partNo", partNo,
				"oem_number", o.OemNumber,
				"price", o.Price)
			return ErrInvalidOEM
		}
	}

	return nil
}
