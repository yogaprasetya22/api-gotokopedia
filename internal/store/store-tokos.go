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

func (s *TokoStore) Create(ctx context.Context, t *Toko) error {
	const query = `INSERT INTO toko (user_id, slug, name, image_profile, country) VALUES ($1, $2, $3, $4, $5) RETURNING id, created_at`

	err := s.db.QueryRowContext(ctx, query, t.UserID, t.Slug, t.Name, t.ImageProfile, t.Country).Scan(&t.ID, &t.CreatedAt)

	if err != nil {
		return err
	}
	return nil
}

func (s *TokoStore) GetByID(ctx context.Context, id int64) (*Toko, error) {
	const query = `SELECT id, user_id, slug, name, image_profile, country, created_at FROM toko WHERE id = $1`

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
