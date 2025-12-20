package notifier

import (
	"encoding/json"
	"log"
	"notifier-service/config"
	"notifier-service/constants"
	"notifier-service/models"

	"github.com/go-mail/mail/v2"
)

func SendEmail(value []byte, config config.Config) error {
	var msg models.EmailMessage

	err := json.Unmarshal(value, &msg)
	if err != nil {
		log.Fatal("SendEmail (notifier) - Wrong email format, failed to unmarshal")
		return err
	}

	emailMsg := models.EmailMessage{
		To:      msg.To,
		Subject: msg.Subject,
		Body:    msg.Body,
	}

	email := mail.NewMessage()

	email.SetHeader("From", constants.SenderMail)
	email.SetHeader("To", emailMsg.To)
	email.SetHeader("Subject", emailMsg.Subject)
	email.SetBody("text/plain", emailMsg.Body)

	dialer := mail.NewDialer(
		config.SMTP.SMTPHost,
		config.SMTP.SMTPPort,
		config.SMTP.SMTPUser,
		config.SMTP.SMTPPass,
	)

	return dialer.DialAndSend(email)
}
