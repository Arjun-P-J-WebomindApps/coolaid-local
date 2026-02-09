package vendor

type VendorRow struct {
	ID          string
	CompanyName string
}

type VendorContactRow struct {
	VendorID      string
	ContactPerson string
	MobileNumber  string
	EmailID       string
}

// Used for LEFT JOIN list
type VendorWithContactRow struct {
	VendorID      string
	CompanyName   string
	ContactPerson *string
	MobileNumber  *string
	EmailID       *string
}
