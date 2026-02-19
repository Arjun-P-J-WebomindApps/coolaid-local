package product

import (
	"time"
)

//
// ============================================================
// 🔹 PRODUCT (Main Table)
// ============================================================
//

type ProductRow struct {
	ID         string
	PartNo     string
	CompanyID  string
	ModelID    string
	BrandID    string
	CategoryID string
	IsActive   bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type ProductResolvedRow struct {
	ID           string
	PartNo       string
	CompanyName  string
	ModelName    string
	BrandName    string
	CategoryName string
	IsActive     bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

//
// ============================================================
// 🔹 MODEL VARIANT
// ============================================================
//

type ModelVariantRow struct {
	ID               string
	PartNo           string
	Type             string
	Gen              *string
	FuelTypes        []string
	HsnCode          *string
	EngineCc         *float64
	TransmissionType []string
	PlatformCodes    []string
	Placement        *string

	Image1Link *string
	Image2Link *string
	Image3Link *string
	Image4Link *string

	Make           *string
	Unicode        []string
	YearStart      *int32
	YearEnd        *int32
	Description    *string
	AdditionalInfo *string

	OemIDs    []string
	VendorIDs []string
}

//
// ============================================================
// 🔹 OEM LISTING
// ============================================================
//

type OEMListingRow struct {
	ID        string
	PartNo    string
	OemNumber string
	Price     float64
}

//
// ============================================================
// 🔹 VENDOR LISTING (NOT master vendor)
// ============================================================
//

type VendorListingRow struct {
	ID            string
	ProductPartNo string
	VendorName    string
	VendorPartNo  string
	VendorPrice   float64
}

//
// ============================================================
// 🔹 PRICING
// ============================================================
//

type PricingRow struct {
	ID            string
	ProductPartID string

	BasicPrice float64
	Freight    float64
	Gst        float64
	Tax        float64

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

	OemMrp                  float64
	UnitMeasure             *string
	MinimumPurchaseQuantity int32

	CreatedAt time.Time
	UpdatedAt time.Time
}

//
// ============================================================
// 🔹 INVENTORY
// ============================================================
//

type InventoryRow struct {
	ID                   string
	PartNo               string
	MinimumOrderLevel    int32
	MaximumOrderLevel    int32
	QtyInStock           int32
	Location             *string
	IsFlash              bool
	IsRequestedForSupply bool
	VendorID             *string
	UpdatedAt            time.Time
}

//
// ============================================================
// 🔹 OFFER
// ============================================================
//

type OfferRow struct {
	ID            string
	PartNo        string
	IsOfferActive bool
	StartDate     time.Time
	EndDate       time.Time

	AcTrader   []string
	MultiBrand []string
	Autotrader []string
	AcWorkshop []string

	CreatedAt time.Time
	UpdatedAt time.Time
}

//
// ============================================================
// 🔹 AGGREGATE READ MODEL (JOINED QUERY)
// Used for GetFilteredProducts / GetProductDetails
// ============================================================
//

type ProductAggregateRow struct {
	// Product
	Product ProductResolvedRow

	// Variant
	ModelVariant ModelVariantRow

	// Pricing
	Pricing PricingRow

	// Inventory
	Inventory InventoryRow

	// Offer
	Offer OfferRow
}

//
// ============================================================
// 🔹 SIMILAR PRICING ROW
// ============================================================
//

type Pricing struct {
	BasicPrice float64
	Freight    float64
	Gst        float64
	Tax        float64

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

	OemMrp                  float64
	UnitMeasure             string
	MinimumPurchaseQuantity int32

	OutstationClassA float64
	OutstationNote   string
}

type SimilarPricingRow struct {
	BrandName string
	Pricing   Pricing
}

//
// ============================================================
// 🔹 Filter Page ROW
// ============================================================
//

type FilterItem struct {
	CompanyName   string
	CompanyImage  string
	ModelName     string
	ModelImage    string
	BrandName     string
	BrandImage    string
	CategoryName  string
	CategoryImage string
	PartNo        string
}

// FilterSelection represents grouped filter results.
// Self: rows matching Unicode filter (or all rows if no filter)
// Unicode: rows not matching Unicode filter
type FilterSelection struct {
	Self    []FilterItem `json:"self"`
	Unicode []FilterItem `json:"unicode"`
}

type FilteredRow struct {
	PartNo        string
	CompanyName   string
	CompanyImage  string
	ModelName     string
	ModelImage    string
	BrandName     string
	BrandImage    string
	CategoryName  string
	CategoryImage string
}
