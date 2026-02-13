package techspec

import (
	"context"

	"github.com/google/uuid"
	"github.com/webomindapps-dev/coolaid-backend/internal/shared/ptr"
)

func ValidateStatorInput(in *StatorInput) error {
	if in == nil {
		return ErrInvalidTechSpec
	}
	return nil
}

func BuildCreateStatorParams(partNo string, in *StatorInput) (CreateStatorParams, error) {
	if partNo == "" {
		return CreateStatorParams{}, ErrInvalidPartNo
	}
	if err := ValidateStatorInput(in); err != nil {
		return CreateStatorParams{}, err
	}

	return CreateStatorParams{
		ID:                uuid.NewString(),
		PartNo:            partNo,
		Voltage:           ptr.StringOrNil(in.Voltage),
		CompressorDetails: ptr.StringOrNil(in.CompressorDetails),
		Notes:             ptr.StringOrNil(in.Notes),
	}, nil
}

func BuildUpdateStatorParams(partNo string, in *StatorInput) (UpdateStatorParams, error) {
	if partNo == "" {
		return UpdateStatorParams{}, ErrInvalidPartNo
	}
	if err := ValidateStatorInput(in); err != nil {
		return UpdateStatorParams{}, err
	}

	return UpdateStatorParams{
		PartNo:            partNo,
		Voltage:           ptr.StringOrNil(in.Voltage),
		CompressorDetails: ptr.StringOrNil(in.CompressorDetails),
		Notes:             ptr.StringOrNil(in.Notes),
	}, nil
}

func CreateStator(ctx context.Context, q Queries, partNo string, in *StatorInput) (*StatorRow, error) {
	p, err := BuildCreateStatorParams(partNo, in)
	if err != nil {
		return nil, err
	}
	row, err := q.CreateStator(ctx, p)
	if err != nil {
		return nil, ErrCreateStatorFailed
	}
	return row, nil
}

func UpdateStator(ctx context.Context, q Queries, partNo string, in *StatorInput) (*StatorRow, error) {
	p, err := BuildUpdateStatorParams(partNo, in)
	if err != nil {
		return nil, err
	}
	row, err := q.UpdateStatorByPartNo(ctx, p)
	if err != nil {
		return nil, ErrUpdateStatorFailed
	}
	return row, nil
}

func DeleteStator(ctx context.Context, q Queries, partNo string) error {
	if partNo == "" {
		return ErrInvalidPartNo
	}
	if err := q.DeleteStatorByPartNo(ctx, partNo); err != nil {
		return ErrDeleteStatorFailed
	}
	return nil
}
