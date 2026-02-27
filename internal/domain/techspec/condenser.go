package techspec

import (
	"context"

	"github.com/google/uuid"
	"github.com/webomindapps-dev/coolaid-backend/internal/shared/ptr"
)

func (s *Service) GetCondenser(ctx context.Context, partNo string) (*CondenserRow, error) {
	if partNo == "" {
		return nil, ErrInvalidPartNo
	}
	return s.DB.Queries().GetCondenserByPartNo(ctx, partNo)
}

func ValidateCondenserInput(in *CondenserInput) error {
	if in == nil {
		return ErrInvalidTechSpec
	}
	return nil
}

func BuildCreateCondenserParams(partNo string, in *CondenserInput) (CreateCondenserParams, error) {
	if partNo == "" {
		return CreateCondenserParams{}, ErrInvalidPartNo
	}
	if err := ValidateCondenserInput(in); err != nil {
		return CreateCondenserParams{}, err
	}

	return CreateCondenserParams{
		ID:             uuid.NewString(),
		PartNo:         partNo,
		Size:           ptr.StringOrNil(in.Size),
		PipeConnector:  ptr.StringOrNil(in.PipeConnector),
		Drier:          ptr.StringOrNil(in.Drier),
		PressureSwitch: ptr.StringOrNil(in.PressureSwitch),
		Notes:          ptr.StringOrNil(in.Notes),
	}, nil
}

func BuildUpdateCondenserParams(partNo string, in *CondenserInput) (UpdateCondenserParams, error) {
	if partNo == "" {
		return UpdateCondenserParams{}, ErrInvalidPartNo
	}
	if err := ValidateCondenserInput(in); err != nil {
		return UpdateCondenserParams{}, err
	}

	return UpdateCondenserParams{
		PartNo:         partNo,
		Size:           ptr.StringOrNil(in.Size),
		PipeConnector:  ptr.StringOrNil(in.PipeConnector),
		Drier:          ptr.StringOrNil(in.Drier),
		PressureSwitch: ptr.StringOrNil(in.PressureSwitch),
		Notes:          ptr.StringOrNil(in.Notes),
	}, nil
}

func CreateCondenser(ctx context.Context, q Queries, partNo string, in *CondenserInput) (*CondenserRow, error) {
	p, err := BuildCreateCondenserParams(partNo, in)
	if err != nil {
		return nil, err
	}

	row, err := q.CreateCondenser(ctx, p)
	if err != nil {
		return nil, ErrCreateCondenserFailed
	}
	return row, nil
}

func UpdateCondenser(ctx context.Context, q Queries, partNo string, in *CondenserInput) (*CondenserRow, error) {
	p, err := BuildUpdateCondenserParams(partNo, in)
	if err != nil {
		return nil, err
	}

	row, err := q.UpdateCondenserByPartNo(ctx, p)
	if err != nil {
		return nil, ErrUpdateCondenserFailed
	}
	return row, nil
}

func DeleteCondenser(ctx context.Context, q Queries, partNo string) error {
	if partNo == "" {
		return ErrInvalidPartNo
	}
	if err := q.DeleteCondenserByPartNo(ctx, partNo); err != nil {
		return ErrDeleteCondenserFailed
	}
	return nil
}
