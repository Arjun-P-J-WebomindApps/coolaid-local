package techspec

import (
	"context"

	"github.com/google/uuid"
	"github.com/webomindapps-dev/coolaid-backend/internal/shared/ptr"
)

func (s *Service) GetRotor(ctx context.Context, partNo string) (*RotorRow, error) {
	if partNo == "" {
		return nil, ErrInvalidPartNo
	}
	return s.DB.Queries().GetRotorByPartNo(ctx, partNo)
}

func ValidateRotorInput(in *RotorInput) error {
	if in == nil {
		return ErrInvalidTechSpec
	}
	return nil
}

func BuildCreateRotorParams(partNo string, in *RotorInput) (CreateRotorParams, error) {
	if partNo == "" {
		return CreateRotorParams{}, ErrInvalidPartNo
	}
	if err := ValidateRotorInput(in); err != nil {
		return CreateRotorParams{}, err
	}

	return CreateRotorParams{
		ID:                uuid.NewString(),
		PartNo:            partNo,
		PulleyRibs:        ptr.StringOrNil(in.PulleyRibs),
		PulleySize:        ptr.StringOrNil(in.PulleySize),
		CompressorDetails: ptr.StringOrNil(in.CompressorDetails),
		Notes:             ptr.StringOrNil(in.Notes),
	}, nil
}

func BuildUpdateRotorParams(partNo string, in *RotorInput) (UpdateRotorParams, error) {
	if partNo == "" {
		return UpdateRotorParams{}, ErrInvalidPartNo
	}
	if err := ValidateRotorInput(in); err != nil {
		return UpdateRotorParams{}, err
	}

	return UpdateRotorParams{
		PartNo:            partNo,
		PulleyRibs:        ptr.StringOrNil(in.PulleyRibs),
		PulleySize:        ptr.StringOrNil(in.PulleySize),
		CompressorDetails: ptr.StringOrNil(in.CompressorDetails),
		Notes:             ptr.StringOrNil(in.Notes),
	}, nil
}

func CreateRotor(ctx context.Context, q Queries, partNo string, in *RotorInput) (*RotorRow, error) {
	p, err := BuildCreateRotorParams(partNo, in)
	if err != nil {
		return nil, err
	}
	row, err := q.CreateRotor(ctx, p)
	if err != nil {
		return nil, ErrCreateRotorFailed
	}
	return row, nil
}

func UpdateRotor(ctx context.Context, q Queries, partNo string, in *RotorInput) (*RotorRow, error) {
	p, err := BuildUpdateRotorParams(partNo, in)
	if err != nil {
		return nil, err
	}
	row, err := q.UpdateRotorByPartNo(ctx, p)
	if err != nil {
		return nil, ErrUpdateRotorFailed
	}
	return row, nil
}

func DeleteRotor(ctx context.Context, q Queries, partNo string) error {
	if partNo == "" {
		return ErrInvalidPartNo
	}
	if err := q.DeleteRotorByPartNo(ctx, partNo); err != nil {
		return ErrDeleteRotorFailed
	}
	return nil
}
