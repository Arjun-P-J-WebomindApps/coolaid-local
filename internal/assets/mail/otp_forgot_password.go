package mail

import "fmt"

// ForgotPasswordOTPEmailData holds data for forgot-password OTP emails.
type ForgotPasswordOTPEmailData struct {
	CommonEmailData
	OTP string
}

// ForgotPasswordOTPTemplate builds the forgot-password OTP email HTML.
func ForgotPasswordOTPTemplate(data ForgotPasswordOTPEmailData) string {
	template := `
	<h2>Password Reset</h2>
	<p>Use the OTP below to reset your password:</p>
	<div style="font-size:32px; font-weight:bold;">%s</div>
	<p>This OTP is valid for 15 minutes.</p>
	<p>&copy; %d CoolAid</p>
	`

	return fmt.Sprintf(template, data.OTP, data.Year)
}
