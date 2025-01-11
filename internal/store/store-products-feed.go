package store

import (
	"context"
	"database/sql"
	"time"

	"github.com/lib/pq"
)

type DetailProduct struct {
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
	Ulasan        Ulasan    `json:"ulasan"` // New field for review stats
}

type Ulasan struct {
	TotalUlasan     int         `json:"total_ulasan"`
	TotalRating     int         `json:"total_rating"`
	RatingBreakdown map[int]int `json:"rating_breakdown"` // Count of each star rating
}

func (s *ProductStore) GetProductByTokos(ctx context.Context, slug_toko string, query PaginatedFeedQuery) ([]*Product, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	products, err := s.productByTokoSlug(ctx, tx, slug_toko, query)
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

func (s *ProductStore) productByTokoSlug(ctx context.Context, tx *sql.Tx, slug_toko string, fq PaginatedFeedQuery) ([]*Product, error) {
	query := `SELECT p.id, p.name, p.slug, p.country, p.description, p.price, p.discount_price, p.discount, p.estimation, p.stock, p.sold, p.is_for_sale, p.is_approved, p.created_at, p.updated_at, p.image_urls,
		c.id, c.name, c.slug, 
		t.id, t.user_id, t.slug, t.name, t.country, t.created_at,
		u.id, u.username, u.email, u.picture, u.created_at, u.is_active
	FROM products p
	JOIN category c ON p.category_id = c.id
	JOIN tokos t ON p.toko_id = t.id
	JOIN users u ON t.user_id = u.id
	WHERE t.slug = $1 AND p.is_approved = true AND (p.name ILIKE '%' || $4 || '%' OR p.description ILIKE '%' || $4 || '%')
	ORDER BY p.created_at ` + fq.Sort + `
	LIMIT $2 OFFSET $3`

	skip := fq.Offset * (fq.Limit)

	rows, err := tx.QueryContext(ctx, query, slug_toko, fq.Limit, skip, fq.Search)
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

func (s *ProductStore) GetProduct(ctx context.Context, slug_toko, slug_product string) (*DetailProduct, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	product, err := s.productBySlug(ctx, tx, slug_toko, slug_product)
	if err != nil {
		return nil, err
	}

	err = s.commentsForProductsDetail(ctx, tx, []*DetailProduct{product})
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return product, nil
}

func (s *ProductStore) productBySlug(ctx context.Context, tx *sql.Tx, slug_toko, slug_product string) (*DetailProduct, error) {
	query := `SELECT p.id, p.name, p.slug, p.country, p.description, p.price, p.discount_price, p.discount, p.estimation, p.stock, p.sold, p.is_for_sale, p.is_approved, p.created_at, p.updated_at, p.image_urls,
		c.id, c.name, c.slug, 
		t.id, t.user_id, t.slug, t.name, t.country, t.created_at,
		u.id, u.username, u.email, u.picture, u.created_at, u.is_active
	FROM products p
	JOIN category c ON p.category_id = c.id
	JOIN tokos t ON p.toko_id = t.id
	JOIN users u ON t.user_id = u.id
	WHERE t.slug = $1 AND p.slug = $2 AND p.is_approved = true`

	p := &DetailProduct{}
	err := tx.QueryRowContext(ctx, query, slug_toko, slug_product).Scan(&p.ID, &p.Name, &p.Slug, &p.Country, &p.Description, &p.Price, &p.DiscountPrice, &p.Discount, &p.Estimation, &p.Stock, &p.Sold, &p.IsForSale, &p.IsApproved, &p.CreatedAt, &p.UpdatedAt, pq.Array(&p.ImageUrls),
		&p.Category.ID, &p.Category.Name, &p.Category.Slug,
		&p.Toko.ID, &p.Toko.UserID, &p.Toko.Slug, &p.Toko.Name, &p.Toko.Country, &p.Toko.CreatedAt,
		&p.Toko.User.ID, &p.Toko.User.Username, &p.Toko.User.Email, &p.Toko.User.Picture, &p.Toko.User.CreatedAt, &p.Toko.User.IsActive)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (s *ProductStore) GetProductFeed(ctx context.Context, categoryIDs []int64, fq PaginatedFeedQuery) ([]*Product, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	products, err := s.productFeed(ctx, tx, categoryIDs, fq)
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

func (s *ProductStore) productFeed(ctx context.Context, tx *sql.Tx, categoryIDs []int64, fq PaginatedFeedQuery) ([]*Product, error) {
	query := `SELECT p.id, p.name, p.slug, p.country, p.description, p.price, p.discount_price, p.discount, p.estimation, p.stock, p.sold, p.is_for_sale, p.is_approved, p.created_at, p.updated_at, p.image_urls, 
           c.id, c.name, c.slug, 
           t.id, t.user_id, t.slug, t.name, t.country, t.created_at, 
           u.id, u.username, u.email, u.picture, u.created_at, u.is_active
    FROM products p
    JOIN category c ON p.category_id = c.id
    JOIN tokos t ON p.toko_id = t.id
    JOIN users u ON t.user_id = u.id
    WHERE p.is_approved = true AND p.category_id = ANY($1) AND (p.name ILIKE '%' || $4 || '%' OR p.description ILIKE '%' || $4 || '%')
   ORDER BY
            p.created_at ` + fq.Sort + `
    LIMIT $2 OFFSET $3`

	skip := fq.Offset * (fq.Limit)

	rows, err := tx.QueryContext(ctx, query, pq.Array(categoryIDs), fq.Limit, skip, fq.Search)
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

func (s *ProductStore) GetProductCategoryFeed(ctx context.Context, fq PaginatedFeedQuery) ([]*Product, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	products, err := s.productsByCategorySlug(ctx, tx, fq)
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

func (s *ProductStore) productsByCategorySlug(ctx context.Context, tx *sql.Tx, fq PaginatedFeedQuery) ([]*Product, error) {
	query := `SELECT p.id, p.name, p.slug, p.country, p.description, p.price, p.discount_price, p.discount, p.estimation, p.stock, p.sold, p.is_for_sale, p.is_approved, p.created_at, p.updated_at, p.image_urls, 
               c.id, c.name, c.slug, 
               t.id, t.user_id, t.slug, t.name, t.country, t.created_at, 
               u.id, u.username, u.email, u.picture, u.created_at, u.is_active
        FROM products p
        JOIN category c ON p.category_id = c.id
        JOIN tokos t ON p.toko_id = t.id
        JOIN users u ON t.user_id = u.id
        WHERE p.is_approved = true AND (SELECT id FROM category WHERE slug = $1) = c.id AND (p.name ILIKE '%' || $4 || '%' OR p.description ILIKE '%' || $4 || '%')
        ORDER BY p.sold ` + fq.Sort + `
        LIMIT $2 OFFSET $3`

	skip := fq.Offset * (fq.Limit)

	rows, err := tx.QueryContext(ctx, query, fq.Category, fq.Limit, skip, fq.Search)
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
