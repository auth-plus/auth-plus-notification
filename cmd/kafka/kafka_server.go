// Package main is the mais file for starting http server or kafka or any trigger
package main

import (
	"auth-plus-notification/api/kafka"
	"auth-plus-notification/config"
	"fmt"
)

func main() {
	env := config.GetEnv()
	url := fmt.Sprintf("%s:%s", env.Kafka.URL, env.Kafka.Port)
	kafka.CreateTopics(url)
	kafka.Server(url)
}
