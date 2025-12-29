package entites

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `db:"id" json:"id"`
	Email     string    `db:"email" json:"email"`
	Password  string    `db:"password" json:"-"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

func (u User) Validate() error {
	if u.ID == uuid.Nil {
		return errors.New("user ID is required")
	}
	if u.Password == "" {
		return errors.New("user Password is required")
	}
	if u.Email == "" {
		return errors.New("user Email is required")
	}
	if u.CreatedAt.IsZero() {
		return errors.New("user CreatedAt is required")
	}
	if u.UpdatedAt.IsZero() {
		return errors.New("user UpdatedAt is required")
	}
	return nil
}
