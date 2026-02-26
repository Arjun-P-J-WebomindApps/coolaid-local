package product

import (
	"context"

	"github.com/webomindapps-dev/coolaid-backend/internal/domain/search"
	"github.com/webomindapps-dev/coolaid-backend/oplog"
)

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

	oplog.Info(ctx, "creating vendors", "count", len(input.Vendors))
	oplog.Info(ctx, "creating oems", "count", len(input.OemNumbers))

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
		},
		oemIDs,
		vendorIDs,
	); err != nil {
		oplog.Error(ctx, "variant creation failed", "partNo=", product.PartNo, "error=", err)
		return nil, err
	}

	//  Create Pricing
	if _, err := s.createPricing(ctx, productQ, string(product.ID), input.Pricing); err != nil {
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

	if s.Indexer != nil {

		searchSvc := s.Indexer

		if err := searchSvc.Index(ctx, search.IndexRequest{
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
	oplog.Info(ctx, "update product started", "partNo", input.Main.PartNo)

	current, err := s.DB.Queries().GetProductByPartNo(ctx, input.Main.PartNo)
	if err != nil {
		oplog.Error(ctx, "load failed", "partNo", input.Main.PartNo, "error", err)
		return nil, ErrInternal
	}
	if current == nil {
		oplog.Warn(ctx, "product not found for update", "partNo", input.Main.PartNo)
		return nil, ErrProductNotFound
	}

	if err := s.validateUpdateProduct(ctx, input, current); err != nil {
		oplog.Error(ctx, "update validation failed", "partNo", input.Main.PartNo, "error", err)
		return nil, err
	}

	// --------------------------------------------------
	// 2️⃣ Begin Transaction
	// --------------------------------------------------

	tx, _, err := s.DB.BeginTx(ctx)
	if err != nil {
		oplog.Error(ctx, "transaction begin failed", "partNo", input.Main.PartNo, "error", err)
		return nil, ErrInternal
	}

	committed := false
	defer func() {
		if !committed {
			tx.Rollback()
			oplog.Warn(ctx, "transaction rolled back on update", "partNo", input.Main.PartNo)
		}
	}()

	productQ := s.DB.NewQueriesFromTx(tx)
	techQ := s.TechnicalService.DB.NewQueriesFromTx(tx)

	partNo := input.Main.PartNo

	// --------------------------------------------------
	// 3️⃣ Replace Vendors
	// --------------------------------------------------

	if err := s.deleteVendors(ctx, productQ, partNo); err != nil {
		oplog.Error(ctx, "failed to delete existing vendors during update",
			"partNo", partNo,
			"error", err,
		)
		return nil, ErrInternal
	}

	vendorIDs, err := s.createVendors(ctx, productQ, partNo, input.Vendors)
	if err != nil {
		oplog.Error(ctx, "failed to create/replace vendors during update",
			"partNo", partNo,
			"incoming_count", len(input.Vendors),
			"error", err,
		)
		return nil, err
	}
	oplog.Info(ctx, "vendors replaced successfully",
		"partNo", partNo,
		"new_count", len(vendorIDs),
	)

	// --------------------------------------------------
	// 4️⃣ Replace OEM
	// --------------------------------------------------

	if err := s.deleteOEMByPartNo(ctx, productQ, partNo); err != nil {
		oplog.Error(ctx, "failed to delete existing OEMs during update",
			"partNo", partNo,
			"error", err,
		)
		return nil, ErrInternal
	}

	oemIDs, err := s.createOEMs(ctx, productQ, partNo, input.OemNumbers)
	if err != nil {
		oplog.Error(ctx, "failed to create/replace OEMs during update",
			"partNo", partNo,
			"incoming_count", len(input.OemNumbers),
			"error", err,
		)
		return nil, err
	}

	oplog.Info(ctx, "OEMs replaced successfully",
		"partNo", partNo,
		"new_count", len(oemIDs),
	)

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
		oplog.Error(ctx, "failed to update model variant",
			"partNo", partNo,
			"error", err,
		)
		return nil, err
	}

	// --------------------------------------------------
	// 6️⃣ Update Pricing
	// --------------------------------------------------

	if _, err := s.updatePricing(ctx, productQ, partNo, input.Pricing); err != nil {
		oplog.Error(ctx, "failed to update pricing",
			"partNo", partNo,
			"error", err,
		)
		return nil, err
	}

	// --------------------------------------------------
	// 7️⃣ Update Inventory
	// --------------------------------------------------

	if _, err := s.updateInventory(ctx, productQ, partNo, input.Inventory); err != nil {
		oplog.Error(ctx, "failed to update inventory",
			"partNo", partNo,
			"error", err,
		)
		return nil, err
	}

	// --------------------------------------------------
	// 8️⃣ Update Offer
	// --------------------------------------------------

	if _, err := s.updateOffer(ctx, productQ, partNo, input.Offer); err != nil {
		oplog.Error(ctx, "failed to update offer",
			"partNo", partNo,
			"error", err,
		)
		return nil, err
	}

	// --------------------------------------------------
	// 9️⃣ Update TechSpec (delegated)
	// --------------------------------------------------

	if err := s.TechnicalService.UpdateTechSpec(ctx, techQ, partNo, &input.TechSpec); err != nil {
		oplog.Error(ctx, "failed to update tech spec",
			"partNo", partNo,
			"error", err,
		)
		return nil, err
	}

	// --------------------------------------------------
	// 🔟 Commit
	// --------------------------------------------------

	if err := tx.Commit(); err != nil {
		oplog.Error(ctx, "transaction commit failed",
			"partNo", partNo,
			"error", err,
		)
		return nil, ErrInternal
	}

	//TODO: Update Typesense only if required
	oplog.Info(ctx, "typesense update skipped", "partNo", partNo)

	if s.Indexer != nil {
		searchSvc := s.Indexer
		err := searchSvc.Update(ctx, search.IndexRequest{
			Collection: "products",
			ID:         partNo,
			Payload: search.ProductSearchDocument{
				ID:       string(current.ID),
				Company:  *input.Main.CompanyName, // note: better to resolve again if changed
				Model:    *input.Main.ModelName,
				PartNo:   partNo,
				Brand:    *input.Main.BrandName,
				Category: *input.Main.CategoryName,
			},
		})
		if err != nil {
			oplog.Error(ctx, "typesense reindex failed after update",
				"partNo", partNo,
				"error", err,
			)
			// do NOT return error — search is non-critical
		} else {
			oplog.Info(ctx, "typesense reindex success after update", "partNo", partNo)
		}
	}

	oplog.Info(ctx, "product update completed successfully", "partNo", partNo)

	return &Product{
		PartNo: partNo,
	}, nil
}

func (s *Service) DeleteProduct(
	ctx context.Context,
	partNo string,
) error {

	oplog.Info(ctx, "delete product started", "partNo", partNo)

	// --------------------------------------------------
	// 1️⃣ Ensure Exists
	// --------------------------------------------------

	product, err := s.DB.Queries().GetProductByPartNo(ctx, partNo)
	if err != nil {
		oplog.Error(ctx, "failed to load product for delete", "partNo", partNo, "error", err)
		return ErrInternal
	}
	if product == nil {
		oplog.Warn(ctx, "product not found for delete", "partNo", partNo)
		return ErrProductNotFound
	}

	tx, _, err := s.DB.BeginTx(ctx)
	if err != nil {
		oplog.Error(ctx, "transaction begin failed for delete", "partNo", partNo, "error", err)
		return ErrInternal
	}
	defer tx.Rollback()

	category, err := s.CategoryService.GetByID(ctx, product.CategoryID)
	if err != nil {
		oplog.Error(ctx, "failed to load category for delete", "partNo", partNo, "categoryID", product.CategoryID, "error", err)
		return ErrCategoryNotFound
	}

	productQ := s.DB.NewQueriesFromTx(tx)
	techQ := s.TechnicalService.DB.NewQueriesFromTx(tx)

	// --------------------------------------------------
	// 2️⃣ Delete Related Data
	// --------------------------------------------------

	if err := s.deleteVendors(ctx, productQ, partNo); err != nil {
		oplog.Error(ctx, "failed to delete vendors", "partNo", partNo, "error", err)
		return ErrInternal
	}

	if err := s.deleteOEMByPartNo(ctx, productQ, partNo); err != nil {
		oplog.Error(ctx, "failed to delete OEMs", "partNo", partNo, "error", err)
		return ErrInternal
	}

	if err := s.deleteInventory(ctx, productQ, partNo); err != nil {
		oplog.Error(ctx, "failed to delete inventory", "partNo", partNo, "error", err)
		return ErrInternal
	}

	if err := s.deleteOffer(ctx, productQ, partNo); err != nil {
		oplog.Error(ctx, "failed to delete offer", "partNo", partNo, "error", err)
		return ErrInternal
	}

	if err := s.deleteVariant(ctx, productQ, partNo); err != nil {
		oplog.Error(ctx, "failed to delete model variant", "partNo", partNo, "error", err)
		return ErrInternal
	}

	if err := s.deleteProduct(ctx, productQ, partNo); err != nil {
		oplog.Error(ctx, "failed to delete main product", "partNo", partNo, "error", err)
		return ErrInternal
	}

	// --------------------------------------------------
	// 3️⃣ Delete TechSpec
	// --------------------------------------------------

	if err := s.TechnicalService.DeleteTechSpec(ctx, techQ, product.PartNo, category.Name); err != nil {
		oplog.Error(ctx, "failed to delete tech spec", "partNo", partNo, "error", err)
		return err
	}

	// --------------------------------------------------
	// 4️⃣ Commit
	// --------------------------------------------------

	if err := tx.Commit(); err != nil {
		oplog.Error(ctx, "transaction commit failed during delete", "partNo", partNo, "error", err)
		return ErrInternal
	}

	// --------------------------------------------------
	// 5️⃣ Async Search Removal (after commit)
	// --------------------------------------------------

	if s.Indexer != nil {
		searchSvc := s.Indexer
		if err := searchSvc.Delete(ctx, "products", partNo); err != nil {
			oplog.Error(ctx, "typesense delete failed after product removal",
				"partNo", partNo,
				"error", err,
			)
			// do NOT fail the delete for search error
		} else {
			oplog.Info(ctx, "typesense delete success", "partNo", partNo)
		}
	}

	oplog.Info(ctx, "product deleted successfully", "partNo", partNo)

	return nil
}
