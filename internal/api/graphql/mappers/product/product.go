package product_mapper

import (
	"context"

	"github.com/webomindapps-dev/coolaid-backend/internal/domain/product"
	"github.com/webomindapps-dev/coolaid-backend/internal/generated/graphql/model"
	"github.com/webomindapps-dev/coolaid-backend/internal/shared/ptr"
	"github.com/webomindapps-dev/coolaid-backend/oplog"
)

func ToFilterOrigin(item product.FilterItem) *model.FilterOrigin {
	return &model.FilterOrigin{
		CompanyName:   item.CompanyName,
		CompanyImage:  item.CompanyImage,
		ModelName:     item.ModelName,
		ModelImage:    item.ModelImage,
		CategoryName:  item.CategoryName,
		CategoryImage: item.CategoryImage,
		BrandName:     item.BrandName,
		BrandImage:    item.BrandImage,
		PartNo:        item.PartNo,
	}
}

func MapCreateMainInput(input *model.ProductPrimaryData) product.CreateProductMainInput {
	return product.CreateProductMainInput{
		PartNo:       input.PartNo,
		CompanyName:  input.CompanyName,
		ModelName:    input.ModelName,
		CategoryName: input.CategoryName,
		BrandName:    input.BrandName,
		BaseData: product.CreateModelVariantInput{
			Type:             input.BaseData.Type,
			Gen:              input.BaseData.Gen,
			FuelTypes:        input.BaseData.FuelTypes,
			HsnCode:          input.BaseData.HsnCode,
			EngineCc:         input.BaseData.EngineCc,
			TransmissionType: input.BaseData.TransmissionType,
			PlatformCodes:    input.BaseData.PlatformCodes,
			Placement:        input.BaseData.Placement,
			Image1Link:       ptr.String(input.BaseData.Image1Link),
			Image2Link:       ptr.String(input.BaseData.Image2Link),
			Image3Link:       ptr.String(input.BaseData.Image3Link),
			Image4Link:       ptr.String(input.BaseData.Image4Link),
			Make:             input.BaseData.Make,
			Unicode:          input.BaseData.Unicode,
			YearStart:        input.BaseData.YearStart,
			YearEnd:          input.BaseData.YearEnd,
			Description:      ptr.String(input.BaseData.Description),
			AdditionalInfo:   ptr.String(input.BaseData.AdditionalInfo),
		},
	}
}

func MapUpdateMainInput(input *model.ProductPrimaryData) product.UpdateProductMainInput {
	return product.UpdateProductMainInput{
		PartNo: input.PartNo,
		BaseData: &product.UpdateModelVariantInput{
			Type:             ptr.StringOrNil(input.BaseData.Type),
			Gen:              ptr.StringOrNil(input.BaseData.Gen),
			FuelTypes:        ptr.StringSlice(input.BaseData.FuelTypes),
			HsnCode:          ptr.StringOrNil(input.BaseData.HsnCode),
			EngineCc:         &input.BaseData.EngineCc,
			TransmissionType: ptr.StringSlice(input.BaseData.TransmissionType),
			PlatformCodes:    ptr.StringSlice(input.BaseData.PlatformCodes),
			Placement:        ptr.StringOrNil(input.BaseData.Placement),
			Image1Link:       input.BaseData.Image1Link,
			Image2Link:       input.BaseData.Image2Link,
			Image3Link:       input.BaseData.Image3Link,
			Image4Link:       input.BaseData.Image4Link,
			Make:             &input.BaseData.Make,
			Unicode:          &input.BaseData.Unicode,
			YearStart:        &input.BaseData.YearStart,
			YearEnd:          &input.BaseData.YearEnd,
			Description:      input.BaseData.Description,
			AdditionalInfo:   input.BaseData.AdditionalInfo,
		},
	}
}

func MapFetchedVariant(ctx context.Context, data *product.ModelVariant) *model.ModelVariant {

	oemList := make([]*model.OEMListing, 0, len(data.OemNumbers))

	for _, i := range data.OemNumbers {
		uid, err := i.ID.ToUUID(ctx)
		if err != nil {
			oplog.Error(ctx, "product validation failed",
				"oemNumber=", i.OemNumber,
				"error=", err)
			continue
		}

		oemList = append(oemList, &model.OEMListing{
			ID:        uid,
			OemNumber: i.OemNumber,
			Price:     i.Price,
		})
	}

	vendorList := make([]*model.VendorList, 0, len(data.Vendors))

	for _, i := range data.Vendors {
		vendorList = append(vendorList, &model.VendorList{
			VendorName:   i.VendorName,
			VendorPartNo: i.VendorPartNo,
			VendorPrice:  i.VendorPrice,
		})
	}

	return &model.ModelVariant{
		Type:             data.Type,
		Gen:              data.Gen,
		FuelTypes:        data.FuelTypes,
		HsnCode:          data.HsnCode,
		EngineCc:         data.EngineCc,
		TransmissionType: data.TransmissionType,
		PlatformCodes:    data.PlatformCodes,
		Placement:        data.Placement,
		Image1Link:       &data.Image1Link,
		Image2Link:       &data.Image2Link,
		Image3Link:       &data.Image3Link,
		Image4Link:       &data.Image4Link,
		Make:             data.Make,
		Unicode:          data.Unicode,
		YearStart:        data.YearStart,
		YearEnd:          data.YearEnd,
		AdditionalInfo:   &data.AdditionalInfo,
		Description:      &data.Description,
		OemNumbers:       oemList,
		Vendor:           vendorList,
	}
}
