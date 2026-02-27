package techspec

import (
	"context"

	"github.com/google/uuid"
	"github.com/webomindapps-dev/coolaid-backend/internal/shared/ptr"
)

func (s *Service) GetRadFanMotor(ctx context.Context, partNo string) (*RadFanMotorRow, error) {
	if partNo == "" {
		return nil, ErrInvalidPartNo
	}
	return s.DB.Queries().GetRadFanMotorByPartNo(ctx, partNo)
}

func ValidateRadFanMotorInput(in *RadFanMotorInput) error {
	if in == nil {
		return ErrInvalidTechSpec
	}
	return nil
}

func BuildCreateRadFanMotorParams(partNo string, in *RadFanMotorInput) (CreateRadFanMotorParams, error) {
	if partNo == "" {
		return CreateRadFanMotorParams{}, ErrInvalidPartNo
	}
	if err := ValidateRadFanMotorInput(in); err != nil {
		return CreateRadFanMotorParams{}, err
	}

	return CreateRadFanMotorParams{
		ID:               uuid.NewString(),
		PartNo:           partNo,
		FanBladeDiameter: ptr.StringOrNil(in.FanBladeDiameter),
		NumberOfBlades:   ptr.Int32OrNil(in.NumberOfBlades),
		Voltage:          ptr.StringOrNil(in.Voltage),
		NumberOfSockets:  ptr.Int32OrNil(in.NumberOfSockets),
		ConnectorType:    ptr.StringOrNil(in.ConnectorType),
		Notes:            ptr.StringOrNil(in.Notes),
	}, nil
}

func BuildUpdateRadFanMotorParams(partNo string, in *RadFanMotorInput) (UpdateRadFanMotorParams, error) {
	if partNo == "" {
		return UpdateRadFanMotorParams{}, ErrInvalidPartNo
	}
	if err := ValidateRadFanMotorInput(in); err != nil {
		return UpdateRadFanMotorParams{}, err
	}

	return UpdateRadFanMotorParams{
		PartNo:           partNo,
		FanBladeDiameter: ptr.StringOrNil(in.FanBladeDiameter),
		NumberOfBlades:   ptr.Int32OrNil(in.NumberOfBlades),
		Voltage:          ptr.StringOrNil(in.Voltage),
		NumberOfSockets:  ptr.Int32OrNil(in.NumberOfSockets),
		ConnectorType:    ptr.StringOrNil(in.ConnectorType),
		Notes:            ptr.StringOrNil(in.Notes),
	}, nil
}

func CreateRadFanMotor(ctx context.Context, q Queries, partNo string, in *RadFanMotorInput) (*RadFanMotorRow, error) {
	p, err := BuildCreateRadFanMotorParams(partNo, in)
	if err != nil {
		return nil, err
	}
	row, err := q.CreateRadFanMotor(ctx, p)
	if err != nil {
		return nil, ErrCreateRadFanMotorFailed
	}
	return row, nil
}

func UpdateRadFanMotor(ctx context.Context, q Queries, partNo string, in *RadFanMotorInput) (*RadFanMotorRow, error) {
	p, err := BuildUpdateRadFanMotorParams(partNo, in)
	if err != nil {
		return nil, err
	}
	row, err := q.UpdateRadFanMotorByPartNo(ctx, p)
	if err != nil {
		return nil, ErrUpdateRadFanMotorFailed
	}
	return row, nil
}

func DeleteRadFanMotor(ctx context.Context, q Queries, partNo string) error {
	if partNo == "" {
		return ErrInvalidPartNo
	}
	if err := q.DeleteRadFanMotorByPartNo(ctx, partNo); err != nil {
		return ErrDeleteRadFanMotorFailed
	}
	return nil
}
