package main

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/mock"
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

	testToken, err := app.authenticator.GenerateToken(nil)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("seharusnya tidak mengizinkan permintaan yang tidak otentikasi", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/api/v1/users/4", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := executeRequest(req, mux)

		checkResponseCode(t, http.StatusUnauthorized, rr.Code)
	})

	t.Run("harus mengizinkan permintaan yang diautentikasi", func(t *testing.T) {
		mockCacheStore := app.cacheStorage.Users.(*cache.MockUserStore)

		mockCacheStore.On("Get", int64(4)).Return(nil, nil).Twice()
		mockCacheStore.On("Set", mock.Anything).Return(nil)

		req, err := http.NewRequest(http.MethodGet, "/api/v1/users/4", nil)
		if err != nil {
			t.Fatal(err)
		}

		req.AddCookie(&http.Cookie{
			Name:  "auth_token",
			Value: testToken, // menggunakan token yang sudah Anda buat sebelumnya
			Path:  "/",
		})

		rr := executeRequest(req, mux)

		checkResponseCode(t, http.StatusOK, rr.Code)

		mockCacheStore.Calls = nil // Reset mock expectations
	})

	t.Run("harus menekan cache terlebih dahulu dan jika tidak ada, itu mengatur pengguna pada cache", func(t *testing.T) {
		mockCacheStore := app.cacheStorage.Users.(*cache.MockUserStore)

		mockCacheStore.On("Get", int64(4)).Return(nil, nil)
		mockCacheStore.On("Get", int64(1)).Return(nil, nil)
		mockCacheStore.On("Set", mock.Anything, mock.Anything).Return(nil)

		req, err := http.NewRequest(http.MethodGet, "/api/v1/users/4", nil)
		if err != nil {
			t.Fatal(err)
		}

		req.AddCookie(&http.Cookie{
			Name:  "auth_token",
			Value: testToken, // menggunakan token yang sudah Anda buat sebelumnya
			Path:  "/",
		})

		rr := executeRequest(req, mux)

		checkResponseCode(t, http.StatusOK, rr.Code)

		mockCacheStore.AssertNumberOfCalls(t, "Get", 2)

		mockCacheStore.Calls = nil // Reset mock expectations
	})

	t.Run("tidak boleh menekan cache jika tidak diaktifkan", func(t *testing.T) {
		withRedis := config{
			redisCfg: redisConfig{
				enabled: false,
			},
		}

		app := newTestApplication(t, withRedis)
		mux := app.mount()

		mockCacheStore := app.cacheStorage.Users.(*cache.MockUserStore)

		req, err := http.NewRequest(http.MethodGet, "/api/v1/users/4", nil)
		if err != nil {
			t.Fatal(err)
		}

		req.AddCookie(&http.Cookie{
			Name:  "auth_token",
			Value: testToken, // menggunakan token yang sudah Anda buat sebelumnya
			Path:  "/",
		})

		rr := executeRequest(req, mux)

		checkResponseCode(t, http.StatusOK, rr.Code)

		mockCacheStore.AssertNotCalled(t, "Get")

		mockCacheStore.Calls = nil // Reset mock expectations
	})
}
