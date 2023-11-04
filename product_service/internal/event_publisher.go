package internal

import (
	"encoding/json"
	"errors"
	"github.com/AminN77/upera_test/product_service/pkg/kafka"
	"github.com/IBM/sarama"
	"log"
	"os"
	"strconv"
	"time"
)

type EventPublisher interface {
	PublishCreatedEvent(e *ProductCreatedEvent) error
	PublishUpdatedEvent(e *ProductUpdatedEvent) error
}

type kafkaEventPublisher struct {
	topic           string
	createPartition int32
	updatePartition int32
	producer        sarama.SyncProducer
}

var (
	ErrPublish = errors.New("some error occurred during publish")
)

func NewKafkaEventPublisher() EventPublisher {
	producer := kafka.NewSyncProducer()

	topic := os.Getenv("KAFKA_TOPIC")
	if topic == "" {
		log.Fatalln("KAFKA_TOPIC is empty")
	}

	updateP, err := strconv.Atoi(os.Getenv("KAFKA_UPDATE_P"))
	if err != nil {
		log.Fatalln(err)
	}

	createP, err := strconv.Atoi(os.Getenv("KAFKA_CREATE_P"))
	if err != nil {
		log.Fatalln(err)
	}

	return &kafkaEventPublisher{
		topic:           topic,
		createPartition: int32(createP),
		updatePartition: int32(updateP),
		producer:        producer,
	}
}

func (kep *kafkaEventPublisher) PublishCreatedEvent(e *ProductCreatedEvent) error {
	data, err := json.Marshal(e)
	if err != nil {
		log.Println("err marshalling, err:", err.Error())
		return ErrPublish
	}

	msg := sarama.ProducerMessage{
		Topic:     kep.topic,
		Key:       nil,
		Value:     sarama.StringEncoder(data),
		Headers:   nil,
		Metadata:  nil,
		Offset:    0,
		Partition: kep.createPartition,
		Timestamp: time.Time{},
	}

	_, _, err = kep.producer.SendMessage(&msg)
	if err != nil {
		log.Println("err sending, err:", err.Error())
		return ErrPublish
	}

	return nil
}

func (kep *kafkaEventPublisher) PublishUpdatedEvent(e *ProductUpdatedEvent) error {
	data, err := json.Marshal(e)
	if err != nil {
		log.Println("err marshalling, err:", err.Error())
		return ErrPublish
	}

	msg := &sarama.ProducerMessage{
		Topic:     kep.topic,
		Key:       nil,
		Value:     sarama.StringEncoder(data),
		Headers:   nil,
		Metadata:  nil,
		Offset:    0,
		Partition: kep.updatePartition,
		Timestamp: time.Time{},
	}

	_, _, err = kep.producer.SendMessage(msg)
	if err != nil {
		log.Println("err sending, err:", err.Error())
		return ErrPublish
	}

	return nil
}
