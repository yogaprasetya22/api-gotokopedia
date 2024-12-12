package store

import (
	"context"
	"database/sql"
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
	Category      Category  `json:"category" `
	Toko          Toko      `json:"toko" `
}

type ProductStore struct {
	db *sql.DB
}

func (s *ProductStore) Create(ctx context.Context, p *Product) error {
	const query = `INSERT INTO product (name, slug, description, price, discount_price, discount, rating, estimation, stock, sold, is_for_sale, is_approved, created_at, updated_at, image_urls, category_id, toko_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17) RETURNING id`
	err := s.db.QueryRowContext(ctx, query, p.Name, p.Slug, p.Description, p.Price, p.DiscountPrice, p.Discount, p.Rating, p.Estimation, p.Stock, p.Sold, p.IsForSale, p.IsApproved, p.CreatedAt, p.UpdatedAt, pq.Array(p.ImageUrls)).Scan(&p.ID)
	if err != nil {
		return err
	}

	return nil
}

func (s *ProductStore) GetByID(ctx context.Context, id int64) (*Product, error) {
	const query = `SELECT id, name, slug, description, price, discount_price, discount, rating, estimation, stock, sold, is_for_sale, is_approved, created_at, updated_at, image_urls, category_id, toko_id FROM product WHERE id = $1`

	p := &Product{}
	err := s.db.QueryRowContext(ctx, query, id).Scan(&p.ID, &p.Name, &p.Slug, &p.Description, &p.Price, &p.DiscountPrice, &p.Discount, &p.Rating, &p.Estimation, &p.Stock, &p.Sold, &p.IsForSale, &p.IsApproved, &p.CreatedAt, &p.UpdatedAt, pq.Array(&p.ImageUrls), &p.Category.ID, &p.Toko.ID)
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
