package cache

import (
	"context"

	"github.com/redis/go-redis/v9"
	"github.com/yogaprasetya22/api-gotokopedia/internal/store"
)

type Storage struct {
	Users interface {
		Get(context.Context, int64) (*store.User, error)
		Set(context.Context, *store.User) error
		Delete(context.Context, int64)
	}
	Products interface {
		Get(context.Context, int64) (*store.Product, error)
		Set(context.Context, *store.Product) error
		Delete(context.Context, int64)
	}
	Carts interface {
		Get(context.Context, int64) (*store.Cart, error)
		Set(context.Context, *store.Cart) error
		Delete(context.Context, int64)
	}
}

func NewRedisStore(rbd *redis.Client) Storage {
	return Storage{
		Users:    &UserStore{rdb: rbd},
		Products: &ProductStore{rdb: rbd},
		Carts:    &CartStore{rdb: rbd},
	}
}
