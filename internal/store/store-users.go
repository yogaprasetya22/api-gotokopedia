package store

import (
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

func (s *UserStore) Create(u *User) error {
	const query = `INSERT INTO user (username, email, password, is_active) VALUES ($1, $2, $3, $4) RETURNING id, created_at`
	err := s.db.QueryRow(query, u.Username, u.Email, u.Password.hash, u.IsActive).Scan(&u.ID, &u.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}
