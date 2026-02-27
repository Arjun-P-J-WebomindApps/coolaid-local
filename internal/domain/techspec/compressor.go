package techspec

import (
	"context"

	"github.com/google/uuid"
	"github.com/webomindapps-dev/coolaid-backend/internal/shared/ptr"
)

func (s *Service) GetCompressor(ctx context.Context, partNo string) (*CompressorRow, error) {
	if partNo == "" {
		return nil, ErrInvalidPartNo
	}
	return s.DB.Queries().GetCompressorByPartNo(ctx, partNo)
}

func ValidateCompressorInput(in *CompressorInput) error {
	if in == nil {
		return ErrInvalidTechSpec
	}
	return nil
}

func BuildCreateCompressorParams(partNo string, in *CompressorInput) (CreateCompressorParams, error) {
	if partNo == "" {
		return CreateCompressorParams{}, ErrInvalidPartNo
	}
	if err := ValidateCompressorInput(in); err != nil {
		return CreateCompressorParams{}, err
	}

	return CreateCompressorParams{
		ID:            uuid.NewString(),
		PartNo:        partNo,
		CompressorID:  ptr.StringOrNil(in.CompressorID),
		Oil:           ptr.StringOrNil(in.Oil),
		Refrigerant:   ptr.StringOrNil(in.Refrigerant),
		Voltage:       ptr.StringOrNil(in.Voltage),
		PulleyRibs:    ptr.StringOrNil(in.PulleyRibs),
		PulleySize:    ptr.StringOrNil(in.PulleySize),
		PipeConnector: ptr.StringOrNil(in.PipeConnector),
		CompType:      ptr.StringOrNil(in.CompType),
		CompMounting:  ptr.StringOrNil(in.CompMounting),
		ConnectorType: ptr.StringOrNil(in.ConnectorType),
		Notes:         ptr.StringOrNil(in.Notes),
	}, nil
}

func BuildUpdateCompressorParams(partNo string, in *CompressorInput) (UpdateCompressorParams, error) {
	if partNo == "" {
		return UpdateCompressorParams{}, ErrInvalidPartNo
	}
	if err := ValidateCompressorInput(in); err != nil {
		return UpdateCompressorParams{}, err
	}

	return UpdateCompressorParams{
		PartNo:        partNo,
		CompressorID:  ptr.StringOrNil(in.CompressorID),
		Oil:           ptr.StringOrNil(in.Oil),
		Refrigerant:   ptr.StringOrNil(in.Refrigerant),
		Voltage:       ptr.StringOrNil(in.Voltage),
		PulleyRibs:    ptr.StringOrNil(in.PulleyRibs),
		PulleySize:    ptr.StringOrNil(in.PulleySize),
		PipeConnector: ptr.StringOrNil(in.PipeConnector),
		CompType:      ptr.StringOrNil(in.CompType),
		CompMounting:  ptr.StringOrNil(in.CompMounting),
		ConnectorType: ptr.StringOrNil(in.ConnectorType),
		Notes:         ptr.StringOrNil(in.Notes),
	}, nil
}

func CreateCompressor(ctx context.Context, q Queries, partNo string, in *CompressorInput) (*CompressorRow, error) {
	p, err := BuildCreateCompressorParams(partNo, in)
	if err != nil {
		return nil, err
	}
	row, err := q.CreateCompressor(ctx, p)
	if err != nil {
		return nil, ErrCreateCompressorFailed
	}
	return row, nil
}

func UpdateCompressor(ctx context.Context, q Queries, partNo string, in *CompressorInput) (*CompressorRow, error) {
	p, err := BuildUpdateCompressorParams(partNo, in)
	if err != nil {
		return nil, err
	}
	row, err := q.UpdateCompressorByPartNo(ctx, p)
	if err != nil {
		return nil, ErrUpdateCompressorFailed
	}
	return row, nil
}

func DeleteCompressor(ctx context.Context, q Queries, partNo string) error {
	if partNo == "" {
		return ErrInvalidPartNo
	}
	if err := q.DeleteCompressorByPartNo(ctx, partNo); err != nil {
		return ErrDeleteCompressorFailed
	}
	return nil
}
