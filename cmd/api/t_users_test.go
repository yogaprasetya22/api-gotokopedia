package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/yogaprasetya22/api-gotokopedia/internal/store/cache"
)

func TestGetUser(t *testing.T) {
	withRedis := config{
		redisCfg: redisConfig{
			enabled: true,
		},
	}

	app := newTestApplication(t, withRedis)
	mux := app.mount()

	// Generate token JWT yang valid
	testToken, err := app.authenticator.GenerateToken(nil)
	require.NoError(t, err)

	t.Run("seharusnya tidak mengizinkan permintaan yang tidak otentikasi", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/api/v1/users/4", nil)
		require.NoError(t, err)

		rr := httptest.NewRecorder()
		mux.ServeHTTP(rr, req)

		log.Printf("token: %s", testToken)
		require.Equal(t, http.StatusUnauthorized, rr.Code)
	})

	t.Run("harus mengizinkan permintaan yang diautentikasi", func(t *testing.T) {
		mockCacheStore := app.cacheStorage.Users.(*cache.MockUserStore)

		// Mock cache get dan set sesuai ekspektasi
		mockCacheStore.On("Get", int64(4)).Return(nil, nil).Twice()
		mockCacheStore.On("Set", mock.Anything).Return(nil)

		// Buat request
		req, err := http.NewRequest(http.MethodGet, "/api/v1/users/4", nil)
		require.NoError(t, err)

		// Tambahkan cookie JWT untuk autentikasi
		addJWTToRequest(req, testToken)

		rr := httptest.NewRecorder()
		mux.ServeHTTP(rr, req)

		require.Equal(t, http.StatusOK, rr.Code)

		// Reset mock
		mockCacheStore.Calls = nil
	})

	t.Run("harus menekan cache terlebih dahulu dan jika tidak ada, itu mengatur pengguna pada cache", func(t *testing.T) {
		mockCacheStore := app.cacheStorage.Users.(*cache.MockUserStore)

		mockCacheStore.On("Get", int64(10)).Return(nil, nil)
		mockCacheStore.On("Get", int64(1)).Return(nil, nil)
		mockCacheStore.On("Set", mock.Anything, mock.Anything).Return(nil)

		req, err := http.NewRequest(http.MethodGet, "/api/v1/users/10", nil)
		require.NoError(t, err)

		// Tambahkan cookie JWT ke request
		addJWTToRequest(req, testToken)

		rr := httptest.NewRecorder()
		mux.ServeHTTP(rr, req)

		require.Equal(t, http.StatusOK, rr.Code)

		mockCacheStore.AssertNumberOfCalls(t, "Get", 2)

		mockCacheStore.Calls = nil
	})

	t.Run("tidak boleh menekan cache jika tidak diaktifkan", func(t *testing.T) {
		withRedisDisabled := config{
			redisCfg: redisConfig{
				enabled: false,
			},
		}

		appWithoutRedis := newTestApplication(t, withRedisDisabled)
		muxWithoutRedis := appWithoutRedis.mount()

		mockCacheStore := appWithoutRedis.cacheStorage.Users.(*cache.MockUserStore)

		req, err := http.NewRequest(http.MethodGet, "/api/v1/users/4", nil)
		require.NoError(t, err)

		addJWTToRequest(req, testToken)

		rr := httptest.NewRecorder()
		muxWithoutRedis.ServeHTTP(rr, req)

		require.Equal(t, http.StatusOK, rr.Code)

		// Pastikan cache tidak dipanggil karena redis nonaktif
		mockCacheStore.AssertNotCalled(t, "Get")

		mockCacheStore.Calls = nil
	})
}
