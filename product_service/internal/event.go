package internal

import (
	"time"
)

type baseEvent struct {
	CreatedAt time.Time
}

type ProductCreatedEvent struct {
	baseEvent
	Data *Product
}

type ProductUpdatedEvent struct {
	baseEvent
	UpdatedAttrs []string
	Before       *Product
	After        *Product
}

func NewProductCreatedEvent(p *Product) *ProductCreatedEvent {
	return &ProductCreatedEvent{
		baseEvent: baseEvent{
			CreatedAt: time.Now().UTC(),
		},
		Data: p,
	}
}

func NewProductUpdatedEvent(before *Product, after *Product, updatedAttrs []string) *ProductUpdatedEvent {
	return &ProductUpdatedEvent{
		baseEvent: baseEvent{
			CreatedAt: time.Now().UTC(),
		},
		Before:       before,
		After:        after,
		UpdatedAttrs: updatedAttrs,
	}
}
