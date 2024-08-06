package common

import (
	"bytes"
	"fmt"
	"html/template"
	"log"

	gomail "gopkg.in/mail.v2"
)

var FORGOT_PASSWORD_EMAIL_TEMPLATE = "forgot-password"

type ForgotPasswordData struct {
	Name       string
	Link       string
	OS         string
	Browser    string
	SupportURL string
}

type EmailData struct {
	Data     interface{}
	Template string
	Body     string
	Subject  string
	To       string
}

func SendEmailHandler(emailData EmailData) {
	tmpl, err := template.ParseFiles(fmt.Sprintf("app/templates/%s.html", emailData.Template))
	if err != nil {
		log.Fatalf("Failed to parse template: %v", err)
	}

	var body bytes.Buffer
	if err := tmpl.Execute(&body, emailData.Data); err != nil {
		log.Fatalf("Failed to execute template: %v", err)
	}

	// Set up the email
	m := gomail.NewMessage()
	m.SetHeader("From", "your-email@example.com")
	m.SetHeader("To", emailData.To)
	m.SetHeader("Subject", emailData.Subject)
	m.SetBody("text/html", body.String())

	// Set up the SMTP server
	d := gomail.NewDialer("localhost", 1025, "", "")

	// Send the email
	if err := d.DialAndSend(m); err != nil {
		log.Fatalf("Failed to send email: %v", err)
	}
}
