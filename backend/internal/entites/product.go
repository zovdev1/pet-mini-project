package entites

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Price       int       `json:"price"`
	Quantity    int       `json:"quantity"`
	UserID      uuid.UUID `json:"user_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (p Product) Validate() error {
	if p.ID == uuid.Nil {
		return errors.New("product ID is required")
	}
	if p.Title == "" {
		return errors.New("product title is required")
	}
	if p.Description == "" {
		return errors.New("product description is required")
	}
	if p.Price < 0 {
		return errors.New("product price is required")
	}
	if p.Quantity < 0 {
		return errors.New("product quantity is required")
	}
	if p.UserID == uuid.Nil {
		return errors.New("product userID is required")
	}
	if p.CreatedAt.IsZero() {
		return errors.New("user CreatedAt is required")
	}
	if p.UpdatedAt.IsZero() {
		return errors.New("user UpdatedAt is required")
	}
	return nil
}
