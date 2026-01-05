package usecase

import (
	"github.com/google/uuid"
	"github.com/zovdev1/mini-app-project/internal/entites"
	"github.com/zovdev1/mini-app-project/internal/usecase/dto"
)

type User interface {
	LogInUser(user dto.SignInInput) (string, error)
	User(input dto.RegisterInput) (uuid.UUID, error)
}

type Product interface {
	Create(nput dto.ProductInput) (entites.Product, error)
	GetAll(limit, offset int) ([]*entites.Product, error)
	GetById(id string) (*entites.Product, error)
	DELETE(userID, productID string) error
}

type Basket interface {
	BasketAdd(idUser uuid.UUID, Product_id uuid.UUID) error
	BasketGet(idUser uuid.UUID) ([]*entites.BasketItem, error)
}
