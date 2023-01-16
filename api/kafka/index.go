// Package kafka contain the function to start to consume Kafka topics
package kafka

import (
	"auth-plus-notification/config"
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

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

		msg := fmt.Sprintf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
		logger.Info(msg)
	}

	if err := r.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}
}
