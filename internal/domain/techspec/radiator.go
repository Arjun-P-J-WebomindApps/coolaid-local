package techspec

import (
	"context"

	"github.com/google/uuid"
	"github.com/webomindapps-dev/coolaid-backend/internal/shared/ptr"
)

func (s *Service) GetRadiator(ctx context.Context, partNo string) (*RadiatorRow, error) {
	if partNo == "" {
		return nil, ErrInvalidPartNo
	}
	return s.DB.Queries().GetRadiatorByPartNo(ctx, partNo)
}
func ValidateRadiatorInput(in *RadiatorInput) error {
	if in == nil {
		return ErrInvalidTechSpec
	}
	return nil
}

func BuildCreateRadiatorParams(partNo string, in *RadiatorInput) (CreateRadiatorParams, error) {
	if partNo == "" {
		return CreateRadiatorParams{}, ErrInvalidPartNo
	}
	if err := ValidateRadiatorInput(in); err != nil {
		return CreateRadiatorParams{}, err
	}

	return CreateRadiatorParams{
		ID:           uuid.NewString(),
		PartNo:       partNo,
		Size:         ptr.StringOrNil(in.Size),
		Transmission: ptr.StringOrNil(in.Transmission),
		TempSensor:   ptr.StringOrNil(in.TempSensor),
		Tank:         ptr.StringOrNil(in.Tank),
		Notes:        ptr.StringOrNil(in.Notes),
	}, nil
}

func BuildUpdateRadiatorParams(partNo string, in *RadiatorInput) (UpdateRadiatorParams, error) {
	if partNo == "" {
		return UpdateRadiatorParams{}, ErrInvalidPartNo
	}
	if err := ValidateRadiatorInput(in); err != nil {
		return UpdateRadiatorParams{}, err
	}

	return UpdateRadiatorParams{
		PartNo:       partNo,
		Size:         ptr.StringOrNil(in.Size),
		Transmission: ptr.StringOrNil(in.Transmission),
		TempSensor:   ptr.StringOrNil(in.TempSensor),
		Tank:         ptr.StringOrNil(in.Tank),
		Notes:        ptr.StringOrNil(in.Notes),
	}, nil
}

func CreateRadiator(ctx context.Context, q Queries, partNo string, in *RadiatorInput) (*RadiatorRow, error) {
	p, err := BuildCreateRadiatorParams(partNo, in)
	if err != nil {
		return nil, err
	}
	row, err := q.CreateRadiator(ctx, p)
	if err != nil {
		return nil, ErrCreateRadiatorFailed
	}
	return row, nil
}

func UpdateRadiator(ctx context.Context, q Queries, partNo string, in *RadiatorInput) (*RadiatorRow, error) {
	p, err := BuildUpdateRadiatorParams(partNo, in)
	if err != nil {
		return nil, err
	}
	row, err := q.UpdateRadiatorByPartNo(ctx, p)
	if err != nil {
		return nil, ErrUpdateRadiatorFailed
	}
	return row, nil
}

func DeleteRadiator(ctx context.Context, q Queries, partNo string) error {
	if partNo == "" {
		return ErrInvalidPartNo
	}
	if err := q.DeleteRadiatorByPartNo(ctx, partNo); err != nil {
		return ErrDeleteRadiatorFailed
	}
	return nil
}
