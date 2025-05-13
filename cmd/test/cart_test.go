package test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yogaprasetya22/api-gotokopedia/internal/store"
)

func TestCartHandler(t *testing.T) {
	ctx := context.Background()
	storeTest, db,_ := NewTestStorage(t)
	require.NotNil(t, storeTest)

	// Ambil user yang sudah ada
	users, err := db.Query("SELECT id, username, email FROM users LIMIT 3")
	require.NoError(t, err)
	defer users.Close()

	var userIDs []int64
	for users.Next() {
		var id int64
		var username, email string
		require.NoError(t, users.Scan(&id, &username, &email))
		userIDs = append(userIDs, id)
	}
	require.NotEmpty(t, userIDs)

	// Ambil produk yang sudah ada
	products, err := db.Query("SELECT id FROM products LIMIT 3")
	require.NoError(t, err)
	defer products.Close()

	var productIDs []int64
	for products.Next() {
		var id int64
		require.NoError(t, products.Scan(&id))
		productIDs = append(productIDs, id)
	}
	require.NotEmpty(t, productIDs)

	for _, userID := range userIDs {
		t.Run("AddToCartTransaction", func(t *testing.T) {
			cart, err := storeTest.Carts.AddToCartTransaction(ctx, userID, productIDs[0], 2)
			require.NoError(t, err)
			require.NotNil(t, cart)
		})

		t.Run("GetCartByUserIDPQ", func(t *testing.T) {
			cart, err := storeTest.Carts.GetCartByUserIDPQ(ctx, userID, store.PaginatedFeedQuery{
				Limit:  10,
				Offset: 0,
				Sort:   "desc",
			})
			require.NoError(t, err)
			require.NotNil(t, cart)
			if len(cart.Cart.Stores) > 0 && len(cart.Cart.Stores[0].Items) > 0 {
				cartStoreItemID := cart.Cart.Stores[0].Items[0].ID
				cartItemID := cart.Cart.Stores[0].Items[0].ID

				t.Run("IncreaseQuantityCartStoreItemTransaction", func(t *testing.T) {
					err := storeTest.Carts.IncreaseQuantityCartStoreItemTransaction(ctx, cartStoreItemID, userID)
					require.NoError(t, err)
				})

				t.Run("DecreaseQuantityCartStoreItemTransaction", func(t *testing.T) {
					err := storeTest.Carts.DecreaseQuantityCartStoreItemTransaction(ctx, cartStoreItemID, userID)
					require.NoError(t, err)
				})

				t.Run("UpdateItemQuantity", func(t *testing.T) {
					quantity := 5 // contoh quantity baru
					err := storeTest.Carts.UpdateItemQuantityByCartItemID(ctx, cartItemID, int64(quantity))
					require.NoError(t, err)
				})

				t.Run("RemoveItem", func(t *testing.T) {
					err := storeTest.Carts.RemoveItemByCartItemID(ctx, cartItemID)
					require.NoError(t, err)
				})
			}
		})

		t.Run("ClearCart", func(t *testing.T) {
			for _, userID := range userIDs {
				t.Run("GetCartByUserIDPQ", func(t *testing.T) {
					cart, err := storeTest.Carts.GetCartByUserIDPQ(ctx, userID, store.PaginatedFeedQuery{
						Limit:  10,
						Offset: 0,
						Sort:   "desc",
					})
					require.NoError(t, err)
					require.NotNil(t, cart)
					if len(cart.Cart.Stores) > 0 && len(cart.Cart.Stores[0].Items) > 0 {
						cartStoreItemID := cart.Cart.Stores[0].Items[0].ID
						// Pastikan cart ada
						cart, err := storeTest.Carts.GetCartByUserIDPQ(ctx, userID, store.PaginatedFeedQuery{
							Limit:  10,
							Offset: 0,
							Sort:   "desc",
						})
						require.NoError(t, err)
						require.NotNil(t, cart)

						err = storeTest.Carts.ClearCartByCartStoreID(ctx, cartStoreItemID, userID)
						require.NoError(t, err)
					}
				})
			}
		})
	}
}
