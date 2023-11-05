package kafka

import (
	"github.com/IBM/sarama"
	"log"
	"os"
	"strings"
	"time"
)

func NewConsumer() sarama.Consumer {
	kafkaConn := os.Getenv("KAFKA_URL")
	if kafkaConn == "" {
		log.Fatal("could not connect to kafka. KAFKA_URL is empty")
	}

	brokersList := strings.Split(kafkaConn, ",")

	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	//retry
	var consumer sarama.Consumer
	var err error
	for i := 1; i <= 10; i++ {
		consumer, err = sarama.NewConsumer(brokersList, config)
		if err != nil {
			log.Printf("Try %d could not connect to kafka", i)
			time.Sleep(3 * time.Second)
			continue
		} else {
			log.Printf("Try %d connected to kafka", i)
			break
		}
	}

	if err != nil {
		log.Fatalln(err)
	}

	return consumer
}
