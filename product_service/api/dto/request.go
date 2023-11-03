package dto

import "github.com/google/uuid"

type AddProductRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Color       string `json:"color"`
	Price       int    `json:"price"`
	ImageUrl    string `json:"imageUrl"`
}

type UpdateProductRequest struct {
	AddProductRequest
	Token uuid.UUID `json:"token"`
}
