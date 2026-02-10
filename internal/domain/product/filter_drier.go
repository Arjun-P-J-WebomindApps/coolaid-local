package techspec

import (
	"context"

	"github.com/google/uuid"
	"github.com/webomindapps-dev/coolaid-backend/internal/shared/ptr"
)

func ValidateFilterDrierInput(in *FilterDrierInput) error {
	if in == nil {
		return ErrInvalidTechSpec
	}
	return nil
}

func BuildCreateFilterDrierParams(partNo string, in *FilterDrierInput) (CreateFilterDrierParams, error) {
	if partNo == "" {
		return CreateFilterDrierParams{}, ErrInvalidPartNo
	}
	if err := ValidateFilterDrierInput(in); err != nil {
		return CreateFilterDrierParams{}, err
	}

	return CreateFilterDrierParams{
		ID:             uuid.NewString(),
		PartNo:         partNo,
		PipeConnector:  ptr.StringOrNil(in.PipeConnector),
		Size:           ptr.StringOrNil(in.Size),
		PressureSwitch: ptr.StringOrNil(in.PressureSwitch),
		Notes:          ptr.StringOrNil(in.Notes),
	}, nil
}

func BuildUpdateFilterDrierParams(partNo string, in *FilterDrierInput) (UpdateFilterDrierParams, error) {
	if partNo == "" {
		return UpdateFilterDrierParams{}, ErrInvalidPartNo
	}
	if err := ValidateFilterDrierInput(in); err != nil {
		return UpdateFilterDrierParams{}, err
	}

	return UpdateFilterDrierParams{
		PartNo:         partNo,
		PipeConnector:  ptr.StringOrNil(in.PipeConnector),
		Size:           ptr.StringOrNil(in.Size),
		PressureSwitch: ptr.StringOrNil(in.PressureSwitch),
		Notes:          ptr.StringOrNil(in.Notes),
	}, nil
}

func CreateFilterDrier(ctx context.Context, q Queries, partNo string, in *FilterDrierInput) (*FilterDrierRow, error) {
	p, err := BuildCreateFilterDrierParams(partNo, in)
	if err != nil {
		return nil, err
	}
	row, err := q.CreateFilterDrier(ctx, p)
	if err != nil {
		return nil, ErrCreateFilterDrierFailed
	}
	return row, nil
}

func UpdateFilterDrier(ctx context.Context, q Queries, partNo string, in *FilterDrierInput) (*FilterDrierRow, error) {
	p, err := BuildUpdateFilterDrierParams(partNo, in)
	if err != nil {
		return nil, err
	}
	row, err := q.UpdateFilterDrierByPartNo(ctx, p)
	if err != nil {
		return nil, ErrUpdateFilterDrierFailed
	}
	return row, nil
}

func DeleteFilterDrier(ctx context.Context, q Queries, partNo string) error {
	if partNo == "" {
		return ErrInvalidPartNo
	}
	if err := q.DeleteFilterDrierByPartNo(ctx, partNo); err != nil {
		return ErrDeleteFilterDrierFailed
	}
	return nil
}
