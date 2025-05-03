package store

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type ShippingAddresses struct {
	ID             uuid.UUID      `json:"id"`
	UserID         int64          `json:"user_id"`
	Label          string         `json:"label"`
	RecipientName  string         `json:"recipient_name"`
	RecipientPhone string         `json:"recipient_phone"`
	AddressLine1   string         `json:"address_line1"`
	NoteForCourier sql.NullString `json:"note_for_courier"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
}

type ShippingAddresStore struct {
	db *sql.DB
}

func (s *ShippingAddresStore) GetByID(ctx context.Context, id uuid.UUID, userID int64) (*ShippingAddresses, error) {
	query := `SELECT id, user_id, label, recipient_name, recipient_phone, address_line1, note_for_courier, created_at, updated_at 
	FROM shipping_addresses WHERE id = $1 AND user_id = $2`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	sa := &ShippingAddresses{}
	err := s.db.QueryRowContext(ctx, query, id, userID).Scan(&sa.ID, &sa.UserID, &sa.Label,
		&sa.RecipientName, &sa.RecipientPhone, &sa.AddressLine1,
		&sa.NoteForCourier, &sa.CreatedAt, &sa.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return sa, nil
}

func (s *ShippingAddresStore) Create(ctx context.Context, sa *ShippingAddresses) error {
	query := `INSERT INTO shipping_addresses (user_id, label, recipient_name, recipient_phone, address_line1, note_for_courier) 
	VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, created_at, updated_at`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	err := s.db.QueryRowContext(ctx, query,
		sa.UserID,
		sa.Label,
		sa.RecipientName,
		sa.RecipientPhone,
		sa.AddressLine1,
		sa.NoteForCourier).Scan(&sa.ID, &sa.CreatedAt, &sa.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (s *ShippingAddresStore) Update(ctx context.Context, sa *ShippingAddresses) error {
	query := `UPDATE shipping_addresses 
	SET label = $1, recipient_name = $2, recipient_phone = $3, address_line1 = $4, note_for_courier = $5, updated_at = NOW() 
	WHERE id = $6 AND user_id = $7`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	_, err := s.db.ExecContext(ctx, query,
		sa.Label,
		sa.RecipientName,
		sa.RecipientPhone,
		sa.AddressLine1,
		sa.NoteForCourier,
		sa.ID,
		sa.UserID)
	if err != nil {
		return err
	}

	return nil
}

func (s *ShippingAddresStore) Delete(ctx context.Context, id uuid.UUID, userID int64) error {
	query := `DELETE FROM shipping_addresses WHERE id = $1 AND user_id = $2`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	_, err := s.db.ExecContext(ctx, query, id, userID)
	if err != nil {
		return err
	}

	return nil
}
