package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/yogaprasetya22/api-gotokopedia/internal/store"
)

func TestOrderHandler(t *testing.T) {
	ctx := context.Background()
	storeTest, db := NewTestStorage(t)
	require.NotNil(t, storeTest)

	// Ambil user yang sudah ada
	users, err := db.Query("SELECT id FROM users LIMIT 2")
	require.NoError(t, err)
	defer users.Close()

	var userIDs []int64
	for users.Next() {
		var id int64
		require.NoError(t, users.Scan(&id))
		userIDs = append(userIDs, id)
	}
	require.NotEmpty(t, userIDs)

	// Ambil payment method
	paymentMethods, err := storeTest.Orders.GetPaymentMethods(ctx)
	require.NoError(t, err)
	require.NotEmpty(t, paymentMethods)
	paymentMethodID := paymentMethods[0].ID

	// Ambil shipping method
	shippingMethods, err := storeTest.Orders.GetShippingMethods(ctx)
	require.NoError(t, err)
	require.NotEmpty(t, shippingMethods)
	shippingMethodID := shippingMethods[0].ID

	for _, userID := range userIDs {
		// Buat Shipping Address
		t.Run("CreateShippingAddress", func(t *testing.T) {
			addr := &store.ShippingAddresses{
				UserID:         userID,
				Label:          "Rumah",
				RecipientName:  fmt.Sprintf("User %d", userID),
				RecipientPhone: fmt.Sprintf("0812345678%d", userID),
				AddressLine1:   fmt.Sprintf("Jl. Alamat Utama No.%d", userID),
			}

			err := storeTest.ShippingAddresses.Create(ctx, addr)
			require.NoError(t, err)
			require.NotEqual(t, uuid.Nil, addr.ID)
		})

		// Ambil shipping address
		//TODO (sql: no rows in result set)  dikarenakan belum ada shipping address pada user tersebut (belum di buat addresses)
		address, err := storeTest.ShippingAddresses.GetDefaultAddress(ctx, userID)
		require.NoError(t, err)
		require.NotNil(t, address)

		// Ambil cart user
		cart, err := storeTest.Carts.GetCartByUserID(ctx, userID, store.PaginatedFeedQuery{
			Limit:  10,
			Offset: 0,
			Sort:   "desc",
		})
		require.NoError(t, err)
		require.NotNil(t, cart)
		if len(cart.Stores) == 0 {
			continue // skip user tanpa cart
		}

		// Buat order dari cart store pertama
		cartStoreID := cart.Stores[0].ID
		t.Run("CreateFromCart", func(t *testing.T) {
			err := storeTest.Orders.CreateFromCart(ctx, cartStoreID, userID, paymentMethodID, shippingMethodID, address.ID, "unit test order")
			require.NoError(t, err)
		})

		// Ambil order by user
		t.Run("GetByUserID", func(t *testing.T) {
			orders, err := storeTest.Orders.GetByUserID(ctx, userID, store.PaginatedFeedQuery{
				Limit:  10,
				Offset: 0,
				Sort:   "desc",
			})
			require.NoError(t, err)
			require.NotEmpty(t, orders)

			// Ambil order by id
			orderID := orders[0].ID
			t.Run("GetByID", func(t *testing.T) {
				order, err := storeTest.Orders.GetByID(ctx, orderID)
				require.NoError(t, err)
				require.NotNil(t, order)
			})

			// Update status order
			t.Run("UpdateStatus", func(t *testing.T) {
				statusID := orders[0].StatusID
				err := storeTest.Orders.UpdateStatus(ctx, orderID, statusID, "update status test")
				require.NoError(t, err)
			})
		})
	}

	// Test GetShippingMethods & GetPaymentMethods
	t.Run("GetShippingMethods", func(t *testing.T) {
		methods, err := storeTest.Orders.GetShippingMethods(ctx)
		require.NoError(t, err)
		require.NotEmpty(t, methods)
	})

	t.Run("GetPaymentMethods", func(t *testing.T) {
		methods, err := storeTest.Orders.GetPaymentMethods(ctx)
		require.NoError(t, err)
		require.NotEmpty(t, methods)
	})
}
