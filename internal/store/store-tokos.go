package store

import (
	"context"
	"database/sql"
	"time"
)

type Toko struct {
	ID           string    `json:"id"`
	Slug         string    `json:"slug"`
	Name         string    `json:"name"`
	ImageProfile *string   `json:"image_profile,omitempty"`
	Country      string    `json:"country"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	UserID       *string   `json:"user_id,omitempty"`
}

type TokoStore struct {
	db *sql.DB
}

func (s *TokoStore) Create(ctx context.Context, t *Toko) error {
	const query = `INSERT INTO Toko (slug, name, image_profile, country, created_at, updated_at, user_id) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`
	err := s.db.QueryRowContext(ctx, query, t.Slug, t.Name, t.ImageProfile, t.Country, t.CreatedAt, t.UpdatedAt, t.UserID).Scan(&t.ID)
	if err != nil {
		return err
	}

	return nil
}
