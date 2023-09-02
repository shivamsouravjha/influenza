package services

import (
	"crypto/tls"
	"fmt"

	"github.com/shivamsouravjha/influenza/config"
	"github.com/shivamsouravjha/influenza/constants"
	gomail "gopkg.in/mail.v2"
)

func SendEmail(to, subject string) error {
	m := gomail.NewMessage()
	// Set E-Mail sender
	m.SetAddressHeader("From", "support@influenza.com", "Influenza")

	// Set E-Mail receivers
	m.SetHeader("To", to)

	// Set E-Mail subject
	m.SetHeader("Subject", subject)

	// Set E-Mail body. You can set plain text or html with text/html
	m.SetBody("text/html", constants.VerificationMail)

	// Settings for SMTP server
	d := gomail.NewDialer("smtp.gmail.com", 587, "shivamsouravjha2@gmail.com", config.Get().EmailPassword)

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		panic(err)
	}

	return nil
}
