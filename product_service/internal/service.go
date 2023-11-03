package internal

import (
	"errors"
	"github.com/google/uuid"
	"log"
)

var (
	ErrRepo      = errors.New("some error with repository")
	ErrPublisher = errors.New("some error with event publisher")
)

// Service is the aggregator of the internal(domain) layer
type Service interface {
	Add(p *Product) error
	Update(up *Product, id int) error
	Fetch(id int) (*Product, error)
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
	p.Token, _ = uuid.NewUUID()
	createP, err := s.repo.Add(p)
	if err != nil {
		log.Println(err)
		return ErrRepo
	}

	e := NewProductCreatedEvent(createP)

	if err := s.publisher.PublishCreatedEvent(e); err != nil {
		log.Println(err)
		return ErrPublisher
	}

	return nil
}

func (s *service) Update(up *Product, id int) error {
	updatedP, changes, err := s.repo.Update(up, id)
	if err != nil {
		log.Println(err)
		return ErrRepo
	}

	e := NewProductUpdatedEvent(updatedP, changes)

	if err := s.publisher.PublishUpdatedEvent(e); err != nil {
		log.Println(err)
		return ErrPublisher
	}

	return nil
}

func (s *service) Fetch(id int) (*Product, error) {
	p, err := s.repo.Fetch(id)
	if err != nil {
		log.Println(err)
		return nil, ErrRepo
	}

	return p, nil
}