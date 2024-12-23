package store

import (
	"context"
	"database/sql"

	"github.com/lib/pq"
)

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
	query := `SELECT p.id, p.name, p.slug, p.country, p.description, p.price, p.discount_price, p.discount, p.rating, p.estimation, p.stock, p.sold, p.is_for_sale, p.is_approved, p.created_at, p.updated_at, p.image_urls, 
           c.id, c.name, c.slug, 
           t.id, t.user_id, t.slug, t.name, t.country, t.created_at, 
           u.id, u.username, u.email, u.created_at, u.is_active
    FROM products p
    JOIN category c ON p.category_id = c.id
    JOIN tokos t ON p.toko_id = t.id
    JOIN users u ON t.user_id = u.id
    WHERE p.is_approved = true AND p.category_id = ANY($1) AND (p.name ILIKE '%' || $4 || '%' OR p.description ILIKE '%' || $4 || '%')
   ORDER BY
            p.created_at ` + fq.Sort + `
    LIMIT $2 OFFSET $3`

	rows, err := tx.QueryContext(ctx, query, pq.Array(categoryIDs), fq.Limit, fq.Offset, fq.Search)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*Product
	for rows.Next() {
		p := &Product{}
		err := rows.Scan(&p.ID, &p.Name, &p.Slug, &p.Country, &p.Description, &p.Price, &p.DiscountPrice, &p.Discount, &p.Rating, &p.Estimation, &p.Stock, &p.Sold, &p.IsForSale, &p.IsApproved, &p.CreatedAt, &p.UpdatedAt, pq.Array(&p.ImageUrls),
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
	query := `SELECT p.id, p.name, p.slug, p.country, p.description, p.price, p.discount_price, p.discount, p.rating, p.estimation, p.stock, p.sold, p.is_for_sale, p.is_approved, p.created_at, p.updated_at, p.image_urls, 
               c.id, c.name, c.slug, 
               t.id, t.user_id, t.slug, t.name, t.country, t.created_at, 
               u.id, u.username, u.email, u.created_at, u.is_active
        FROM products p
        JOIN category c ON p.category_id = c.id
        JOIN tokos t ON p.toko_id = t.id
        JOIN users u ON t.user_id = u.id
        WHERE p.is_approved = true AND (SELECT id FROM category WHERE slug = $1) = c.id AND (p.name ILIKE '%' || $4 || '%' OR p.description ILIKE '%' || $4 || '%')
        ORDER BY p.created_at ` + fq.Sort + `
        LIMIT $2 OFFSET $3`

	rows, err := tx.QueryContext(ctx, query, fq.Category, fq.Limit, fq.Offset, fq.Search)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*Product
	for rows.Next() {
		p := &Product{}
		err := rows.Scan(&p.ID, &p.Name, &p.Slug, &p.Country, &p.Description, &p.Price, &p.DiscountPrice, &p.Discount, &p.Rating, &p.Estimation, &p.Stock, &p.Sold, &p.IsForSale, &p.IsApproved, &p.CreatedAt, &p.UpdatedAt, pq.Array(&p.ImageUrls),
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

func (s *ProductStore) commentsForProducts(ctx context.Context, tx *sql.Tx, products []*Product) error {
	productIDs := make([]int64, len(products))
	for i, p := range products {
		productIDs[i] = p.ID
	}

	const commentsQuery = `
    SELECT c.product_id, c.id, c.user_id, c.content, c.created_at,
           u.id, u.username, u.email, u.created_at, u.is_active
    FROM comments c
    JOIN users u ON c.user_id = u.id
    WHERE c.product_id = ANY($1)`

	commentRows, err := tx.QueryContext(ctx, commentsQuery, pq.Array(productIDs))
	if err != nil {
		return err
	}
	defer commentRows.Close()

	commentsMap := make(map[int64][]Comment)
	for commentRows.Next() {
		var c Comment
		var user User
		var productID int64
		err := commentRows.Scan(&productID, &c.ID, &c.UserID, &c.Content, &c.CreatedAt,
			&user.ID, &user.Username, &user.Email, &user.CreatedAt, &user.IsActive)
		if err != nil {
			return err
		}
		c.User = user
		commentsMap[productID] = append(commentsMap[productID], c)
	}

	if err := commentRows.Err(); err != nil {
		return err
	}

	for _, p := range products {
		p.Comments = commentsMap[p.ID]
	}

	return nil
}
