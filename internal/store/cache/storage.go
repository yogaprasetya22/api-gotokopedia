package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/yogaprasetya22/api-gotokopedia/internal/store"
)

const (
	checkoutExpTime = 60 * time.Minute
	lockExpTime     = 60 * time.Minute
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
	Checkout interface {
		StartCheckoutSession(ctx context.Context, userID int64, cartStore []store.CartStores) (*store.CheckoutSession, error)
		GetCheckoutSession(ctx context.Context, sessionID string) (*store.CheckoutSession, error)
		CompleteCheckout(ctx context.Context, sessionID string) error
		sessionKey(sessionID string) string
		inventoryLockKey(productID int64) string
	}
}

func NewRedisStore(rbd *redis.Client) Storage {
	return Storage{
		Users:    &UserStore{rdb: rbd},
		Products: &ProductStore{rdb: rbd},
		Carts:    &CartStore{rdb: rbd},
		Checkout: &CheckoutStore{rdb: rbd},
	}
}
