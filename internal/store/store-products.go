package store

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/lib/pq"
)

type Product struct {
	ID            int64     `json:"id" `
	Name          string    `json:"name" `
	Slug          string    `json:"slug" `
	Description   string    `json:"description,omitempty" `
	Price         float64   `json:"price" `
	DiscountPrice float64   `json:"discount_price" `
	Discount      float64   `json:"discount" `
	Rating        float64   `json:"rating" `
	Estimation    string    `json:"estimation" `
	Stock         int       `json:"stock" `
	Sold          int       `json:"sold" `
	IsForSale     bool      `json:"is_for_sale" `
	IsApproved    bool      `json:"is_approved" `
	CreatedAt     time.Time `json:"created_at" `
	UpdatedAt     time.Time `json:"updated_at" `
	ImageUrls     []string  `json:"image_urls" `
	Version      int       `json:"version"`
	Category      Category  `json:"category" `
	Toko          Toko      `json:"toko" `
	Comments      []Comment `json:"comments" `
}

type ProductStore struct {
	db *sql.DB
}

func (s *ProductStore) Create(ctx context.Context, p *Product) error {
	const query = `INSERT INTO product (name, slug, description, price, discount_price, discount, rating, estimation, stock, sold, is_for_sale, is_approved, image_urls, category_id, toko_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15) RETURNING id, created_at, updated_at`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	err := s.db.QueryRowContext(ctx, query, p.Name, p.Slug, p.Description, p.Price, p.DiscountPrice, p.Discount, p.Rating, p.Estimation, p.Stock, p.Sold, p.IsForSale, p.IsApproved, pq.Array(p.ImageUrls), p.Category.ID, p.Toko.ID).Scan(&p.ID, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (s *ProductStore) GetByID(ctx context.Context, id int64) (*Product, error) {
	const query = `SELECT id, name, slug, description, price, discount_price, discount, rating, estimation, stock, sold, is_for_sale, is_approved, created_at, updated_at, image_urls, category_id, toko_id , version FROM product WHERE id = $1`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	p := &Product{}
	err := s.db.QueryRowContext(ctx, query, id).Scan(&p.ID, &p.Name, &p.Slug, &p.Description, &p.Price, &p.DiscountPrice, &p.Discount, &p.Rating, &p.Estimation, &p.Stock, &p.Sold, &p.IsForSale, &p.IsApproved, &p.CreatedAt, &p.UpdatedAt, pq.Array(&p.ImageUrls), &p.Category.ID, &p.Toko.ID, &p.Version)
	if err != nil {
		switch {
		case err == sql.ErrNoRows:
			return nil, ErrNotFound
		default:
			return nil, err
		}
	}

	return p, nil
}

func (s *ProductStore) Update(ctx context.Context, p *Product) error {
	query := `UPDATE product SET name = $1, slug = $2, description = $3, price = $4, discount_price = $5, discount = $6, rating = $7, estimation = $8, stock = $9, sold = $10, is_for_sale = $11, is_approved = $12, image_urls = $13, version = version + 1, updated_at = now() WHERE id = $14 AND version = $15 RETURNING version, updated_at`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	err := s.db.QueryRowContext(ctx, query, p.Name, p.Slug, p.Description, p.Price, p.DiscountPrice, p.Discount, p.Rating, p.Estimation, p.Stock, p.Sold, p.IsForSale, p.IsApproved, pq.Array(p.ImageUrls), p.ID, p.Version).Scan(&p.Version, &p.UpdatedAt)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrNotFound
		default:
			return err
		}
	}

	return nil
}

func (s *ProductStore) Delete(ctx context.Context, id int64) error {
	const query = `DELETE FROM product WHERE id = $1`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	res, err := s.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrNotFound
	}

	return nil

}
