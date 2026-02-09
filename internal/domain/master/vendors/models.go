package vendor

type Vendor struct {
	ID             string
	CompanyName    string
	VendorContacts []VendorContact
}

type VendorContact struct {
	ContactPerson string
	MobileNumber  string
	EmailID       string
}
