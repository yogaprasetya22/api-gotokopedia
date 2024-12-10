package store

import (
	"context"
	"database/sql"
)

type Category struct {
	ID          string  `json:"id" `
	Name        string  `json:"name" `
	Slug        string  `json:"slug" `
	Description *string `json:"description,omitempty" `
}

type CategoryStore struct {
	db *sql.DB
}

func (s *CategoryStore) Create(ctx context.Context, c *Category) error {
	const query = `INSERT INTO categories (name, slug, description) VALUES ($1, $2, $3) RETURNING id`
	err := s.db.QueryRowContext(ctx, query, c.Name, c.Slug, c.Description).Scan(&c.ID)
	if err != nil {
		return err
	}

	return nil
}
