package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/sessions"
	"github.com/yogaprasetya22/api-gotokopedia/internal/auth"
	"github.com/yogaprasetya22/api-gotokopedia/internal/ratelimiter"
	"github.com/yogaprasetya22/api-gotokopedia/internal/store"
	"github.com/yogaprasetya22/api-gotokopedia/internal/store/cache"
	"go.uber.org/zap"
)

func newTestApplication(t *testing.T, cfg config) *application {
	t.Helper()

	// logger := zap.Must(zap.NewDevelopment()).Sugar()
	logger := zap.NewExample().Sugar()
	mockStore := store.NewMockStore()
	mockCacheStore := cache.NewMockStore()

	testAuth := &auth.TestAuthenticator{}

	// Tambahkan mock session store
	mockSessionStore := sessions.NewCookieStore([]byte("jagreskuy112"))
	mockSessionStore.MaxAge(86400 * 30)

	mockSessionStore.Options.Path = "/"
	mockSessionStore.Options.HttpOnly = true // HttpOnly should always be enabled
	mockSessionStore.Options.Secure = true
	mockSessionStore.Options.SameSite = http.SameSiteNoneMode

	// Rate limiter
	rateLimiter := ratelimiter.NewFixedWindowLimiter(
		cfg.rateLimiter.RequestsPerTimeFrame,
		cfg.rateLimiter.TimeFrame,
	)

	return &application{
		logger:        logger,
		store:         mockStore,
		cacheStorage:  mockCacheStore,
		authenticator: testAuth,
		config:        cfg,
		rateLimiter:   rateLimiter,
	}
}

func addJWTToRequest(req *http.Request, token string) {
	req.AddCookie(&http.Cookie{
		Name:  "auth_token",
		Value: token,
		Path:  "/",
	})
}

func executeRequest(req *http.Request, mux http.Handler) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	mux.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("kode respons yang diharapkan %d, tetapi mendapat %d", expected, actual)
	}
}

func addSessionToRequestContext(r *http.Request, session *sessions.Session) *http.Request {
	ctx := r.Context()
	ctx = context.WithValue(ctx, "session", session)
	return r.WithContext(ctx)
}
