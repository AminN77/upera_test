package internal

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string    `json:"name" gorm:"size:256"`
	Description string    `json:"description"`
	Color       string    `json:"color" gorm:"size:128"`
	Price       int       `json:"price"`
	ImageUrl    string    `json:"imageUrl"`
	Token       uuid.UUID `json:"token"`
}
