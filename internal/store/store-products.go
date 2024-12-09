package store

import (
	"context"
	"database/sql"
)

type Products struct {
	ID      int64  `json:"id"`
	Content string `json:"content"`
	Title   string `json:"title"`
	UserID  int64  `json:"user_id"`
}

type ProductStore struct {
	db *sql.DB
}

func (s *ProductStore) Create(ctx context.Context, p *Products) error {
	const query = `INSERT INTO products (title, content, user_id) VALUES ($1, $2, $3) RETURNING id`
	err := s.db.QueryRowContext(ctx, query, p.Title, p.Content, p.UserID).Scan(&p.ID)
	if err != nil {
		return err
	}

	return nil
}

func (s *ProductStore) GetByID(ctx context.Context, id int64) (*Products, error) {
	const query = `SELECT id, title, content, user_id FROM products WHERE id = $1`
	p := &Products{}
	err := s.db.QueryRowContext(ctx, query, id).Scan(&p.ID, &p.Title, &p.Content, &p.UserID)
	if err != nil {
		return nil, err
	}

	return p, nil
}
