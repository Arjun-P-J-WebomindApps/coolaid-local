package techspec

import (
	"context"

	"github.com/google/uuid"
	"github.com/webomindapps-dev/coolaid-backend/internal/shared/ptr"
)

func (s *Service) GetRadFanAssy(ctx context.Context, partNo string) (*RadFanAssyRow, error) {
	if partNo == "" {
		return nil, ErrInvalidPartNo
	}
	return s.DB.Queries().GetRadFanAssyByPartNo(ctx, partNo)
}

func ValidateRadFanAssyInput(in *RadFanAssyInput) error {
	if in == nil {
		return ErrInvalidTechSpec
	}
	return nil
}

func BuildCreateRadFanAssyParams(partNo string, in *RadFanAssyInput) (CreateRadFanAssyParams, error) {
	if partNo == "" {
		return CreateRadFanAssyParams{}, ErrInvalidPartNo
	}
	if err := ValidateRadFanAssyInput(in); err != nil {
		return CreateRadFanAssyParams{}, err
	}

	return CreateRadFanAssyParams{
		ID:               uuid.NewString(),
		PartNo:           partNo,
		Voltage:          ptr.StringOrNil(in.Voltage),
		MotorType:        ptr.StringOrNil(in.MotorType),
		Resistance:       ptr.StringOrNil(in.Resistance),
		NumberOfSockets:  ptr.Int32OrNil(in.NumberOfSockets),
		Shroud:           ptr.StringOrNil(in.Shroud),
		ConnectorType:    ptr.StringOrNil(in.ConnectorType),
		FanBladeDiameter: ptr.StringOrNil(in.FanBladeDiameter),
		NumberOfBlades:   ptr.Int32OrNil(in.NumberOfBlades),
		Notes:            ptr.StringOrNil(in.Notes),
	}, nil
}

func BuildUpdateRadFanAssyParams(partNo string, in *RadFanAssyInput) (UpdateRadFanAssyParams, error) {
	if partNo == "" {
		return UpdateRadFanAssyParams{}, ErrInvalidPartNo
	}
	if err := ValidateRadFanAssyInput(in); err != nil {
		return UpdateRadFanAssyParams{}, err
	}

	return UpdateRadFanAssyParams{
		PartNo:           partNo,
		Voltage:          ptr.StringOrNil(in.Voltage),
		MotorType:        ptr.StringOrNil(in.MotorType),
		Resistance:       ptr.StringOrNil(in.Resistance),
		NumberOfSockets:  ptr.Int32OrNil(in.NumberOfSockets),
		Shroud:           ptr.StringOrNil(in.Shroud),
		ConnectorType:    ptr.StringOrNil(in.ConnectorType),
		FanBladeDiameter: ptr.StringOrNil(in.FanBladeDiameter),
		NumberOfBlades:   ptr.Int32OrNil(in.NumberOfBlades),
		Notes:            ptr.StringOrNil(in.Notes),
	}, nil
}

func CreateRadFanAssy(ctx context.Context, q Queries, partNo string, in *RadFanAssyInput) (*RadFanAssyRow, error) {
	p, err := BuildCreateRadFanAssyParams(partNo, in)
	if err != nil {
		return nil, err
	}
	row, err := q.CreateRadFanAssy(ctx, p)
	if err != nil {
		return nil, ErrCreateRadFanAssyFailed
	}
	return row, nil
}

func UpdateRadFanAssy(ctx context.Context, q Queries, partNo string, in *RadFanAssyInput) (*RadFanAssyRow, error) {
	p, err := BuildUpdateRadFanAssyParams(partNo, in)
	if err != nil {
		return nil, err
	}
	row, err := q.UpdateRadFanAssyByPartNo(ctx, p)
	if err != nil {
		return nil, ErrUpdateRadFanAssyFailed
	}
	return row, nil
}

func DeleteRadFanAssy(ctx context.Context, q Queries, partNo string) error {
	if partNo == "" {
		return ErrInvalidPartNo
	}
	if err := q.DeleteRadFanAssyByPartNo(ctx, partNo); err != nil {
		return ErrDeleteRadFanAssyFailed
	}
	return nil
}
