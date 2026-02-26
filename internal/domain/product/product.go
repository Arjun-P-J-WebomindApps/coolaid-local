package product

import (
	"context"
	"time"

	"github.com/google/uuid"
)

//
// ============================================================
// 🔹 PRODUCT (Main Table Only)
// Internal helpers — not exported outside domain
// ============================================================
//

func (s *Service) getProductByPartNo(
	ctx context.Context,
	Q Queries,
	partNo string,
) (*Product, error) {

	if partNo == "" {
		return nil, ErrInvalidInput
	}

	row, err := Q.GetProductByPartNo(ctx, partNo)
	if err != nil {
		return nil, ErrProductNotFound
	}

	return mapProductRowToModel(row), nil
}

func (s *Service) createProduct(
	ctx context.Context,
	Q Queries,
	partNo string,
	companyID string,
	modelID string,
	brandID string,
	categoryID string,
) (*Product, error) {

	if partNo == "" ||
		companyID == "" ||
		modelID == "" ||
		brandID == "" ||
		categoryID == "" {
		return nil, ErrInvalidInput
	}

	// check duplicate
	existing, err := Q.GetProductByPartNo(ctx, partNo)
	if err == nil && existing != nil {
		return nil, ErrProductExists
	}

	now := time.Now()

	row, err := Q.CreateProductPart(ctx, CreateProductParams{
		ID:         uuid.NewString(),
		PartNo:     partNo,
		CompanyID:  companyID,
		ModelID:    modelID,
		BrandID:    brandID,
		CategoryID: categoryID,
		IsActive:   true,
		CreatedAt:  now,
		UpdatedAt:  now,
	})
	if err != nil {
		return nil, err
	}

	return mapProductRowToModel(row), nil
}

// Main details are not to be updated as per current configuration
func (s *Service) updateProduct(
	ctx context.Context,
	Q Queries,
	id string,
	input UpdateProductParams,
) (*Product, error) {

	if id == "" {
		return nil, ErrInvalidInput
	}

	now := time.Now().String()
	input.ID = id

	input.UpdatedAt = &now

	row, err := Q.UpdateProductPart(ctx, input)
	if err != nil {
		return nil, err
	}

	return mapProductRowToModel(row), nil
}

func (s *Service) deleteProduct(
	ctx context.Context,
	Q Queries,
	partNo string,
) error {

	if partNo == "" {
		return ErrInvalidInput
	}

	_, err := Q.GetProductByPartNo(ctx, partNo)
	if err != nil {
		return ErrProductNotFound
	}

	return Q.DeleteProductPart(ctx, partNo)
}
