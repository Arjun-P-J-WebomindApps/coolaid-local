package techspec

import (
	"context"

	"github.com/google/uuid"
	"github.com/webomindapps-dev/coolaid-backend/internal/shared/ptr"
)

func ValidatePressureSwitchInput(in *PressureSwitchInput) error {
	if in == nil {
		return ErrInvalidTechSpec
	}
	return nil
}

func BuildCreatePressureSwitchParams(partNo string, in *PressureSwitchInput) (CreatePressureSwitchParams, error) {
	if partNo == "" {
		return CreatePressureSwitchParams{}, ErrInvalidPartNo
	}
	if err := ValidatePressureSwitchInput(in); err != nil {
		return CreatePressureSwitchParams{}, err
	}

	return CreatePressureSwitchParams{
		ID:            uuid.NewString(),
		PartNo:        partNo,
		ConnectorType: ptr.StringOrNil(in.ConnectorType),
		ThreadType:    ptr.StringOrNil(in.ThreadType),
		Notes:         ptr.StringOrNil(in.Notes),
	}, nil
}

func BuildUpdatePressureSwitchParams(partNo string, in *PressureSwitchInput) (UpdatePressureSwitchParams, error) {
	if partNo == "" {
		return UpdatePressureSwitchParams{}, ErrInvalidPartNo
	}
	if err := ValidatePressureSwitchInput(in); err != nil {
		return UpdatePressureSwitchParams{}, err
	}

	return UpdatePressureSwitchParams{
		PartNo:        partNo,
		ConnectorType: ptr.StringOrNil(in.ConnectorType),
		ThreadType:    ptr.StringOrNil(in.ThreadType),
		Notes:         ptr.StringOrNil(in.Notes),
	}, nil
}

func CreatePressureSwitch(ctx context.Context, q Queries, partNo string, in *PressureSwitchInput) (*PressureSwitchRow, error) {
	p, err := BuildCreatePressureSwitchParams(partNo, in)
	if err != nil {
		return nil, err
	}
	row, err := q.CreatePressureSwitch(ctx, p)
	if err != nil {
		return nil, ErrCreatePressureSwitchFailed
	}
	return row, nil
}

func UpdatePressureSwitch(ctx context.Context, q Queries, partNo string, in *PressureSwitchInput) (*PressureSwitchRow, error) {
	p, err := BuildUpdatePressureSwitchParams(partNo, in)
	if err != nil {
		return nil, err
	}
	row, err := q.UpdatePressureSwitchByPartNo(ctx, p)
	if err != nil {
		return nil, ErrUpdatePressureSwitchFailed
	}
	return row, nil
}

func DeletePressureSwitch(ctx context.Context, q Queries, partNo string) error {
	if partNo == "" {
		return ErrInvalidPartNo
	}
	if err := q.DeletePressureSwitchByPartNo(ctx, partNo); err != nil {
		return ErrDeletePressureSwitchFailed
	}
	return nil
}
