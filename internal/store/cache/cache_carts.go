package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/yogaprasetya22/api-gotokopedia/internal/store"
)

type CartStore struct {
	rdb *redis.Client
}

const cartExpTime = time.Minute

func (p *CartStore) Get(ctx context.Context, cartID int64) (*store.Cart, error) {
	cacheKey := fmt.Sprintf("cart-%d", cartID)

	data, err := p.rdb.Get(ctx, cacheKey).Result()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	var cart store.Cart
	if data != "" {
		err := json.Unmarshal([]byte(data), &cart)
		if err != nil {
			return nil, err
		}
	}

	return &cart, nil
}

func (s *CartStore) Set(ctx context.Context, cart *store.Cart) error {
	cacheKey := fmt.Sprintf("cart-%d", cart.ID)

	json, err := json.Marshal(cart)
	if err != nil {
		return err
	}

	return s.rdb.Set(ctx, cacheKey, json, cartExpTime).Err()
}

func (s *CartStore) Delete(ctx context.Context, cartID int64) {
	cacheKey := fmt.Sprintf("cart-%d", cartID)
	s.rdb.Del(ctx, cacheKey)
}
