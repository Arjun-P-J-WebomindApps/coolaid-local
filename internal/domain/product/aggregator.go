package product

import (
	"context"

	"github.com/webomindapps-dev/coolaid-backend/internal/domain/search"
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

func (s *Service) CreateProduct(
	ctx context.Context,
	input CreateProductInput,
) (*Product, error) {

	oplog.Info(ctx, "create product started", "partNo=", input.Main.PartNo)

	//  Validate Entire Input First
	if err := s.validateCreateProduct(input); err != nil {
		oplog.Error(ctx, "product validation failed",
			"partNo=", input.Main.PartNo,
			"error=", err,
		)
		return nil, err
	}
	//  Resolve dependencies via services (READ ONLY)
	company, err := s.CompanyService.DB.Queries().GetCompanyByName(ctx, input.Main.CompanyName)
	if err != nil {
		oplog.Error(ctx, "company resolution failed",
			"companyName=", input.Main.CompanyName,
			"error=", err,
		)
		return nil, err
	}

	model, err := s.ModelService.DB.Queries().GetModelByName(ctx, input.Main.ModelName)
	if err != nil {
		oplog.Error(ctx, "model resolution failed",
			"modelName=", input.Main.ModelName,
			"error=", err,
		)
		return nil, err
	}

	brand, err := s.BrandService.DB.Queries().GetBrandByName(ctx, input.Main.BrandName)
	if err != nil {
		oplog.Error(ctx, "brand resolution failed",
			"brandName=", input.Main.BrandName,
			"error=", err,
		)
		return nil, err
	}

	category, err := s.CategoryService.DB.Queries().GetCategoryByName(ctx, input.Main.CategoryName)
	if err != nil {
		oplog.Error(ctx, "category resolution failed",
			"categoryName=", input.Main.CategoryName,
			"error=", err,
		)
		return nil, err
	}

	tx, _, err := s.DB.BeginTx(ctx)
	if err != nil {
		oplog.Error(ctx, "transaction begin failed", "error=", err)
		return nil, err
	}

	committed := false
	defer func() {
		if !committed {
			tx.Rollback()
			oplog.Warn(ctx, "transaction rolled back", "partNo=", input.Main.PartNo)
		}
	}()

	productQ := s.DB.NewQueriesFromTx(tx)
	techQ := s.TechnicalService.DB.NewQueriesFromTx(tx)

	product, err := s.createProduct(ctx, productQ, input.Main.PartNo, company.ID, model.ID, brand.ID, category.ID)
	if err != nil {
		oplog.Error(ctx, "product creation failed",
			"partNo=", input.Main.PartNo,
			"error=", err,
		)
		return nil, err
	}

	//Create OEM Listings
	oemIDs, err := s.createOEMs(ctx, productQ, input.Main.PartNo, input.OemNumbers)
	if err != nil {
		oplog.Error(ctx, "oem creation failed", "partNo=", product.PartNo, "error=", err)
		return nil, err
	}

	//Create Vendor Listings
	vendorIDs, err := s.createVendors(ctx, productQ, input.Main.PartNo, input.Vendors)
	if err != nil {
		oplog.Error(ctx, "vendor creation failed", "partNo=", product.PartNo, "error=", err)
		return nil, err
	}

	//  Create Base Data / Variant
	if _, err := s.createVariant(
		ctx,
		productQ,
		product.PartNo,
		CreateModelVariantInput{
			Type:             input.Main.BaseData.Type,
			Gen:              input.Main.BaseData.Gen,
			FuelTypes:        input.Main.BaseData.FuelTypes,
			HsnCode:          input.Main.BaseData.HsnCode,
			EngineCc:         input.Main.BaseData.EngineCc,
			TransmissionType: input.Main.BaseData.TransmissionType,
			PlatformCodes:    input.Main.BaseData.PlatformCodes,
			Placement:        input.Main.BaseData.Placement,

			Image1Link: input.Main.BaseData.Image1Link,
			Image2Link: input.Main.BaseData.Image2Link,
			Image3Link: input.Main.BaseData.Image3Link,
			Image4Link: input.Main.BaseData.Image4Link,

			Make:           input.Main.BaseData.Make,
			Unicode:        input.Main.BaseData.Unicode,
			YearStart:      input.Main.BaseData.YearStart,
			YearEnd:        input.Main.BaseData.YearEnd,
			Description:    input.Main.BaseData.Description,
			AdditionalInfo: input.Main.BaseData.AdditionalInfo,

			OemNumbers: oemIDs,
			Vendors:    vendorIDs,
		},
	); err != nil {
		oplog.Error(ctx, "variant creation failed", "partNo=", product.PartNo, "error=", err)
		return nil, err
	}

	//  Create Pricing
	if _, err := s.createPricing(ctx, productQ, product.PartNo, input.Pricing); err != nil {
		oplog.Error(ctx, "pricing creation failed", "partNo=", product.PartNo, "error=", err)
		return nil, err
	}

	//  Create Inventory
	if _, err := s.createInventory(ctx, productQ, product.PartNo, input.Inventory); err != nil {
		oplog.Error(ctx, "inventory creation failed", "partNo=", product.PartNo, "error=", err)
		return nil, err
	}

	//  Create Offer
	if _, err := s.createOffer(ctx, productQ, product.PartNo, input.Offer); err != nil {
		oplog.Error(ctx, "offer creation failed", "partNo=", product.PartNo, "error=", err)
		return nil, err
	}

	//  Create Technical Specs (if table exits)
	if err := s.TechnicalService.CreateTechSpec(ctx, techQ, product.PartNo, &input.TechSpec); err != nil {
		oplog.Error(ctx, "techspec creation failed", "partNo=", product.PartNo, "error=", err)
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		oplog.Error(ctx, "transaction commit failed", "partNo=", product.PartNo, "error=", err)
		return nil, err
	}
	committed = true

	oplog.Info(ctx, "product created successfully", "partNo=", product.PartNo)

	if s.SearchService != nil {
		if err := s.SearchService.Index(ctx, search.IndexRequest{
			Collection: "products",
			ID:         product.PartNo,
			Payload: search.ProductSearchDocument{
				ID:       string(product.ID),
				Company:  input.Main.CompanyName,
				Model:    input.Main.ModelName,
				PartNo:   input.Main.PartNo,
				Brand:    input.Main.BrandName,
				Category: input.Main.CategoryName,
			},
		}); err != nil {
			oplog.Error(ctx, "typesense index failed", "partNo=", product.PartNo, "error=", err)
		} else {
			oplog.Info(ctx, "typesense index success", "partNo=", product.PartNo)
		}
	}

	return product, nil
}

func (s *Service) UpdateProduct(
	ctx context.Context,
	input UpdateProductInput,
) (*Product, error) {

	// --------------------------------------------------
	// 1️⃣ Ensure Product Exists
	// --------------------------------------------------

	exists, err := s.DB.Queries().GetProductByPartNo(ctx, input.Main.PartNo)
	if err != nil {
		return nil, ErrInternal
	}
	if exists == nil {
		return nil, ErrProductNotFound
	}

	// --------------------------------------------------
	// 2️⃣ Begin Transaction
	// --------------------------------------------------

	tx, q, err := s.DB.BeginTx(ctx)
	if err != nil {
		return nil, ErrInternal
	}
	defer tx.Rollback()

	committed := false
	defer func() {
		if !committed {
			tx.Rollback()
			oplog.Warn(ctx, "transaction rolled back", "partNo=", input.Main.PartNo)
		}
	}()

	productQ := s.DB.NewQueriesFromTx(tx)
	techQ := s.TechnicalService.DB.NewQueriesFromTx(tx)

	partNo := input.Main.PartNo

	// --------------------------------------------------
	// 3️⃣ Replace Vendors
	// --------------------------------------------------

	if err := q.DeleteVendorsByPartNo(ctx, partNo); err != nil {
		return nil, ErrInternal
	}

	vendorIDs, err := s.createVendors(ctx, productQ, partNo, input.Vendors)
	if err != nil {
		return nil, err
	}

	// --------------------------------------------------
	// 4️⃣ Replace OEM
	// --------------------------------------------------

	if err := q.DeleteOEMByPartNo(ctx, partNo); err != nil {
		return nil, ErrInternal
	}

	oemIDs, err := s.createOEMs(ctx, productQ, partNo, input.OemNumbers)
	if err != nil {
		return nil, err
	}

	// --------------------------------------------------
	// 5️⃣ Update Model Variant
	// --------------------------------------------------

	if _, err := s.updateVariant(
		ctx,
		productQ,
		input.Main.PartNo,
		UpdateModelVariantParams{
			PartNo:           &input.Main.PartNo,
			FuelTypes:        input.Main.BaseData.FuelTypes,
			Gen:              input.Main.BaseData.Gen,
			EngineCc:         input.Main.BaseData.EngineCc,
			TransmissionType: input.Main.BaseData.TransmissionType,
			PlatformCodes:    input.Main.BaseData.PlatformCodes,
			Placement:        input.Main.BaseData.Placement,
			Image1Link:       input.Main.BaseData.Image1Link,
			Image2Link:       input.Main.BaseData.Image2Link,
			Image3Link:       input.Main.BaseData.Image3Link,
			Image4Link:       input.Main.BaseData.Image4Link,
			HsnCode:          input.Main.BaseData.HsnCode,
			Unicode:          input.Main.BaseData.Unicode,
			Description:      input.Main.BaseData.Description,
			Make:             input.Main.BaseData.Make,
			OemIds:           &oemIDs,
			Type:             input.Main.BaseData.Type,
			YearStart:        input.Main.BaseData.YearStart,
			YearEnd:          input.Main.BaseData.YearEnd,
			VendorID:         &vendorIDs,
			AdditionalInfo:   input.Main.BaseData.AdditionalInfo,
		},
	); err != nil {
		return nil, err
	}

	// --------------------------------------------------
	// 6️⃣ Update Pricing
	// --------------------------------------------------

	if _, err := s.updatePricing(ctx, productQ, partNo, input.Pricing); err != nil {
		return nil, err
	}

	// --------------------------------------------------
	// 7️⃣ Update Inventory
	// --------------------------------------------------

	if _, err := s.updateInventory(ctx, productQ, partNo, input.Inventory); err != nil {
		return nil, err
	}

	// --------------------------------------------------
	// 8️⃣ Update Offer
	// --------------------------------------------------

	if _, err := s.updateOffer(ctx, productQ, partNo, input.Offer); err != nil {
		return nil, err
	}

	// --------------------------------------------------
	// 9️⃣ Update TechSpec (delegated)
	// --------------------------------------------------

	if err := s.TechnicalService.UpdateTechSpec(ctx, techQ, partNo, &input.TechSpec); err != nil {
		return nil, err
	}

	// --------------------------------------------------
	// 🔟 Commit
	// --------------------------------------------------

	if err := tx.Commit(); err != nil {
		return nil, ErrInternal
	}

	return &Product{
		PartNo: partNo,
	}, nil
}

func (s *Service) DeleteProduct(
	ctx context.Context,
	partNo string,
) error {

	// --------------------------------------------------
	// 1️⃣ Ensure Exists
	// --------------------------------------------------

	product, err := s.DB.Queries().GetProductByPartNo(ctx, partNo)
	if err != nil {
		return ErrInternal
	}
	if product == nil {
		return ErrProductNotFound
	}

	tx, q, err := s.DB.BeginTx(ctx)
	if err != nil {
		return ErrInternal
	}
	defer tx.Rollback()

	categroy, err := s.CategoryService.GetByID(ctx, product.CategoryID)

	// --------------------------------------------------
	// 2️⃣ Delete Related Data
	// --------------------------------------------------

	if err := q.DeleteVendorsByPartNo(ctx, partNo); err != nil {
		return ErrInternal
	}

	if err := q.DeleteOEMByPartNo(ctx, partNo); err != nil {
		return ErrInternal
	}

	if err := q.DeleteInventory(ctx, partNo); err != nil {
		return ErrInternal
	}

	if err := q.DeleteOffer(ctx, partNo); err != nil {
		return ErrInternal
	}

	if err := q.DeleteModelVariant(ctx, partNo); err != nil {
		return ErrInternal
	}

	if err := q.DeleteProduct(ctx, partNo); err != nil {
		return ErrInternal
	}

	// --------------------------------------------------
	// 3️⃣ Delete TechSpec
	// --------------------------------------------------

	if err := s.TechnicalService.DeleteTechSpec(ctx, q, product.PartNo); err != nil {
		return err
	}

	// --------------------------------------------------
	// 4️⃣ Commit
	// --------------------------------------------------

	if err := tx.Commit(); err != nil {
		return ErrInternal
	}

	// --------------------------------------------------
	// 5️⃣ Async Search Removal (after commit)
	// --------------------------------------------------

	go s.SearchService.Delete(partNo)

	return nil
}
