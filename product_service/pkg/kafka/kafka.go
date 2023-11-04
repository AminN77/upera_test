package kafka

import (
	"github.com/IBM/sarama"
	"log"
	"os"
	"strings"
)

func NewSyncProducer() sarama.SyncProducer {
	kafkaConn := os.Getenv("KAFKA_URL")

	//
	log.Println(os.Getenv("KAFKA_URL"))
	log.Println(os.Getenv("KAFKA_TOPIC"))
	log.Println(os.Getenv("KAFKA_CREATE_P"))
	//

	if kafkaConn == "" {
		log.Fatal("could not connect to kafka. KAFKA_URL is empty")
	}

	brokersList := strings.Split(kafkaConn, ",")
	kafkaAdminConfig := sarama.NewConfig()
	log.Println("before cluster admin")
	kafkaAdminClient, err := sarama.NewClusterAdmin(brokersList, kafkaAdminConfig)
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

	log.Println("before topic")

	if err := kafkaAdminClient.CreateTopic(topic, topicDetail, false); err != nil {
		if !strings.ContainsAny(strings.ToLower(err.Error()), strings.ToLower("Topic with this name already exists")) {
			log.Fatal(err)
		}
	}

	log.Println("after topic")

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
