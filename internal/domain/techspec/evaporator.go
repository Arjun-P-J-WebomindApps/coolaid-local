package techspec

import (
	"context"

	"github.com/google/uuid"
	"github.com/webomindapps-dev/coolaid-backend/internal/shared/ptr"
)

func (s *Service) GetEvaporator(ctx context.Context, partNo string) (*EvaporatorRow, error) {
	if partNo == "" {
		return nil, ErrInvalidPartNo
	}
	return s.DB.Queries().GetEvaporatorByPartNo(ctx, partNo)
}

func ValidateEvaporatorInput(in *EvaporatorInput) error {
	if in == nil {
		return ErrInvalidTechSpec
	}
	return nil
}

func BuildCreateEvaporatorParams(partNo string, in *EvaporatorInput) (CreateEvaporatorParams, error) {
	if partNo == "" {
		return CreateEvaporatorParams{}, ErrInvalidPartNo
	}
	if err := ValidateEvaporatorInput(in); err != nil {
		return CreateEvaporatorParams{}, err
	}

	return CreateEvaporatorParams{
		ID:             uuid.NewString(),
		PartNo:         partNo,
		Mounting:       ptr.StringOrNil(in.Mounting),
		ExpValve:       ptr.StringOrNil(in.ExpValve),
		AdditionalInfo: ptr.StringOrNil(in.AdditionalInfo),
		Dimensions:     ptr.StringOrNil(in.Dimensions),
		PipeConnector:  ptr.StringOrNil(in.PipeConnector),
		Notes:          ptr.StringOrNil(in.Notes),
	}, nil
}

func BuildUpdateEvaporatorParams(partNo string, in *EvaporatorInput) (UpdateEvaporatorParams, error) {
	if partNo == "" {
		return UpdateEvaporatorParams{}, ErrInvalidPartNo
	}
	if err := ValidateEvaporatorInput(in); err != nil {
		return UpdateEvaporatorParams{}, err
	}

	return UpdateEvaporatorParams{
		PartNo:         partNo,
		Mounting:       ptr.StringOrNil(in.Mounting),
		ExpValve:       ptr.StringOrNil(in.ExpValve),
		AdditionalInfo: ptr.StringOrNil(in.AdditionalInfo),
		Dimensions:     ptr.StringOrNil(in.Dimensions),
		PipeConnector:  ptr.StringOrNil(in.PipeConnector),
		Notes:          ptr.StringOrNil(in.Notes),
	}, nil
}

func CreateEvaporator(ctx context.Context, q Queries, partNo string, in *EvaporatorInput) (*EvaporatorRow, error) {
	p, err := BuildCreateEvaporatorParams(partNo, in)
	if err != nil {
		return nil, err
	}
	row, err := q.CreateEvaporator(ctx, p)
	if err != nil {
		return nil, ErrCreateEvaporatorFailed
	}
	return row, nil
}

func UpdateEvaporator(ctx context.Context, q Queries, partNo string, in *EvaporatorInput) (*EvaporatorRow, error) {
	p, err := BuildUpdateEvaporatorParams(partNo, in)
	if err != nil {
		return nil, err
	}
	row, err := q.UpdateEvaporatorByPartNo(ctx, p)
	if err != nil {
		return nil, ErrUpdateEvaporatorFailed
	}
	return row, nil
}

func DeleteEvaporator(ctx context.Context, q Queries, partNo string) error {
	if partNo == "" {
		return ErrInvalidPartNo
	}
	if err := q.DeleteEvaporatorByPartNo(ctx, partNo); err != nil {
		return ErrDeleteEvaporatorFailed
	}
	return nil
}
