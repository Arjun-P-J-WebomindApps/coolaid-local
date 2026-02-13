package product

import (
	"context"

	"github.com/google/uuid"
)

//
// ============================================================
// 🔹 INVENTORY (Single Table Lifecycle)
// ============================================================
//

//
// 🔹 GET
//

func (s *Service) getInventoryByPartNo(
	ctx context.Context,
	Q Queries,
	partNo string,
) (*ProductInventory, error) {

	if partNo == "" {
		return nil, ErrInvalidInput
	}

	row, err := Q.GetInventoryByPartNo(ctx, partNo)
	if err != nil {
		return nil, ErrInventoryNotFound
	}

	return mapInventoryRowToModel(row), nil
}

//
// 🔹 CREATE
//

func (s *Service) createInventory(
	ctx context.Context,
	Q Queries,
	partNo string,
	input CreateProductInventoryInput,
) (*ProductInventory, error) {

	if partNo == "" {
		return nil, ErrInvalidInput
	}

	row, err := Q.CreateInventory(ctx, CreateProductInventoryParams{
		ID:                   uuid.NewString(),
		PartNo:               partNo,
		MinimumOrderLevel:    input.MinimumOrderLevel,
		MaximumOrderLevel:    input.MaximumOrderLevel,
		QtyInStock:           input.QtyInStock,
		Location:             input.Location,
		IsFlash:              input.IsFlash,
		IsRequestedForSupply: false,
		VendorID:             nil,
	})
	if err != nil {
		return nil, err
	}

	return mapInventoryRowToModel(row), nil
}

//
// 🔹 UPDATE
//

func (s *Service) updateInventory(
	ctx context.Context,
	Q Queries,
	partNo string,
	input UpdateProductInventoryInput,
) (*ProductInventory, error) {

	if partNo == "" {
		return nil, ErrInvalidInput
	}

	var vendorID *string

	// -------------------------------------------------
	// Resolve VendorName -> VendorID (if provided)
	// -------------------------------------------------
	if input.VendorName != nil && *input.VendorName != "" {
		vendor, err := s.VendorService.DB.Queries().GetVendorByCompanyName(ctx, *input.VendorName)
		if err != nil {
			return nil, ErrInvalidVendor
		}
		vID := vendor.ID
		vendorID = &vID
	}

	row, err := Q.UpdateInventory(ctx, UpdateProductInventoryParams{
		PartNo:            partNo,
		MinimumOrderLevel: input.MinimumOrderLevel,
		MaximumOrderLevel: input.MaximumOrderLevel,
		QtyInStock:        input.QtyInStock,
		Location:          input.Location,
		IsFlash:           input.IsFlash,
		VendorID:          vendorID,
	})
	if err != nil {
		return nil, err
	}

	return mapInventoryRowToModel(row), nil
}

//
// 🔹 DELETE
//

func (s *Service) deleteInventory(
	ctx context.Context,
	Q Queries,
	partNo string,
) error {

	if partNo == "" {
		return ErrInvalidInput
	}

	_, err := Q.GetInventoryByPartNo(ctx, partNo)
	if err != nil {
		return ErrInventoryNotFound
	}

	return Q.DeleteInventory(ctx, partNo)
}
