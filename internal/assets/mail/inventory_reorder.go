package mail

import "fmt"

// InventoryOrderEmailData holds data for order confirmation emails.
type InventoryOrderEmailData struct {
	CommonEmailData
	OrderID      string
	CustomerName string
	ItemsCount   int
}

// InventoryOrderTemplate builds order confirmation email HTML.
func InventoryOrderTemplate(data InventoryOrderEmailData) string {
	template := `
	<h2>Order Confirmed</h2>
	<p>Hi %s,</p>
	<p>Your order <strong>%s</strong> has been placed successfully.</p>
	<p>Items: %d</p>
	<p>&copy; %d CoolAid</p>
	`

	return fmt.Sprintf(
		template,
		data.CustomerName,
		data.OrderID,
		data.ItemsCount,
		data.Year,
	)
}
