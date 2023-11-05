package kafka

import (
	"github.com/IBM/sarama"
	"log"
	"os"
	"strings"
)

func NewConsumer() sarama.Consumer {
	kafkaConn := os.Getenv("KAFKA_URL")
	if kafkaConn == "" {
		log.Fatal("could not connect to kafka. KAFKA_URL is empty")
	}

	brokersList := strings.Split(kafkaConn, ",")

	// Create a new Sarama configuration and set it to consume from the specified partition.
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	// Create a new Kafka consumer and connect to the Kafka brokers.
	consumer, err := sarama.NewConsumer(brokersList, config)
	if err != nil {
		log.Fatal(err)
	}

	return consumer
}
