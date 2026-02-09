package mailer

import smpt_mailer "github.com/webomindapps-dev/coolaid-backend/pkg/mailer"

// Service implements the auth.Mailer interface.
//
// This is a thin adapter responsible for sending emails.
// It delegates the actual email delivery to the SMTP utility.
//
// This layer must NOT contain:
// - business logic
// - template logic
// - OTP or auth rules
type Service struct{}

// NewService creates a new Mailer service instance.
//
// The mailer service is stateless and safe to reuse
// across multiple requests.
func NewService() *Service {
	return &Service{}
}

// Send sends an email to the specified recipient.
//
// Parameters:
// - to: recipient email address
// - subject: email subject line
// - html: HTML email body
//
// This method delegates the actual email sending
// to the underlying SMTP implementation.
func (s *Service) Send(to, subject, html string) error {
	return smpt_mailer.SendMail(to, subject, html)
}
