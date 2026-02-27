package techspec

import (
	"context"

	"github.com/google/uuid"
	"github.com/webomindapps-dev/coolaid-backend/internal/shared/ptr"
)

// =======================================================
// ACTUATOR
// =======================================================

func (s *Service) GetActuator(ctx context.Context, partNo string) (*ActuatorRow, error) {
	if partNo == "" {
		return nil, ErrInvalidPartNo
	}
	return s.DB.Queries().GetActuatorByPartNo(ctx, partNo)
}

// ValidateActuatorInput performs minimal domain validation.
// Keep this intentionally light.
func ValidateActuatorInput(in *ActuatorInput) error {
	if in == nil {
		return ErrInvalidTechSpec
	}
	return nil
}

// BuildCreateActuatorParams maps input → DB create params
func BuildCreateActuatorParams(
	partNo string,
	in *ActuatorInput,
) (CreateActuatorParams, error) {

	if partNo == "" {
		return CreateActuatorParams{}, ErrInvalidPartNo
	}

	if err := ValidateActuatorInput(in); err != nil {
		return CreateActuatorParams{}, err
	}

	return CreateActuatorParams{
		ID:            uuid.NewString(),
		PartNo:        partNo,
		ConnectorType: ptr.StringOrNil(in.ConnectorType),
		Mounting:      ptr.StringOrNil(in.Mounting),
		Voltage:       ptr.StringOrNil(in.Voltage),
		RotationAngle: ptr.StringOrNil(in.RotationAngle),
		Notes:         ptr.StringOrNil(in.Notes),
	}, nil
}

// BuildUpdateActuatorParams maps input → DB update params
func BuildUpdateActuatorParams(
	partNo string,
	in *ActuatorInput,
) (UpdateActuatorParams, error) {

	if partNo == "" {
		return UpdateActuatorParams{}, ErrInvalidPartNo
	}

	if err := ValidateActuatorInput(in); err != nil {
		return UpdateActuatorParams{}, err
	}

	return UpdateActuatorParams{
		PartNo:        partNo,
		ConnectorType: ptr.StringOrNil(in.ConnectorType),
		Mounting:      ptr.StringOrNil(in.Mounting),
		Voltage:       ptr.StringOrNil(in.Voltage),
		RotationAngle: ptr.StringOrNil(in.RotationAngle),
		Notes:         ptr.StringOrNil(in.Notes),
	}, nil
}

// CreateActuator executes the create use-case
func CreateActuator(
	ctx context.Context,
	q Queries,
	partNo string,
	in *ActuatorInput,
) (*ActuatorRow, error) {

	params, err := BuildCreateActuatorParams(partNo, in)
	if err != nil {
		return nil, err
	}

	row, err := q.CreateActuator(ctx, params)
	if err != nil {
		return nil, ErrCreateActuatorFailed
	}

	return row, nil
}

// UpdateActuator executes the update use-case
func UpdateActuator(
	ctx context.Context,
	q Queries,
	partNo string,
	in *ActuatorInput,
) (*ActuatorRow, error) {

	params, err := BuildUpdateActuatorParams(partNo, in)
	if err != nil {
		return nil, err
	}

	row, err := q.UpdateActuatorByPartNo(ctx, params)
	if err != nil {
		return nil, ErrUpdateActuatorFailed
	}

	return row, nil
}

// DeleteActuator deletes an actuator by part number
func DeleteActuator(
	ctx context.Context,
	q Queries,
	partNo string,
) error {

	if partNo == "" {
		return ErrInvalidPartNo
	}

	if err := q.DeleteActuatorByPartNo(ctx, partNo); err != nil {
		return ErrDeleteActuatorFailed
	}

	return nil
}
