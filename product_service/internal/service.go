package internal

import (
	"errors"
	"github.com/AminN77/upera_test/product_service/internal/event"
	"log"
)

var (
	ErrRepo      = errors.New("some error with repository")
	ErrPublisher = errors.New("some error with event pulisher")
)

// Service is the aggregator of the internal(domain) layer
type Service interface {
	Add(p *Product) error
}

type service struct {
	repo      Repository
	publisher event.Publisher
}

func NewService(repo Repository, publisher event.Publisher) Service {
	return &service{
		repo:      repo,
		publisher: publisher,
	}
}

func (s *service) Add(p *Product) error {
	if err := s.repo.Add(p); err != nil {
		log.Println(err)
		return ErrRepo
	}

	e := event.NewProductCreatedEvent(p)

	if err := s.publisher.PublishCreatedEvent(e); err != nil {
		log.Println(err)
		return ErrPublisher
	}

	return nil
}
