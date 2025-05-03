package store

import (
	"context"
	"database/sql"
	"time"
)

type Toko struct {
	ID           int64       `json:"id,omitempty"`
	UserID       int64       `json:"user_id,omitempty"`
	Slug         string      `json:"slug"`
	Name         string      `json:"name"`
	ImageProfile string      `json:"image_profile"`
	Country      string      `json:"country"`
	CreatedAt    time.Time   `json:"created_at"`
	User         *SingleUser `json:"user"`
}

type TokoStore struct {
	db *sql.DB
}

func (s *TokoStore) Exists(ctx context.Context, slug string) (bool, error) {
	const query = `SELECT 1 FROM tokos WHERE slug = $1`

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
	exists, err := s.Exists(ctx, t.Slug)
	if err != nil {
		return err
	}

	if exists {
		return nil
	}

	const query = `INSERT INTO tokos (user_id, slug, name, image_profile, country) VALUES ($1, $2, $3, $4, $5) RETURNING id, created_at`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	err = s.db.QueryRowContext(ctx, query, t.UserID, t.Slug, t.Name, t.ImageProfile, t.Country).Scan(&t.ID, &t.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (s *TokoStore) GetByID(ctx context.Context, id int64) (*Toko, error) {
	const query = `
    SELECT t.id, t.user_id, t.slug, t.name, t.image_profile, t.country, t.created_at,
           u.id, u.username, u.email, u.created_at
    FROM tokos t
    JOIN users u ON t.user_id = u.id
    WHERE t.id = $1`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	t := &Toko{}
	user := &SingleUser{}
	err := s.db.QueryRowContext(ctx, query, id).Scan(
		&t.ID, &t.UserID, &t.Slug, &t.Name, &t.ImageProfile, &t.Country, &t.CreatedAt,
		&user.ID, &user.Username, &user.Email, &user.CreatedAt)
	if err != nil {
		switch {
		case err == sql.ErrNoRows:
			return nil, ErrNotFound
		default:
			return nil, err
		}
	}

	t.User = user
	return t, nil
}

func (s *TokoStore) GetBySlug(ctx context.Context, slug string) (*Toko, error) {
	const query = `SELECT t.id, t.user_id, t.slug, t.name, t.image_profile, t.country, t.created_at,
				u.id, u.username, u.email, u.picture, u.created_at
				FROM tokos t
				JOIN users u ON t.user_id = u.id
				WHERE t.slug = $1`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	t := &Toko{}
	user := &SingleUser{}
	err := s.db.QueryRowContext(ctx, query, slug).Scan(
		&t.ID, &t.UserID, &t.Slug, &t.Name, &t.ImageProfile, &t.Country, &t.CreatedAt,
		&user.ID, &user.Username, &user.Email, &user.Picture, &user.CreatedAt)
	if err != nil {
		switch {
		case err == sql.ErrNoRows:
			return nil, ErrNotFound
		default:
			return nil, err
		}
	}

	t.User = user
	return t, nil
}
