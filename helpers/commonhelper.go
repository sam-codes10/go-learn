package helpers

import (
	"auth-service/kafka"
	"auth-service/loggerconfig"
	"auth-service/models"
	"context"
	"crypto/rand"
	"math/big"
)

func GenerateOTP(digitCount int) (string, error) {
	const digits = "0123456789"
	length := digitCount
	b := make([]byte, length)
	for i := 0; i < length; i++ {
		randomIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(digits))))
		if err != nil {
			return "", err
		}
		b[i] = digits[randomIndex.Int64()]
	}
	return string(b), nil
}

func SendEmail(email, content string) error {
	producer := kafka.GetProducer()

	msg := models.EmailMessage{
		SenderEmail: email,
		Content:     content,
	}
	err := producer.SendEmail(context.Background(), msg)
	if err != nil {
		loggerconfig.Error("Failed to send email to kafka with error: ", err)
		return err
	}
	return nil
}

func SendNotification(userId, email, content string) error {
	producer := kafka.GetProducer()
	msg := models.NotificationMessage{
		UserId:  userId,
		Email:   email,
		Content: content,
	}

	err := producer.SendNotification(context.Background(), msg)
	if err != nil {
		loggerconfig.Error("Failed to send notification to kafka with error: ", err)
		return err
	}

	return nil
}
