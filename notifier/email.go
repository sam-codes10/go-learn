package notifier

import (
	"encoding/json"
	"log"
	"notifier-service/config"
	"notifier-service/models"

	"github.com/go-mail/mail/v2"
)

func SendEmail(value []byte, config config.EnvConfig, senderMail string) error {
	var msg models.EmailMessage

	err := json.Unmarshal(value, &msg)
	if err != nil {
		log.Fatal("SendEmail (notifier) - Wrong email format, failed to unmarshal")
		return err
	}

	email := mail.NewMessage()

	email.SetHeader("From", senderMail)
	email.SetHeader("To", msg.To)
	email.SetHeader("Subject", msg.Subject)
	email.SetBody("text/plain", msg.Body)

	dialer := mail.NewDialer(
		config.SMTP.SMTPHost,
		config.SMTP.SMTPPort,
		config.SMTP.SMTPUser,
		config.SMTP.SMTPPass,
	)

	return dialer.DialAndSend(email)
}
