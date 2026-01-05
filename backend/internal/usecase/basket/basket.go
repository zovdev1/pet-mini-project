package basket

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/zovdev1/mini-app-project/internal/entites"
	"github.com/zovdev1/mini-app-project/internal/repo"
)

type BasketUseCase struct {
	repo repo.BasketUseRepo
	time repo.BasketTimeUseRepo
}

func New(repo repo.BasketUseRepo, time repo.BasketTimeUseRepo) *BasketUseCase {
	return &BasketUseCase{
		repo: repo,
		time: time,
	}
}

func (r *BasketUseCase) BasketAdd(idUser uuid.UUID, Product_id uuid.UUID) error {

	ctx := context.Background()

	newTime := time.Now()

	basket, err := r.repo.Find(idUser)

	if errors.Is(err, sql.ErrNoRows) {
		CreateBasket := entites.Basket{
			ID:        uuid.Must(uuid.NewV7()),
			User_id:   idUser,
			CreatedAt: newTime,
			UpdatedAt: newTime,
		}

		err := r.repo.Add(ctx, CreateBasket)

		return err

		// basket = &CreateBasket
	}

	item, err := r.time.FindByItem(basket.ID, Product_id)

	if errors.Is(err, sql.ErrNoRows) {
		newItem := entites.BasketItem{
			ID:         uuid.Must(uuid.NewV7()),
			Product_id: Product_id,
			Basket_id:  basket.ID,
			Quantity:   1,
			Price:      100,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		}
		// if err := r.time.Create(newItem); err != nil {
		// 	return err

		_, err := r.time.Create(newItem)

		if err != nil {
			return err
		}
	}

	item.Quantity += 1
	item.UpdatedAt = newTime
	_, err = r.time.Update(ctx, *item)
	return err
}

func (r *BasketUseCase) BasketGet(idUser uuid.UUID) ([]*entites.BasketItem, error) {
	basket, err := r.time.GetAll()

	if err != nil {
		return nil, err
	}

	return basket, err
}
