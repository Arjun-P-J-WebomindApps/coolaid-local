package smpt_mailer

import (
	"crypto/tls"
	"fmt"
	"net/smtp"
	"strings"

	"github.com/webomindapps-dev/coolaid-backend/config"
)

func SendMail(toEmail string, subject string, htmlBody string) error {
	emailfrom := "support@ecoolaid.com"
	from := "CoolAid <" + emailfrom + ">"

	headers := map[string]string{
		"From":         from,
		"To":           toEmail,
		"Subject":      subject,
		"Mime-Version": "1.0",
		"Content-Type": "text/html; charset=\"UTF-8\"",
	}

	var msg strings.Builder

	for k, v := range headers {
		msg.WriteString(fmt.Sprintf("%s:%s\r\n", k, v))
	}

	msg.WriteString("\r\n" + htmlBody)

	addr := config.SMTP.Host + ":" + config.SMTP.Port
	auth := smtp.PlainAuth("", config.SMTP.Username, config.SMTP.Password, config.SMTP.Host)

	if config.SMTP.Secure == "tls" {

		conn, err := smtp.Dial(addr)
		if err != nil {
			return err
		}
		defer conn.Close()

		tlsConfig := &tls.Config{
			ServerName: config.SMTP.Host,
		}

		if err = conn.StartTLS(tlsConfig); err != nil {
			return fmt.Errorf("starttls failed %s", err.Error())
		}

		if err = conn.Auth(auth); err != nil {
			return fmt.Errorf("auth failed: %s", err.Error())
		}

		if err = conn.Mail(emailfrom); err != nil {
			return fmt.Errorf("email from is invalid %s", err.Error())
		}

		if err = conn.Rcpt(toEmail); err != nil {
			return fmt.Errorf("to email is invalid %s", err.Error())
		}

		w, err := conn.Data()

		if err != nil {
			return fmt.Errorf("data start failed %s", err.Error())
		}

		_, err = w.Write([]byte(msg.String()))
		if err != nil {
			return fmt.Errorf("write failed %s", err.Error())
		}

		err = w.Close()
		if err != nil {
			return fmt.Errorf("data close failed %s", err.Error())
		}

		// return smtp.SendMail(addr, auth, emailfrom, []string{toEmail}, []byte(msg.String()))

		return conn.Quit()
	}

	if config.SMTP.Secure == "ssl" {
		tlsConfig := &tls.Config{
			InsecureSkipVerify: true,
			ServerName:         config.SMTP.Host,
		}

		conn, err := tls.Dial("tcp", addr, tlsConfig)
		if err != nil {
			return err
		}

		c, err := smtp.NewClient(conn, config.SMTP.Host)
		if err != nil {
			return err
		}

		defer c.Quit()

		if err = c.Auth(auth); err != nil {
			return err
		}

		if err = c.Mail(from); err != nil {
			return err
		}

		if err = c.Rcpt(toEmail); err != nil {
			return err
		}

		w, err := c.Data()

		if err != nil {
			return err
		}

		_, errWrite := w.Write([]byte(msg.String()))

		if errWrite != nil {
			return err
		}

		return w.Close()
	}

	return fmt.Errorf("unsupported smtpSecure %s", config.SMTP.Secure)
}
