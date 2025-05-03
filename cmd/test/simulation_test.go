package test

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yogaprasetya22/api-gotokopedia/internal/store"
)

var usernames = []string{
	"yogaprasetyaa1",
	"yogaprasetyaa2",
	"yogaprasetyaa3",
	"yogaprasetyaa4",
	"yogaprasetyaa5",
	"yogaprasetyaa6",
	"yogaprasetyaa7",
	"yogaprasetyaa8",
	"yogaprasetyaa9",
}

func TestUserHandler(t *testing.T) {
	AMOUNT := 20
	ctx := context.Background()

	storeTest, db := NewTestStorage(t)
	if storeTest == nil {
		t.Fatal("store is nil")
	}
	
	var result_users []*store.User
	t.Run("CreateUser", func(t *testing.T) {
		users := generateUsers(AMOUNT)

		for _, user := range users {
			var createdUser *store.User
			var err error

			WithTransaction(t, db, func(tx *sql.Tx) error {
				createdUser, err = storeTest.Users.CreateWithActive(ctx, tx, user)
				return err
			})

			if err != nil {
				// Jika gagal karena email sudah ada, ambil user di luar transaksi
				if err.Error() == "a user with that email already exists" {
					existingUser, getErr := storeTest.Users.GetByEmail(ctx, user.Email)
					require.NoError(t, getErr)
					require.NotNil(t, existingUser)

					// Simpan user yang sudah ada ke dalam result_users
					result_users = append(result_users, existingUser)
					continue
				}
				require.NoError(t, err)
			} else {
				require.NotNil(t, createdUser)
				result_users = append(result_users, createdUser)
			}
		}

		require.Equal(t, len(users), len(result_users))
		for i, user := range users {
			require.Equal(t, user.Username, result_users[i].Username)
		}
	})

	var result_cart []*store.Cart
	t.Run("CreateCartByUser", func(t *testing.T) {
		for _, user := range result_users {
			var cart *store.Cart
			var err error
			cart, err = storeTest.Carts.AddToCartTransaction(ctx, user.ID, RandomInt(1, 400), 1)
			require.NoError(t, err)
			require.NotNil(t, cart)
			result_cart = append(result_cart, cart)
			fmt.Printf("User %s (%s) created with ID %d\n", user.Username, user.Email, user.ID)
		}
	})

	var result_cart_store []*store.CartStores
	t.Run("GetCartByUserID", func(t *testing.T) {
		for _, cart := range result_cart {
			cartStore, err := storeTest.Carts.GetCartByUserID(ctx, cart.UserID, store.PaginatedFeedQuery{
				Limit:  AMOUNT,
				Offset: 0,
				Sort:   "desc",
				Search: "",
			})
			require.NoError(t, err)
			require.NotNil(t, cartStore)
			for i := range cartStore.Stores {
				result_cart_store = append(result_cart_store, &cartStore.Stores[i])
			}
			fmt.Printf("Cart %d (%d) created with ID %d\n", cart.ID, cart.UserID, cart.ID)
		}
	})

	t.Run("AddingQuantityCartStoreItemTransaction", func(t *testing.T) {
		for _, cart_store := range result_cart_store {
			// Cari cart yang sesuai berdasarkan CartID
			var userID int64
			for _, cart := range result_cart {
				if cart.ID == cart_store.CartID {
					userID = cart.UserID
					break
				}
			}
			if userID == 0 {
				t.Fatalf("Tidak ditemukan userID untuk cart_store dengan CartID %d", cart_store.CartID)
			}
			// First addition
			err := storeTest.Carts.AddingQuantityCartStoreItemTransaction(ctx, cart_store.ID, userID)
			require.NoError(t, err)
			fmt.Printf("CartStore %d (UserID: %d) (first add)\n", cart_store.ID, userID)

			// Second addition
			err = storeTest.Carts.AddingQuantityCartStoreItemTransaction(ctx, cart_store.ID, userID)
			require.NoError(t, err)
			fmt.Printf("CartStore %d (UserID: %d) (second add)\n", cart_store.ID, userID)
		}
	})

	t.Run("RemovingQuantityCartStoreItemTransaction", func(t *testing.T) {
		for _, cart_store := range result_cart_store {
			// Cari cart yang sesuai berdasarkan CartID
			var userID int64
			for _, cart := range result_cart {
				if cart.ID == cart_store.CartID {
					userID = cart.UserID
					break
				}
			}
			if userID == 0 {
				t.Fatalf("Tidak ditemukan userID untuk cart_store dengan CartID %d", cart_store.CartID)
			}
			err := storeTest.Carts.RemovingQuantityCartStoreItemTransaction(ctx, cart_store.ID, userID)
			require.NoError(t, err)
			fmt.Printf("CartStore %d (UserID: %d) (remove)\n", cart_store.ID, userID)
		}
	})

	t.Run("CreateOrderFromCart", func(t *testing.T) {
		for _, cart_store := range result_cart_store {
			// Cari cart yang sesuai berdasarkan CartID
			var userID int64
			for _, cart := range result_cart {
				if cart.ID == cart_store.CartID {
					userID = cart.UserID
					break
				}
			}
			if userID == 0 {
				t.Fatalf("Tidak ditemukan userID untuk cart_store dengan CartID %d", cart_store.CartID)
			}
			// Buat order dari cart
			err := storeTest.Orders.CreateFromCart(ctx, cart_store.ID, userID, 1, 1, "address", "notes")
			require.NoError(t, err)
			fmt.Printf("Order created from CartStore %d (UserID: %d)\n", cart_store.ID, userID)
		}
	})

}

func generateUsers(num int) []*store.User {
	users := make([]*store.User, num)

	for i := 0; i < num; i++ {
		users[i] = &store.User{
			Username: usernames[i%len(usernames)] + fmt.Sprintf("%d", i),
			Email:    usernames[i%len(usernames)] + fmt.Sprintf("%d", i) + "@example.com",
			Role: store.Role{
				Name: "user",
			},
		}
		if err := users[i].Password.Set("password"); err != nil {
			log.Println("Error setting password:", err)
			return nil
		}
	}

	return users
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}
