package kafka

import (
	"auth-service/resourceConfig"

	"github.com/segmentio/kafka-go"
)

type Producer struct {
	OTPWriter *kafka.Writer
}

var kafkaProducer *Producer

func NewProducer(cfg resourceConfig.Config, env, topic string) {
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
		OTPWriter: &kafka.Writer{
			Addr:     kafka.TCP(kafkaConfig.Brokers[0]),
			Topic:    topic,
			Balancer: &kafka.Hash{},
		},
	}
}

func GetProducer() *Producer {
	return kafkaProducer
}
