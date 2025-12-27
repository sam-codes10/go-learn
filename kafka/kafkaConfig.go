package kafka

import (
	"auth-service/resourceConfig"

	"github.com/segmentio/kafka-go"
)

type Producer struct {
	EmailWriter        *kafka.Writer
	NotificationWriter *kafka.Writer
}

var kafkaProducer *Producer

func NewProducer(cfg resourceConfig.Config, env string) {
	if env == "" {
		env = "local"
	}

	var kafkaConfig resourceConfig.KafkaConfig

	switch env {
	case "local":
		kafkaConfig = cfg.LocalConfig.Kafka
	default:
		kafkaConfig = cfg.LocalConfig.Kafka
	}

	kafkaProducer = &Producer{
		EmailWriter: kafka.NewWriter(kafka.WriterConfig{
			Brokers:  kafkaConfig.Brokers,
			Topic:    "otp",
			Balancer: &kafka.Hash{},
		}),
		NotificationWriter: kafka.NewWriter(kafka.WriterConfig{
			Brokers:  kafkaConfig.Brokers,
			Topic:    "notification",
			Balancer: &kafka.Hash{},
		}),
	}

}

func GetProducer() *Producer {
	return kafkaProducer
}
