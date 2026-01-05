package repo

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/zovdev1/mini-app-project/internal/entites"
	"github.com/zovdev1/mini-app-project/internal/usecase/dto"
)

type UseRepo interface {
	Create(user entites.User) (uuid.UUID, error)
	FindByUser(user dto.SignInInput) (*entites.User, error)
}

type ProductUseRepo interface {
	Create(input entites.Product) (*entites.Product, error)
	GetAllProduct(limit, offset int) ([]*entites.Product, error)
	GetById(id uuid.UUID) (*entites.Product, error)
	DELETE(userID, productID uuid.UUID) error
}

type RedisUseRepo interface {
	GetProduct(ctx context.Context, productID uuid.UUID) (*entites.Product, error)
	UpdateCache(ctx context.Context, product *entites.Product, expiration time.Duration) error
}

type BasketUseRepo interface {
	Add(ctx context.Context, basket entites.Basket) error
	Find(user uuid.UUID) (*entites.Basket, error)
}

type BasketTimeUseRepo interface {
	Create(basketItem entites.BasketItem) (*entites.BasketItem, error)
	Update(ctx context.Context, item entites.BasketItem) (*entites.BasketItem, error)
	FindByItem(basketId uuid.UUID, productId uuid.UUID) (*entites.BasketItem, error)
	GetAll() ([]*entites.BasketItem, error)
}
