package store

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type CheckoutSession struct {
	SessionID       string             `json:"session_id"`
	UserID          int64              `json:"user_id"`
	ID              uuid.UUID          `json:"id"`
	CartStore       []CartStores       `json:"cart_store"`
	TotalPrice      float64            `json:"total_price"`
	ShippingMethod  *ShippingMethod    `json:"shipping_method,omitempty"`
	PaymentMethod   *PaymentMethod     `json:"payment_method,omitempty"`
	ShippingAddress *ShippingAddresses `json:"shipping_address,omitempty"`
	Notes           string             `json:"notes,omitempty"`
	CreatedAt       time.Time          `json:"created_at"`
	ExpiresAt       time.Time          `json:"expires_at"`
}

type CheckoutStore struct {
	db *sql.DB
}

func (s *CheckoutStore) CreateOrderFromCheckout(ctx context.Context, checkout *CheckoutSession) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Validasi komponen penting
	if checkout.ShippingMethod == nil || checkout.PaymentMethod == nil || checkout.ShippingAddress == nil {
		return fmt.Errorf("missing required checkout components")
	}

	for _, cartStore := range checkout.CartStore {
		var orderID int64

		err := tx.QueryRowContext(ctx, `
			SELECT create_order_from_cart($1, $2, $3, $4, $5, $6)
		`,
			checkout.UserID,             // $1: p_user_id
			cartStore.ID,                // $2: p_cart_store_id (UUID)
			checkout.PaymentMethod.ID,   // $3: p_payment_method_id
			checkout.ShippingMethod.ID,  // $4: p_shipping_method_id
			checkout.ShippingAddress.ID, // $5: p_shipping_addresses_id (UUID)
			checkout.Notes,              // $6: p_notes
		).Scan(&orderID)

		if err != nil {
			return fmt.Errorf("failed to create order from cart store %s: %w", cartStore.ID, err)
		}
	}

	return tx.Commit()
}
