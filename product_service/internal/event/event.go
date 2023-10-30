package event

import (
	"github.com/AminN77/upera_test/product_service/internal"
	"time"
)

type baseEvent struct {
	CreatedAt time.Time
	Data      *internal.Product
}

type ProductCreatedEvent struct {
	baseEvent
}

type ProductUpdatedEvent struct {
	baseEvent
	UpdatedAttrs []string
}

func NewProductCreatedEvent(p *internal.Product) *ProductCreatedEvent {
	return &ProductCreatedEvent{
		baseEvent: baseEvent{
			CreatedAt: time.Now().UTC(),
			Data:      p,
		},
	}
}

func NewProductUpdatedEvent(p *internal.Product, updatedAttrs []string) *ProductUpdatedEvent {
	return &ProductUpdatedEvent{
		baseEvent: baseEvent{
			CreatedAt: time.Now().UTC(),
			Data:      p,
		},
		UpdatedAttrs: updatedAttrs,
	}
}
