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

// Order merepresentasikan pesanan
type Order struct {
	ID                  int64     `json:"id"`
	UserID              int64     `json:"user_id"`
	ShippingAddressesID uuid.UUID `json:"shipping_addresses_id"`
	OrderNumber         string    `json:"order_number"`
	StatusID            int64     `json:"status_id"`
	PaymentMethodID     int64     `json:"payment_method_id"`
	ShippingMethodID    int64     `json:"shipping_method_id"`
	ShippingCost        float64   `json:"shipping_cost"`
	TotalPrice          float64   `json:"total_price"`
	FinalPrice          float64   `json:"final_price"`
	Notes               string    `json:"notes,omitempty"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`

	// Relasi
	Status            *OrderStatus       `json:"status,omitempty"`
	ShippingMethod    *ShippingMethod    `json:"shipping_method,omitempty"`
	PaymentMethod     *PaymentMethod     `json:"payment_method,omitempty"`
	Items             []*OrderItem       `json:"items,omitempty"`
	Tracking          []*OrderTracking   `json:"tracking,omitempty"`
	ShippingAddresses *ShippingAddresses `json:"shipping_addresses,omitempty"`
}

// OrderItem merepresentasikan item dalam pesanan
type OrderItem struct {
	ID            int64     `json:"id"`
	OrderID       int64     `json:"order_id"`
	ProductID     int64     `json:"product_id"`
	TokoID        int64     `json:"toko_id"`
	Quantity      int       `json:"quantity"`
	Price         float64   `json:"price"`
	DiscountPrice float64   `json:"discount_price"`
	Discount      float64   `json:"discount"`
	Subtotal      float64   `json:"subtotal"`
	CreatedAt     time.Time `json:"created_at"`

	// Relasi
	Product *Product `json:"product,omitempty"`
	Toko    *Toko    `json:"toko,omitempty"`
}

// OrderTracking merepresentasikan tracking status pesanan
type OrderTracking struct {
	ID        int64     `json:"id"`
	OrderID   int64     `json:"order_id"`
	StatusID  int64     `json:"status_id"`
	Notes     string    `json:"notes,omitempty"`
	CreatedAt time.Time `json:"created_at"`

	// Relasi
	Status *OrderStatus `json:"status,omitempty"`
}

// OrderStatus merepresentasikan status pesanan
type OrderStatus struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

// ShippingMethod merepresentasikan metode pengiriman
type ShippingMethod struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description,omitempty"`
	Price       float64 `json:"price"`
	IsActive    bool    `json:"is_active"`
}

// PaymentMethod merepresentasikan metode pembayaran
type PaymentMethod struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	IsActive    bool   `json:"is_active"`
}

// OrderDetail merepresentasikan view order_details
type OrderDetail struct {
	ID              int64     `json:"id"`
	OrderNumber     string    `json:"order_number"`
	CustomerName    string    `json:"customer_name"`
	CustomerEmail   string    `json:"customer_email"`
	Status          string    `json:"status"`
	PaymentMethod   string    `json:"payment_method"`
	ShippingMethod  string    `json:"shipping_method"`
	ShippingCost    float64   `json:"shipping_cost"`
	TotalPrice      float64   `json:"total_price"`
	Discount        float64   `json:"discount"`
	FinalPrice      float64   `json:"final_price"`
	Notes           string    `json:"notes,omitempty"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

// OrderStore merupakan implementasi store untuk order
type OrderStore struct {
	db *sql.DB
}

// CreateFromCart membuat pesanan dari keranjang belanja
func (s *OrderStore) CreateFromCart(ctx context.Context, cartStoreID uuid.UUID, userID, paymentMethodID, shippingMethodID int64, shippingAddressesID uuid.UUID, notes string) error {
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	_, err := s.db.ExecContext(ctx,
		`SELECT create_order_from_cart($1, $2, $3, $4, $5, $6)`,
		userID, cartStoreID, paymentMethodID, shippingMethodID, shippingAddressesID, notes)

	return err
}

// UpdateStatus memperbarui status pesanan
func (s *OrderStore) UpdateStatus(ctx context.Context, orderID, statusID int64, notes string) error {
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	_, err := s.db.ExecContext(ctx,
		`SELECT update_order_status($1, $2, $3)`,
		orderID, statusID, notes)

	return err
}

// GetByUserID mendapatkan semua pesanan berdasarkan user ID dengan penanganan error yang lebih baik
func (s *OrderStore) GetByUserID(ctx context.Context, userID int64, fq PaginatedFeedQuery) ([]*Order, error) {
	// Validasi input
	if userID <= 0 {
		return nil, errors.New("userID harus lebih besar dari 0")
	}

	// Set default values untuk pagination
	if fq.Limit <= 0 {
		fq.Limit = 10
	}
	if fq.Offset < 0 {
		fq.Offset = 0
	}

	// Mulai transaksi dengan timeout
	txCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	tx, err := s.db.BeginTx(txCtx, &sql.TxOptions{
		Isolation: sql.LevelReadCommitted,
		ReadOnly:  true,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}

	// Gunakan defer dengan pengecekan error
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p) // re-throw panic setelah rollback
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	// Pisahkan query utama menjadi lebih sederhana
	orders, err := s.getUserOrders(txCtx, tx, userID, fq)
	if err != nil {
		return nil, fmt.Errorf("failed to get user orders: %w", err)
	}

	// Isi relasi untuk setiap order
	for _, order := range orders {
		if err := s.getOrderRelations(txCtx, tx, order); err != nil {
			return nil, fmt.Errorf("failed to get order relations: %w", err)
		}
	}

	return orders, nil
}

// getUserOrders mendapatkan daftar order tanpa relasi
func (s *OrderStore) getUserOrders(ctx context.Context, tx *sql.Tx, userID int64, fq PaginatedFeedQuery) ([]*Order, error) {
	query := `
		SELECT o.id, o.user_id, o.order_number, o.status_id, o.payment_method_id, 
			o.shipping_method_id, o.shipping_cost, o.total_price,
			o.final_price, o.notes, o.created_at, o.updated_at,
			o.shipping_addresses_id,
			os.id, os.name, os.description,
			sm.id, sm.name, sm.description, sm.price, sm.is_active,
			pm.id, pm.name, pm.description, pm.is_active
		FROM orders o
		JOIN order_status os ON o.status_id = os.id
		JOIN shipping_methods sm ON o.shipping_method_id = sm.id
		JOIN payment_methods pm ON o.payment_method_id = pm.id
		WHERE o.user_id = $1
		ORDER BY o.created_at DESC
		LIMIT $2 OFFSET $3`

	rows, err := tx.QueryContext(ctx, query, userID, fq.Limit, fq.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []*Order
	for rows.Next() {
		var order Order
		var shippingAddressesID sql.NullString

		// Initialize pointer fields before scanning
		order.Status = &OrderStatus{}
		order.ShippingMethod = &ShippingMethod{}
		order.PaymentMethod = &PaymentMethod{}

		err := rows.Scan(
			&order.ID, &order.UserID, &order.OrderNumber, &order.StatusID, &order.PaymentMethodID,
			&order.ShippingMethodID, &order.ShippingCost, &order.TotalPrice,
			&order.FinalPrice, &order.Notes, &order.CreatedAt, &order.UpdatedAt,
			&shippingAddressesID,
			&order.Status.ID, &order.Status.Name, &order.Status.Description,
			&order.ShippingMethod.ID, &order.ShippingMethod.Name, &order.ShippingMethod.Description,
			&order.ShippingMethod.Price, &order.ShippingMethod.IsActive,
			&order.PaymentMethod.ID, &order.PaymentMethod.Name, &order.PaymentMethod.Description,
			&order.PaymentMethod.IsActive,
		)
		if err != nil {
			return nil, err
		}
		if shippingAddressesID.Valid {
			order.ShippingAddressesID, err = uuid.Parse(shippingAddressesID.String)
			if err != nil {
				return nil, fmt.Errorf("failed to parse shipping addresses ID: %w", err)
			}
		} else {
			order.ShippingAddressesID = uuid.Nil
		}
		orders = append(orders, &order)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return orders, nil
}

// getOrderRelations mengisi relasi untuk sebuah order
func (s *OrderStore) getOrderRelations(ctx context.Context, tx *sql.Tx, order *Order) error {
	// Dapatkan status order
	if err := s.getOrderStatus(ctx, tx, order); err != nil {
		return err
	}

	// Dapatkan shipping method
	if err := s.getShippingMethod(ctx, tx, order); err != nil {
		return err
	}

	// Dapatkan payment method
	if err := s.getPaymentMethod(ctx, tx, order); err != nil {
		return err
	}

	// Dapatkan order items
	if err := s.getOrderItemsTx(ctx, tx, order); err != nil {
		return err
	}

	// Dapatkan latest tracking
	return s.getLatestOrderTrackingTx(ctx, tx, order)
}

// getOrderStatus mengisi data status order
func (s *OrderStore) getOrderStatus(ctx context.Context, tx *sql.Tx, order *Order) error {
	order.Status = &OrderStatus{}
	query := `SELECT id, name, description FROM order_status WHERE id = $1`
	return tx.QueryRowContext(ctx, query, order.StatusID).Scan(
		&order.Status.ID, &order.Status.Name, &order.Status.Description,
	)
}

// getShippingMethod mengisi data shipping method
func (s *OrderStore) getShippingMethod(ctx context.Context, tx *sql.Tx, order *Order) error {
	order.ShippingMethod = &ShippingMethod{}
	query := `SELECT id, name, description, price, is_active FROM shipping_methods WHERE id = $1`
	return tx.QueryRowContext(ctx, query, order.ShippingMethodID).Scan(
		&order.ShippingMethod.ID, &order.ShippingMethod.Name,
		&order.ShippingMethod.Description, &order.ShippingMethod.Price,
		&order.ShippingMethod.IsActive,
	)
}

// getPaymentMethod mengisi data payment method
func (s *OrderStore) getPaymentMethod(ctx context.Context, tx *sql.Tx, order *Order) error {
	order.PaymentMethod = &PaymentMethod{}
	query := `SELECT id, name, description, is_active FROM payment_methods WHERE id = $1`
	return tx.QueryRowContext(ctx, query, order.PaymentMethodID).Scan(
		&order.PaymentMethod.ID, &order.PaymentMethod.Name,
		&order.PaymentMethod.Description, &order.PaymentMethod.IsActive,
	)
}

// getLatestOrderTrackingTx mendapatkan status tracking terbaru untuk order
func (s *OrderStore) getLatestOrderTrackingTx(ctx context.Context, tx *sql.Tx, order *Order) error {
	query := `
		SELECT ot.id, ot.status_id, ot.notes, ot.created_at,
			   os.id, os.name, os.description
		FROM order_tracking ot
		JOIN order_status os ON ot.status_id = os.id
		WHERE ot.order_id = $1
		ORDER BY ot.created_at DESC
		LIMIT 1`

	tracking := &OrderTracking{
		OrderID: order.ID,
		Status:  &OrderStatus{},
	}

	err := tx.QueryRowContext(ctx, query, order.ID).Scan(
		&tracking.ID, &tracking.StatusID, &tracking.Notes, &tracking.CreatedAt,
		&tracking.Status.ID, &tracking.Status.Name, &tracking.Status.Description,
	)

	if err != nil && err != sql.ErrNoRows {
		return err
	}

	if err == nil {
		order.Tracking = []*OrderTracking{tracking}
	}

	return nil
}

// GetByID mendapatkan pesanan berdasarkan ID dengan transaksi untuk memenuhi semua relasi
func (s *OrderStore) GetByID(ctx context.Context, id int64) (*Order, error) {
	// Mulai transaksi
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	// Query utama untuk mendapatkan data order
	order := &Order{
		Status:            &OrderStatus{},
		ShippingMethod:    &ShippingMethod{},
		PaymentMethod:     &PaymentMethod{},
		ShippingAddresses: &ShippingAddresses{},
	}

	query := `
		SELECT o.id, o.user_id, o.order_number, o.status_id, o.payment_method_id, 
			o.shipping_method_id, o.shipping_cost, o.total_price,
			o.final_price, o.notes, o.created_at, o.updated_at,
			o.shipping_addresses_id,
			os.id, os.name, os.description,
			sm.id, sm.name, sm.description, sm.price, sm.is_active,
			pm.id, pm.name, pm.description, pm.is_active
		FROM orders o
		JOIN order_status os ON o.status_id = os.id
		JOIN shipping_methods sm ON o.shipping_method_id = sm.id
		JOIN payment_methods pm ON o.payment_method_id = pm.id
		WHERE o.id = $1`

	var shippingAddressesID sql.NullString
	err = tx.QueryRowContext(ctx, query, id).Scan(
		&order.ID, &order.UserID, &order.OrderNumber, &order.StatusID, &order.PaymentMethodID,
		&order.ShippingMethodID, &order.ShippingCost, &order.TotalPrice,
		&order.FinalPrice, &order.Notes, &order.CreatedAt, &order.UpdatedAt,
		&shippingAddressesID,
		&order.Status.ID, &order.Status.Name, &order.Status.Description,
		&order.ShippingMethod.ID, &order.ShippingMethod.Name, &order.ShippingMethod.Description,
		&order.ShippingMethod.Price, &order.ShippingMethod.IsActive,
		&order.PaymentMethod.ID, &order.PaymentMethod.Name, &order.PaymentMethod.Description,
		&order.PaymentMethod.IsActive,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}
		return nil, err
	}

	// Get shipping addresses
	if shippingAddressesID.Valid {
		order.ShippingAddresses, err = s.getShippingAddressesByID(ctx, tx, uuid.MustParse(shippingAddressesID.String))
		if err != nil {
			return nil, err
		}
	} else {
		order.ShippingAddresses = nil
	}

	// Get order items
	if err = s.getOrderItemsTx(ctx, tx, order); err != nil {
		return nil, err
	}

	// Get order tracking history
	if err = s.getOrderTrackingTx(ctx, tx, order); err != nil {
		return nil, err
	}

	// Commit transaksi jika semua berhasil
	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return order, nil
}

func (s *OrderStore) getShippingAddressesByID(ctx context.Context, tx *sql.Tx, id uuid.UUID) (*ShippingAddresses, error) {
	query := `SELECT id, user_id, label, recipient_name, recipient_phone, address_line1, note_for_courier, created_at, updated_at 
	FROM shipping_addresses WHERE id = $1`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	sa := &ShippingAddresses{}
	var noteForCourier sql.NullString
	err := s.db.QueryRowContext(ctx, query, id).Scan(&sa.ID, &sa.UserID, &sa.Label,
		&sa.RecipientName, &sa.RecipientPhone, &sa.AddressLine1,
		&noteForCourier, &sa.CreatedAt, &sa.UpdatedAt)
	if err != nil {
		return nil, err
	}
	if noteForCourier.Valid {
		sa.NoteForCourier = noteForCourier.String
	} else {
		sa.NoteForCourier = ""
	}

	return sa, nil
}

// getOrderItemsTx mendapatkan item-item pesanan dalam transaksi
func (s *OrderStore) getOrderItemsTx(ctx context.Context, tx *sql.Tx, order *Order) error {
	query := `
		SELECT oi.id, oi.product_id, oi.toko_id, oi.quantity, oi.price, 
			oi.discount_price, oi.discount, oi.subtotal, oi.created_at,
			p.id, p.name, p.slug, p.description, p.price as product_price,
			p.discount_price as product_discount_price, p.discount as product_discount,
			p.image_urls, p.stock, p.sold, p.is_for_sale, p.is_approved, p.created_at as product_created_at,
			p.updated_at as product_updated_at, p.version,
			t.id, t.user_id, t.name, t.slug, t.image_profile, t.country, t.created_at as toko_created_at
		FROM order_items oi
		JOIN products p ON oi.product_id = p.id
		JOIN tokos t ON oi.toko_id = t.id
		WHERE oi.order_id = $1`

	rows, err := tx.QueryContext(ctx, query, order.ID)
	if err != nil {
		return err
	}
	defer rows.Close()

	var items []*OrderItem
	for rows.Next() {
		item := &OrderItem{
			OrderID: order.ID,
			Product: &Product{},
			Toko:    &Toko{},
		}

		var (
			imageUrls        []string
			productCreatedAt time.Time
			productUpdatedAt time.Time
			tokoCreatedAt    time.Time
		)

		err := rows.Scan(
			&item.ID, &item.ProductID, &item.TokoID, &item.Quantity, &item.Price,
			&item.DiscountPrice, &item.Discount, &item.Subtotal, &item.CreatedAt,
			&item.Product.ID, &item.Product.Name, &item.Product.Slug, &item.Product.Description, &item.Product.Price,
			&item.Product.DiscountPrice, &item.Product.Discount, pq.Array(&imageUrls),
			&item.Product.Stock, &item.Product.Sold, &item.Product.IsForSale, &item.Product.IsApproved,
			&productCreatedAt, &productUpdatedAt, &item.Product.Version,
			&item.Toko.ID, &item.Toko.UserID, &item.Toko.Name, &item.Toko.Slug,
			&item.Toko.ImageProfile, &item.Toko.Country, &tokoCreatedAt,
		)

		if err != nil {
			return err
		}

		item.Product.ImageUrls = imageUrls
		item.Product.CreatedAt = productCreatedAt
		item.Product.UpdatedAt = productUpdatedAt
		item.Toko.CreatedAt = tokoCreatedAt
		items = append(items, item)
	}

	order.Items = items
	return nil
}

// getOrderTrackingTx mendapatkan riwayat status pesanan dalam transaksi
func (s *OrderStore) getOrderTrackingTx(ctx context.Context, tx *sql.Tx, order *Order) error {
	query := `
		SELECT ot.id, ot.status_id, ot.notes, ot.created_at,
			os.id, os.name, os.description
		FROM order_tracking ot
		JOIN order_status os ON ot.status_id = os.id
		WHERE ot.order_id = $1
		ORDER BY ot.created_at DESC`

	rows, err := tx.QueryContext(ctx, query, order.ID)
	if err != nil {
		return err
	}
	defer rows.Close()

	var trackings []*OrderTracking
	for rows.Next() {
		tracking := &OrderTracking{
			OrderID: order.ID,
			Status:  &OrderStatus{},
		}

		err := rows.Scan(
			&tracking.ID, &tracking.StatusID, &tracking.Notes, &tracking.CreatedAt,
			&tracking.Status.ID, &tracking.Status.Name, &tracking.Status.Description,
		)

		if err != nil {
			return err
		}

		trackings = append(trackings, tracking)
	}

	order.Tracking = trackings
	return nil
}


// GetShippingMethods mendapatkan semua metode pengiriman
func (s *OrderStore) GetShippingMethods(ctx context.Context) ([]*ShippingMethod, error) {
	query := `SELECT id, name, description, price, is_active FROM shipping_methods WHERE is_active = true`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var methods []*ShippingMethod
	for rows.Next() {
		method := &ShippingMethod{}
		err := rows.Scan(&method.ID, &method.Name, &method.Description, &method.Price, &method.IsActive)
		if err != nil {
			return nil, err
		}
		methods = append(methods, method)
	}

	return methods, nil
}

// GetPaymentMethods mendapatkan semua metode pembayaran
func (s *OrderStore) GetPaymentMethods(ctx context.Context) ([]*PaymentMethod, error) {
	query := `SELECT id, name, description, is_active FROM payment_methods WHERE is_active = true`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var methods []*PaymentMethod
	for rows.Next() {
		method := &PaymentMethod{}
		err := rows.Scan(&method.ID, &method.Name, &method.Description, &method.IsActive)
		if err != nil {
			return nil, err
		}
		methods = append(methods, method)
	}

	return methods, nil
}
