package product

import "time"

//
// ============================================================
// 🔹 CREATE PRODUCT PARAMS
// ============================================================
//

type CreateProductParams struct {
	ID         string
	CompanyID  string
	ModelID    string
	BrandID    string
	CategoryID string
	PartNo     string
	IsActive   bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

//
// ============================================================
// 🔹 UPDATE PRODUCT PARAMS
// ============================================================
//

type UpdateProductParams struct {
	ID         string
	CompanyID  *string
	ModelID    *string
	BrandID    *string
	CategoryID *string
	PartNo     *string
	IsActive   *bool
	UpdatedAt  *string
}

//
// ============================================================
// 🔹 CREATE MODEL VARIANT PARAMS
// ============================================================
//

type CreateModelVariantParams struct {
	ID               string
	PartNo           string
	FuelTypes        []string
	Gen              *string
	EngineCc         *float64
	TransmissionType []string
	PlatformCodes    []string
	Placement        string
	Image1Link       *string
	Image2Link       *string
	Image3Link       *string
	Image4Link       *string
	HsnCode          *string
	Unicode          []string
	Description      *string
	Make             string
	OemIds           []string
	Type             string
	YearStart        *int32
	YearEnd          *int32
	VendorID         []string
	AdditionalInfo   *string
}

//
// ============================================================
// 🔹 UPDATE MODEL VARIANT PARAMS
// ============================================================
//

type UpdateModelVariantParams struct {
	ID               string
	PartNo           *string
	FuelTypes        *[]string
	Gen              *string
	EngineCc         *float64
	TransmissionType *[]string
	PlatformCodes    *[]string
	Placement        *string
	Image1Link       *string
	Image2Link       *string
	Image3Link       *string
	Image4Link       *string
	HsnCode          *string
	Unicode          *[]string
	Description      *string
	Make             *string
	OemIds           *[]string
	Type             *string
	YearStart        *int32
	YearEnd          *int32
	VendorID         *[]string
	AdditionalInfo   *string
}

//
// ============================================================
// 🔹 CREATE PRODUCT PRICING PARAMS
// ============================================================
//

type CreateProductPricingParams struct {
	ID            string
	ProductPartID string
	BasicPrice    float64
	Freight       float64
	Gst           float64

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
	OutstationNote   *string

	OemMrp      float64
	UnitMeasure *string

	MinimumPurchaseQuantity int32

	CreatedAt time.Time
	UpdatedAt time.Time
}

//
// ============================================================
// 🔹 UPDATE PRODUCT PRICING PARAMS
// ============================================================
//

type UpdateProductPricingParams struct {
	PartNo string

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

	OemMrp      *float64
	UnitMeasure *string

	MinimumPurchaseQuantity *int32

	UpdatedAt *time.Time
}

//
// ============================================================
// 🔹 CREATE PRODUCT INVENTORY PARAMS
// ============================================================
//

type CreateProductInventoryParams struct {
	ID                   string
	PartNo               string
	MinimumOrderLevel    int32
	MaximumOrderLevel    int32
	QtyInStock           int32
	Location             string
	IsFlash              bool
	IsRequestedForSupply bool
	VendorID             *string
}

//
// ============================================================
// 🔹 UPDATE PRODUCT INVENTORY PARAMS
// ============================================================
//

type UpdateProductInventoryParams struct {
	PartNo               string
	MinimumOrderLevel    *int32
	MaximumOrderLevel    *int32
	QtyInStock           *int32
	Location             *string
	IsFlash              *bool
	IsRequestedForSupply *bool
	VendorID             *string
	UpdatedAt            *string
}

//
// ============================================================
// 🔹 CREATE PRODUCT OFFER PARAMS
// ============================================================
//

type CreateProductOfferParams struct {
	ID            string
	PartNo        string
	IsOfferActive bool
	StartDate     string
	EndDate       string
	AcTrader      []string
	MultiBrand    []string
	Autotrader    []string
	AcWorkshop    []string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

//
// ============================================================
// 🔹 UPDATE PRODUCT OFFER PARAMS
// ============================================================
//

type UpdateProductOfferParams struct {
	ID            string
	PartNo        string
	IsOfferActive *bool
	StartDate     *string
	EndDate       *string
	AcTrader      *[]string
	MultiBrand    *[]string
	Autotrader    *[]string
	AcWorkshop    *[]string
	UpdatedAt     *time.Time
}

//
// ============================================================
// 🔹 CREATE PRODUCT OEM PARAMS
// ============================================================
//

type CreateOEMParams struct {
	ID        string
	PartNo    string
	OemNumber string
	Price     float64
}

//
// ============================================================
// 🔹 CREATE PRODUCT VENDOR LISTING PARAMS
// ============================================================
//
type CreateVendorListingParams struct {
	ID            string
	ProductPartNo string
	VendorName    string
	VendorPartNo  string
	VendorPrice   float64
}

type FilterSelectionParams struct {
	Company  *string
	Model    *string
	Brand    *string
	Category *string
	Unicode  *string
}

type SimilarPricingParams struct {
	Company  string
	Model    string
	Category string
	Type     string
}

type ProductsFilterParams struct {
	Company *string `json:"company,omitempty"`
	Model   *string `json:"model,omitempty"`
	PartNo  *string `json:"part_no,omitempty"`

	Categories        []string `json:"categories,omitempty"`
	Brands            []string `json:"brands,omitempty"`
	Gen               []string `json:"gen,omitempty"`
	FuelType          []string `json:"fuel_type,omitempty"`
	Mark              []string `json:"mark,omitempty"`
	UnicodeCategories []string `json:"unicode_categories,omitempty"`
}
