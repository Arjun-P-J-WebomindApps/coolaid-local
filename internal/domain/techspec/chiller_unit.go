package techspec

import (
	"context"

	"github.com/google/uuid"
	"github.com/webomindapps-dev/coolaid-backend/internal/shared/ptr"
)

func (s *Service) GetChillerUnit(ctx context.Context, partNo string) (*ChillerUnitRow, error) {
	if partNo == "" {
		return nil, ErrInvalidPartNo
	}
	return s.DB.Queries().GetChillerUnitByPartNo(ctx, partNo)
}

func ValidateChillerUnitInput(in *ChillerUnitInput) error {
	if in == nil {
		return ErrInvalidTechSpec
	}
	return nil
}

func BuildCreateChillerUnitParams(partNo string, in *ChillerUnitInput) (CreateChillerUnitParams, error) {
	if partNo == "" {
		return CreateChillerUnitParams{}, ErrInvalidPartNo
	}
	if err := ValidateChillerUnitInput(in); err != nil {
		return CreateChillerUnitParams{}, err
	}

	return CreateChillerUnitParams{
		ID:      uuid.NewString(),
		PartNo:  partNo,
		Type:    ptr.StringOrNil(in.Type),
		Voltage: ptr.StringOrNil(in.Voltage),
		Notes:   ptr.StringOrNil(in.Notes),
	}, nil
}

func BuildUpdateChillerUnitParams(partNo string, in *ChillerUnitInput) (UpdateChillerUnitParams, error) {
	if partNo == "" {
		return UpdateChillerUnitParams{}, ErrInvalidPartNo
	}
	if err := ValidateChillerUnitInput(in); err != nil {
		return UpdateChillerUnitParams{}, err
	}

	return UpdateChillerUnitParams{
		PartNo:  partNo,
		Type:    ptr.StringOrNil(in.Type),
		Voltage: ptr.StringOrNil(in.Voltage),
		Notes:   ptr.StringOrNil(in.Notes),
	}, nil
}

func CreateChillerUnit(ctx context.Context, q Queries, partNo string, in *ChillerUnitInput) (*ChillerUnitRow, error) {
	p, err := BuildCreateChillerUnitParams(partNo, in)
	if err != nil {
		return nil, err
	}
	row, err := q.CreateChillerUnit(ctx, p)
	if err != nil {
		return nil, ErrCreateChillerUnitFailed
	}
	return row, nil
}

func UpdateChillerUnit(ctx context.Context, q Queries, partNo string, in *ChillerUnitInput) (*ChillerUnitRow, error) {
	p, err := BuildUpdateChillerUnitParams(partNo, in)
	if err != nil {
		return nil, err
	}
	row, err := q.UpdateChillerUnitByPartNo(ctx, p)
	if err != nil {
		return nil, ErrUpdateChillerUnitFailed
	}
	return row, nil
}

func DeleteChillerUnit(ctx context.Context, q Queries, partNo string) error {
	if err := q.DeleteChillerUnitByPartNo(ctx, partNo); err != nil {
		return ErrDeleteChillerUnitFailed
	}
	return nil
}
