package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/yogaprasetya22/api-gotokopedia/internal/store"
)

type ProductStore struct {
	rdb *redis.Client
}


const productExpTime = time.Minute

func (p *ProductStore) Get(ctx context.Context, productID int64) (*store.Product, error) {
	cacheKey := fmt.Sprintf("product-%d", productID)

	data, err := p.rdb.Get(ctx, cacheKey).Result()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	var product store.Product
	if data != "" {
		err := json.Unmarshal([]byte(data), &product)
		if err != nil {
			return nil, err
		}
	}

	return &product, nil
}

func (s *ProductStore) Set(ctx context.Context, product *store.Product) error {
	cacheKey := fmt.Sprintf("product-%d", product.ID)

	json, err := json.Marshal(product)
	if err != nil {
		return err
	}

	return s.rdb.Set(ctx, cacheKey, json, productExpTime).Err()
}

func (s *ProductStore) Delete(ctx context.Context, productID int64) {
	cacheKey := fmt.Sprintf("product-%d", productID)
	s.rdb.Del(ctx, cacheKey)
}
