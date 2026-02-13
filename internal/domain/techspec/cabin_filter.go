package techspec

import (
	"context"

	"github.com/google/uuid"
	"github.com/webomindapps-dev/coolaid-backend/internal/shared/ptr"
)

func ValidateCabinFilterInput(in *CabinFilterInput) error {
	if in == nil {
		return ErrInvalidTechSpec
	}
	return nil
}

func BuildCreateCabinFilterParams(partNo string, in *CabinFilterInput) (CreateCabinFilterParams, error) {
	if partNo == "" {
		return CreateCabinFilterParams{}, ErrInvalidPartNo
	}
	if err := ValidateCabinFilterInput(in); err != nil {
		return CreateCabinFilterParams{}, err
	}

	return CreateCabinFilterParams{
		ID:         uuid.NewString(),
		PartNo:     partNo,
		Type:       ptr.StringOrNil(in.Type),
		Dimensions: ptr.StringOrNil(in.Dimensions),
		Material:   ptr.StringOrNil(in.Material),
		Notes:      ptr.StringOrNil(in.Notes),
	}, nil
}

func BuildUpdateCabinFilterParams(partNo string, in *CabinFilterInput) (UpdateCabinFilterParams, error) {
	if partNo == "" {
		return UpdateCabinFilterParams{}, ErrInvalidPartNo
	}
	if err := ValidateCabinFilterInput(in); err != nil {
		return UpdateCabinFilterParams{}, err
	}

	return UpdateCabinFilterParams{
		PartNo:     partNo,
		Type:       ptr.StringOrNil(in.Type),
		Dimensions: ptr.StringOrNil(in.Dimensions),
		Material:   ptr.StringOrNil(in.Material),
		Notes:      ptr.StringOrNil(in.Notes),
	}, nil
}

func CreateCabinFilter(ctx context.Context, q Queries, partNo string, in *CabinFilterInput) (*CabinFilterRow, error) {
	p, err := BuildCreateCabinFilterParams(partNo, in)
	if err != nil {
		return nil, err
	}

	row, err := q.CreateCabinFilter(ctx, p)
	if err != nil {
		return nil, ErrCreateCabinFilterFailed
	}
	return row, nil
}

func UpdateCabinFilter(ctx context.Context, q Queries, partNo string, in *CabinFilterInput) (*CabinFilterRow, error) {
	p, err := BuildUpdateCabinFilterParams(partNo, in)
	if err != nil {
		return nil, err
	}

	row, err := q.UpdateCabinFilterByPartNo(ctx, p)
	if err != nil {
		return nil, ErrUpdateCabinFilterFailed
	}
	return row, nil
}

func DeleteCabinFilter(ctx context.Context, q Queries, partNo string) error {
	if partNo == "" {
		return ErrInvalidPartNo
	}
	if err := q.DeleteCabinFilterByPartNo(ctx, partNo); err != nil {
		return ErrDeleteCabinFilterFailed
	}
	return nil
}
