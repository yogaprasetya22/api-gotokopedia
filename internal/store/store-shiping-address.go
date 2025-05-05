package store

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type ShippingAddresses struct {
	ID             uuid.UUID `json:"id"`
	UserID         int64     `json:"user_id"`
	IsDefault      bool      `json:"is_default,omitempty"` // Tambahkan ini
	Label          string    `json:"label"`
	RecipientName  string    `json:"recipient_name"`
	RecipientPhone string    `json:"recipient_phone"`
	AddressLine1   string    `json:"address_line1"`
	NoteForCourier string    `json:"note_for_courier,omitempty"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type ShippingAddresStore struct {
	db *sql.DB
}

// SetDefaultAddress mengatur alamat tertentu sebagai default
func (s *ShippingAddresStore) SetDefaultAddress(ctx context.Context, addressID uuid.UUID, userID int64) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// 1. Reset semua alamat user ke non-default
	_, err = tx.ExecContext(ctx,
		`UPDATE shipping_addresses SET is_default = false WHERE user_id = $1`,
		userID)
	if err != nil {
		return err
	}

	// 2. Set alamat yang dipilih sebagai default
	_, err = tx.ExecContext(ctx,
		`UPDATE shipping_addresses SET is_default = true WHERE id = $1 AND user_id = $2`,
		addressID, userID)
	if err != nil {
		return err
	}

	// 3. Update referensi di tabel users
	_, err = tx.ExecContext(ctx,
		`UPDATE users SET default_shipping_address_id = $1 WHERE id = $2`,
		addressID, userID)
	if err != nil {
		return err
	}

	return tx.Commit()
}

// GetDefaultAddress mengambil alamat default user
func (s *ShippingAddresStore) GetDefaultAddress(ctx context.Context, userID int64) (*ShippingAddresses, error) {
	query := `SELECT id, user_id, label, recipient_name, recipient_phone, 
                     address_line1, note_for_courier, is_default, created_at, updated_at 
              FROM shipping_addresses 
              WHERE user_id = $1 AND is_default = true`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	sa := &ShippingAddresses{}
	var noteForCourier sql.NullString
	err := s.db.QueryRowContext(ctx, query, userID).Scan(
		&sa.ID, &sa.UserID, &sa.Label,
		&sa.RecipientName, &sa.RecipientPhone, &sa.AddressLine1,
		&noteForCourier,
		&sa.IsDefault, &sa.CreatedAt, &sa.UpdatedAt)
	if err != nil {
		return nil, err
	}
	if noteForCourier.Valid {
		sa.NoteForCourier = noteForCourier.String
	} else {
		sa.NoteForCourier = ""
	}

	return sa, nil
}

func (s *ShippingAddresStore) GetByID(ctx context.Context, id uuid.UUID, userID int64) (*ShippingAddresses, error) {
	query := `SELECT id, user_id, label, recipient_name, recipient_phone, 
                     address_line1, note_for_courier, is_default, created_at, updated_at 
              FROM shipping_addresses WHERE id = $1 AND user_id = $2`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	sa := &ShippingAddresses{}
	err := s.db.QueryRowContext(ctx, query, id, userID).Scan(
		&sa.ID, &sa.UserID, &sa.Label,
		&sa.RecipientName, &sa.RecipientPhone, &sa.AddressLine1,
		&sa.NoteForCourier, &sa.IsDefault, &sa.CreatedAt, &sa.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return sa, nil
}

func (s *ShippingAddresStore) ListByUser(ctx context.Context, userID int64) ([]*ShippingAddresses, error) {
	query := `SELECT id, user_id, label, recipient_name, recipient_phone, 
                     address_line1, note_for_courier, is_default, created_at, updated_at 
              FROM shipping_addresses 
              WHERE user_id = $1 
              ORDER BY is_default DESC, created_at DESC`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	rows, err := s.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var addresses []*ShippingAddresses
	for rows.Next() {
		sa := &ShippingAddresses{}
		err := rows.Scan(
			&sa.ID, &sa.UserID, &sa.Label,
			&sa.RecipientName, &sa.RecipientPhone, &sa.AddressLine1,
			&sa.NoteForCourier, &sa.IsDefault, &sa.CreatedAt, &sa.UpdatedAt)
		if err != nil {
			return nil, err
		}
		addresses = append(addresses, sa)
	}

	return addresses, nil
}

func (s *ShippingAddresStore) Create(ctx context.Context, sa *ShippingAddresses) error {
	// Mulai transaksi
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	// 1. Cek apakah ini alamat pertama user
	var count int
	err = tx.QueryRowContext(ctx,
		`SELECT COUNT(*) FROM shipping_addresses WHERE user_id = $1`,
		sa.UserID).Scan(&count)
	if err != nil {
		return fmt.Errorf("failed to count user addresses: %w", err)
	}

	// 2. Jika ini alamat pertama, otomatis jadikan default
	if count == 0 {
		sa.IsDefault = true
	}

	// 3. Jika ingin dijadikan default dan sudah ada alamat default sebelumnya,
	//    kita perlu menonaktifkan alamat default yang lama
	if sa.IsDefault {
		_, err = tx.ExecContext(ctx,
			`UPDATE shipping_addresses SET is_default = false WHERE user_id = $1 AND is_default = true`,
			sa.UserID)
		if err != nil {
			return fmt.Errorf("failed to reset default addresses: %w", err)
		}
	}

	// 4. Insert alamat baru
	query := `INSERT INTO shipping_addresses 
              (user_id, label, recipient_name, recipient_phone, 
               address_line1, note_for_courier, is_default) 
              VALUES ($1, $2, $3, $4, $5, $6, $7) 
              RETURNING id, created_at, updated_at, is_default`

	err = tx.QueryRowContext(ctx, query,
		sa.UserID,
		sa.Label,
		sa.RecipientName,
		sa.RecipientPhone,
		sa.AddressLine1,
		sa.NoteForCourier,
		sa.IsDefault).Scan(&sa.ID, &sa.CreatedAt, &sa.UpdatedAt, &sa.IsDefault)
	if err != nil {
		return fmt.Errorf("failed to insert shipping address: %w", err)
	}

	// 5. Jika default, update juga tabel users
	if sa.IsDefault {
		_, err = tx.ExecContext(ctx,
			`UPDATE users SET default_shipping_address_id = $1 WHERE id = $2`,
			sa.ID, sa.UserID)
		if err != nil {
			return fmt.Errorf("failed to update user default address: %w", err)
		}
	}

	// Commit transaksi
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
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
	// Cek apakah alamat yang akan dihapus adalah default
	var isDefault bool
	err := s.db.QueryRowContext(ctx,
		`SELECT is_default FROM shipping_addresses WHERE id = $1 AND user_id = $2`,
		id, userID).Scan(&isDefault)
	if err != nil {
		return err
	}

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Hapus alamat
	_, err = tx.ExecContext(ctx,
		`DELETE FROM shipping_addresses WHERE id = $1 AND user_id = $2`,
		id, userID)
	if err != nil {
		return err
	}

	// Jika alamat yang dihapus adalah default
	if isDefault {
		// Cari alamat lain untuk dijadikan default
		var newDefaultID uuid.UUID
		err = tx.QueryRowContext(ctx,
			`SELECT id FROM shipping_addresses 
             WHERE user_id = $1 LIMIT 1`,
			userID).Scan(&newDefaultID)

		if err == nil {
			// Set alamat baru sebagai default
			_, err = tx.ExecContext(ctx,
				`UPDATE shipping_addresses SET is_default = true 
                 WHERE id = $1 AND user_id = $2`,
				newDefaultID, userID)
			if err != nil {
				return err
			}

			// Update reference di users
			_, err = tx.ExecContext(ctx,
				`UPDATE users SET default_shipping_address_id = $1 WHERE id = $2`,
				newDefaultID, userID)
			if err != nil {
				return err
			}
		} else if err == sql.ErrNoRows {
			// Tidak ada alamat lain, set ke NULL di users
			_, err = tx.ExecContext(ctx,
				`UPDATE users SET default_shipping_address_id = NULL WHERE id = $1`,
				userID)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}

	return tx.Commit()
}
