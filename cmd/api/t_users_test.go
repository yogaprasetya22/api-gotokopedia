package main

import (
	"log"
	"net/http"
	"net/http/httptest"
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

		log.Printf("token: %s", testToken)
		checkResponseCode(t, http.StatusUnauthorized, rr.Code)
	})

	t.Run("harus mengizinkan permintaan yang diautentikasi", func(t *testing.T) {
		mockCacheStore := app.cacheStorage.Users.(*cache.MockUserStore)

		mockCacheStore.On("Get", int64(4)).Return(nil, nil).Twice()
		mockCacheStore.On("Set", mock.Anything).Return(nil)

		// Buat request
		req, err := http.NewRequest(http.MethodGet, "/api/v1/users/4", nil)
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

	t.Run("harus menekan cache terlebih dahulu dan jika tidak ada, itu mengatur pengguna pada cache", func(t *testing.T) {
		mockCacheStore := app.cacheStorage.Users.(*cache.MockUserStore)

		mockCacheStore.On("Get", int64(10)).Return(nil, nil)
		mockCacheStore.On("Get", int64(1)).Return(nil, nil)
		mockCacheStore.On("Set", mock.Anything, mock.Anything).Return(nil)

		req, err := http.NewRequest(http.MethodGet, "/api/v1/users/10", nil)
		if err != nil {
			t.Fatal(err)
		}

		session, _ := app.session.Get(req, "auth_token")
		session.Values["auth_token"] = testToken

		// Simpan session ke request context
		req = addSessionToRequestContext(req, session)

		rr := executeRequest(req, mux)

		if err := session.Save(req, rr); err != nil {
			t.Fatal(err)
		}

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

		mockCacheStore := app.cacheStorage.Users.(*cache.MockUserStore)

		req, err := http.NewRequest(http.MethodGet, "/api/v1/users/4", nil)
		if err != nil {
			t.Fatal(err)
		}

		session, _ := app.session.Get(req, "auth_token")
		session.Values["auth_token"] = testToken

		// Simpan session ke request context
		req = addSessionToRequestContext(req, session)

		rr := httptest.NewRecorder()

		if err := session.Save(req, rr); err != nil {
			t.Fatal(err)
		}

		checkResponseCode(t, http.StatusOK, rr.Code)

		mockCacheStore.AssertNotCalled(t, "Get")

		mockCacheStore.Calls = nil // Reset mock expectations
	})
}
