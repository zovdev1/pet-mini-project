package persistent

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/zovdev1/mini-app-project/internal/entites"
	"github.com/zovdev1/mini-app-project/internal/usecase/dto"
)

type UseRepo struct {
	db *sqlx.DB
}

func NewUseRepo(db *sqlx.DB) *UseRepo {
	return &UseRepo{db: db}
}

func (r *UseRepo) Create(input entites.User) (uuid.UUID, error) {
	query := `INSERT INTO users (id, email, password, created_at, updated_at) 
    VALUES ($1, $2, $3, $4, $5) RETURNING id`

	var userId uuid.UUID

	row := r.db.QueryRow(
		query,
		input.ID,
		input.Email,
		input.Password,
		input.CreatedAt,
		input.UpdatedAt,
	)

	err := row.Scan(&userId)

	if err != nil {
		return uuid.UUID{}, err
	}

	return userId, nil

}

func (r *UseRepo) FindByUser(input dto.SignInInput) (*entites.User, error) {
	query := `SELECT id, email, password, created_at, updated_at FROM users WHERE email = $1`

	row := r.db.QueryRow(query, input.Email)

	var user entites.User

	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found with email: %s", input.Email)
		}
		return nil, fmt.Errorf("failed to find user: %w", err)
	}

	return &user, nil
}
