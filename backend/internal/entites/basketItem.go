package entites

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type BasketItem struct {
	ID         uuid.UUID `db:"id" json:"id"`
	Basket_id  uuid.UUID `db:"basket_id" json:"basket_id"`
	Product_id uuid.UUID `db:"product_id" json:"product_id"`
	Quantity   int       `db:"quantity" json:"quantity"`
	Price      int       `db:"price" json:"price"`
	CreatedAt  time.Time `db:"created_at" json:"created_at"`
	UpdatedAt  time.Time `db:"updated_at" json:"updated_at"`
}

func (b *BasketItem) Validate() error {
	if b.ID == uuid.Nil {
		return errors.New("basketItem ID is required")
	}
	if b.Basket_id == uuid.Nil {
		return errors.New("basketItem basket_id is required")
	}
	if b.Product_id == uuid.Nil {
		return errors.New("basketItem product_id is required")
	}
	if b.Price < 0 {
		return errors.New("basketItem price is required")
	}
	if b.CreatedAt.IsZero() {
		return errors.New("basketItem CreatedAt is required")
	}
	if b.UpdatedAt.IsZero() {
		return errors.New("basketItem UpdatedAt is required")
	}
	return nil
}
