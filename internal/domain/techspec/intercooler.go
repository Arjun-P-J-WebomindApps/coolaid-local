package techspec

import (
	"context"

	"github.com/google/uuid"
	"github.com/webomindapps-dev/coolaid-backend/internal/shared/ptr"
)

func ValidateIntercoolerInput(in *IntercoolerInput) error {
	if in == nil {
		return ErrInvalidTechSpec
	}
	return nil
}

func BuildCreateIntercoolerParams(partNo string, in *IntercoolerInput) (CreateIntercoolerParams, error) {
	if partNo == "" {
		return CreateIntercoolerParams{}, ErrInvalidPartNo
	}
	if err := ValidateIntercoolerInput(in); err != nil {
		return CreateIntercoolerParams{}, err
	}

	return CreateIntercoolerParams{
		ID:         uuid.NewString(),
		PartNo:     partNo,
		Size:       ptr.StringOrNil(in.Size),
		TempSensor: ptr.StringOrNil(in.TempSensor),
		Notes:      ptr.StringOrNil(in.Notes),
	}, nil
}

func BuildUpdateIntercoolerParams(partNo string, in *IntercoolerInput) (UpdateIntercoolerParams, error) {
	if partNo == "" {
		return UpdateIntercoolerParams{}, ErrInvalidPartNo
	}
	if err := ValidateIntercoolerInput(in); err != nil {
		return UpdateIntercoolerParams{}, err
	}

	return UpdateIntercoolerParams{
		PartNo:     partNo,
		Size:       ptr.StringOrNil(in.Size),
		TempSensor: ptr.StringOrNil(in.TempSensor),
		Notes:      ptr.StringOrNil(in.Notes),
	}, nil
}

func CreateIntercooler(ctx context.Context, q Queries, partNo string, in *IntercoolerInput) (*IntercoolerRow, error) {
	p, err := BuildCreateIntercoolerParams(partNo, in)
	if err != nil {
		return nil, err
	}
	row, err := q.CreateIntercooler(ctx, p)
	if err != nil {
		return nil, ErrCreateIntercoolerFailed
	}
	return row, nil
}

func UpdateIntercooler(ctx context.Context, q Queries, partNo string, in *IntercoolerInput) (*IntercoolerRow, error) {
	p, err := BuildUpdateIntercoolerParams(partNo, in)
	if err != nil {
		return nil, err
	}
	row, err := q.UpdateIntercoolerByPartNo(ctx, p)
	if err != nil {
		return nil, ErrUpdateIntercoolerFailed
	}
	return row, nil
}

func DeleteIntercooler(ctx context.Context, q Queries, partNo string) error {
	if partNo == "" {
		return ErrInvalidPartNo
	}
	if err := q.DeleteIntercoolerByPartNo(ctx, partNo); err != nil {
		return ErrDeleteIntercoolerFailed
	}
	return nil
}
