package vendor

func mapRowToModel(v *VendorRow) *Vendor {
	return &Vendor{
		ID:          v.ID,
		CompanyName: v.CompanyName,
	}
}
