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
