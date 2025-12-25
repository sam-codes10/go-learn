package kafka

import (
	"auth-service/dbops"
	"auth-service/loggerconfig"
	"auth-service/models"
	"context"
	"encoding/json"
	"fmt"

	"github.com/segmentio/kafka-go"
)

type Producer struct {
	Writer *kafka.Writer
}

func NewProducer(cfg dbops.Config, env string) *Producer {
	if env == "" {
		env = "local"
	}

	var kafkaConfig dbops.KafkaConfig
	switch env {
	case "local":
		kafkaConfig = cfg.LocalConfig.Kafka
	default:
		kafkaConfig = cfg.LocalConfig.Kafka
	}

	writer := &kafka.Writer{
		Addr:     kafka.TCP(kafkaConfig.Brokers[0]),
		Topic:    "otp",
		Balancer: &kafka.Hash{},
	}

	return &Producer{Writer: writer}
}

func (p *Producer) SendOTP(to, content, sendType string) error {
	var payload models.ConsumerMessageOTP

	payload.Content = content
	payload.Subject = "OTP for verification"

	switch sendType {
	case "email":
		payload.Email = true
		payload.RecieverMail = to
	case "sms":
		payload.SMS = true
		payload.RecieverContact = to
	default:
		return fmt.Errorf("invalid send type: %s", sendType)
	}

	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	msg := kafka.Message{
		Key:   []byte(sendType),
		Value: data,
	}

	if err := p.Writer.WriteMessages(context.Background(), msg); err != nil {
		loggerconfig.Error("kafka write failed", err)
		return err
	}

	return nil
}
