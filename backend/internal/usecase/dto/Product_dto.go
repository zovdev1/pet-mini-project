package dto

import "github.com/google/uuid"

type ProductInput struct {
	Title       string
	Description string
	Price       int
	Quantity    int
	UserID      uuid.UUID
}
