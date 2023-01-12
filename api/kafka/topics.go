// Package kafka contain the function to start to consume Kafka topics
package kafka

import (
	"net"
	"strconv"

	"github.com/segmentio/kafka-go"
)

// TOPICLIST is a array o topic created
var TOPICLIST = [...]string{
	"2FA_EMAIL_CREATED",
	"2FA_PHONE_CREATED",
	"2FA_EMAIL_SENT",
	"2FA_PHONE_SENT",
	"USER_CREATED",
	"ORGANIZATION_CREATED",
}

// CreateTopics is a function to create topics on kafka in case they don't exist
func CreateTopics(url string) {

	conn, err := kafka.Dial("tcp", url)
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	controller, err := conn.Controller()
	if err != nil {
		panic(err.Error())
	}
	var controllerConn *kafka.Conn
	controllerConn, err = kafka.Dial("tcp", net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port)))
	if err != nil {
		panic(err.Error())
	}
	defer controllerConn.Close()

	topicListConfigs := []kafka.TopicConfig{}

	for _, value := range TOPICLIST {
		config := kafka.TopicConfig{
			Topic:             value,
			NumPartitions:     1,
			ReplicationFactor: 1,
		}
		topicListConfigs = append(topicListConfigs, config)
	}

	err = controllerConn.CreateTopics(topicListConfigs...)
	if err != nil {
		panic(err.Error())
	}
}
