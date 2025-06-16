package store

import (
    "context"
    "database/sql"
    "errors"
    "time"
)

type ShippingMethod struct {
    ID          int64   `json:"id"`
    Name        string  `json:"name"`
    Description string  `json:"description,omitempty"`
    Price       float64 `json:"price"`
    IsActive    bool    `json:"is_active"`
}

type ShippingMethodStore struct {
    db *sql.DB
}

func (s *ShippingMethodStore) Create(ctx context.Context, sm *ShippingMethod) error {
    query := `INSERT INTO shipping_methods (name, description, price, is_active) VALUES ($1, $2, $3, $4) RETURNING id`
    ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
    defer cancel()
    return s.db.QueryRowContext(ctx, query, sm.Name, sm.Description, sm.Price, sm.IsActive).Scan(&sm.ID)
}

func (s *ShippingMethodStore) GetAll(ctx context.Context) ([]*ShippingMethod, error) {
	query := `SELECT id, name, description, price, is_active FROM shipping_methods`
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var shippingMethods []*ShippingMethod
	for rows.Next() {
		sm := &ShippingMethod{}
		err := rows.Scan(&sm.ID, &sm.Name, &sm.Description, &sm.Price, &sm.IsActive)
		if err != nil {
			return nil, err
		}
		shippingMethods = append(shippingMethods, sm)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return shippingMethods, nil
}

func (s *ShippingMethodStore) GetByID(ctx context.Context, id int64) (*ShippingMethod, error) {
    query := `SELECT id, name, description, price, is_active FROM shipping_methods WHERE id = $1`
    ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
    defer cancel()
    sm := &ShippingMethod{}
    err := s.db.QueryRowContext(ctx, query, id).Scan(&sm.ID, &sm.Name, &sm.Description, &sm.Price, &sm.IsActive)
    if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return nil, ErrNotFound
        }
        return nil, err
    }
    return sm, nil
}

func (s *ShippingMethodStore) Update(ctx context.Context, sm *ShippingMethod) error {
    query := `UPDATE shipping_methods SET name = $1, description = $2, price = $3, is_active = $4 WHERE id = $5`
    ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
    defer cancel()
    res, err := s.db.ExecContext(ctx, query, sm.Name, sm.Description, sm.Price, sm.IsActive, sm.ID)
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

func (s *ShippingMethodStore) Delete(ctx context.Context, id int64) error {
    query := `DELETE FROM shipping_methods WHERE id = $1`
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

func (s *ShippingMethodStore) List(ctx context.Context) ([]*ShippingMethod, error) {
    query := `SELECT id, name, description, price, is_active FROM shipping_methods`
    ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
    defer cancel()
    rows, err := s.db.QueryContext(ctx, query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var methods []*ShippingMethod
    for rows.Next() {
        sm := &ShippingMethod{}
        err := rows.Scan(&sm.ID, &sm.Name, &sm.Description, &sm.Price, &sm.IsActive)
        if err != nil {
            return nil, err
        }
        methods = append(methods, sm)
    }
    return methods, nil
}