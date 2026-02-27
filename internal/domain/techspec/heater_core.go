package techspec

import (
	"context"

	"github.com/google/uuid"
	"github.com/webomindapps-dev/coolaid-backend/internal/shared/ptr"
)

func (s *Service) GetHeaterCore(ctx context.Context, partNo string) (*HeaterCoreRow, error) {
	if partNo == "" {
		return nil, ErrInvalidPartNo
	}
	return s.DB.Queries().GetHeaterCoreByPartNo(ctx, partNo)
}

func ValidateHeaterCoreInput(in *HeaterCoreInput) error {
	if in == nil {
		return ErrInvalidTechSpec
	}
	return nil
}

func BuildCreateHeaterCoreParams(partNo string, in *HeaterCoreInput) (CreateHeaterCoreParams, error) {
	if partNo == "" {
		return CreateHeaterCoreParams{}, ErrInvalidPartNo
	}
	if err := ValidateHeaterCoreInput(in); err != nil {
		return CreateHeaterCoreParams{}, err
	}

	return CreateHeaterCoreParams{
		ID:     uuid.NewString(),
		PartNo: partNo,
		Size:   ptr.StringOrNil(in.Size),
		Pipe:   ptr.StringOrNil(in.Pipe),
		Type:   ptr.StringOrNil(in.Type),
		Notes:  ptr.StringOrNil(in.Notes),
	}, nil
}

func BuildUpdateHeaterCoreParams(partNo string, in *HeaterCoreInput) (UpdateHeaterCoreParams, error) {
	if partNo == "" {
		return UpdateHeaterCoreParams{}, ErrInvalidPartNo
	}
	if err := ValidateHeaterCoreInput(in); err != nil {
		return UpdateHeaterCoreParams{}, err
	}

	return UpdateHeaterCoreParams{
		PartNo: partNo,
		Size:   ptr.StringOrNil(in.Size),
		Pipe:   ptr.StringOrNil(in.Pipe),
		Type:   ptr.StringOrNil(in.Type),
		Notes:  ptr.StringOrNil(in.Notes),
	}, nil
}

func CreateHeaterCore(ctx context.Context, q Queries, partNo string, in *HeaterCoreInput) (*HeaterCoreRow, error) {
	p, err := BuildCreateHeaterCoreParams(partNo, in)
	if err != nil {
		return nil, err
	}
	row, err := q.CreateHeaterCore(ctx, p)
	if err != nil {
		return nil, ErrCreateHeaterCoreFailed
	}
	return row, nil
}

func UpdateHeaterCore(ctx context.Context, q Queries, partNo string, in *HeaterCoreInput) (*HeaterCoreRow, error) {
	p, err := BuildUpdateHeaterCoreParams(partNo, in)
	if err != nil {
		return nil, err
	}
	row, err := q.UpdateHeaterCoreByPartNo(ctx, p)
	if err != nil {
		return nil, ErrUpdateHeaterCoreFailed
	}
	return row, nil
}

func DeleteHeaterCore(ctx context.Context, q Queries, partNo string) error {
	if partNo == "" {
		return ErrInvalidPartNo
	}
	if err := q.DeleteHeaterCoreByPartNo(ctx, partNo); err != nil {
		return ErrDeleteHeaterCoreFailed
	}
	return nil
}
