package notifier

import (
	"emailer-service/config"
	"emailer-service/constants"
	"emailer-service/models"
	"time"

	"github.com/go-mail/mail/v2"
)

func SendEmail(consumerMessage models.ConsumerMessage, config config.Config) error {

	emailMsg := models.EmailMessage{
		To:      consumerMessage.RecieverMail,
		Subject: consumerMessage.Subject,
		Body:    consumerMessage.Content,
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

	dialer.Timeout = 20*time.Second

	return dialer.DialAndSend(email)
}
