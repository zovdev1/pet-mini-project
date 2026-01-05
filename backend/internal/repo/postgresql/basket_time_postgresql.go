package persistent

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/zovdev1/mini-app-project/internal/entites"
)

type BasketTimeUseRepo struct {
	db *sqlx.DB
}

func NewBasketTimeUseRepo(db *sqlx.DB) *BasketTimeUseRepo {
	return &BasketTimeUseRepo{db: db}
}

func (r *BasketTimeUseRepo) Create(item entites.BasketItem) (*entites.BasketItem, error) {
	query := `
        INSERT INTO basket_items (id, basket_id, product_id, quantity, price, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, $6, $7)`

	var input entites.BasketItem

	row := r.db.QueryRow(
		query,
		item.ID,
		item.Basket_id,
		item.Product_id,
		item.Quantity,
		item.Price,
		item.CreatedAt,
		item.UpdatedAt,
	)

	err := row.Scan(
		&input.ID,
		&input.Basket_id,
		&input.Product_id,
		&input.Quantity,
		&input.Price,
		&input.CreatedAt,
		&input.UpdatedAt,
	)

	return &input, err
}

func (r *BasketTimeUseRepo) FindByItem(basketId uuid.UUID, productId uuid.UUID) (*entites.BasketItem, error) {
	query := `SELECT id, basket_id, product_id, quantity, price, created_at, updated_at FROM basket_items WHERE basket_id = $1 AND product_id = $2`

	item := &entites.BasketItem{}
	err := r.db.QueryRow(query, basketId, productId).Scan(
		&item.ID,
		&item.Basket_id,
		&item.Product_id,
		&item.Quantity,
		&item.Price,
		&item.CreatedAt,
		&item.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (r *BasketTimeUseRepo) GetAll() ([]*entites.BasketItem, error) {
	query := `SELECT id, basket_id, product_id, quantity, price, created_at, updated_at FROM  basket_items`

	var basketItem []*entites.BasketItem

	row, err := r.db.Query(query)

	if err != nil {
		return nil, err
	}

	defer row.Close()

	for row.Next() {
		var basket entites.BasketItem
		err := row.Scan(
			&basket.ID,
			&basket.Basket_id,
			&basket.Product_id,
			&basket.Quantity,
			&basket.Price,
			&basket.CreatedAt,
			&basket.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		basketItem = append(basketItem, &basket)
	}

	if err = row.Err(); err != nil {
		return nil, err
	}

	return basketItem, nil
}

func (r *BasketTimeUseRepo) Update(ctx context.Context, item entites.BasketItem) (*entites.BasketItem, error) {
	query := `UPDATE basket_items SET quantity = $1, updated_at = $2 WHERE id = $3`
	res, err := r.db.ExecContext(ctx, query, item.Quantity, item.UpdatedAt, item.ID)

	if err != nil {
		return nil, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		return nil, sql.ErrNoRows
	}

	return &item, nil
}

func (r *BasketTimeUseRepo) Delete() error {
	return nil
}
