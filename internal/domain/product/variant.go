package product

import (
	"context"

	"github.com/google/uuid"
)

//
// ============================================================
// 🔹 MODEL VARIANT (Single Table Lifecycle)
// ============================================================
//

//
// 🔹 GET
//

func (s *Service) getVariantByPartNo(
	ctx context.Context,
	Q Queries,
	partNo string,
) (*ModelVariant, error) {

	if partNo == "" {
		return nil, ErrInvalidInput
	}

	row, err := Q.GetModelVariantByPartNo(ctx, partNo)
	if err != nil {
		return nil, ErrModelNotFound
	}

	return mapModelVariantRowToModel(row), nil
}

//
// 🔹 CREATE
//

func (s *Service) createVariant(
	ctx context.Context,
	Q Queries,
	partNo string,
	input CreateModelVariantInput,
	OemIds []string,
	VendorIDs []string,
) (*ModelVariant, error) {

	if partNo == "" || input.Type == "" {
		return nil, ErrInvalidInput
	}

	params := CreateModelVariantParams{
		ID:               uuid.NewString(),
		PartNo:           partNo,
		Type:             input.Type,
		FuelTypes:        input.FuelTypes,
		TransmissionType: input.TransmissionType,
		PlatformCodes:    input.PlatformCodes,
		Placement:        input.Placement,
		Make:             input.Make,
		Unicode:          input.Unicode,
		YearStart:        &input.YearStart,
		YearEnd:          &input.YearEnd,
		Gen:              &input.Gen,
		EngineCc:         &input.EngineCc,
		Image1Link:       &input.Image1Link,
		Image2Link:       &input.Image2Link,
		Image3Link:       &input.Image3Link,
		Image4Link:       &input.Image4Link,
		HsnCode:          &input.HsnCode,
		Description:      &input.Description,
		AdditionalInfo:   &input.AdditionalInfo,
		OemIds:           OemIds,
		VendorID:         VendorIDs,
	}

	// Optional pointer fields — only set if not empty

	if input.Gen != "" {
		params.Gen = &input.Gen
	}

	if input.HsnCode != "" {
		params.HsnCode = &input.HsnCode
	}

	if input.EngineCc != 0 {
		params.EngineCc = &input.EngineCc
	}

	if input.Image1Link != "" {
		params.Image1Link = &input.Image1Link
	}
	if input.Image2Link != "" {
		params.Image2Link = &input.Image2Link
	}
	if input.Image3Link != "" {
		params.Image3Link = &input.Image3Link
	}
	if input.Image4Link != "" {
		params.Image4Link = &input.Image4Link
	}

	if input.Description != "" {
		params.Description = &input.Description
	}

	if input.AdditionalInfo != "" {
		params.AdditionalInfo = &input.AdditionalInfo
	}

	row, err := Q.CreateModelVariant(ctx, params)
	if err != nil {
		return nil, err
	}

	return mapModelVariantRowToModel(row), nil
}

//
// 🔹 UPDATE
//

func (s *Service) updateVariant(
	ctx context.Context,
	Q Queries,
	partNo string,
	input UpdateModelVariantParams,
) (*ModelVariant, error) {

	if partNo == "" {
		return nil, ErrInvalidInput
	}

	row, err := Q.UpdateModelVariant(ctx, input)
	if err != nil {
		return nil, err
	}

	return mapModelVariantRowToModel(row), nil
}

//
// 🔹 DELETE
//

func (s *Service) deleteVariant(
	ctx context.Context,
	Q Queries,
	partNo string,
) error {

	if partNo == "" {
		return ErrInvalidInput
	}

	_, err := Q.GetModelVariantByPartNo(ctx, partNo)
	if err != nil {
		return ErrModelNotFound
	}

	return Q.DeleteModelVariant(ctx, partNo)
}
