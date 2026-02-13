package product

import "github.com/webomindapps-dev/coolaid-backend/internal/shared/ptr"

// ------------------------------------------------------------
// 🔹 PRODUCT
// ------------------------------------------------------------

func mapProductRowToModel(r *ProductRow) *Product {
	if r == nil {
		return nil
	}

	return &Product{
		ID:         ID(r.ID),
		PartNo:     r.PartNo,
		CompanyID:  r.CompanyID,
		ModelID:    r.ModelID,
		BrandID:    r.BrandID,
		CategoryID: r.CategoryID,
		IsActive:   r.IsActive,
		CreatedAt:  r.CreatedAt,
		UpdatedAt:  r.UpdatedAt,
	}
}

// ------------------------------------------------------------
// 🔹 MODEL VARIANT
// ------------------------------------------------------------

func mapModelVariantRowToModel(r *ModelVariantRow) *ModelVariant {
	if r == nil {
		return nil
	}

	return &ModelVariant{
		ID:               ID(r.ID),
		PartNo:           r.PartNo,
		Type:             r.Type,
		Gen:              ptr.String(r.Gen),
		FuelTypes:        r.FuelTypes,
		HsnCode:          ptr.String(r.HsnCode),
		EngineCc:         ptr.Float64Value(r.EngineCc),
		TransmissionType: r.TransmissionType,
		PlatformCodes:    r.PlatformCodes,
		Placement:        ptr.String(r.Placement),
		Image1Link:       ptr.String(r.Image1Link),
		Image2Link:       ptr.String(r.Image2Link),
		Image3Link:       ptr.String(r.Image3Link),
		Image4Link:       ptr.String(r.Image4Link),
		Make:             ptr.String(r.Make),
		Unicode:          r.Unicode,
		YearStart:        ptr.Int32Value(r.YearStart),
		YearEnd:          ptr.Int32Value(r.YearEnd),
		Description:      ptr.String(r.Description),
		AdditionalInfo:   ptr.String(r.AdditionalInfo),
	}
}

// ------------------------------------------------------------
// 🔹 PRICING
// ------------------------------------------------------------

func mapPricingRowToModel(r *PricingRow) *ProductPricing {
	if r == nil {
		return nil
	}

	return &ProductPricing{
		ID:                      r.ID,
		ProductPartID:           r.ProductPartID,
		BasicPrice:              r.BasicPrice,
		Freight:                 r.Freight,
		Gst:                     r.Gst,
		Tax:                     r.Tax,
		AcWorkshop:              r.AcWorkshop,
		AcWorkshopPer:           r.AcWorkshopPer,
		AcWorkshopAmt:           r.AcWorkshopAmt,
		MultibrandWorkshop:      r.MultibrandWorkshop,
		MultibrandWorkshopPer:   r.MultibrandWorkshopPer,
		MultibrandWorkshopAmt:   r.MultibrandWorkshopAmt,
		AutoTrader:              r.AutoTrader,
		AutoTraderPer:           r.AutoTraderPer,
		AutoTraderAmt:           r.AutoTraderAmt,
		AcTrader:                r.AcTrader,
		AcTraderPer:             r.AcTraderPer,
		AcTraderAmt:             r.AcTraderAmt,
		OutstationClassA:        r.OutstationClassA,
		OutstationNote:          ptr.String(r.OutstationNote),
		OemMrp:                  r.OemMrp,
		UnitMeasure:             ptr.String(r.UnitMeasure),
		MinimumPurchaseQuantity: r.MinimumPurchaseQuantity,
		CreatedAt:               r.CreatedAt,
		UpdatedAt:               r.UpdatedAt,
	}
}

// ------------------------------------------------------------
// 🔹 INVENTORY
// ------------------------------------------------------------

func mapInventoryRowToModel(r *InventoryRow) *ProductInventory {
	if r == nil {
		return nil
	}

	return &ProductInventory{
		ID:                   ID(r.ID),
		PartNo:               r.PartNo,
		MinimumOrderLevel:    r.MinimumOrderLevel,
		MaximumOrderLevel:    r.MaximumOrderLevel,
		QtyInStock:           r.QtyInStock,
		Location:             ptr.String(r.Location),
		IsFlash:              r.IsFlash,
		IsRequestedForSupply: r.IsRequestedForSupply,
		VendorID:             ptr.String(r.VendorID),
	}
}

// ------------------------------------------------------------
// 🔹 OFFER
// ------------------------------------------------------------

func mapOfferRowToModel(r *OfferRow) *ProductOffer {
	if r == nil {
		return nil
	}

	return &ProductOffer{
		ID:            ID(r.ID),
		PartNo:        r.PartNo,
		IsOfferActive: r.IsOfferActive,
		StartDate:     r.StartDate,
		EndDate:       r.EndDate,
		AcTrader:      r.AcTrader,
		MultiBrand:    r.MultiBrand,
		Autotrader:    r.Autotrader,
		AcWorkshop:    r.AcWorkshop,
	}
}
