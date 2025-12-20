// internal/kafka/consumer.go
package kafka

import (
	"context"
	"encoding/json"
	"notifier-service/config"
	"notifier-service/loggerconfig"
	"notifier-service/models"
	"notifier-service/notifier"
	"sync"

	// "time"

	"github.com/segmentio/kafka-go"
)

type Consumer struct {
	reader *kafka.Reader
}

func NewConsumer(topic, group string, config config.Config) *Consumer {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  config.Kafka.KafkaBrokers,
		GroupID:  group,
		Topic:    topic,
		MinBytes: 1,
		MaxBytes: 10e6,
	})
	return &Consumer{reader: r}
}

func (c *Consumer) Start(ctx context.Context, cfg config.Config) error {
	defer c.reader.Close()

	for {
		msg, err := c.reader.ReadMessage(ctx)
		if err != nil {
			if ctx.Err() != nil {
				loggerconfig.Error("kafka-consumer: context cancelled, stopping consumer")
				return nil
			}
			loggerconfig.Error("kafka-read error:", err)
			// time.Sleep(time.Second)
			continue
		}

		loggerconfig.Info("Kafka message: topic=%s partition=%d offset=%d key=%s\n", msg.Topic, msg.Partition, msg.Offset, string(msg.Key))

		switch msg.Topic {
		case "otp":
			loggerconfig.Info("Processed OTP message: %s", string(msg.Value))
			var consumerMessageOTP models.ConsumerMessageOTP
			err := json.Unmarshal(msg.Value, &consumerMessageOTP)
			if err != nil {
				loggerconfig.Error("kafka-consumer: failed to unmarshal OTP message:", err)
				continue
			}
			var wg *sync.WaitGroup
			if consumerMessageOTP.SMS {
				wg.Add(1)
				go notifier.SendSMS(json.RawMessage(consumerMessageOTP.Content))
			}
			if consumerMessageOTP.Email {
				wg.Add(1)
				go notifier.SendEmail(json.RawMessage(consumerMessageOTP.Content), cfg)
			}
			wg.Wait()
		default:
			loggerconfig.Info("No specific processing for topic: %s", msg.Topic)
		}
	}
}
