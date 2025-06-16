package test

// import (
// 	"context"
// 	"database/sql"
// 	"fmt"
// 	"sync"
// 	"testing"
// 	"time"

// 	"github.com/google/uuid"
// 	"github.com/stretchr/testify/require"
// 	"github.com/yogaprasetya22/api-gotokopedia/internal/store"
// )

// func TestCheckoutHandler(t *testing.T) {
// 	ctx := context.Background()
// 	storeTest, db, cache := NewTestStorage(t)
// 	require.NotNil(t, storeTest)

// 	// Setup test data
// 	users, err := db.Query("SELECT id, username, email FROM users")
// 	require.NoError(t, err)
// 	defer users.Close()

// 	var userID []int64
// 	for users.Next() {
// 		var id int64
// 		var username, email string
// 		require.NoError(t, users.Scan(&id, &username, &email))
// 		userID = append(userID, id)
// 	}
// 	require.NotEmpty(t, userID)

// 	for _, userID := range userID {
// 		product := createTestProduct(t, db, storeTest, 1) // Stok hanya 1
// 		createTestCart(t, storeTest, userID, product.ID, 1)

// 		t.Run("HappyPathCheckout", func(t *testing.T) {
// 			// Get cart items
// 			cart, err := storeTest.Carts.GetCartByUserID(ctx, userID)
// 			require.NoError(t, err)
// 			require.NotEmpty(t, cart.Stores)

// 			// Start checkout session
// 			checkoutSession, err := cache.Checkout.StartCheckoutSession(ctx, userID, cart.Stores)
// 			require.NoError(t, err)
// 			require.NotNil(t, checkoutSession)

// 			// Create shipping address
// 			addr := &store.ShippingAddresses{
// 				UserID:         userID,
// 				Label:          "Rumah",
// 				RecipientName:  "Test User",
// 				RecipientPhone: "08123456789",
// 				AddressLine1:   "Jl. Test No.1",
// 			}
// 			err = storeTest.ShippingAddresses.Create(ctx, addr)
// 			require.NoError(t, err)

// 			// Get payment and shipping methods
// 			paymentMethods, err := storeTest.Orders.GetPaymentMethods(ctx)
// 			require.NoError(t, err)
// 			require.NotEmpty(t, paymentMethods)

// 			shippingMethods, err := storeTest.Orders.GetShippingMethods(ctx)
// 			require.NoError(t, err)
// 			require.NotEmpty(t, shippingMethods)

// 			// Complete checkout
// 			err = storeTest.Checkout.CreateOrderFromCheckout(ctx, &store.CheckoutSession{
// 				SessionID:       checkoutSession.SessionID,
// 				UserID:          userID,
// 				CartStore:       cart.Stores,
// 				ShippingMethod:  shippingMethods[0],
// 				PaymentMethod:   paymentMethods[0],
// 				ShippingAddress: addr,
// 			})
// 			require.NoError(t, err)

// 			// Verify stock reduced
// 			updatedProduct, err := storeTest.Products.GetByID(ctx, product.ID)
// 			require.NoError(t, err)
// 			require.Equal(t, int64(0), updatedProduct.Stock)
// 		})

// 		t.Run("ConcurrentCheckout", func(t *testing.T) {
// 			// Setup 2 users with same product in cart
// 			user1 := createTestUser(t, db, storeTest)
// 			user2 := createTestUser(t, db, storeTest)
// 			product := createTestProduct(t, db, storeTest, 1) // Stok hanya 1

// 			createTestCart(t, storeTest, user1.ID, product.ID, 1)
// 			createTestCart(t, storeTest, user2.ID, product.ID, 1)

// 			var wg sync.WaitGroup
// 			successCount := 0
// 			var lock sync.Mutex

// 			// Simulate 2 users trying to checkout same product
// 			for i := 0; i < 2; i++ {
// 				wg.Add(1)
// 				go func(userNum int, userID int64) {
// 					defer wg.Done()

// 					// Get cart
// 					cart, err := storeTest.Carts.GetCartByUserID(ctx, userID)
// 					if err != nil {
// 						t.Logf("User %d failed to get cart: %v", userNum, err)
// 						return
// 					}

// 					// Start checkout
// 					session, err := cache.Checkout.StartCheckoutSession(ctx, userID, cart.Stores)
// 					if err != nil {
// 						t.Logf("User %d failed to start checkout: %v", userNum, err)
// 						return
// 					}

// 					// Create shipping address
// 					addr := &store.ShippingAddresses{
// 						UserID:         userID,
// 						Label:          "Rumah",
// 						RecipientName:  fmt.Sprintf("User %d", userNum),
// 						RecipientPhone: fmt.Sprintf("0812345678%d", userNum),
// 						AddressLine1:   fmt.Sprintf("Jl. Test No.%d", userNum),
// 					}
// 					err = storeTest.ShippingAddresses.Create(ctx, addr)
// 					if err != nil {
// 						t.Logf("User %d failed to create address: %v", userNum, err)
// 						return
// 					}

// 					// Get payment and shipping methods
// 					paymentMethods, err := storeTest.Orders.GetPaymentMethods(ctx)
// 					if err != nil || len(paymentMethods) == 0 {
// 						t.Logf("User %d failed to get payment methods: %v", userNum, err)
// 						return
// 					}

// 					shippingMethods, err := storeTest.Orders.GetShippingMethods(ctx)
// 					if err != nil || len(shippingMethods) == 0 {
// 						t.Logf("User %d failed to get shipping methods: %v", userNum, err)
// 						return
// 					}

// 					// Complete checkout
// 					err = storeTest.Checkout.CreateOrderFromCheckout(ctx, &store.CheckoutSession{
// 						SessionID:       session.SessionID,
// 						UserID:          userID,
// 						CartStore:       cart.Stores,
// 						ShippingMethod:  shippingMethods[0],
// 						PaymentMethod:   paymentMethods[0],
// 						ShippingAddress: addr,
// 					})

// 					lock.Lock()
// 					if err == nil {
// 						successCount++
// 					} else {
// 						t.Logf("User %d failed to complete checkout: %v", userNum, err)
// 					}
// 					lock.Unlock()
// 				}(i+1, []int64{user1.ID, user2.ID}[i])
// 			}

// 			wg.Wait()

// 			// Verify only 1 user succeeded
// 			require.Equal(t, 1, successCount, "Only 1 user should be able to checkout")

// 			// Verify stock is now 0
// 			updatedProduct, err := storeTest.Products.GetByID(ctx, product.ID)
// 			require.NoError(t, err)
// 			require.Equal(t, int64(0), updatedProduct.Stock)
// 		})

// 		t.Run("ExpiredSession", func(t *testing.T) {
// 			product := createTestProduct(t, db, storeTest, 1)
// 			createTestCart(t, storeTest, userID, product.ID, 1)

// 			// Get cart
// 			cart, err := storeTest.Carts.GetCartByUserID(ctx, userID)
// 			require.NoError(t, err)

// 			// Start checkout session
// 			session, err := cache.Checkout.StartCheckoutSession(ctx, userID, cart.Stores)
// 			require.NoError(t, err)

// 			// Wait for session to expire (simulate expiration by waiting longer than default expiration if possible)
// 			time.Sleep(2 * time.Second)

// 			// Try to complete checkout
// 			err = storeTest.Checkout.CreateOrderFromCheckout(ctx, &store.CheckoutSession{
// 				SessionID: session.SessionID,
// 				UserID:    userID,
// 				CartStore: cart.Stores,
// 			})
// 			require.Error(t, err)
// 			require.Contains(t, err.Error(), "not found", "Should fail with expired session")
// 		})
// 	}
// }

// // Helper functions
// func createTestUser(t *testing.T, db *sql.DB, storeTest *store.Storage) *store.User {
// 	ctx := context.Background()
// 	user := &store.User{
// 		Username: "testuser-" + uuid.New().String()[:8],
// 		Email:    "test-" + uuid.New().String()[:8] + "@example.com",
// 	}
// 	require.NoError(t, user.Password.Set("password"))
// 	var createdUser *store.User
// 	WithTransaction(t, db, func(tx *sql.Tx) error {
// 		var createErr error
// 		createdUser, createErr = storeTest.Users.CreateWithActive(ctx, tx, user)
// 		return createErr
// 	})
// 	return createdUser
// }

// func createTestProduct(t *testing.T, db *sql.DB, storeTest *store.Storage, stock int64) *store.Product {
// 	// Get category
// 	categories, err := storeTest.Categoris.GetAll(context.Background())
// 	require.NoError(t, err)
// 	require.NotEmpty(t, categories)

// 	// Get user for toko
// 	users, err := db.Query("SELECT id FROM users LIMIT 1")
// 	require.NoError(t, err)
// 	defer users.Close()
// 	require.True(t, users.Next())
// 	var userID int64
// 	require.NoError(t, users.Scan(&userID))

// 	// Create toko
// 	toko := &store.Toko{
// 		UserID:  userID,
// 		Name:    "Test Toko",
// 		Country: "Indonesia",
// 	}
// 	require.NoError(t, storeTest.Tokos.Create(context.Background(), toko))

// 	product := &store.Product{
// 		Name:        "Test Product",
// 		Description: "Test Description",
// 		Price:       10000,
// 		Stock:       int(stock),
// 		Category:    categories[0],
// 		Toko:        toko,
// 	}
// 	require.NoError(t, storeTest.Products.Create(context.Background(), product))
// 	return product
// }

// func createTestCart(t *testing.T, storeTest *store.Storage, userID, productID, quantity int64) {
// 	cart, err := storeTest.Carts.AddToCartTransaction(context.Background(), userID, productID, quantity)
// 	require.NoError(t, err)
// 	require.NotNil(t, cart)
// }
