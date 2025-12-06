// internal/kafka/consumer.go
package kafka

import (
	"context"
	"log"
	"notifier-service/config"
	"time"

	"github.com/segmentio/kafka-go"
)

type Consumer struct {
	reader *kafka.Reader
}

func NewConsumer(config config.EnvConfig) *Consumer {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  config.Kafka.KafkaBrokers,
		GroupID:  config.Kafka.KafkaGroup,
		Topic:    config.Kafka.KafkaTopic,
		MinBytes: 1,
		MaxBytes: 10e6,
	})
	return &Consumer{reader: r}
}

func (c *Consumer) Start(ctx context.Context, handler func(key []byte, value []byte) error) error {
	defer c.reader.Close()

	for {
		m, err := c.reader.ReadMessage(ctx)
		if err != nil {
			if ctx.Err() != nil {
				return nil
			}
			log.Println("kafka-read error:", err)
			time.Sleep(time.Second)
			continue
		}

		log.Printf("Kafka message: topic=%s partition=%d offset=%d key=%s\n",
			m.Topic, m.Partition, m.Offset, string(m.Key))

		if err := handler(m.Key, m.Value); err != nil {
			log.Println("handler error:", err)
		}
	}
}
