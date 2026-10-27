package main

import (
	"log"

	"github.com/wneessen/go-mail"
)

func main() {
	m := mail.NewMsg()

	if err := m.From("admin@foo.com"); err != nil {
		log.Fatal("Failed to set from address", err)
	}

	if err := m.To("guest@foo.com"); err != nil {
		log.Fatal("Failed to set to address", err)
	}

	m.Subject("Test email")
	m.SetBodyString(mail.TypeTextPlain, "This is a test!")

	c, err := mail.NewClient("localhost",
		mail.WithPort(1025),
		mail.WithTLSPolicy(mail.TLSOpportunistic))

	if err != nil {
		log.Fatalf("Failed to create mail client %v", err)
	}

	if err := c.DialAndSend(m); err != nil {
		log.Fatalf("Failed to send email %v", err)
	}
	log.Println("SMTP send succeeded")
}
