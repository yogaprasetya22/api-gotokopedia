package store

import (
	"context"
	"database/sql"
	"time"

	"github.com/lib/pq"
)

type Product struct {
	ID            int64     `json:"id,omitempty" `
	Name          string    `json:"name" `
	Slug          string    `json:"slug" `
	Country       string    `json:"country" `
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
	Version       int       `json:"version"`
	Category      Category  `json:"category" `
	Toko          Toko      `json:"toko" `
}

type ProductStore struct {
	db *sql.DB
}

func (s *ProductStore) Create(ctx context.Context, p *Product) error {
	const query = `INSERT INTO products (name, slug, country, description, price, discount_price, discount, estimation, stock, sold, is_for_sale, is_approved, image_urls, category_id, toko_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15) RETURNING id, created_at, updated_at`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	err := s.db.QueryRowContext(ctx, query, p.Name, p.Slug, p.Country, p.Description, p.Price, p.DiscountPrice, p.Discount, p.Estimation, p.Stock, p.Sold, p.IsForSale, p.IsApproved, pq.Array(p.ImageUrls), p.Category.ID, p.Toko.ID).Scan(&p.ID, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (s *ProductStore) GetAllProduct(ctx context.Context, fq PaginatedFeedQuery) ([]*Product, error) {
	query := `SELECT p.id, p.name, p.slug, p.country, p.description, p.price, p.discount_price, p.discount, p.estimation, p.stock, p.sold, p.is_for_sale, p.is_approved, p.created_at, p.updated_at, p.image_urls, 
               c.id, c.name, c.slug, 
               t.id, t.user_id, t.slug, t.name, t.country, t.created_at, 
               u.id, u.username, u.email, u.created_at, u.is_active
        FROM products p
        JOIN category c ON p.category_id = c.id
        JOIN tokos t ON p.toko_id = t.id
        JOIN users u ON t.user_id = u.id
        WHERE p.is_approved = true
        ORDER BY p.sold DESC
        LIMIT $1 OFFSET $2`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	rows, err := s.db.QueryContext(ctx, query, fq.Limit, fq.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*Product
	for rows.Next() {
		p := &Product{}
		err := rows.Scan(&p.ID, &p.Name, &p.Slug, &p.Country, &p.Description, &p.Price, &p.DiscountPrice, &p.Discount, &p.Estimation, &p.Stock, &p.Sold, &p.IsForSale, &p.IsApproved, &p.CreatedAt, &p.UpdatedAt, pq.Array(&p.ImageUrls),
			&p.Category.ID, &p.Category.Name, &p.Category.Slug,
			&p.Toko.ID, &p.Toko.UserID, &p.Toko.Slug, &p.Toko.Name, &p.Toko.Country, &p.Toko.CreatedAt,
			&p.Toko.User.ID, &p.Toko.User.Username, &p.Toko.User.Email, &p.Toko.User.CreatedAt, &p.Toko.User.IsActive)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func (s *ProductStore) GetByID(ctx context.Context, id int64) (*Product, error) {
	const query = `SELECT id, name, slug, country, description, price, discount_price, discount, estimation, stock, sold, is_for_sale, is_approved, created_at, updated_at, image_urls, category_id, toko_id , version FROM products WHERE id = $1`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	product := &Product{}
	err := s.db.QueryRowContext(ctx, query, id).Scan(&product.ID, &product.Name, &product.Slug, &product.Country, &product.Description, &product.Price, &product.DiscountPrice, &product.Discount, &product.Estimation, &product.Stock, &product.Sold, &product.IsForSale, &product.IsApproved, &product.CreatedAt, &product.UpdatedAt, pq.Array(&product.ImageUrls), &product.Category.ID, &product.Toko.ID, &product.Version)
	if err != nil {
		switch {
		case err == sql.ErrNoRows:
			return nil, ErrNotFound
		default:
			return nil, err
		}
	}

	return product, nil
}

func (s *ProductStore) Update(ctx context.Context, product *Product) error {
	query := `UPDATE products SET name = $1, slug = $2, country = $3, description = $4, price = $5, discount_price = $6, discount = $7, estimation = $8, stock = $9, sold = $10, is_for_sale = $11, is_approved = $12, image_urls = $13, category_id = $14, toko_id = $15, version = version + 1 WHERE id = $16 AND version = $17 RETURNING updated_at`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	err := s.db.QueryRowContext(ctx, query, product.Name, product.Slug, product.Country, product.Description, product.Price, product.DiscountPrice, product.Discount, product.Estimation, product.Stock, product.Sold, product.IsForSale, product.IsApproved, pq.Array(product.ImageUrls), product.Category.ID, product.Toko.ID, product.ID, product.Version).Scan(&product.UpdatedAt)
	if err != nil {
		switch {
		case err == sql.ErrNoRows:
			return ErrNotFound
		default:
			return err
		}
	}

	return nil

}

func (s *ProductStore) Delete(ctx context.Context, id int64) error {
	const query = `DELETE FROM products WHERE id = $1`

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

func (s *ProductStore) GetByTokoID(ctx context.Context, id int64) ([]*Product, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	products, err := s.productByTokoID(ctx, tx, id)
	if err != nil {
		return nil, err
	}

	err = s.commentsForProducts(ctx, tx, products)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return products, nil
}

func (s *ProductStore) productByTokoID(ctx context.Context, tx *sql.Tx, id int64) ([]*Product, error) {
	query := `SELECT p.id, p.name, p.slug, p.country, p.description, p.price, p.discount_price, p.discount, p.estimation, p.stock, p.sold, p.is_for_sale, p.is_approved, p.created_at, p.updated_at, p.image_urls, 
		   c.id, c.name, c.slug, 
		   t.id, t.user_id, t.slug, t.name, t.country, t.image_profile, t.created_at, 
		   u.id, u.username, u.email, u.picture, u.created_at, u.is_active
	FROM products p
	JOIN category c ON p.category_id = c.id
	JOIN tokos t ON p.toko_id = t.id
	JOIN users u ON t.user_id = u.id
		WHERE t.id = $1`

	rows, err := tx.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var products []*Product
	for rows.Next() {
		p := &Product{}
		err := rows.Scan(&p.ID, &p.Name, &p.Slug, &p.Country, &p.Description, &p.Price, &p.DiscountPrice, &p.Discount, &p.Estimation, &p.Stock, &p.Sold, &p.IsForSale, &p.IsApproved, &p.CreatedAt, &p.UpdatedAt, pq.Array(&p.ImageUrls),
			&p.Category.ID, &p.Category.Name, &p.Category.Slug,
			&p.Toko.ID, &p.Toko.UserID, &p.Toko.Slug, &p.Toko.Name, &p.Toko.Country, &p.Toko.ImageProfile, &p.Toko.CreatedAt,
			&p.Toko.User.ID, &p.Toko.User.Username, &p.Toko.User.Email, &p.Toko.User.Picture, &p.Toko.User.CreatedAt, &p.Toko.User.IsActive)

		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func (s *ProductStore) commentsForProducts(ctx context.Context, tx *sql.Tx, products []*Product) error {
	productIDs := make([]int64, len(products))
	for i, p := range products {
		productIDs[i] = p.ID
	}

	const commentsQuery = `
    SELECT c.product_id, c.id, c.user_id, c.content, c.rating, c.created_at,
           u.id, u.username, u.email, u.picture, u.created_at, u.is_active
    FROM comments c
    JOIN users u ON c.user_id = u.id
    WHERE c.product_id = ANY($1)`

	commentRows, err := tx.QueryContext(ctx, commentsQuery, pq.Array(productIDs))
	if err != nil {
		return err
	}
	defer commentRows.Close()

	commentsMap := make(map[int64][]Comment)
	ratingSumMap := make(map[int64]float64)
	ratingCountMap := make(map[int64]int)

	for commentRows.Next() {
		var c Comment
		var user User
		var productID int64
		err := commentRows.Scan(&productID, &c.ID, &c.UserID, &c.Content, &c.Rating, &c.CreatedAt,
			&user.ID, &user.Username, &user.Email, &user.Picture, &user.CreatedAt, &user.IsActive)
		if err != nil {
			return err
		}
		c.User = user
		commentsMap[productID] = append(commentsMap[productID], c)
		ratingSumMap[productID] += float64(c.Rating)
		ratingCountMap[productID]++
	}

	if err := commentRows.Err(); err != nil {
		return err
	}

	for _, p := range products {
		// p.Comments = commentsMap[p.ID]
		if ratingCountMap[p.ID] > 0 {
			p.Rating = ratingSumMap[p.ID] / float64(ratingCountMap[p.ID])
		} else {
			p.Rating = 0
		}
	}

	return nil
}

func (s *ProductStore) commentsForProductsDetail(ctx context.Context, tx *sql.Tx, products []*DetailProduct) error {
	productIDs := make([]int64, len(products))
	for i, p := range products {
		productIDs[i] = p.ID
	}

	const commentsQuery = `
    SELECT c.product_id, c.id, c.user_id, c.content, c.rating, c.created_at,
           u.id, u.username, u.email, u.picture, u.created_at, u.is_active
    FROM comments c
    JOIN users u ON c.user_id = u.id
    WHERE c.product_id = ANY($1)`

	commentRows, err := tx.QueryContext(ctx, commentsQuery, pq.Array(productIDs))
	if err != nil {
		return err
	}
	defer commentRows.Close()

	commentsMap := make(map[int64][]Comment)
	ratingSumMap := make(map[int64]float64)
	ratingCountMap := make(map[int64]int)
	ratingBreakdownMap := make(map[int64]map[int]int) // Breakdown per rating

	for commentRows.Next() {
		var c Comment
		var user User
		var productID int64
		err := commentRows.Scan(&productID, &c.ID, &c.UserID, &c.Content, &c.Rating, &c.CreatedAt,
			&user.ID, &user.Username, &user.Email, &user.Picture, &user.CreatedAt, &user.IsActive)
		if err != nil {
			return err
		}
		c.User = user
		commentsMap[productID] = append(commentsMap[productID], c)
		ratingSumMap[productID] += float64(c.Rating)
		ratingCountMap[productID]++

		// Update breakdown
		if _, exists := ratingBreakdownMap[productID]; !exists {
			ratingBreakdownMap[productID] = make(map[int]int)
		}
		ratingBreakdownMap[productID][int(c.Rating)]++
	}

	if err := commentRows.Err(); err != nil {
		return err
	}

	for _, p := range products {
		// p.Comments = commentsMap[p.ID]
		if ratingCountMap[p.ID] > 0 {
			p.Rating = ratingSumMap[p.ID] / float64(ratingCountMap[p.ID])
		} else {
			p.Rating = 0
		}

		// Populate Ulasan
		ratingBreakdown := ratingBreakdownMap[p.ID]
		totalRating := 0
		for rating := 1; rating <= 5; rating++ { // Ensure all ratings from 1 to 5 are present
			if _, exists := ratingBreakdown[rating]; !exists {
				ratingBreakdown[rating] = 0
			}
			totalRating += rating * ratingBreakdown[rating]
		}

		p.Ulasan = Ulasan{
			TotalUlasan:    ratingCountMap[p.ID],
			TotalRating:    totalRating,
			RatingBreakdown: ratingBreakdown,
		}
	}

	return nil
}
