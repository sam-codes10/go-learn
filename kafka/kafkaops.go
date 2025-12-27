package kafka

import (
	"auth-service/models"
	"context"
	"encoding/json"
	"time"

	"github.com/segmentio/kafka-go"
)

func (p *Producer) SendEmail(ctx context.Context, msg models.EmailMessage) error {
	payload, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	kafkaMsg := kafka.Message{
		Key:   []byte(msg.SenderEmail), // ensures same user goes to same partition
		Value: payload,
		Time:  time.Now(),
		Topic: p.EmailWriter.Topic,
	}

	return p.EmailWriter.WriteMessages(ctx, kafkaMsg)
}

func (p *Producer) SendNotification(ctx context.Context, msg models.NotificationMessage) error {
	payload, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	kafkaMsg := kafka.Message{
		Key:   []byte(msg.UserId), // ensures same user goes to same partition
		Value: payload,
		Time:  time.Now(),
		Topic: p.NotificationWriter.Topic,
	}

	return p.NotificationWriter.WriteMessages(ctx, kafkaMsg)
}
