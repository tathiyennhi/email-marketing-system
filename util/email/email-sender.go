package email

import (
	"fmt"
	"golang-email-marketing-system/model"
	"net/smtp"
)

func SendEmails(users []model.UserJSON) error {
	for _, user := range users {
		if err := SendEmail(user); err != nil {
			return err
		}
	}
	return nil
}

func SendEmail(user model.UserJSON) error {
	// Sender data.
	from, password := GetENV()

	// Receiver email address.
	to := []string{
		user.To,
	}

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Message.
	mailTemplate := fmt.Sprintf("MIME-version: 1.0;\nContent-Type: %s; charset=\"UTF-8\";\r\n", user.MimeType)
	mailTemplate = fmt.Sprintf(
		"From: %s\r\nTo: %s\r\nSubject: %s\r\n \r\n%s\r\n",
		user.From, user.To, user.Subject, user.Body,
	)

	message := []byte(mailTemplate)

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		return err
	}

	fmt.Printf("Email Sent: %s  Successfully!\n", user.To)
	return nil
}
