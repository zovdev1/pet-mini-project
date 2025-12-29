package product

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/zovdev1/mini-app-project/internal/entites"
	"github.com/zovdev1/mini-app-project/internal/repo"
	redisdb "github.com/zovdev1/mini-app-project/internal/repo/redisDB"
	"github.com/zovdev1/mini-app-project/internal/usecase/dto"
)

type ProductUseCase struct {
	repo   repo.ProductUseRepo
	client *redisdb.RedisUseRepo
}

func New(repo repo.ProductUseRepo, client *redisdb.RedisUseRepo) *ProductUseCase {
	return &ProductUseCase{
		repo:   repo,
		client: client,
	}
}

func (u *ProductUseCase) Create(input dto.ProductInput) (entites.Product, error) {

	new := time.Now()

	newProduct := entites.Product{
		ID:          uuid.Must(uuid.NewV7()),
		Title:       input.Title,
		Description: input.Description,
		Price:       input.Price,
		Quantity:    input.Quantity,
		UserID:      input.UserID,
		CreatedAt:   new,
		UpdatedAt:   new,
	}

	err := newProduct.Validate()

	if err != nil {
		return entites.Product{}, fmt.Errorf("validation failed: %w", err)
	}

	savedProduct, err := u.repo.Create(newProduct)

	if err != nil {
		return entites.Product{}, fmt.Errorf("failed to create product: %w", err)
	}

	return *savedProduct, nil
}

func (u *ProductUseCase) GetAll(limit, offset int) ([]*entites.Product, error) {
	productAll, err := u.repo.GetAllProduct(limit, offset)

	if err != nil {
		return nil, err
	}

	return productAll, nil
}

func (u *ProductUseCase) GetById(id string) (*entites.Product, error) {

	parse, err := uuid.Parse(id)
	ctx := context.Background()

	if err != nil {
		return &entites.Product{}, err
	}

	cachedProduct, err := u.client.GetProduct(ctx, parse)

	if err == nil {
		return cachedProduct, nil
	}

	product, err := u.repo.GetById(parse)

	if err != nil {
		return &entites.Product{}, err
	}

	go u.client.UpdateCache(ctx, product, 30*time.Minute)

	return product, nil
}

func (u *ProductUseCase) DELETE(userID, productID string) error {
	parseUser, err := uuid.Parse(userID)
	if err != nil {
		return fmt.Errorf("invalid user ID: %w", err)
	}

	parseProduct, err := uuid.Parse(productID)
	if err != nil {
		return fmt.Errorf("invalid product ID: %w", err)
	}

	fmt.Println(parseProduct, parseUser)

	err = u.repo.DELETE(parseUser, parseProduct)

	if err != nil {
		return err
	}

	return nil
}
