package vendor

type CreateVendorInput struct {
	CompanyName    string
	VendorContacts []VendorContactInput
}

type UpdateVendorInput struct {
	CompanyName    string
	VendorContacts []VendorContactInput
}

type VendorContactInput struct {
	VendorContactPerson string
	MobileNumber        string
	EmailID             string
}
