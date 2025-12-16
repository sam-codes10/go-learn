package main

import (
	"context"
	"notifier-service/config"
	"notifier-service/kafka"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	env := os.Getenv("env")
	if env == "" {
		env = "local"
	}

	cfg := config.Start(env)

	otpConsumer := kafka.NewConsumer("otp", "otp-notifier-group", cfg)
	go otpConsumer.Start(context.Background(), cfg)
}
