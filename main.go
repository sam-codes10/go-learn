package main

import (
	"context"
	"emailer-service/config"
	"emailer-service/kafka"
	"emailer-service/loggerconfig"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	env := os.Getenv("env")
	if env == "" {
		env = "local"
	}

	loggerconfig.InitLogrus()

	cfg := config.Start(env)

	emailConsumer := kafka.NewConsumer("emailer", "emailer-service", cfg)
	go emailConsumer.Start(context.Background(), cfg)
	select {}
}
