package store

import (
	"context"
	"database/sql"
)

type Category struct {
	ID          int64  `json:"id" `
	Name        string `json:"name" `
	Slug        string `json:"slug" `
	Description string `json:"description,omitempty" `
}

type CategoryStore struct {
	db *sql.DB
}

func (s *CategoryStore) Create(ctx context.Context, c *Category) error {
	const query = `INSERT INTO category (name, slug, description) VALUES ($1, $2, $3) RETURNING id`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	err := s.db.QueryRowContext(ctx, query, c.Name, c.Slug, c.Description).Scan(&c.ID)
	if err != nil {
		return err
	}

	return nil
}

func (s *CategoryStore) GetByID(ctx context.Context, id int64) (*Category, error) {
	const query = `SELECT id, name, slug, description FROM category WHERE id = $1`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	c := &Category{}
	err := s.db.QueryRowContext(ctx, query, id).Scan(&c.ID, &c.Name, &c.Slug, &c.Description)
	if err != nil {
		switch {
		case err == sql.ErrNoRows:
			return nil, ErrNotFound
		default:
			return nil, err
		}
	}

	return c, nil
}

func (s *CategoryStore) GetBySlug(ctx context.Context, slug string) (*Category, error) {
    const query = `SELECT id, name, slug, description FROM category WHERE slug = $1`

    ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
    defer cancel()

    c := &Category{}
    err := s.db.QueryRowContext(ctx, query, slug).Scan(&c.ID, &c.Name, &c.Slug, &c.Description)
    if err != nil {
        switch {
        case err == sql.ErrNoRows:
            return nil, ErrNotFound
        default:
            return nil, err
        }
    }

    return c, nil
}
