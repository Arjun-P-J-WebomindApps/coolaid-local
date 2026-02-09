package vendor

type CreateVendorParams struct {
	ID          string
	CompanyName string
}

type CreateVendorContactParams struct {
	ID            string
	VendorID      string
	ContactPerson string
	MobileNumber  string
	EmailID       string
}
