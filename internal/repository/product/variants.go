package productrepo

import (
	"context"

	"github.com/google/uuid"
	domain "github.com/webomindapps-dev/coolaid-backend/internal/domain/product"
	"github.com/webomindapps-dev/coolaid-backend/internal/generated/sqlc"
	"github.com/webomindapps-dev/coolaid-backend/internal/shared/ptr"
	"github.com/webomindapps-dev/coolaid-backend/internal/shared/sqlnull"
)

func (p *productQueries) GetModelVariantByPartNo(
	ctx context.Context,
	partNo string,
) (*domain.ModelVariantRow, error) {

	row, err := p.q.GetModelVariantByPartNo(ctx, partNo)
	if err != nil {
		return nil, err
	}

	return mapVariant(row), nil
}

func (p *productQueries) CreateModelVariant(
	ctx context.Context,
	params domain.CreateModelVariantParams,
) (*domain.ModelVariantRow, error) {

	id, err := uuid.Parse(params.ID)
	if err != nil {
		return nil, err
	}

	row, err := p.q.CreateModelVariant(ctx, sqlc.CreateModelVariantParams{
		ID:               id,
		PartNo:           params.PartNo,
		Type:             params.Type,
		FuelTypes:        params.FuelTypes,
		TransmissionType: params.TransmissionType,
		PlatformCodes:    params.PlatformCodes,
		Placement:        params.Placement,
		Make:             params.Make,
		Unicode:          params.Unicode,
		YearStart:        sqlnull.Int32(params.YearStart),
		YearEnd:          sqlnull.Int32(params.YearEnd),
		Gen:              sqlnull.String(params.Gen),
		EngineCc:         sqlnull.Float64(params.EngineCc),
		Image1Link:       sqlnull.String(params.Image1Link),
		Image2Link:       sqlnull.String(params.Image2Link),
		Image3Link:       sqlnull.String(params.Image3Link),
		Image4Link:       sqlnull.String(params.Image4Link),
		HsnCode:          sqlnull.String(params.HsnCode),
		Description:      sqlnull.String(params.Description),
		AdditionalInfo:   sqlnull.String(params.AdditionalInfo),
		OemIds:           params.OemIds,
		VendorID:         params.VendorID,
	})
	if err != nil {
		return nil, err
	}

	return mapVariant(row), nil
}

func mapVariant(row sqlc.ModelVariant) *domain.ModelVariantRow {
	return &domain.ModelVariantRow{
		ID:               row.ID.String(),
		PartNo:           row.PartNo,
		Type:             row.Type,
		Gen:              sqlnull.StringPtr(row.Gen),
		FuelTypes:        row.FuelTypes,
		HsnCode:          sqlnull.StringPtr(row.HsnCode),
		EngineCc:         sqlnull.Float64Ptr(row.EngineCc),
		TransmissionType: row.TransmissionType,
		PlatformCodes:    row.PlatformCodes,
		Placement:        ptr.StringPtr(&row.Placement),
		Image1Link:       sqlnull.StringPtr(row.Image1Link),
		Image2Link:       sqlnull.StringPtr(row.Image2Link),
		Image3Link:       sqlnull.StringPtr(row.Image3Link),
		Image4Link:       sqlnull.StringPtr(row.Image4Link),
		Make:             ptr.StringOrNil(row.Make),
		Unicode:          row.Unicode,
		YearStart:        sqlnull.Int32Ptr(row.YearStart),
		YearEnd:          sqlnull.Int32Ptr(row.YearEnd),
		Description:      sqlnull.StringPtr(row.Description),
		AdditionalInfo:   sqlnull.StringPtr(row.AdditionalInfo),
		OemIDs:           row.OemIds,
		VendorIDs:        row.VendorID,
	}
}
