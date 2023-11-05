package internal

import (
	"context"
	"github.com/AminN77/upera_test/history_service/pkg/kafka"
	"github.com/IBM/sarama"
	"github.com/goccy/go-json"
	"github.com/google/uuid"
	"log"
	"os"
	"strconv"
	"sync"
)

type EventSubscriber interface {
	Subscribe() error
	UnSubscribe() error
}

type kafkaEventSubscriber struct {
	ctx    context.Context
	cancel func()
	repo   Repository
}

func NewKafkaEventSubscriber(repo Repository) EventSubscriber {
	ctx, cancel := context.WithCancel(context.Background())
	return &kafkaEventSubscriber{
		ctx:    ctx,
		cancel: cancel,
		repo:   repo,
	}
}

func (kes *kafkaEventSubscriber) Subscribe() error {
	go kes.startAgent()
	return nil
}

func (kes *kafkaEventSubscriber) UnSubscribe() error {
	kes.cancel()
	return nil
}

func (kes *kafkaEventSubscriber) startAgent() {
	consumer := kafka.NewConsumer()

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

	createConsumer, err := consumer.ConsumePartition(topic, int32(createP), sarama.OffsetNewest)
	if err != nil {
		log.Fatal("created consumer ", err)
	}
	defer func() {
		if err := createConsumer.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	updateConsumer, err := consumer.ConsumePartition(topic, int32(updateP), sarama.OffsetNewest)
	if err != nil {
		log.Fatal("updated consumer ", err)
	}
	defer func() {
		if err := updateConsumer.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	var wg sync.WaitGroup

	wg.Add(2)
	go func(w *sync.WaitGroup) {
		defer w.Done()

		for {
			select {
			case <-kes.ctx.Done():
				return

			case newP := <-createConsumer.Messages():
				log.Println("new created event arrived!")
				if err := kes.saveCreatedEvent(newP.Value); err != nil {
					log.Println(err)
				}
			}

		}

	}(&wg)

	go func(w *sync.WaitGroup) {
		defer w.Done()

		for {
			select {
			case <-kes.ctx.Done():
				return

			case revP := <-updateConsumer.Messages():
				log.Println("new updated event arrived!")
				if err := kes.saveUpdatedEvent(revP.Value); err != nil {
					log.Println(err)
				}
			}
		}

	}(&wg)

	wg.Wait()
}

func (kes *kafkaEventSubscriber) saveCreatedEvent(e []byte) error {
	var createEvent ProductCreatedEvent
	if err := json.Unmarshal(e, &createEvent); err != nil {
		log.Fatal(err)
		return err
	}

	r := Revision{
		RevisionNumber: uuid.NewString(),
		ProductID:      createEvent.Data.ID,
		ChangedAttr:    nil,
		PrevValue:      nil,
		NewValue:       createEvent.Data,
		CreatedAt:      createEvent.CreatedAt,
	}

	if err := kes.repo.Insert(&r); err != nil {
		log.Println(err)
	}

	return nil
}

func (kes *kafkaEventSubscriber) saveUpdatedEvent(e []byte) error {
	var updateEvent ProductUpdatedEvent
	if err := json.Unmarshal(e, &updateEvent); err != nil {
		return err
	}

	r := Revision{
		RevisionNumber: uuid.NewString(),
		ProductID:      updateEvent.Before.ID,
		ChangedAttr:    updateEvent.UpdatedAttrs,
		PrevValue:      updateEvent.Before,
		NewValue:       updateEvent.After,
		CreatedAt:      updateEvent.CreatedAt,
	}

	return kes.repo.Insert(&r)
}
