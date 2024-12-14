package store

import (
	"context"
	"database/sql"
)

type Toko struct {
	ID           int64  `json:"id"`                      // Primary key
	UserID       int64  `json:"user_id"`                 // Foreign key to users table
	Slug         string `json:"slug"`                    // Unique slug for the store
	Name         string `json:"name"`                    // Store name
	ImageProfile string `json:"image_profile,omitempty"` // Optional profile image
	Country      string `json:"country"`                 // Store's country
	CreatedAt    string `json:"created_at"`              // Timestamp of creation
}

type TokoStore struct {
	db *sql.DB
}

func (s *TokoStore) Exists(ctx context.Context, slug string) (bool, error) {
	const query = `SELECT 1 FROM toko WHERE slug = $1`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	var exists int
	err := s.db.QueryRowContext(ctx, query, slug).Scan(&exists)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (s *TokoStore) Create(ctx context.Context, t *Toko) error {
	// Periksa apakah toko dengan slug yang sama sudah ada
	exists, err := s.Exists(ctx, t.Slug)
	if err != nil {
		return err
	}

	if exists {
		return nil // Jika sudah ada, tidak melakukan operasi create
	}

	const query = `INSERT INTO toko (user_id, slug, name, image_profile, country) VALUES ($1, $2, $3, $4, $5) RETURNING id, created_at`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	err = s.db.QueryRowContext(ctx, query, t.UserID, t.Slug, t.Name, t.ImageProfile, t.Country).Scan(&t.ID, &t.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (s *TokoStore) GetByID(ctx context.Context, id int64) (*Toko, error) {
	const query = `SELECT id, user_id, slug, name, image_profile, country, created_at FROM toko WHERE id = $1`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	t := &Toko{}
	err := s.db.QueryRowContext(ctx, query, id).Scan(&t.ID, &t.UserID, &t.Slug, &t.Name, &t.ImageProfile, &t.Country, &t.CreatedAt)
	if err != nil {
		switch {
		case err == sql.ErrNoRows:
			return nil, ErrNotFound
		default:
			return nil, err
		}
	}

	return t, nil
}

func (s *TokoStore) GetBySlug(ctx context.Context, slug string) (*Toko, error) {
	const query = `SELECT id, user_id, slug, name, image_profile, country FROM toko WHERE slug = $1`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	toko := &Toko{}
	err := s.db.QueryRowContext(ctx, query, slug).Scan(&toko.ID, &toko.UserID, &toko.Slug, &toko.Name, &toko.ImageProfile, &toko.Country)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return toko, nil
}
