package persistent

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/zovdev1/mini-app-project/internal/entites"
)

type ProductUseRepo struct {
	db *sqlx.DB
}

func NewProductUseRepo(db *sqlx.DB) *ProductUseRepo {
	return &ProductUseRepo{db: db}
}

func (r *ProductUseRepo) Create(input entites.Product) (*entites.Product, error) {

	query := `INSERT INTO products (id, title, description, price, quantity, user_id, created_at, updated_at) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8) 
	RETURNING id, title, description, price, quantity, user_id, created_at, updated_at`

	var product entites.Product

	row := r.db.QueryRow(
		query,
		input.ID,
		input.Title,
		input.Description,
		input.Price,
		input.Quantity,
		input.UserID,
		input.CreatedAt,
		input.UpdatedAt,
	)

	err := row.Scan(
		&product.ID,
		&product.Title,
		&product.Description,
		&product.Price,
		&product.Quantity,
		&product.UserID,
		&product.CreatedAt,
		&product.UpdatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to create product: %w", err)
	}

	return &product, nil
}

func (r *ProductUseRepo) GetAllProduct(limit, offset int) ([]*entites.Product, error) {

	query := `SELECT id, title, description, price, quantity, user_id, created_at, updated_at FROM products ORDER BY created_at DESC LIMIT $1 OFFSET $2`
	var product []*entites.Product

	row, err := r.db.Query(query, limit, offset)

	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}

	defer row.Close()

	for row.Next() {
		var p entites.Product
		err := row.Scan(
			&p.ID,
			&p.Title,
			&p.Description,
			&p.Price,
			&p.Quantity,
			&p.UserID,
			&p.CreatedAt,
			&p.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("scan failed: %w", err)
		}
		product = append(product, &p)
	}

	if err = row.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return product, nil

}

func (r *ProductUseRepo) GetById(id uuid.UUID) (*entites.Product, error) {

	query := `SELECT id, title, description, price, quantity, user_id, created_at, updated_at FROM products WHERE id = $1`

	row := r.db.QueryRow(query, id)

	var product entites.Product

	err := row.Scan(
		&product.ID,
		&product.Title,
		&product.Description,
		&product.Price,
		&product.Quantity,
		&product.UserID,
		&product.CreatedAt,
		&product.UpdatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("scan failed: %w", err)
	}

	return &product, nil
}

func (r *ProductUseRepo) DELETE(userID, productID uuid.UUID) error {
	query := `DELETE FROM products WHERE id = $1 AND user_id = $2`

	result, err := r.db.Exec(query, productID, userID)
	if err != nil {
		return fmt.Errorf("database delete failed: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("could not get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
