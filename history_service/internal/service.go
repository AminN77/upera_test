package internal

import (
	"context"
	"errors"
	"log"
)

var (
	ErrRepo = errors.New("some error occurred on repository")
)

// Service is the aggregator of the internal(domain) layer
type Service interface {
	FetchRevision(revisionNumber string, ctx context.Context) (*Product, error)
	FetchRevisionsOfOneProduct(pageSize, pageIndex, productID int64, ctx context.Context) ([]*Revision, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) FetchRevision(revisionNumber string, ctx context.Context) (*Product, error) {
	rev, err := s.repo.GetRevisionByRevisionNumber(revisionNumber, ctx)
	if err != nil {
		log.Println(err)
		return nil, ErrRepo
	}

	return rev.NewValue, nil
}

func (s *service) FetchRevisionsOfOneProduct(pageSize, pageIndex, productID int64, ctx context.Context) ([]*Revision, error) {
	rev, err := s.repo.GetRevisionsOfOneProduct(pageSize, pageIndex, productID, ctx)
	if err != nil {
		log.Println(err)
		return nil, ErrRepo
	}

	return rev, nil
}
