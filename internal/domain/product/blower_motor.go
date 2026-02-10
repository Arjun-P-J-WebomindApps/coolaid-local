package techspec

import (
	"context"

	"github.com/google/uuid"
	"github.com/webomindapps-dev/coolaid-backend/internal/shared/ptr"
)

// =======================================================
// BLOWER MOTOR
// =======================================================

func ValidateBlowerMotorInput(in *BlowerMotorInput) error {
	if in == nil {
		return ErrInvalidTechSpec
	}
	return nil
}

func BuildCreateBlowerMotorParams(
	partNo string,
	in *BlowerMotorInput,
) (CreateBlowerMotorParams, error) {

	if partNo == "" {
		return CreateBlowerMotorParams{}, ErrInvalidPartNo
	}
	if err := ValidateBlowerMotorInput(in); err != nil {
		return CreateBlowerMotorParams{}, err
	}

	return CreateBlowerMotorParams{
		ID:            uuid.NewString(),
		PartNo:        partNo,
		Mounting:      ptr.StringOrNil(in.Mounting),
		ConnectorType: ptr.StringOrNil(in.ConnectorType),
		Impeller:      ptr.StringOrNil(in.Impeller),
		Resistance:    ptr.StringOrNil(in.Resistance),
		MotorMounting: ptr.StringOrNil(in.MotorMounting),
		MotorType:     ptr.StringOrNil(in.MotorType),
		Voltage:       ptr.StringOrNil(in.Voltage),
		Notes:         ptr.StringOrNil(in.Notes),
	}, nil
}

func BuildUpdateBlowerMotorParams(
	partNo string,
	in *BlowerMotorInput,
) (UpdateBlowerMotorParams, error) {

	if partNo == "" {
		return UpdateBlowerMotorParams{}, ErrInvalidPartNo
	}
	if err := ValidateBlowerMotorInput(in); err != nil {
		return UpdateBlowerMotorParams{}, err
	}

	return UpdateBlowerMotorParams{
		PartNo:        partNo,
		Mounting:      ptr.StringOrNil(in.Mounting),
		ConnectorType: ptr.StringOrNil(in.ConnectorType),
		Impeller:      ptr.StringOrNil(in.Impeller),
		Resistance:    ptr.StringOrNil(in.Resistance),
		MotorMounting: ptr.StringOrNil(in.MotorMounting),
		MotorType:     ptr.StringOrNil(in.MotorType),
		Voltage:       ptr.StringOrNil(in.Voltage),
		Notes:         ptr.StringOrNil(in.Notes),
	}, nil
}

func CreateBlowerMotor(
	ctx context.Context,
	q Queries,
	partNo string,
	in *BlowerMotorInput,
) (*BlowerMotorRow, error) {

	params, err := BuildCreateBlowerMotorParams(partNo, in)
	if err != nil {
		return nil, err
	}

	row, err := q.CreateBlowerMotor(ctx, params)
	if err != nil {
		return nil, ErrCreateBlowerMotorFailed
	}

	return row, nil
}

func UpdateBlowerMotor(
	ctx context.Context,
	q Queries,
	partNo string,
	in *BlowerMotorInput,
) (*BlowerMotorRow, error) {

	params, err := BuildUpdateBlowerMotorParams(partNo, in)
	if err != nil {
		return nil, err
	}

	row, err := q.UpdateBlowerMotorByPartNo(ctx, params)
	if err != nil {
		return nil, ErrUpdateBlowerMotorFailed
	}

	return row, nil
}

func DeleteBlowerMotor(
	ctx context.Context,
	q Queries,
	partNo string,
) error {

	if partNo == "" {
		return ErrInvalidPartNo
	}

	if err := q.DeleteBlowerMotorByPartNo(ctx, partNo); err != nil {
		return ErrDeleteBlowerMotorFailed
	}

	return nil
}
