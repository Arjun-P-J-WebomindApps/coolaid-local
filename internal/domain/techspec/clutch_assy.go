package techspec

import (
	"context"

	"github.com/google/uuid"
	"github.com/webomindapps-dev/coolaid-backend/internal/shared/ptr"
)

func ValidateClutchAssyInput(in *ClutchAssyInput) error {
	if in == nil {
		return ErrInvalidTechSpec
	}
	return nil
}

func BuildCreateClutchAssyParams(partNo string, in *ClutchAssyInput) (CreateClutchAssyParams, error) {
	if partNo == "" {
		return CreateClutchAssyParams{}, ErrInvalidPartNo
	}
	if err := ValidateClutchAssyInput(in); err != nil {
		return CreateClutchAssyParams{}, err
	}

	return CreateClutchAssyParams{
		ID:                uuid.NewString(),
		PartNo:            partNo,
		PulleyRibs:        ptr.StringOrNil(in.PulleyRibs),
		PulleySize:        ptr.StringOrNil(in.PulleySize),
		CompressorDetails: ptr.StringOrNil(in.CompressorDetails),
		ConnectorType:     ptr.StringOrNil(in.ConnectorType),
		Voltage:           ptr.StringOrNil(in.Voltage),
		ShaftType:         ptr.StringOrNil(in.ShaftType),
		Notes:             ptr.StringOrNil(in.Notes),
	}, nil
}

func BuildUpdateClutchAssyParams(partNo string, in *ClutchAssyInput) (UpdateClutchAssyParams, error) {
	if partNo == "" {
		return UpdateClutchAssyParams{}, ErrInvalidPartNo
	}
	if err := ValidateClutchAssyInput(in); err != nil {
		return UpdateClutchAssyParams{}, err
	}

	return UpdateClutchAssyParams{
		PartNo:            partNo,
		PulleyRibs:        ptr.StringOrNil(in.PulleyRibs),
		PulleySize:        ptr.StringOrNil(in.PulleySize),
		CompressorDetails: ptr.StringOrNil(in.CompressorDetails),
		ConnectorType:     ptr.StringOrNil(in.ConnectorType),
		Voltage:           ptr.StringOrNil(in.Voltage),
		ShaftType:         ptr.StringOrNil(in.ShaftType),
		Notes:             ptr.StringOrNil(in.Notes),
	}, nil
}

func CreateClutchAssy(ctx context.Context, q Queries, partNo string, in *ClutchAssyInput) (*ClutchAssyRow, error) {
	p, err := BuildCreateClutchAssyParams(partNo, in)
	if err != nil {
		return nil, err
	}
	row, err := q.CreateClutchAssy(ctx, p)
	if err != nil {
		return nil, ErrCreateClutchAssyFailed
	}
	return row, nil
}

func UpdateClutchAssy(ctx context.Context, q Queries, partNo string, in *ClutchAssyInput) (*ClutchAssyRow, error) {
	p, err := BuildUpdateClutchAssyParams(partNo, in)
	if err != nil {
		return nil, err
	}
	row, err := q.UpdateClutchAssyByPartNo(ctx, p)
	if err != nil {
		return nil, ErrUpdateClutchAssyFailed
	}
	return row, nil
}

func DeleteClutchAssy(ctx context.Context, q Queries, partNo string) error {
	if err := q.DeleteClutchAssyByPartNo(ctx, partNo); err != nil {
		return ErrDeleteClutchAssyFailed
	}
	return nil
}
