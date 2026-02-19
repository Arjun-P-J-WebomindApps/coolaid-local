package productrepo

import (
	"context"
	"time"

	domain "github.com/webomindapps-dev/coolaid-backend/internal/domain/product"
	"github.com/webomindapps-dev/coolaid-backend/internal/generated/sqlc"
	"github.com/webomindapps-dev/coolaid-backend/internal/shared/ptr"
	"github.com/webomindapps-dev/coolaid-backend/internal/shared/sqlnull"
)

func (p *productQueries) GetProductDetailsList(
	ctx context.Context,
	params domain.ProductsFilterParams,
) ([]*domain.ProductAggregateRow, error) {

	// ---------- Handle Categories ----------
	var categories []string
	if len(params.Categories) > 0 {
		categories = params.Categories
	} else {
		categories = nil
	}

	// ---------- Handle Unicode Categories ----------
	var unicodeCategories []string
	var unicode string

	if len(params.UnicodeCategories) > 0 {
		unicodeCategories = params.UnicodeCategories

		if params.Model != nil {
			unicode = *params.Model
		} else {
			unicode = ""
		}
	} else {
		unicodeCategories = nil
		unicode = ""
	}

	rows, err := p.q.GetFilteredProductsWithAllDetails(
		ctx,
		sqlc.GetFilteredProductsWithAllDetailsParams{
			Companyname:         ptr.String(params.Company),
			Modelname:           ptr.String(params.Model),
			Partno:              ptr.String(params.PartNo),
			Categorylist:        categories,
			Brandlist:           params.Brands,
			Genlist:             params.Gen,
			Fuellist:            params.FuelType,
			Makelist:            params.Mark,
			Unicode:             unicode,
			Unicodecategorylist: unicodeCategories,
		},
	)
	if err != nil {
		return nil, err
	}

	result := make([]*domain.ProductAggregateRow, 0, len(rows))

	for _, r := range rows {

		result = append(result, &domain.ProductAggregateRow{

			// ----------------------------
			// PRODUCT
			// ----------------------------
			Product: domain.ProductResolvedRow{
				PartNo:       r.PartNo,
				CompanyName:  r.CompanyName,
				ModelName:    r.ModelName,
				BrandName:    r.BrandName,
				CategoryName: r.CategoryName,
			},

			// ----------------------------
			// VARIANT
			// ----------------------------
			ModelVariant: domain.ModelVariantRow{
				PartNo:           r.PartNo,
				Type:             sqlnull.StringValueOrEmpty(r.Type),
				Gen:              sqlnull.StringPtr(r.Gen),
				FuelTypes:        r.FuelTypes,
				HsnCode:          sqlnull.StringPtr(r.HsnCode),
				EngineCc:         sqlnull.Float64Ptr(r.EngineCc),
				TransmissionType: r.TransmissionType,
				PlatformCodes:    r.PlatformCodes,
				Placement:        sqlnull.StringPtr(r.Placement),
				Image1Link:       sqlnull.StringPtr(r.Image1Link),
				Image2Link:       sqlnull.StringPtr(r.Image2Link),
				Image3Link:       sqlnull.StringPtr(r.Image3Link),
				Image4Link:       sqlnull.StringPtr(r.Image4Link),
				Make:             sqlnull.StringPtr(r.Make),
				Unicode:          r.Unicode,
				YearStart:        sqlnull.Int32Ptr(r.YearStart),
				YearEnd:          sqlnull.Int32Ptr(r.YearEnd),
				Description:      sqlnull.StringPtr(r.Description),
				AdditionalInfo:   sqlnull.StringPtr(r.AdditionalInfo),
				OemIDs:           r.OemIds,
				VendorIDs:        r.VendorID,
			},

			// ----------------------------
			// PRICING
			// ----------------------------
			Pricing: domain.PricingRow{
				BasicPrice: sqlnull.Float64Value(r.BasicPrice),
				Freight:    sqlnull.Float64Value(r.Freight),
				Gst:        sqlnull.Float64Value(r.Gst),
				Tax:        sqlnull.Float64Value(r.Tax),

				AcWorkshop:    sqlnull.Float64Value(r.AcWorkshop),
				AcWorkshopPer: sqlnull.Float64Value(r.AcWorkshopPer),
				AcWorkshopAmt: sqlnull.Float64Value(r.AcWorkshopAmt),

				MultibrandWorkshop:    sqlnull.Float64Value(r.MultibrandWorkshop),
				MultibrandWorkshopPer: sqlnull.Float64Value(r.MultibrandWorkshopPer),
				MultibrandWorkshopAmt: sqlnull.Float64Value(r.MultibrandWorkshopAmt),

				AutoTrader:    sqlnull.Float64Value(r.AutoTrader),
				AutoTraderPer: sqlnull.Float64Value(r.AutoTraderPer),
				AutoTraderAmt: sqlnull.Float64Value(r.AutoTraderAmt),

				AcTrader:    sqlnull.Float64Value(r.AcTrader),
				AcTraderPer: sqlnull.Float64Value(r.AcTraderPer),
				AcTraderAmt: sqlnull.Float64Value(r.AcTraderAmt),

				OutstationClassA: sqlnull.Float64Value(r.OutstationClassA),
				OutstationNote:   sqlnull.StringPtr(r.OutstationNote),

				OemMrp:                  sqlnull.Float64Value(r.OemMrp),
				UnitMeasure:             sqlnull.StringPtr(r.UnitMeasure),
				MinimumPurchaseQuantity: sqlnull.Int32OrZero(r.MinimumPurchaseQuantity),
			},

			// ----------------------------
			// INVENTORY
			// ----------------------------
			Inventory: domain.InventoryRow{
				PartNo: r.PartNo, // available

				MinimumOrderLevel: sqlnull.Int32OrZero(r.MinimumOrderLevel),
				MaximumOrderLevel: sqlnull.Int32OrZero(r.MaximumOrderLevel),
				QtyInStock:        sqlnull.Int32OrZero(r.QtyInStock),

				Location: sqlnull.StringPtr(r.Location),

				IsFlash:              sqlnull.BoolValue(r.IsFlash),
				IsRequestedForSupply: sqlnull.BoolValue(r.IsRequestedForSupply),

				// Your query returns []string for VendorID
				// but InventoryRow expects *string
				// so take first element if exists
				VendorID: func() *string {
					if len(r.VendorID) > 0 {
						return &r.VendorID[0]
					}
					return nil
				}(),
			},

			// ----------------------------
			// OFFER
			// ----------------------------
			Offer: domain.OfferRow{
				ID:     "",       // not returned in this query
				PartNo: r.PartNo, // available

				IsOfferActive: sqlnull.BoolValue(r.OfferStatus),

				StartDate: func() time.Time {
					if r.OfferStartDate.Valid {
						return r.OfferStartDate.Time
					}
					return time.Time{}
				}(),

				EndDate: func() time.Time {
					if r.OfferEndDate.Valid {
						return r.OfferEndDate.Time
					}
					return time.Time{}
				}(),

				AcTrader:   r.AcTraderPrice,
				MultiBrand: r.MultiBrandPrice,
				Autotrader: r.AutotraderPrice,
				AcWorkshop: r.AcWorkshopPrice,

				CreatedAt: time.Time{}, // not available in this query
				UpdatedAt: time.Time{}, // not available in this query
			},
		})
	}

	return result, nil
}
