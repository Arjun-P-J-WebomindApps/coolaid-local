package techspec

import (
	"context"

	"github.com/google/uuid"
	"github.com/webomindapps-dev/coolaid-backend/internal/shared/ptr"
)

func ValidateExpansionValveInput(in *ExpansionValveInput) error {
	if in == nil {
		return ErrInvalidTechSpec
	}
	return nil
}

func BuildCreateExpansionValveParams(partNo string, in *ExpansionValveInput) (CreateExpansionValveParams, error) {
	if partNo == "" {
		return CreateExpansionValveParams{}, ErrInvalidPartNo
	}
	if err := ValidateExpansionValveInput(in); err != nil {
		return CreateExpansionValveParams{}, err
	}

	return CreateExpansionValveParams{
		ID:          uuid.NewString(),
		PartNo:      partNo,
		Type:        ptr.StringOrNil(in.Type),
		Material:    ptr.StringOrNil(in.Material),
		Refrigerant: ptr.StringOrNil(in.Refrigerant),
		Notes:       ptr.StringOrNil(in.Notes),
	}, nil
}

func BuildUpdateExpansionValveParams(partNo string, in *ExpansionValveInput) (UpdateExpansionValveParams, error) {
	if partNo == "" {
		return UpdateExpansionValveParams{}, ErrInvalidPartNo
	}
	if err := ValidateExpansionValveInput(in); err != nil {
		return UpdateExpansionValveParams{}, err
	}

	return UpdateExpansionValveParams{
		PartNo:      partNo,
		Type:        ptr.StringOrNil(in.Type),
		Material:    ptr.StringOrNil(in.Material),
		Refrigerant: ptr.StringOrNil(in.Refrigerant),
		Notes:       ptr.StringOrNil(in.Notes),
	}, nil
}

func CreateExpansionValve(ctx context.Context, q Queries, partNo string, in *ExpansionValveInput) (*ExpansionValveRow, error) {
	p, err := BuildCreateExpansionValveParams(partNo, in)
	if err != nil {
		return nil, err
	}
	row, err := q.CreateExpansionValve(ctx, p)
	if err != nil {
		return nil, ErrCreateExpansionValveFailed
	}
	return row, nil
}

func UpdateExpansionValve(ctx context.Context, q Queries, partNo string, in *ExpansionValveInput) (*ExpansionValveRow, error) {
	p, err := BuildUpdateExpansionValveParams(partNo, in)
	if err != nil {
		return nil, err
	}
	row, err := q.UpdateExpansionValveByPartNo(ctx, p)
	if err != nil {
		return nil, ErrUpdateExpansionValveFailed
	}
	return row, nil
}

func DeleteExpansionValve(ctx context.Context, q Queries, partNo string) error {
	if partNo == "" {
		return ErrInvalidPartNo
	}
	if err := q.DeleteExpansionValveByPartNo(ctx, partNo); err != nil {
		return ErrDeleteExpansionValveFailed
	}
	return nil
}
