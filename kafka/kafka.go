// internal/kafka/consumer.go
package kafka

import (
	"context"
	"encoding/json"
	"emailer-service/config"
	"emailer-service/loggerconfig"
	"emailer-service/models"
	"emailer-service/emailer"

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

	loggerconfig.Info("Kafka message: topic=", msg.Topic, " partition=", msg.Partition, " offset=", msg.Offset, " key=msg.Key")

			var consumerMessage models.ConsumerMessage
			err = json.Unmarshal(msg.Value, &consumerMessage)
			if err != nil {
				loggerconfig.Error("kafka-consumer: failed to unmarshal email:", err)
				continue
			}
			err = notifier.SendEmail(consumerMessage, cfg)
			if err!=nil{
				loggerconfig.Error("kafka-consumer: failed to send email : ", err)
		}	
	}
}
