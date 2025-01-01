package main

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/yogaprasetya22/api-gotokopedia/internal/store"
	"golang.org/x/oauth2"
)

func (app *application) googleLoginHandler(w http.ResponseWriter, r *http.Request) {
	url := app.googleOauthConfig.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func (app *application) googleCallbackHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := app.session.Get(r, "auth_token")
	state := r.FormValue("state")
	if state != "state-token" {
		http.Error(w, "Invalid state", http.StatusBadRequest)
		return
	}

	code := r.FormValue("code")
	token, err := app.googleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	client := app.googleOauthConfig.Client(context.Background(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}
	defer resp.Body.Close()

	userInfo := struct {
		Sub      string `json:"sub"`
		Email    string `json:"email"`
		Name     string `json:"name"`
		PhotoURL string `json:"picture"`
	}{}
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	// Periksa apakah pengguna sudah ada berdasarkan GoogleID atau Email
	existingUser, err := app.store.Users.GetByGoogleID(r.Context(), userInfo.Sub)
	if err != nil && err != store.ErrNotFound {
		app.internalServerError(w, r, err)
		return
	}

	if existingUser == nil {
		// Buat pengguna baru jika belum ada
		newUser := &store.User{
			GoogleID: userInfo.Sub,
			Email:    userInfo.Email,
			Username: userInfo.Name,
			Picture:  userInfo.PhotoURL,
			Role: store.Role{
				Name: "user",
			},
		}

		if err := app.store.Users.CreateByOAuth(r.Context(), newUser); err != nil {
			switch err {
			case store.ErrDuplicateEmail:
				app.badRequestResponse(w, r, err)
			case store.ErrDuplicateUsername:
				app.badRequestResponse(w, r, err)
			default:
				app.internalServerError(w, r, err)
			}
			return
		}

		existingUser = newUser
	}

	// Generate JWT token
	claims := jwt.MapClaims{
		"sub": existingUser.ID,
		"exp": time.Now().Add(app.config.auth.token.exp).Unix(),
		"iat": time.Now().Unix(),
		"nbf": time.Now().Unix(),
		"iss": app.config.auth.token.iss,
		"aud": app.config.auth.token.iss,
	}

	jwtToken, err := app.authenticator.GenerateToken(claims)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	// Set token in cookie
	session.Values["auth_token"] = jwtToken
	if err := session.Save(r, w); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	// Redirect to home
	http.Redirect(w, r, app.config.frontendURL, http.StatusSeeOther)

}
