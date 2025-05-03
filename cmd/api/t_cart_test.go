package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/yogaprasetya22/api-gotokopedia/internal/store"
	"github.com/yogaprasetya22/api-gotokopedia/internal/store/cache"
)

func TestGetCart(t *testing.T) {
	withRedis := config{
		redisCfg: redisConfig{
			enabled: true,
		},
	}

	app := newTestApplication(t, withRedis)
	mux := app.mount()

	// Setup mock cache user store
	mockCacheStore := app.cacheStorage.Users.(*cache.MockUserStore)
	mockCacheStore.On("Get", int64(10)).Return(nil, nil)
	mockCacheStore.On("Set", mock.Anything).Return(nil)
	mockCacheStore.On("Delete", mock.Anything).Return()

	// Setup mock store cart jika diperlukan
	mockCart := &store.Cart{ID: 1, UserID: 10}
	app.store.Carts = &store.MockCartStore{}
	app.store.Carts.(*store.MockCartStore).On("AddToCartTransaction", mock.Anything, int64(10), int64(258), int64(1)).Return(mockCart, nil)

	testToken, err := app.authenticator.GenerateToken(nil)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("seharusnya tidak mengizinkan permintaan yang tidak otentikasi", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/api/v1/cart", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := executeRequest(req, mux)

		log.Printf("token: %s", testToken)
		checkResponseCode(t, http.StatusUnauthorized, rr.Code)
	})

	t.Run("harus mengizinkan permintaan yang diautentikasi", func(t *testing.T) {
		mockCacheStore := app.cacheStorage.Carts.(*cache.MockCartStore)

		mockCacheStore.On("Get", int64(5)).Return(nil, nil).Twice()
		mockCacheStore.On("Set", mock.Anything).Return(nil)

		// Buat request
		req, err := http.NewRequest(http.MethodGet, "/api/v1/cart", nil)
		if err != nil {
			t.Fatal(err)
		}

		// Tambahkan session dengan token sebelum eksekusi
		session, _ := app.session.Get(req, "auth_token")
		session.Values["auth_token"] = testToken

		// Simpan session ke request context
		req = addSessionToRequestContext(req, session)

		rr := httptest.NewRecorder()

		if err := session.Save(req, rr); err != nil {
			t.Fatal(err)
		}

		checkResponseCode(t, http.StatusOK, rr.Code)

		mockCacheStore.Calls = nil // Reset mock expectations
	})

	t.Run("membuat cart baru", func(t *testing.T) {
		bodycart := strings.NewReader(`{"product_id":` + strconv.Itoa(258) + `,"quantity":` + strconv.Itoa(1) + `}`)
		req, err := http.NewRequest(http.MethodPost, "/api/v1/cart", bodycart)
		if err != nil {
			t.Fatalf("error creating request: %v", err)
		}

		session, _ := app.session.Get(req, "auth_token")
		session.Values["auth_token"] = testToken

		// Simpan session ke request context
		req = addSessionToRequestContext(req, session)

		rr := executeRequest(req, mux)

		if err := session.Save(req, rr); err != nil {
			t.Fatal(err)
		}

		fmt.Printf("session: %v\n", session.Values["auth_token"])

		// Panggil handler setelah semua setup selesai
		app.createCartHandler(rr, req)

		mux.ServeHTTP(rr, req)

		checkResponseCode(t, http.StatusCreated, rr.Code)
	})
}
