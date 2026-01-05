package entites

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Basket struct {
	ID        uuid.UUID `db:"id" json:"id"`
	User_id   uuid.UUID `db:"user_id" json:"user_id"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

func (b *Basket) Validate() error {
	if b.ID == uuid.Nil {
		return errors.New("basket ID is required")
	}
	if b.User_id == uuid.Nil {
		return errors.New("basket user_id is required")
	}
	if b.CreatedAt.IsZero() {
		return errors.New("basketItem CreatedAt is required")
	}
	if b.UpdatedAt.IsZero() {
		return errors.New("basketItem UpdatedAt is required")
	}
	return nil
}
