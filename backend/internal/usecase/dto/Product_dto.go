package dto

import "github.com/google/uuid"

type ProductInput struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Price       int       `json:"price"`
	Quantity    int       `json:"quantity"`
	UserID      uuid.UUID `json:"user_id"`
}
