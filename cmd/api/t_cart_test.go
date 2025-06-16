package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/yogaprasetya22/api-gotokopedia/internal/store"
	"github.com/yogaprasetya22/api-gotokopedia/internal/store/cache"
	// import lain sesuai project kamu, misal app, store, cache, dll
)

func TestCartAndLogout(t *testing.T) {
	app := newTestApplication(t, config{
		redisCfg: redisConfig{
			enabled: true,
		},
	})

	mux := app.mount()

	// Mock cache user store
	mockCacheUser := app.cacheStorage.Users.(*cache.MockUserStore)
	mockCacheUser.On("Get", int64(10)).Return(nil, nil)
	mockCacheUser.On("Set", mock.Anything).Return(nil)
	mockCacheUser.On("Delete", mock.Anything).Return()

	// Mock cart store
	mockCart := &store.Cart{ID: 1, UserID: 10}
	app.store.Carts = &store.MockCartStore{}
	app.store.Carts.(*store.MockCartStore).On("AddToCartTransaction", mock.Anything, int64(10), int64(258), int64(1)).Return(mockCart, nil)

	// Generate valid JWT token
	testToken, err := app.authenticator.GenerateToken(nil)
	require.NoError(t, err)

	// Test 1: request GET /api/v1/cart tanpa token harus 401
	t.Run("Unauthorized GET /api/v1/cart without token", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/api/v1/cart", nil)
		require.NoError(t, err)

		rr := httptest.NewRecorder()
		mux.ServeHTTP(rr, req)

		require.Equal(t, http.StatusUnauthorized, rr.Code)
	})

	// Test 2: request GET /api/v1/cart dengan token harus berhasil
	t.Run("Authorized GET /api/v1/cart with token", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/api/v1/cart", nil)
		require.NoError(t, err)

		addJWTToRequest(req, testToken)

		rr := httptest.NewRecorder()
		mux.ServeHTTP(rr, req)

		require.Equal(t, http.StatusOK, rr.Code)
	})

	// Test 3: POST /api/v1/cart dengan token harus berhasil buat cart baru
	t.Run("Authorized POST /api/v1/cart create cart", func(t *testing.T) {
		body := strings.NewReader(`{"product_id":258,"quantity":1}`)
		req, err := http.NewRequest(http.MethodPost, "/api/v1/cart", body)
		require.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")

		addJWTToRequest(req, testToken)

		rr := httptest.NewRecorder()
		mux.ServeHTTP(rr, req)

		require.Equal(t, http.StatusCreated, rr.Code)
	})

	// Test 4: Logout harus hapus cookie auth_token
	t.Run("Logout clears auth_token cookie", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodPost, "/logout", nil)
		require.NoError(t, err)

		addJWTToRequest(req, testToken)

		rr := httptest.NewRecorder()
		app.logoutHandler(rr, req)

		require.Equal(t, http.StatusOK, rr.Code)

		cookies := rr.Result().Cookies()
		found := false
		for _, c := range cookies {
			if c.Name == "auth_token" {
				found = true
				// MaxAge -1 artinya cookie dihapus di browser
				require.Equal(t, -1, c.MaxAge)
			}
		}
		require.True(t, found, "auth_token cookie must be cleared on logout")
	})
}
