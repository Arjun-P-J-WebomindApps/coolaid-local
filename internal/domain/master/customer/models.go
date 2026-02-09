package customer

type Customer struct {
	ID                  string
	CustomerCompanyName string
	ContactPerson       string
	Mobile              string
	Type                string
	CustomerDesignation string
	Address             *string
	Flat                *string
	Street              *string
	City                string
	State               string
	Pincode             *string
	PaymentMode         string
}
