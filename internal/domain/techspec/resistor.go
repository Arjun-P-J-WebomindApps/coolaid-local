package techspec

import (
	"context"

	"github.com/google/uuid"
	"github.com/webomindapps-dev/coolaid-backend/internal/shared/ptr"
)

func ValidateResistorInput(in *ResistorInput) error {
	if in == nil {
		return ErrInvalidTechSpec
	}
	return nil
}

func BuildCreateResistorParams(partNo string, in *ResistorInput) (CreateResistorParams, error) {
	if partNo == "" {
		return CreateResistorParams{}, ErrInvalidPartNo
	}
	if err := ValidateResistorInput(in); err != nil {
		return CreateResistorParams{}, err
	}

	return CreateResistorParams{
		ID:            uuid.NewString(),
		PartNo:        partNo,
		Type:          ptr.StringOrNil(in.Type),
		ConnectorType: ptr.StringOrNil(in.ConnectorType),
		Voltage:       ptr.StringOrNil(in.Voltage),
		Notes:         ptr.StringOrNil(in.Notes),
	}, nil
}

func BuildUpdateResistorParams(partNo string, in *ResistorInput) (UpdateResistorParams, error) {
	if partNo == "" {
		return UpdateResistorParams{}, ErrInvalidPartNo
	}
	if err := ValidateResistorInput(in); err != nil {
		return UpdateResistorParams{}, err
	}

	return UpdateResistorParams{
		PartNo:        partNo,
		Type:          ptr.StringOrNil(in.Type),
		ConnectorType: ptr.StringOrNil(in.ConnectorType),
		Voltage:       ptr.StringOrNil(in.Voltage),
		Notes:         ptr.StringOrNil(in.Notes),
	}, nil
}

func CreateResistor(ctx context.Context, q Queries, partNo string, in *ResistorInput) (*ResistorRow, error) {
	p, err := BuildCreateResistorParams(partNo, in)
	if err != nil {
		return nil, err
	}
	row, err := q.CreateResistor(ctx, p)
	if err != nil {
		return nil, ErrCreateResistorFailed
	}
	return row, nil
}

func UpdateResistor(ctx context.Context, q Queries, partNo string, in *ResistorInput) (*ResistorRow, error) {
	p, err := BuildUpdateResistorParams(partNo, in)
	if err != nil {
		return nil, err
	}
	row, err := q.UpdateResistorByPartNo(ctx, p)
	if err != nil {
		return nil, ErrUpdateResistorFailed
	}
	return row, nil
}

func DeleteResistor(ctx context.Context, q Queries, partNo string) error {
	if partNo == "" {
		return ErrInvalidPartNo
	}
	if err := q.DeleteResistorByPartNo(ctx, partNo); err != nil {
		return ErrDeleteResistorFailed
	}
	return nil
}
