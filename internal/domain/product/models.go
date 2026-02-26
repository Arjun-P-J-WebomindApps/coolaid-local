package product

import "time"

//
// 🔹 Root Product
//

type Product struct {
	ID         ID
	PartNo     string
	CompanyID  string
	ModelID    string
	BrandID    string
	CategoryID string
	IsActive   bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type ProductResolved struct {
	ID           ID
	PartNo       string
	CompanyName  string
	ModelName    string
	BrandName    string
	BrandImage   string
	CategoryName string
}

//
// 🔹 Aggregate Product Details (Read Model)
//

type ProductDetails struct {
	Product      *ProductResolved
	ModelVariant *ModelVariant
	Pricing      *ProductPricing
	Inventory    *ProductInventory
	Offer        *ProductOffer
}

//
// 🔹 Model Variant
//

type ModelVariant struct {
	ID               ID
	PartNo           string
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

	OemNumbers []OEMListing
	Vendors    []VendorListing
}

//
// 🔹 OEM
//

type OEMListing struct {
	ID        ID
	OemNumber string
	Price     float64
}

//
// 🔹 Vendor Listing (NOT master vendor)
//

type VendorListing struct {
	ID           ID
	VendorName   string
	VendorPartNo string
	VendorPrice  float64
}

//
// 🔹 Pricing
//

type ProductPricing struct {
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
	OutstationNote   string

	OemMrp                  float64
	UnitMeasure             string
	MinimumPurchaseQuantity int32

	CreatedAt time.Time
	UpdatedAt time.Time
}

//
// 🔹 Inventory
//

type ProductInventory struct {
	ID                   ID
	PartNo               string
	MinimumOrderLevel    int32
	MaximumOrderLevel    int32
	QtyInStock           int32
	Location             string
	IsFlash              bool
	IsRequestedForSupply bool
	VendorID             string
}

//
// 🔹 Offer
//

type ProductOffer struct {
	ID            ID
	PartNo        string
	IsOfferActive bool
	StartDate     time.Time
	EndDate       time.Time

	AcTrader   []string
	MultiBrand []string
	Autotrader []string
	AcWorkshop []string
}
