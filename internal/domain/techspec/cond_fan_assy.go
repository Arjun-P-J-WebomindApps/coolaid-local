package techspec

import (
	"context"

	"github.com/google/uuid"
	"github.com/webomindapps-dev/coolaid-backend/internal/shared/ptr"
)

// =======================================================
// CONDENSER FAN ASSY
// =======================================================

func (s *Service) GetCondFanAssy(ctx context.Context, partNo string) (*CondFanAssyRow, error) {
	if partNo == "" {
		return nil, ErrInvalidPartNo
	}
	return s.DB.Queries().GetCondFanAssyByPartNo(ctx, partNo)
}

func ValidateCondFanAssyInput(in *CondFanAssyInput) error {
	if in == nil {
		return ErrInvalidTechSpec
	}
	return nil
}

func BuildCreateCondFanAssyParams(
	partNo string,
	in *CondFanAssyInput,
) (CreateCondFanAssyParams, error) {

	if partNo == "" {
		return CreateCondFanAssyParams{}, ErrInvalidPartNo
	}
	if err := ValidateCondFanAssyInput(in); err != nil {
		return CreateCondFanAssyParams{}, err
	}

	return CreateCondFanAssyParams{
		ID:               uuid.NewString(),
		PartNo:           partNo,
		Voltage:          ptr.StringOrNil(in.Voltage),
		MotorType:        ptr.StringOrNil(in.MotorType),
		Resistance:       ptr.StringOrNil(in.Resistance),
		FanBladeDiameter: ptr.StringOrNil(in.FanBladeDiameter),
		NumberOfBlades:   ptr.Int32OrNil(in.NumberOfBlades),
		Shroud:           ptr.StringOrNil(in.Shroud),
		ConnectorType:    ptr.StringOrNil(in.ConnectorType),
		Notes:            ptr.StringOrNil(in.Notes),
	}, nil
}

func BuildUpdateCondFanAssyParams(
	partNo string,
	in *CondFanAssyInput,
) (UpdateCondFanAssyParams, error) {

	if partNo == "" {
		return UpdateCondFanAssyParams{}, ErrInvalidPartNo
	}
	if err := ValidateCondFanAssyInput(in); err != nil {
		return UpdateCondFanAssyParams{}, err
	}

	return UpdateCondFanAssyParams{
		PartNo:           partNo,
		Voltage:          ptr.StringOrNil(in.Voltage),
		MotorType:        ptr.StringOrNil(in.MotorType),
		Resistance:       ptr.StringOrNil(in.Resistance),
		FanBladeDiameter: ptr.StringOrNil(in.FanBladeDiameter),
		NumberOfBlades:   ptr.Int32OrNil(in.NumberOfBlades),
		Shroud:           ptr.StringOrNil(in.Shroud),
		ConnectorType:    ptr.StringOrNil(in.ConnectorType),
		Notes:            ptr.StringOrNil(in.Notes),
	}, nil
}

func CreateCondFanAssy(
	ctx context.Context,
	q Queries,
	partNo string,
	in *CondFanAssyInput,
) (*CondFanAssyRow, error) {

	params, err := BuildCreateCondFanAssyParams(partNo, in)
	if err != nil {
		return nil, err
	}

	row, err := q.CreateCondFanAssy(ctx, params)
	if err != nil {
		return nil, ErrCreateCondFanAssyFailed
	}

	return row, nil
}

func UpdateCondFanAssy(
	ctx context.Context,
	q Queries,
	partNo string,
	in *CondFanAssyInput,
) (*CondFanAssyRow, error) {

	params, err := BuildUpdateCondFanAssyParams(partNo, in)
	if err != nil {
		return nil, err
	}

	row, err := q.UpdateCondFanAssyByPartNo(ctx, params)
	if err != nil {
		return nil, ErrUpdateCondFanAssyFailed
	}

	return row, nil
}

func DeleteCondFanAssy(
	ctx context.Context,
	q Queries,
	partNo string,
) error {

	if partNo == "" {
		return ErrInvalidPartNo
	}

	if err := q.DeleteCondFanAssyByPartNo(ctx, partNo); err != nil {
		return ErrDeleteCondFanAssyFailed
	}

	return nil
}
