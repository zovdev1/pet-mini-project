package redisdb

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"github.com/zovdev1/mini-app-project/internal/entites"
)

type RedisUseRepo struct {
	client *redis.Client
}

func NewRedisUseRepo(client *redis.Client) *RedisUseRepo {
	return &RedisUseRepo{client: client}
}

func (storage *RedisUseRepo) GetProduct(ctx context.Context, productID uuid.UUID) (*entites.Product, error) {
	key := fmt.Sprintf("product:%s", productID)

	data, err := storage.client.Get(ctx, key).Result()

	if err != nil {
		if err == redis.Nil {
			return nil, fmt.Errorf("product not found in cache")
		}
		return nil, err
	}

	var product entites.Product
	if err := json.Unmarshal([]byte(data), &product); err != nil {
		return nil, err
	}

	return &product, nil
}

func (storage *RedisUseRepo) UpdateCache(ctx context.Context, product *entites.Product, expiration time.Duration) error {
	key := fmt.Sprintf("product:%s", product.ID)

	productJSON, err := json.Marshal(product)

	if err != nil {
		return err
	}

	return storage.client.Set(ctx, key, productJSON, expiration).Err()
}
