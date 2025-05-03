package store

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type CartDetailResponse struct {
	ID        int64             `json:"id"`
	UserID    int64             `json:"user_id"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
	Stores    []CartStoreDetail `json:"stores"`
}

type CartStoreDetail struct {
	ID       uuid.UUID        `json:"id"`
	Toko     TokoResponse     `json:"toko"`
	Items    []CartItemDetail `json:"items"`
	Subtotal float64          `json:"subtotal"`
}

type CartItemDetail struct {
	ID            uuid.UUID `json:"id"`
	ProductID     int64     `json:"product_id"`
	Name          string    `json:"name"`
	ImageURL      []string  `json:"image_url"`
	Price         float64   `json:"price"`
	Discount      float64   `json:"discount"`
	DiscountPrice float64   `json:"discount_price" `
	Quantity      int       `json:"quantity"`
	TotalPrice    float64   `json:"total_price"`
}

type TokoResponse struct {
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	Slug         string `json:"slug"`
	ImageProfile string `json:"image_profile"`
}

func (cs *CartStore) GetDetailCartByCartStoreID(ctx context.Context, cartStoreID uuid.UUID, userID int64) (*CartDetailResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	tx, err := cs.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// Get main cart information
	const cartQuery = `
        SELECT id, user_id, created_at, updated_at
        FROM carts
        WHERE id = (SELECT cart_id FROM cart_stores WHERE id = $1)
        AND user_id = $2
    `
	var cart CartDetailResponse
	if err := tx.QueryRowContext(ctx, cartQuery, cartStoreID, userID).Scan(
		&cart.ID, &cart.UserID, &cart.CreatedAt, &cart.UpdatedAt,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("cart not found")
		}
		return nil, err
	}

	// Get cart_store information
	const cartStoreQuery = `
        SELECT id, toko_id
        FROM cart_stores
        WHERE id = $1
        AND cart_id = $2
    `
	var cartStore CartStoreDetail
	if err := tx.QueryRowContext(ctx, cartStoreQuery, cartStoreID, cart.ID).Scan(
		&cartStore.ID, &cartStore.Toko.ID,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("cart store not found")
		}
		return nil, err
	}

	// Get toko information
	const tokoQuery = `
        SELECT name, slug, image_profile
        FROM tokos
        WHERE id = $1
    `
	if err := tx.QueryRowContext(ctx, tokoQuery, cartStore.Toko.ID).Scan(
		&cartStore.Toko.Name, &cartStore.Toko.Slug, &cartStore.Toko.ImageProfile,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("toko not found")
		}
		return nil, err
	}

	// Get items for the cart_store
	const itemQuery = `
        SELECT 
            ci.id, 
            ci.product_id, 
            p.name, 
            p.image_urls, 
            p.price,
            p.discount_price,
            p.discount, 
            ci.quantity
        FROM cart_items ci
        JOIN products p ON p.id = ci.product_id
        WHERE ci.cart_store_id = $1
    `
	itemRows, err := tx.QueryContext(ctx, itemQuery, cartStore.ID)
	if err != nil {
		return nil, err
	}
	defer itemRows.Close()

	var items []CartItemDetail
	var subtotal float64

	for itemRows.Next() {
		var item CartItemDetail
		var price, discountPrice, discount float64

		if err := itemRows.Scan(
			&item.ID,
			&item.ProductID,
			&item.Name,
			pq.Array(&item.ImageURL),
			&price,
			&discountPrice,
			&discount,
			&item.Quantity,
		); err != nil {
			return nil, err
		}

		// Determine which price to use
		item.Price = price
		item.DiscountPrice = discountPrice
		item.Discount = discount

		item.TotalPrice = price * float64(item.Quantity)
		subtotal += item.TotalPrice

		items = append(items, item)
	}

	cartStore.Items = items
	cartStore.Subtotal = subtotal
	cart.Stores = []CartStoreDetail{cartStore}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &cart, nil
}
