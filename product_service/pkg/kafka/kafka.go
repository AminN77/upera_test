package kafka

import (
	"github.com/IBM/sarama"
	"log"
	"os"
	"strings"
	"time"
)

func NewSyncProducer() sarama.SyncProducer {
	kafkaConn := os.Getenv("KAFKA_URL")

	if kafkaConn == "" {
		log.Fatal("could not connect to kafka. KAFKA_URL is empty")
	}

	brokersList := strings.Split(kafkaConn, ",")
	kafkaAdminConfig := sarama.NewConfig()

	//retry
	var kafkaAdminClient sarama.ClusterAdmin
	var err error
	for i := 1; i <= 10; i++ {
		kafkaAdminClient, err = sarama.NewClusterAdmin(brokersList, kafkaAdminConfig)
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
		log.Fatal(err)
	}

	topicDetail := &sarama.TopicDetail{
		NumPartitions:     2,
		ReplicationFactor: 1,
	}

	topic := os.Getenv("KAFKA_TOPIC")
	if topic == "" {
		log.Fatalln("KAFKA_TOPIC is empty")
	}

	if err := kafkaAdminClient.CreateTopic(topic, topicDetail, false); err != nil {
		if !strings.ContainsAny(strings.ToLower(err.Error()), strings.ToLower("Topic with this name already exists")) {
			log.Fatal(err)
		}
	}

	producerConfig := sarama.NewConfig()
	producerConfig.Producer.Return.Errors = true
	producerConfig.Producer.Return.Successes = true
	producerConfig.Producer.Partitioner = sarama.NewManualPartitioner

	p, err := sarama.NewSyncProducer(brokersList, producerConfig)
	if err != nil {
		log.Fatal(err)
	}

	return p
}
