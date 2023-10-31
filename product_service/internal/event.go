package internal

import (
	"time"
)

type baseEvent struct {
	CreatedAt time.Time
	Data      *Product
}

type ProductCreatedEvent struct {
	baseEvent
}

type ProductUpdatedEvent struct {
	baseEvent
	UpdatedAttrs []string
}

func NewProductCreatedEvent(p *Product) *ProductCreatedEvent {
	return &ProductCreatedEvent{
		baseEvent: baseEvent{
			CreatedAt: time.Now().UTC(),
			Data:      p,
		},
	}
}

func NewProductUpdatedEvent(p *Product, updatedAttrs []string) *ProductUpdatedEvent {
	return &ProductUpdatedEvent{
		baseEvent: baseEvent{
			CreatedAt: time.Now().UTC(),
			Data:      p,
		},
		UpdatedAttrs: updatedAttrs,
	}
}
