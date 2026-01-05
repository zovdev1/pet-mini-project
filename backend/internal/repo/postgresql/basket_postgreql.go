package persistent

import (
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/zovdev1/mini-app-project/internal/entites"
)

type BasketUseRepo struct {
	db *sqlx.DB
}

func NewBasketUseRepo(db *sqlx.DB) *BasketUseRepo {
	return &BasketUseRepo{db: db}
}

func (r *BasketUseRepo) Add(ctx context.Context, basket entites.Basket) error {
	query := `
		INSERT INTO baskets (id, user_id, created_at, updated_at)
		VALUES ($1, $2, $3, $4)
	`

	_, err := r.db.ExecContext(
		ctx,
		query,
		basket.ID,
		basket.User_id,
		basket.CreatedAt,
		basket.UpdatedAt,
	)

	return err
}

func (r *BasketUseRepo) Find(user uuid.UUID) (*entites.Basket, error) {
	query := `SELECT  id, user_id, created_at, updated_at FROM baskets WHERE user_id = $1`

	row := r.db.QueryRow(query, user)

	var input entites.Basket

	err := row.Scan(
		&input.ID,
		&input.User_id,
		&input.CreatedAt,
		&input.UpdatedAt,
	)

	return &input, err
}

// func (r *BasketUseRepo) Create() error {

// }

// func (r *BasketUseRepo) Get() {

// }
