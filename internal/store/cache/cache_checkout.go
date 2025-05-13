package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"github.com/yogaprasetya22/api-gotokopedia/internal/store"
)

type CheckoutStore struct {
	rdb *redis.Client
}

func (c *CheckoutStore) StartCheckoutSession(ctx context.Context, userID int64, cartStore []store.CartStores) (*store.CheckoutSession, error) {
	// Cek existing session untuk semua store di cart
	for _, cs := range cartStore {
		existing, err := c.GetActiveSessionByUserAndStore(ctx, userID, cs.TokoID)
		if err == nil && existing != nil {
			return existing, nil
		}
	}

	sessionID := uuid.New().String()
	now := time.Now()
	expiresAt := now.Add(checkoutExpTime)

	checkout := &store.CheckoutSession{
		SessionID: sessionID,
		UserID:    userID,
		CartStore: cartStore,
		CreatedAt: now,
		ExpiresAt: expiresAt,
	}

	jsonData, err := json.Marshal(checkout)
	if err != nil {
		return nil, err
	}

	// Gunakan transaction untuk atomic operation
	maxRetries := 3
	for i := 0; i < maxRetries; i++ {
		err = c.rdb.Watch(ctx, func(tx *redis.Tx) error {
			// Double check existing session
			for _, cs := range cartStore {
				if keys, _ := tx.Keys(ctx, c.userStoreSessionPattern(userID, cs.TokoID)).Result(); len(keys) > 0 {
					return fmt.Errorf("existing session found")
				}
			}

			// Pre-check inventory sebelum membuat session
			for _, cs := range cartStore {
				for _, item := range cs.Items {
					lockKey := c.inventoryLockKey(item.ProductID)
					current, err := tx.Get(ctx, lockKey).Int64()
					if err != nil && err != redis.Nil {
						return err
					}
					if int64(item.Quantity) > (int64(item.Product.Stock) - current) {
						return fmt.Errorf("insufficient stock for product %d", item.ProductID)
					}
				}
			}

			// Mulai transaction
			_, err := tx.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
				// Set session
				for _, cs := range cartStore {
					sessionKey := c.userStoreSessionKey(userID, cs.TokoID, sessionID)
					if err := pipe.Set(ctx, sessionKey, jsonData, checkoutExpTime).Err(); err != nil {
						return err
					}
				}

				// Set lock
				for _, cs := range cartStore {
					for _, item := range cs.Items {
						lockKey := c.inventoryLockKey(item.ProductID)
						if err := pipe.IncrBy(ctx, lockKey, item.Quantity).Err(); err != nil {
							return err
						}
						pipe.Expire(ctx, lockKey, lockExpTime)
					}
				}
				return nil
			})
			return err
		})

		if err == nil {
			return checkout, nil
		}

		if strings.Contains(err.Error(), "insufficient stock") {
			return nil, err // Langsung return jika stok tidak cukup
		}

		time.Sleep(time.Duration(i+1) * 100 * time.Millisecond) // Exponential backoff
	}

	return nil, fmt.Errorf("max retries reached")
}

func (c *CheckoutStore) CompleteCheckout(ctx context.Context, sessionID string) error {
	// Dapatkan session terlebih dahulu
	session, err := c.GetCheckoutSession(ctx, sessionID)
	if err != nil {
		return err
	}

	// Hapus session dan lock stok
	pipe := c.rdb.TxPipeline()
	pipe.Del(ctx, c.sessionKey(sessionID))

	// Hapus semua inventory lock untuk produk dalam session ini
	for _, store := range session.CartStore {
		for _, item := range store.Items {
			pipe.Del(ctx, c.inventoryLockKey(item.ProductID))
		}
	}

	_, err = pipe.Exec(ctx)
	return err
}

// Helper function untuk key yang konsisten
func (c *CheckoutStore) userStoreSessionKey(userID int64, storeID int64, sessionID string) string {
	return fmt.Sprintf("checkout:user:%d:store:%d:%s", userID, storeID, sessionID)
}

func (c *CheckoutStore) GetCheckoutSession(ctx context.Context, sessionID string) (*store.CheckoutSession, error) {
	// Cari di semua keys pattern checkout:*:sessionID
	keys, err := c.rdb.Keys(ctx, fmt.Sprintf("checkout:*:%s", sessionID)).Result()
	if err != nil || len(keys) == 0 {
		return nil, store.ErrNotFound
	}

	data, err := c.rdb.Get(ctx, keys[0]).Result()
	if err != nil {
		return nil, err
	}

	var session store.CheckoutSession
	if err := json.Unmarshal([]byte(data), &session); err != nil {
		return nil, err
	}
	return &session, nil
}

func (c *CheckoutStore) GetActiveSessionByUserAndStore(ctx context.Context, userID int64, storeID int64) (*store.CheckoutSession, error) {
	// Pattern: checkout:user:{userID}:store:{storeID}:*
	pattern := fmt.Sprintf("checkout:user:%d:store:%d:*", userID, storeID)

	keys, err := c.rdb.Keys(ctx, pattern).Result()
	if err != nil {
		return nil, err
	}

	if len(keys) == 0 {
		return nil, store.ErrNotFound
	}

	// Ambil session terbaru
	return c.GetCheckoutSession(ctx, extractSessionID(keys[0]))
}

func extractSessionID(key string) string {
	parts := strings.Split(key, ":")
	return parts[len(parts)-1] // Ambil bagian terakhir sebagai sessionID
}

func (c *CheckoutStore) userStoreSessionPattern(userID int64, storeID int64) string {
	return fmt.Sprintf("checkout:user:%d:store:%d:*", userID, storeID)
}

func (c *CheckoutStore) sessionKey(sessionID string) string {
	return fmt.Sprintf("checkout:%s", sessionID)
}

func (c *CheckoutStore) inventoryLockKey(productID int64) string {
	return fmt.Sprintf("inventory_lock:product:%d", productID)
}
