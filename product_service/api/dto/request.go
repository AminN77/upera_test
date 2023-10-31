package dto

type AddProductRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Color       string `json:"color"`
	Price       int    `json:"price"`
	ImageUrl    string `json:"imageUrl"`
}
