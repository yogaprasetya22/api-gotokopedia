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

// Helper function to compare two CartStores slices
func compareCartStores(a, b []store.CartStores) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i].TokoID != b[i].TokoID {
			return false
		}

		if len(a[i].Items) != len(b[i].Items) {
			return false
		}

		for j := range a[i].Items {
			if a[i].Items[j].ProductID != b[i].Items[j].ProductID ||
				a[i].Items[j].Quantity != b[i].Items[j].Quantity {
				return false
			}
		}
	}

	return true
}

func (c *CheckoutStore) StartCheckoutSession(ctx context.Context, userID int64, cartStore []store.CartStores) (*store.CheckoutSession, error) {
	// Get all active sessions for this user
	allSessions, err := c.getAllActiveSessionsForUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	// Check if any existing session matches exactly with the current cartStore
	for _, session := range allSessions {
		if compareCartStores(session.CartStore, cartStore) {
			return session, nil
		}
	}

	// If no exact match found, delete all overlapping sessions
	for _, session := range allSessions {
		if hasOverlappingStores(session.CartStore, cartStore) {
			if err := c.CompleteCheckout(ctx, session.SessionID); err != nil {
				return nil, err
			}
		}
	}

	// Total Price calculation
	var totalPrice float64
	for _, cs := range cartStore {
		for _, item := range cs.Items {
			// Assuming item.Price is the price per unit and item.Quantity is the number of units
			totalPrice += item.Product.Price * float64(item.Quantity)
		}
	}

	// Create new session
	sessionID := uuid.New().String()
	now := time.Now()
	expiresAt := now.Add(checkoutExpTime)

	checkout := &store.CheckoutSession{
		SessionID:  sessionID,
		UserID:     userID,
		CartStore:  cartStore,
		TotalPrice: totalPrice,
		CreatedAt:  now,
		ExpiresAt:  expiresAt,
	}

	jsonData, err := json.Marshal(checkout)
	if err != nil {
		return nil, err
	}

	// Use transaction for atomic operation
	maxRetries := 3
	for i := 0; i < maxRetries; i++ {
		err = c.rdb.Watch(ctx, func(tx *redis.Tx) error {
			// Pre-check inventory before creating session
			for _, cs := range cartStore {
				for _, item := range cs.Items {
					lockKey := c.inventoryLockKey(item.ProductID)
					current, err := tx.Get(ctx, lockKey).Int64()
					if err != nil && err != redis.Nil {
						return err
					}
					if int64(item.Quantity) > (int64(item.Product.Stock) - current) {
						return &store.StockError{
							ProductID: item.ProductID,
						}

					}
				}
			}

			// Start transaction
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
			return nil, err // Direct return if stock insufficient
		}

		time.Sleep(time.Duration(i+1) * 100 * time.Millisecond) // Exponential backoff
	}

	return nil, fmt.Errorf("max retries reached")
}

// Helper to get all active sessions for a user
func (c *CheckoutStore) getAllActiveSessionsForUser(ctx context.Context, userID int64) ([]*store.CheckoutSession, error) {
	pattern := fmt.Sprintf("checkout:user:%d:store:*", userID)
	keys, err := c.rdb.Keys(ctx, pattern).Result()
	if err != nil {
		return nil, err
	}

	var sessions []*store.CheckoutSession
	for _, key := range keys {
		data, err := c.rdb.Get(ctx, key).Result()
		if err != nil {
			continue
		}

		var session store.CheckoutSession
		if err := json.Unmarshal([]byte(data), &session); err != nil {
			continue
		}

		sessions = append(sessions, &session)
	}

	return sessions, nil
}

// Helper to check if two cart stores have overlapping toko IDs
func hasOverlappingStores(a, b []store.CartStores) bool {
	tokoMap := make(map[int64]bool)
	for _, cs := range a {
		tokoMap[cs.TokoID] = true
	}

	for _, cs := range b {
		if tokoMap[cs.TokoID] {
			return true
		}
	}

	return false
}

// Existing functions remain the same...
func (c *CheckoutStore) CompleteCheckout(ctx context.Context, sessionID string) error {
	// Ambil session terlebih dahulu
	session, err := c.GetCheckoutSession(ctx, sessionID)
	if err != nil {
		return err
	}

	// Cari semua key yang berkaitan dengan sessionID
	keys, err := c.rdb.Keys(ctx, fmt.Sprintf("checkout:user:%d:store:*:%s", session.UserID, sessionID)).Result()
	if err != nil {
		return err
	}

	pipe := c.rdb.TxPipeline()

	// Hapus semua key sesi
	for _, key := range keys {
		pipe.Del(ctx, key)
	}

	// Hapus semua inventory lock
	for _, store := range session.CartStore {
		for _, item := range store.Items {
			pipe.Del(ctx, c.inventoryLockKey(item.ProductID))
		}
	}

	_, err = pipe.Exec(ctx)
	fmt.Printf("Checkout completed for session %s, deleted keys: %v\n", sessionID, keys)
	return err
}


func (c *CheckoutStore) userStoreSessionKey(userID int64, storeID int64, sessionID string) string {
	return fmt.Sprintf("checkout:user:%d:store:%d:%s", userID, storeID, sessionID)
}

func (c *CheckoutStore) GetCheckoutSession(ctx context.Context, sessionID string) (*store.CheckoutSession, error) {
	keys, err := c.rdb.Keys(ctx, fmt.Sprintf("checkout:*:%s", sessionID)).Result()
	if err != nil || len(keys) == 0 {
		return nil, store.ErrNotFound
	}

	data, err := c.rdb.Get(ctx, keys[0]).Result()
	if err == redis.Nil {
		return nil, store.ErrSessionExpired
	}
	if err != nil {
		return nil, err
	}

	var session store.CheckoutSession
	if err := json.Unmarshal([]byte(data), &session); err != nil {
		return nil, err
	}

	// Check if session is expired based on ExpiresAt
	if time.Now().After(session.ExpiresAt) {
		return nil, store.ErrSessionExpired
	}

	return &session, nil
}

func (c *CheckoutStore) sessionKey(sessionID string) string {
	return fmt.Sprintf("checkout:%s", sessionID)
}

func (c *CheckoutStore) inventoryLockKey(productID int64) string {
	return fmt.Sprintf("inventory_lock:product:%d", productID)
}
