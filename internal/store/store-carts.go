package store

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

// Cart mewakili keranjang belanja pengguna
type Cart struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// Relasi
	Stores []CartStores `json:"stores,omitempty"`
}

// CartStore mewakili toko dalam keranjang
type CartStores struct {
	ID        uuid.UUID `json:"id"`
	CartID    int64     `json:"cart_id"`
	TokoID    int64     `json:"toko_id"`
	CreatedAt time.Time `json:"created_at"`

	// Relasi
	Toko  *Toko      `json:"toko,omitempty"`
	Items []CartItem `json:"items,omitempty"`
}

// CartItem mewakili item produk dalam keranjang
type CartItem struct {
	ID          uuid.UUID `json:"id"`
	CartID      int64     `json:"cart_id"`
	CartStoreID uuid.UUID `json:"cart_store_id,omitempty"` // Nullable di database
	ProductID   int64     `json:"product_id"`
	Quantity    int64     `json:"quantity"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`

	// Relasi
	Product *Product `json:"product,omitempty"`
}

type MetaCart struct {
	Cart       *Cart   `json:"cart"`
	TotalItems int64   `json:"total_items"`
	TotalPrice float64 `json:"total_price"`
}

type CartStore struct {
	db *sql.DB
}

func (s *CartStore) GetCartByUserIDPQ(ctx context.Context, userID int64, query PaginatedFeedQuery) (*MetaCart, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// 1. Dapatkan cart utama
	cart, err := s.getUserCart(ctx, tx, userID)
	if err != nil {
		return nil, err
	}

	// 3. Dapatkan stores dengan pagination
	stores, err := s.getCartStores(ctx, tx, cart.ID, query)
	if err != nil {
		return nil, err
	}

	// 2. Dapatkan total items dan harga sekaligus
	totalItems, totalPrice, err := s.GetCartTotals(ctx, cart.ID)
	if err != nil {
		return nil, err
	}

	// 4. Untuk setiap store, dapatkan items-nya
	for i := range stores {
		items, err := s.getCartItemsByStore(ctx, tx, cart.ID, stores[i].ID)
		if err != nil {
			return nil, err
		}
		stores[i].Items = items
	}

	cart.Stores = stores

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &MetaCart{
		Cart:       cart,
		TotalItems: totalItems,
		TotalPrice: totalPrice,
	}, nil
}

// getUserCart mendapatkan cart utama user
func (s *CartStore) GetCartByUserID(ctx context.Context, userID int64) (*Cart, error) {
	const query = `
		SELECT id, user_id, created_at, updated_at 
		FROM carts 
		WHERE user_id = $1
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	var cart Cart
	err := s.db.QueryRowContext(ctx, query, userID).Scan(
		&cart.ID,
		&cart.UserID,
		&cart.CreatedAt,
		&cart.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &cart, nil
}

// getUserCart mendapatkan cart utama user
func (s *CartStore) getUserCart(ctx context.Context, tx *sql.Tx, userID int64) (*Cart, error) {
	const query = `
		SELECT id, user_id, created_at, updated_at 
		FROM carts 
		WHERE user_id = $1
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	var cart Cart
	err := tx.QueryRowContext(ctx, query, userID).Scan(
		&cart.ID,
		&cart.UserID,
		&cart.CreatedAt,
		&cart.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// Jika cart belum ada, buat baru
			return s.Create(ctx, tx, userID)
		}
		return nil, err
	}

	return &cart, nil
}

func (s *CartStore) GetCartTotals(ctx context.Context, cartID int64) (totalItems int64, totalPrice float64, err error) {
	query := `
		SELECT 
			COALESCE(SUM(ci.quantity), 0) AS total_items,
			COALESCE(SUM(ci.quantity * p.price), 0) AS total_price
		FROM 
			cart_items ci
		JOIN 
			products p ON ci.product_id = p.id
		WHERE 
			ci.cart_id = $1`

	err = s.db.QueryRowContext(ctx, query, cartID).Scan(&totalItems, &totalPrice)
	if err != nil {
		return 0, 0, err
	}
	return totalItems, totalPrice, nil
}

// GetCartStoresByID mendapatkan cart store berdasarkan []ID
func (s *CartStore) GetCartStoresByID(ctx context.Context, cartStoreIDs []uuid.UUID) ([]CartStores, error) {
	if len(cartStoreIDs) == 0 {
		return []CartStores{}, nil
	}

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	query := `
		SELECT cs.id, cs.cart_id, cs.toko_id, cs.created_at,
			   t.id, t.user_id, t.slug, t.name, t.image_profile, t.country, t.created_at
		FROM cart_stores cs
		JOIN tokos t ON cs.toko_id = t.id
		WHERE cs.id = ANY($1)
	ORDER BY cs.created_at`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	rows, err := tx.QueryContext(ctx, query, pq.Array(cartStoreIDs))
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var stores []CartStores
	for rows.Next() {
		var store CartStores
		var toko Toko
		err := rows.Scan(
			&store.ID,
			&store.CartID,
			&store.TokoID,
			&store.CreatedAt,
			&toko.ID,
			&toko.UserID,
			&toko.Slug,
			&toko.Name,
			&toko.ImageProfile,
			&toko.Country,
			&toko.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		store.Toko = &toko
		stores = append(stores, store)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	for i := range stores {
		items, err := s.getCartItemsByStore(ctx, tx, stores[i].CartID, stores[i].ID)
		if err != nil {
			return nil, err
		}
		stores[i].Items = items
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return stores, nil
}

// getCartStores mendapatkan semua store dalam cart
func (s *CartStore) getCartStores(ctx context.Context, tx *sql.Tx, cartID int64, fq PaginatedFeedQuery) ([]CartStores, error) {
	query := `
		SELECT cs.id, cs.cart_id, cs.toko_id, cs.created_at,
               t.id, t.user_id, t.slug, t.name, t.image_profile, t.country, t.created_at
        FROM cart_stores cs
        JOIN tokos t ON cs.toko_id = t.id
        WHERE cs.cart_id = $1
        ORDER BY cs.created_at ` + fq.Sort + `
        LIMIT $2 OFFSET $3`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	offset := fq.Offset
	if fq.Offset > 0 && fq.Limit > 0 {
		offset = fq.Offset * fq.Limit
	}

	rows, err := tx.QueryContext(ctx, query, cartID, fq.Limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stores []CartStores
	for rows.Next() {
		var store CartStores
		var toko Toko
		err := rows.Scan(
			&store.ID,
			&store.CartID,
			&store.TokoID,
			&store.CreatedAt,
			&toko.ID,
			&toko.UserID,
			&toko.Slug,
			&toko.Name,
			&toko.ImageProfile,
			&toko.Country,
			&toko.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		store.Toko = &toko
		stores = append(stores, store)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return stores, nil
}

// getCartItemsByStore mendapatkan semua item dalam cart_store tertentu
func (s *CartStore) getCartItemsByStore(ctx context.Context, tx *sql.Tx, cartID int64, cartStoreID uuid.UUID) ([]CartItem, error) {
	const query = `
        SELECT 
            ci.id, ci.cart_id, ci.cart_store_id, ci.product_id, ci.quantity, 
            ci.created_at, ci.updated_at,
            p.id, p.name, p.slug, p.description, p.country, 
            p.price, p.discount_price, p.discount, p.estimation, 
            p.stock, p.sold, p.is_for_sale, p.is_approved, 
            p.image_urls, p.created_at, p.updated_at, p.version,
            c.id, c.name, c.slug, c.description,
            t.id, t.user_id, t.slug, t.name, t.image_profile, t.country, t.created_at,
			u.id, u.picture, u.username, u.email, u.is_active
		FROM cart_items ci
		JOIN products p ON ci.product_id = p.id
		JOIN category c ON p.category_id = c.id
		JOIN tokos t ON p.toko_id = t.id
		JOIN users u ON t.user_id = u.id
		WHERE ci.cart_id = $1 AND ci.cart_store_id = $2
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	rows, err := tx.QueryContext(ctx, query, cartID, cartStoreID)
	if err != nil {
		return nil, fmt.Errorf("failed to query cart items: %v and cartID: %d and storeID: %s", err, cartID, cartStoreID.String())
	}
	defer rows.Close()

	var items []CartItem
	for rows.Next() {
		var item CartItem
		var product Product
		var imageUrls []string
		var productCreatedAt, productUpdatedAt time.Time
		var category Category
		var toko Toko
		var user SingleUser

		err := rows.Scan(
			// Cart Item fields
			&item.ID,
			&item.CartID,
			&item.CartStoreID,
			&item.ProductID,
			&item.Quantity,
			&item.CreatedAt,
			&item.UpdatedAt,

			// Product fields
			&product.ID,
			&product.Name,
			&product.Slug,
			&product.Description,
			&product.Country,
			&product.Price,
			&product.DiscountPrice,
			&product.Discount,
			&product.Estimation,
			&product.Stock,
			&product.Sold,
			&product.IsForSale,
			&product.IsApproved,
			pq.Array(&imageUrls),
			&productCreatedAt,
			&productUpdatedAt,
			&product.Version,

			// Category fields
			&category.ID,
			&category.Name,
			&category.Slug,
			&category.Description,

			// Toko fields
			&toko.ID,
			&toko.UserID,
			&toko.Slug,
			&toko.Name,
			&toko.ImageProfile,
			&toko.Country,
			&toko.CreatedAt,

			// User fields
			&user.ID,
			&user.Picture,
			&user.Username,
			&user.Email,
			&user.IsActive,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan cart item: %v", err)
		}

		// Set additional fields
		product.ImageUrls = imageUrls
		product.CreatedAt = productCreatedAt
		product.UpdatedAt = productUpdatedAt
		product.Category = &category

		// Convert time to string for Toko and User if needed
		toko.User = &user

		product.Toko = &toko
		item.Product = &product

		items = append(items, item)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %v", err)
	}

	return items, nil
}

// Create membuat keranjang baru
func (s *CartStore) Create(ctx context.Context, tx *sql.Tx, userID int64) (*Cart, error) {
	query := `INSERT INTO carts (user_id, created_at, updated_at) 
              VALUES ($1, $2, $2) 
              RETURNING id, user_id, created_at, updated_at`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	var cart Cart
	err := tx.QueryRowContext(ctx, query, userID, time.Now()).Scan(
		&cart.ID,
		&cart.UserID,
		&cart.CreatedAt,
		&cart.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &cart, nil
}

// FindOrCreateCart mencari atau membuat cart baru jika tidak ada
func (s *CartStore) FindOrCreateCart(ctx context.Context, tx *sql.Tx, userID int64) (*Cart, error) {
	query := `SELECT id, user_id, created_at, updated_at 
              FROM carts 
              WHERE user_id = $1 
              ORDER BY created_at DESC 
              LIMIT 1`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	var cart Cart
	err := tx.QueryRowContext(ctx, query, userID).Scan(
		&cart.ID,
		&cart.UserID,
		&cart.CreatedAt,
		&cart.UpdatedAt,
	)

	switch {
	case err == sql.ErrNoRows:
		return s.Create(ctx, tx, userID)
	case err != nil:
		return nil, err
	default:
		return &cart, nil
	}
}

// findOrCreateCartStore mencari atau membuat cart store baru
func (s *CartStore) findOrCreateCartStore(ctx context.Context, tx *sql.Tx, cartID, tokoID int64) (*CartStores, error) {
	query := `SELECT id, cart_id, toko_id, created_at 
              FROM cart_stores 
              WHERE cart_id = $1 AND toko_id = $2 
              LIMIT 1`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	var store CartStores
	err := tx.QueryRowContext(ctx, query, cartID, tokoID).Scan(
		&store.ID,
		&store.CartID,
		&store.TokoID,
		&store.CreatedAt,
	)

	switch {
	case err == sql.ErrNoRows:
		// Buat cart store baru
		insertQuery := `INSERT INTO cart_stores (cart_id, toko_id, created_at)
                        VALUES ($1, $2, $3)
                        RETURNING id, cart_id, toko_id, created_at`
		err = tx.QueryRowContext(ctx, insertQuery, cartID, tokoID, time.Now()).Scan(
			&store.ID,
			&store.CartID,
			&store.TokoID,
			&store.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		return &store, nil
	case err != nil:
		return nil, err
	default:
		return &store, nil
	}
}

// AddItem menambahkan item ke keranjang dengan grouping by toko
func (s *CartStore) AddItem(ctx context.Context, tx *sql.Tx, cartID, productID, quantity int64) error {
	// Dapatkan info produk untuk mengetahui toko_id
	product, err := s.getProduct(ctx, tx, productID)
	if err != nil {
		return err
	}

	// Cari atau buat cart store berdasarkan toko produk
	cartStore, err := s.findOrCreateCartStore(ctx, tx, cartID, product.Toko.ID)
	if err != nil {
		return err
	}

	// Cek apakah item sudah ada di cart
	query := `SELECT id, quantity FROM cart_items 
              WHERE cart_id = $1 AND product_id = $2 AND cart_store_id = $3`

	var itemID uuid.UUID
	var currentQty int64
	err = tx.QueryRowContext(ctx, query, cartID, productID, cartStore.ID).Scan(&itemID, &currentQty)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// Item belum ada, tambahkan item baru
			insertQuery := `INSERT INTO cart_items (id, cart_id, cart_store_id, product_id, quantity, created_at, updated_at)
                          VALUES ($1, $2, $3, $4, $5, $6, $6)
                          RETURNING id`
			newID := uuid.New()
			return tx.QueryRowContext(ctx, insertQuery,
				newID,
				cartID,
				cartStore.ID,
				productID,
				quantity,
				time.Now(),
			).Scan(&itemID)
		}
		return err
	}

	// Item sudah ada, update quantity
	updateQuery := `UPDATE cart_items SET quantity = $1, updated_at = $2 
                    WHERE id = $3`
	_, err = tx.ExecContext(ctx, updateQuery, currentQty+quantity, time.Now(), itemID)
	return err
}

func (s *CartStore) getProduct(ctx context.Context, tx *sql.Tx, id int64) (*Product, error) {
	const query = `SELECT id, name, slug, country, description, price, discount_price, 
                          discount, estimation, stock, sold, is_for_sale, is_approved, 
                          created_at, updated_at, image_urls, category_id, toko_id, version 
                   FROM products WHERE id = $1`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	product := &Product{
		Category: &Category{},
		Toko:     &Toko{User: &SingleUser{}},
	}

	// Scan semua field dasar produk
	err := tx.QueryRowContext(ctx, query, id).Scan(
		&product.ID,
		&product.Name,
		&product.Slug,
		&product.Country,
		&product.Description,
		&product.Price,
		&product.DiscountPrice,
		&product.Discount,
		&product.Estimation,
		&product.Stock,
		&product.Sold,
		&product.IsForSale,
		&product.IsApproved,
		&product.CreatedAt,
		&product.UpdatedAt,
		pq.Array(&product.ImageUrls),
		&product.Category.ID, // Hanya ambil category_id
		&product.Toko.ID,     // Hanya ambil toko_id
		&product.Version,
	)

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

// AddToCartTransaction menambahkan item ke cart dengan transaksi
func (s *CartStore) AddToCartTransaction(ctx context.Context, userID, productID, quantity int64) (*Cart, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// Cari atau buat cart
	cart, err := s.FindOrCreateCart(ctx, tx, userID)
	if err != nil {
		return nil, err
	}

	// Tambahkan item ke cart
	err = s.AddItem(ctx, tx, cart.ID, productID, quantity)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return cart, nil
}

// AddingQuantityCartStoreItem menambahkan quantity item di cart store
func (s *CartStore) IncreaseQuantityCartStoreItemTransaction(ctx context.Context, cartstoreItemID uuid.UUID, userID int64) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Tambahkan quantity item ke cart store
	query := `UPDATE cart_items SET quantity = quantity + 1 WHERE id = $1 AND cart_id = (SELECT id FROM carts WHERE user_id = $2)`
	_, err = tx.ExecContext(ctx, query, cartstoreItemID, userID)
	if err != nil {
		return nil
	}

	// Commit transaksi
	return tx.Commit()
}

// RemoveQuantityCartStoreItem mengurangi quantity item di cart store
func (s *CartStore) DecreaseQuantityCartStoreItemTransaction(ctx context.Context, cartstoreItemID uuid.UUID, userID int64) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Kurangi quantity item di cart store
	query := `UPDATE cart_items SET quantity = quantity - 1 WHERE id = $1 AND cart_id = (SELECT id FROM carts WHERE user_id = $2)`
	res, err := tx.ExecContext(ctx, query, cartstoreItemID, userID)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		// Tidak ada item yang diupdate, kemungkinan item sudah tidak ada
		// Anggap sukses (idempotent), atau bisa return error khusus jika ingin
		return tx.Commit()
	}

	// Cek apakah quantity menjadi kurang dari atau sama dengan 0
	var quantity int
	checkQuery := `SELECT quantity FROM cart_items WHERE id = $1`
	err = tx.QueryRowContext(ctx, checkQuery, cartstoreItemID).Scan(&quantity)
	if err != nil {
		if err == sql.ErrNoRows {
			// Item sudah tidak ada, anggap sukses
			return tx.Commit()
		}
		return err
	}

	if quantity <= 0 {
		// Hapus item jika quantity <= 0
		deleteQuery := `DELETE FROM cart_items WHERE id = $1`
		_, err = tx.ExecContext(ctx, deleteQuery, cartstoreItemID)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

// UpdateItemQuantityByCartItemID mengupdate jumlah item di keranjang berdasarkan cartItemID
func (s *CartStore) UpdateItemQuantityByCartItemID(ctx context.Context, cartItemID uuid.UUID, quantity int64) error {
	if quantity <= 0 {
		// Jika quantity <= 0, hapus item
		tx, err := s.db.BeginTx(ctx, nil)
		if err != nil {
			return err
		}
		defer tx.Rollback()

		_, err = tx.ExecContext(ctx, `DELETE FROM cart_items WHERE id = $1`, cartItemID)
		if err != nil {
			return err
		}

		return tx.Commit()
	}

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	query := `UPDATE cart_items SET quantity = $1, updated_at = $2 WHERE id = $3`
	_, err = tx.ExecContext(ctx, query, quantity, time.Now(), cartItemID)
	if err != nil {
		return err
	}

	return tx.Commit()
}

// RemoveItemByCartItemID menghapus item dari keranjang berdasarkan cartItemID
func (s *CartStore) RemoveItemByCartItemID(ctx context.Context, cartItemID uuid.UUID) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Dapatkan cart_id dan cart_store_id dari cart_items
	var cartID int64
	var cartStoreID uuid.UUID
	query := `SELECT cart_id, cart_store_id FROM cart_items WHERE id = $1`
	err = tx.QueryRowContext(ctx, query, cartItemID).Scan(&cartID, &cartStoreID)
	if err != nil {
		if err == sql.ErrNoRows {
			// Item sudah tidak ada, anggap sukses
			return nil
		}
		return err
	}

	// Hapus item
	_, err = tx.ExecContext(ctx,
		`DELETE FROM cart_items WHERE id = $1`,
		cartItemID,
	)
	if err != nil {
		return err
	}

	// Hapus cart_store jika sudah tidak ada item lagi
	_, err = tx.ExecContext(ctx,
		`DELETE FROM cart_stores 
		 WHERE id = $1 
		 AND NOT EXISTS (
			 SELECT 1 FROM cart_items WHERE cart_store_id = $1
		 )`,
		cartStoreID,
	)
	if err != nil {
		return err
	}

	return tx.Commit()
}

// ClearCartByCartStoreID menghapus semua item dari cart_store tertentu dan cart_store-nya
func (s *CartStore) ClearCartByCartStoreID(ctx context.Context, cartStoreID uuid.UUID, userID int64) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Hapus semua item di cart_store, pastikan cart_store milik user
	_, err = tx.ExecContext(ctx,
		`DELETE FROM cart_items 
		 WHERE cart_store_id = $1 
		 AND cart_id = (SELECT id FROM carts WHERE user_id = $2)`,
		cartStoreID,
		userID,
	)
	if err != nil {
		return err
	}

	// Hapus cart_store jika memang milik user
	_, err = tx.ExecContext(ctx,
		`DELETE FROM cart_stores 
		 WHERE id = $1 
		 AND cart_id = (SELECT id FROM carts WHERE user_id = $2)`,
		cartStoreID,
		userID,
	)
	if err != nil {
		return err
	}

	return tx.Commit()
}

// RemoveCartItemByID menghapus satu item dari cart berdasarkan cartItemID
func (s *CartStore) RemoveCartItemByID(ctx context.Context, cartItemID uuid.UUID) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Dapatkan cart_store_id dari cart_items
	var cartStoreID uuid.UUID
	query := `SELECT cart_store_id FROM cart_items WHERE id = $1`
	err = tx.QueryRowContext(ctx, query, cartItemID).Scan(&cartStoreID)
	if err != nil {
		if err == sql.ErrNoRows {
			// Item sudah tidak ada, anggap sukses
			return nil
		}
		return err
	}

	// Hapus item
	_, err = tx.ExecContext(ctx,
		`DELETE FROM cart_items WHERE id = $1`,
		cartItemID,
	)
	if err != nil {
		return err
	}

	// Hapus cart_store jika sudah tidak ada item lagi
	_, err = tx.ExecContext(ctx,
		`DELETE FROM cart_stores 
		 WHERE id = $1 
		 AND NOT EXISTS (
			 SELECT 1 FROM cart_items WHERE cart_store_id = $1
		 )`,
		cartStoreID,
	)
	if err != nil {
		return err
	}

	return tx.Commit()
}
