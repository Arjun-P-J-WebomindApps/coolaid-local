package techspec

import (
	"context"

	"github.com/google/uuid"
	"github.com/webomindapps-dev/coolaid-backend/internal/shared/ptr"
)

func (s *Service) GetCompressorValve(ctx context.Context, partNo string) (*CompressorValveRow, error) {
	if partNo == "" {
		return nil, ErrInvalidPartNo
	}
	return s.DB.Queries().GetCompressorValveByPartNo(ctx, partNo)
}

func ValidateCompressorValveInput(in *CompressorValveInput) error {
	if in == nil {
		return ErrInvalidTechSpec
	}
	return nil
}

func BuildCreateCompressorValveParams(partNo string, in *CompressorValveInput) (CreateCompressorValveParams, error) {
	if partNo == "" {
		return CreateCompressorValveParams{}, ErrInvalidPartNo
	}
	if err := ValidateCompressorValveInput(in); err != nil {
		return CreateCompressorValveParams{}, err
	}

	return CreateCompressorValveParams{
		ID:                uuid.NewString(),
		PartNo:            partNo,
		Type:              ptr.StringOrNil(in.Type),
		Voltage:           ptr.StringOrNil(in.Voltage),
		ConnectorType:     ptr.StringOrNil(in.ConnectorType),
		CompressorDetails: ptr.StringOrNil(in.CompressorDetails),
		Notes:             ptr.StringOrNil(in.Notes),
	}, nil
}

func BuildUpdateCompressorValveParams(partNo string, in *CompressorValveInput) (UpdateCompressorValveParams, error) {
	if partNo == "" {
		return UpdateCompressorValveParams{}, ErrInvalidPartNo
	}
	if err := ValidateCompressorValveInput(in); err != nil {
		return UpdateCompressorValveParams{}, err
	}

	return UpdateCompressorValveParams{
		PartNo:            partNo,
		Type:              ptr.StringOrNil(in.Type),
		Voltage:           ptr.StringOrNil(in.Voltage),
		ConnectorType:     ptr.StringOrNil(in.ConnectorType),
		CompressorDetails: ptr.StringOrNil(in.CompressorDetails),
		Notes:             ptr.StringOrNil(in.Notes),
	}, nil
}

func CreateCompressorValve(ctx context.Context, q Queries, partNo string, in *CompressorValveInput) (*CompressorValveRow, error) {
	p, err := BuildCreateCompressorValveParams(partNo, in)
	if err != nil {
		return nil, err
	}

	row, err := q.CreateCompressorValve(ctx, p)
	if err != nil {
		return nil, ErrCreateCompressorValveFailed
	}
	return row, nil
}

func UpdateCompressorValve(ctx context.Context, q Queries, partNo string, in *CompressorValveInput) (*CompressorValveRow, error) {
	p, err := BuildUpdateCompressorValveParams(partNo, in)
	if err != nil {
		return nil, err
	}

	row, err := q.UpdateCompressorValveByPartNo(ctx, p)
	if err != nil {
		return nil, ErrUpdateCompressorValveFailed
	}
	return row, nil
}

func DeleteCompressorValve(ctx context.Context, q Queries, partNo string) error {
	if partNo == "" {
		return ErrInvalidPartNo
	}
	if err := q.DeleteCompressorValveByPartNo(ctx, partNo); err != nil {
		return ErrDeleteCompressorValveFailed
	}
	return nil
}
