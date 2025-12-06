package main

import (
	"context"
	"fmt"
	"notifier-service/config"
	"notifier-service/kafka"
	"notifier-service/notifier"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	env := os.Getenv("env")
	if env == "" {
		env = "local"
	}
	senderMail := os.Getenv("mail")

	config := config.Start(env)

	consumer := kafka.NewConsumer(config)

	handler := func(key, value []byte) error {
		switch string(key) {
		case "sms":
			return notifier.SendSMS(value)
		case "email":
			return notifier.SendEmail(value, config, senderMail)
		default:
			return fmt.Errorf("unknown key: %s", key)
		}
	}

	consumer.Start(context.Background(), handler)
}
