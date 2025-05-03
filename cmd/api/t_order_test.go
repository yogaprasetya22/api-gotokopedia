package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"net/http/httptest"

// 	"strconv"
// 	"strings"
// 	"testing"
// 	"time"

// 	"github.com/golang-jwt/jwt/v5"
// 	"github.com/stretchr/testify/mock"
// 	"github.com/yogaprasetya22/api-gotokopedia/internal/store/cache"
// )

// func TestOrderFlow(t *testing.T) {
// 	withRedis := config{
// 		redisCfg: redisConfig{
// 			enabled: true,
// 		},
// 	}

// 	app := newTestApplication(t, withRedis)
// 	mux := app.mount()

// 	userIDs := make([]int64, 10)
// 	// Step 1: Create 10 new accounts
// 	t.Run("Create 10 new accounts with mock", func(t *testing.T) {
// 		// Setup mock
// 		mockCacheStore := app.cacheStorage.Users.(*cache.MockUserStore)

// 		mockCacheStore.On("Get", mock.Anything).Return(func(_ int64) error {
// 			newID := int64(len(userIDs) + 1)
// 			_ = newID // Use newID as needed
// 			return nil
// 		})

// 		mockCacheStore.On("Set", mock.Anything).Return(nil)

// 		for i := 0; i < 10; i++ {
// 			body := fmt.Sprintf(
// 				`{"email":"testuser%d@example.com","password":"password","username":"testuser%d"}`,
// 				i, i,
// 			)
// 			req := httptest.NewRequest(http.MethodPost, "/api/v1/authentication/user", strings.NewReader(body))
// 			req.Header.Set("Content-Type", "application/json")

// 			rr := httptest.NewRecorder()
// 			mux.ServeHTTP(rr, req)

// 			userIDs[i] = int64(i+1) // Simpan ID yang benar dari response
// 		}
// 	})


// 	// Step 2: Generate tokens for each user
// 	tokens := make([]string, 10)
// 	t.Run("Generate tokens for each user", func(t *testing.T) {
// 		for i, userID := range userIDs {
// 			claims := jwt.MapClaims{
// 				"sub": userID,
// 				"exp": time.Now().Add(app.config.auth.token.exp).Unix(),
// 				"iat": time.Now().Unix(),
// 				"nbf": time.Now().Unix(),
// 				"iss": app.config.auth.token.iss,
// 				"aud": app.config.auth.token.iss,
// 			}

// 			token, err := app.authenticator.GenerateToken(claims)

// 			if err != nil {
// 				t.Fatalf("failed to generate token for user %d: %v", userID, err)
// 			}

// 			tokens[i] = token

// 			// Ensure the log is printed
// 			t.Logf("Generated token for user %d: %s", userID, token)
// 			log.Println("Order flow test completed successfully")
// 		}
// 	})

// 	// Step 3: Create carts and add items for each user
// 	t.Run("Retrieve cart IDs and include in order body", func(t *testing.T) {
// 		for i, token := range tokens {
// 			bodycart := strings.NewReader(`{"product_id":` + strconv.Itoa(258) + `,"quantity":` + strconv.Itoa(1) + `}`)
// 			req, err := http.NewRequest(http.MethodPost, "/api/v1/cart", bodycart)
// 			if err != nil {
// 				t.Fatalf("failed to retrieve cart for user %d: %v", userIDs[i], err)
// 			}

// 			session, _ := app.session.Get(req, "auth_token")
// 			session.Values["auth_token"] = token

// 			// Simpan session ke request context
// 			req = addSessionToRequestContext(req, session)

// 			rr := httptest.NewRecorder()
// 			fmt.Println(rr.Body)

// 			if err := session.Save(req, rr); err != nil {
// 				t.Fatal(err)
// 			}

// 			mux.ServeHTTP(rr, req)

// 			if rr.Code != http.StatusCreated {
// 				t.Fatalf("failed to retrieve cart for user %d, got status: %d", userIDs[i], rr.Code)
// 			}

// 			// Parse the response to extract cart IDs
// 			// var cartResponse struct {
// 			// 	Carts []struct {
// 			// 		ID int64 `json:"id"`
// 			// 	} `json:"carts"`
// 			// }

// 			// if err := json.NewDecoder(rr.Body).Decode(&cartResponse); err != nil {
// 			// 	t.Fatalf("failed to decode cart response for user %d: %v", userIDs[i], err)
// 			// }

// 			// if len(cartResponse.Carts) == 0 {
// 			// 	t.Fatalf("no carts found for user %d", userIDs[i])
// 			// }

// 			// // Use the first cart ID for the order
// 			// cartID := cartResponse.Carts[0].ID

// 			// // Create the order body with the cart ID
// 			// bodyorder := strings.NewReader(`{"cart_id":` + strconv.FormatInt(cartID, 10) + `,"payment_method":"credit_card","shipping_address":"123 Test Street"}`)
// 			// req, err = http.NewRequest(http.MethodPost, "/api/v1/orders", bodyorder)
// 			// if err != nil {
// 			// 	t.Fatalf("failed to create order request with cart ID for user %d: %v", userIDs[i], err)
// 			// }

// 			// session, _ = app.session.Get(req, "auth_token")
// 			// session.Values["auth_token"] = token

// 			// // Simpan session ke request context
// 			// req = addSessionToRequestContext(req, session)

// 			// rr = httptest.NewRecorder()

// 			// if err := session.Save(req, rr); err != nil {
// 			// 	t.Fatal(err)
// 			// }

// 			// mux.ServeHTTP(rr, req)

// 			// checkResponseCode(t, http.StatusCreated, rr.Code)

// 			// if rr.Code != http.StatusCreated {
// 			// 	t.Fatalf("failed to create order with cart ID for user %d, got status: %d", userIDs[i], rr.Code)
// 			// }
// 		}
// 	})
// 	log.Println("Order flow test completed successfully")
// }
