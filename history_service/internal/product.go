package internal

import (
	"github.com/google/uuid"
	"time"
)

type Product struct {
	ID          uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Color       string    `json:"color"`
	Price       int       `json:"price"`
	ImageUrl    string    `json:"imageUrl"`
	Token       uuid.UUID `json:"token"`
}
