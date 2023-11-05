package internal

import (
	"context"
	"time"
)

type mockRepository struct {
}

func (*mockRepository) Insert(r *Revision) error {
	return nil
}

func (*mockRepository) InsertBatch(r []*Revision) error {
	return nil
}

func (*mockRepository) GetRevisionsOfOneProduct(pageSize, pageIndex,
	productID int64, ctx context.Context) ([]*Revision, error) {
	res := []*Revision{
		{
			RevisionNumber: "1",
			ProductID:      uint(productID),
			CreatedAt:      time.Now().UTC().Add(-1 * time.Hour),
		},
		{
			RevisionNumber: "2",
			ProductID:      uint(productID),
			CreatedAt:      time.Now().UTC().Add(-2 * time.Hour),
		},
		{
			RevisionNumber: "3",
			ProductID:      uint(productID),
			CreatedAt:      time.Now().UTC().Add(-3 * time.Hour),
		},
	}
	return res, nil
}

func (*mockRepository) GetRevisionByRevisionNumber(revisionNumber string,
	ctx context.Context) (*Revision, error) {
	res := &Revision{
		RevisionNumber: revisionNumber,
		CreatedAt:      time.Now().UTC().Add(-1 * time.Hour),
		NewValue:       &Product{},
	}

	return res, nil
}

type errMockRepository struct {
}

func (*errMockRepository) Insert(r *Revision) error {
	return ErrRepo
}

func (*errMockRepository) InsertBatch(r []*Revision) error {
	return ErrRepo
}

func (*errMockRepository) GetRevisionsOfOneProduct(pageSize, pageIndex,
	productID int64, ctx context.Context) ([]*Revision, error) {
	return nil, ErrRepo
}

func (*errMockRepository) GetRevisionByRevisionNumber(revisionNumber string,
	ctx context.Context) (*Revision, error) {

	return nil, ErrRepo
}
