package store

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

type PaymentMethod struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	IsActive    bool   `json:"is_active"`
}

type PaymentMethodStore struct {
	db *sql.DB
}

func (s *PaymentMethodStore) Create(ctx context.Context, pm *PaymentMethod) error {
	query := `INSERT INTO payment_methods (name, description, is_active) VALUES ($1, $2, $3) RETURNING id`
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	return s.db.QueryRowContext(ctx, query, pm.Name, pm.Description, pm.IsActive).Scan(&pm.ID)
}

func (s *PaymentMethodStore) GetAll(ctx context.Context) ([]*PaymentMethod, error) {
	query := `SELECT id, name, description, is_active FROM payment_methods`
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var paymentMethods []*PaymentMethod
	for rows.Next() {
		pm := &PaymentMethod{}
		err := rows.Scan(&pm.ID, &pm.Name, &pm.Description, &pm.IsActive)
		if err != nil {
			return nil, err
		}
		paymentMethods = append(paymentMethods, pm)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return paymentMethods, nil
}

func (s *PaymentMethodStore) GetByID(ctx context.Context, id int64) (*PaymentMethod, error) {
	query := `SELECT id, name, description, is_active FROM payment_methods WHERE id = $1`
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	pm := &PaymentMethod{}
	err := s.db.QueryRowContext(ctx, query, id).Scan(&pm.ID, &pm.Name, &pm.Description, &pm.IsActive)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return pm, nil
}

func (s *PaymentMethodStore) Update(ctx context.Context, pm *PaymentMethod) error {
	query := `UPDATE payment_methods SET name = $1, description = $2, is_active = $3 WHERE id = $4`
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	res, err := s.db.ExecContext(ctx, query, pm.Name, pm.Description, pm.IsActive, pm.ID)
	if err != nil {
		return err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if affected == 0 {
		return ErrNotFound
	}
	return nil
}

func (s *PaymentMethodStore) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM payment_methods WHERE id = $1`
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	res, err := s.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if affected == 0 {
		return ErrNotFound
	}
	return nil
}
