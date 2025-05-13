package store

import (
	"context"
	"database/sql"
	"time"
)

type CheckoutSession struct {
	SessionID       string             `json:"session_id"`
	UserID          int64              `json:"user_id"`
	CartStore       []CartStores       `json:"cart_store"`
	ShippingMethod  *ShippingMethod    `json:"shipping_method,omitempty"`
	PaymentMethod   *PaymentMethod     `json:"payment_method,omitempty"`
	ShippingAddress *ShippingAddresses `json:"shipping_address,omitempty"`
	CreatedAt       time.Time          `json:"created_at"`
	ExpiresAt       time.Time          `json:"expires_at"`
}

type CheckoutStore struct {
	db *sql.DB
}

func (s *CheckoutStore) CreateOrderFromCheckout(ctx context.Context, checkout *CheckoutSession) error {
	// Implementasi pembuatan order permanen di database
	// Ini akan dipanggil setelah pembayaran berhasil
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Implementasi logika pembuatan order disini
	// ...

	return tx.Commit()
}
