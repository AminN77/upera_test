package internal

import (
	"errors"
	"log"
)

var (
	ErrRepo      = errors.New("some error with repository")
	ErrPublisher = errors.New("some error with event publisher")
)

// Service is the aggregator of the internal(domain) layer
type Service interface {
	Add(p *Product) error
}

type service struct {
	repo      Repository
	publisher EventPublisher
}

func NewService(repo Repository, publisher EventPublisher) Service {
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

	e := NewProductCreatedEvent(p)

	if err := s.publisher.PublishCreatedEvent(e); err != nil {
		log.Println(err)
		return ErrPublisher
	}

	return nil
}
