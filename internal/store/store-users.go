package store

import (
	"context"
	"database/sql"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        int64    `json:"id"`
	Username  string   `json:"username"`
	Email     string   `json:"email"`
	Password  password `json:"-"`
	CreatedAt string   `json:"created_at"`
	IsActive  bool     `json:"is_active"`
}

type password struct {
	text *string
	hash []byte
}

func (p *password) Set(text string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(text), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	p.text = &text
	p.hash = hash

	return nil
}

func (p *password) Compare(text string) error {
	return bcrypt.CompareHashAndPassword(p.hash, []byte(text))
}

type UserStore struct {
	db *sql.DB
}

func (s *UserStore) Create(ctx context.Context, u *User) error {
	const query = `INSERT INTO users (username, email, password,is_active) VALUES ($1, $2, $3, $4) RETURNING id, created_at`

	err := s.db.QueryRowContext(ctx, query, u.Username, u.Email, u.Password.hash, u.IsActive).Scan(&u.ID, &u.CreatedAt)

	if err != nil {
		return err
	}
	
	return nil
}

func (s *UserStore) GetByID(ctx context.Context, id int64) (*User, error) {
	const query = `SELECT id, username, email, password, created_at, is_active FROM users WHERE id = $1`

	u := &User{}
	err := s.db.QueryRowContext(ctx, query, id).Scan(&u.ID, &u.Username, &u.Email, &u.Password.hash, &u.CreatedAt, &u.IsActive)
	if err != nil {
		switch {
		case err == sql.ErrNoRows:
			return nil, ErrNotFound
		default:
			return nil, err
		}
	}

	return u, nil
}
