// Package kafka contain the function to start to consume Kafka topics
package kafka

import (
	"auth-plus-notification/config"
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
	"go.opentelemetry.io/otel"
)

type KafkaHeaderCarrier struct {
	msg *kafka.Message
}

func (c KafkaHeaderCarrier) Get(key string) string {
	for _, h := range c.msg.Headers {
		if h.Key == key {
			return string(h.Value)
		}
	}
	return ""
}

func (c KafkaHeaderCarrier) Set(key string, value string) {
	c.msg.Headers = append(c.msg.Headers, kafka.Header{
		Key:   key,
		Value: []byte(value),
	})
}

func (c KafkaHeaderCarrier) Keys() []string {
	keys := make([]string, len(c.msg.Headers))
	for i, h := range c.msg.Headers {
		keys[i] = h.Key
	}
	return keys
}

// Server for initiate kafka consumer server
func Server(url string) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     []string{url},
		Topic:       TOPICLIST[0],
		GroupTopics: TOPICLIST[:],
	})
	logger := config.GetLogger()
	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			break
		}

		ctx := otel.GetTextMapPropagator().Extract(context.Background(), KafkaHeaderCarrier{msg: &m})
		_, span := otel.Tracer("kafka-consumer").Start(ctx, fmt.Sprintf("consume %s", m.Topic))

		msg := fmt.Sprintf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
		logger.Info(msg)

		span.End()
	}

	if err := r.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}
}
