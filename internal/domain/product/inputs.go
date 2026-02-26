package product

import techspec "github.com/webomindapps-dev/coolaid-backend/internal/domain/techspec"

//
// ============================================================
// 🔹 CREATE PRODUCT (Aggregate Input)
// ============================================================
//

type CreateProductInput struct {
	Main      CreateProductMainInput
	Pricing   CreateProductPricingInput
	Inventory CreateProductInventoryInput
	Offer     CreateProductOfferInput
	TechSpec  techspec.TechnicalSpecsInput

	OemNumbers []CreateOEMInput
	Vendors    []CreateVendorListingInput
}

type UpdateProductInput struct {
	Main      UpdateProductMainInput
	Pricing   UpdateProductPricingInput
	Inventory UpdateProductInventoryInput
	Offer     UpdateProductOfferInput
	TechSpec  techspec.TechnicalSpecsInput

	OemNumbers []CreateOEMInput
	Vendors    []CreateVendorListingInput
}

type DeleteProductInput struct {
	PartNo string
}

//
// ------------------------------------------------------------
// 🔹 MAIN PRODUCT
// ------------------------------------------------------------
//

type CreateProductMainInput struct {
	PartNo       string
	CompanyName  string
	ModelName    string
	BrandName    string
	CategoryName string

	BaseData CreateModelVariantInput
}

type UpdateProductMainInput struct {
	PartNo string // required to identify product

	CompanyName  *string
	ModelName    *string
	BrandName    *string
	CategoryName *string

	BaseData *UpdateModelVariantInput
}

//
// ------------------------------------------------------------
// 🔹 MODEL VARIANT INPUT
// ------------------------------------------------------------
//

type CreateModelVariantInput struct {
	Type             string
	Gen              string
	FuelTypes        []string
	HsnCode          string
	EngineCc         float64
	TransmissionType []string
	PlatformCodes    []string
	Placement        string

	Image1Link string
	Image2Link string
	Image3Link string
	Image4Link string

	Make           string
	Unicode        []string
	YearStart      int32
	YearEnd        int32
	Description    string
	AdditionalInfo string
}

type UpdateModelVariantInput struct {
	Type             *string
	Gen              *string
	FuelTypes        *[]string
	HsnCode          *string
	EngineCc         *float64
	TransmissionType *[]string
	PlatformCodes    *[]string
	Placement        *string

	Image1Link *string
	Image2Link *string
	Image3Link *string
	Image4Link *string

	Make           *string
	Unicode        *[]string
	YearStart      *int32
	YearEnd        *int32
	Description    *string
	AdditionalInfo *string

	OemNumbers *[]CreateOEMInput
	Vendors    *[]CreateVendorListingInput
}

//
// ------------------------------------------------------------
// 🔹 OEM INPUT
// ------------------------------------------------------------
//

type CreateOEMInput struct {
	OemNumber string
	Price     float64
}

//
// ------------------------------------------------------------
// 🔹 Vendor Listing Input (not master vendor)
// ------------------------------------------------------------
//

type CreateVendorListingInput struct {
	VendorName   string
	VendorPartNo string
	VendorPrice  float64
}

//
// ------------------------------------------------------------
// 🔹 PRICING INPUT
// ------------------------------------------------------------
//

type CreateProductPricingInput struct {
	BasicPrice float64
	Freight    float64
	Gst        float64

	AcWorkshop    float64
	AcWorkshopPer float64
	AcWorkshopAmt float64

	MultibrandWorkshop    float64
	MultibrandWorkshopPer float64
	MultibrandWorkshopAmt float64

	AutoTrader    float64
	AutoTraderPer float64
	AutoTraderAmt float64

	AcTrader    float64
	AcTraderPer float64
	AcTraderAmt float64

	OutstationClassA float64
	OutstationNote   string

	OemMrp                  float64
	UnitMeasure             string
	MinimumPurchaseQuantity int32
}

type UpdateProductPricingInput struct {
	BasicPrice *float64
	Freight    *float64
	Gst        *float64

	AcWorkshop    *float64
	AcWorkshopPer *float64
	AcWorkshopAmt *float64

	MultibrandWorkshop    *float64
	MultibrandWorkshopPer *float64
	MultibrandWorkshopAmt *float64

	AutoTrader    *float64
	AutoTraderPer *float64
	AutoTraderAmt *float64

	AcTrader    *float64
	AcTraderPer *float64
	AcTraderAmt *float64

	OutstationClassA *float64
	OutstationNote   *string

	OemMrp                  *float64
	UnitMeasure             *string
	MinimumPurchaseQuantity *int32
}

//
// ------------------------------------------------------------
// 🔹 INVENTORY INPUT
// ------------------------------------------------------------
//

type CreateProductInventoryInput struct {
	MinimumOrderLevel int32
	MaximumOrderLevel int32
	QtyInStock        int32
	Location          string
	IsFlash           bool
	VendorName        string
}

type UpdateProductInventoryInput struct {
	MinimumOrderLevel *int32
	MaximumOrderLevel *int32
	QtyInStock        *int32
	Location          *string
	IsFlash           *bool
	VendorName        *string
}

//
// ------------------------------------------------------------
// 🔹 OFFER INPUT
// ------------------------------------------------------------
//

type CreateProductOfferInput struct {
	IsOfferActive bool
	StartDate     string // parsed in service
	EndDate       string

	AcTrader   []string
	MultiBrand []string
	Autotrader []string
	AcWorkshop []string
}

type UpdateProductOfferInput struct {
	IsOfferActive *bool
	StartDate     *string
	EndDate       *string

	AcTrader   *[]string
	MultiBrand *[]string
	Autotrader *[]string
	AcWorkshop *[]string
}
