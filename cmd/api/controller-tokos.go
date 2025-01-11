package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/yogaprasetya22/api-gotokopedia/internal/store"
)

type TokoKey string

const TokoCtx TokoKey = "toko"

type CreateTokoRequest struct {
	Slug         string `json:"slug" validate:"required"`
	Name         string `json:"name" validate:"required"`
	ImageProfile string `json:"image_profile,omitempty" validate:"required"`
	Country      string `json:"country" validate:"required"`
}

type TokoToProduct struct {
	UserID       *int64           `json:"user_id"`
	Slug         *string          `json:"slug"`
	Name         *string          `json:"name"`
	ImageProfile *string          `json:"image_profile,omitempty"`
	Country      *string          `json:"country"`
	CreatedAt    *string          `json:"created_at"`
	User         *store.User      `json:"user"`
	Products     []*store.Product `json:"products"`
}

// GetProductToko gdoc
//
//	@Summary		fetch a product toko
//	@Description	fetch a product toko by slug toko and slug product
//	@Tags			toko
//	@Accept			json
//	@Produce		json
//	@Param			slug_toko	path		string	true	"slug toko"
//	@Param			limit		query		int		false	"limit"
//	@Param			offset		query		int		false	"offset"
//	@Param			sort		query		string	false	"sort"
//	@Param			search		query		string	false	"search"
//	@Success		200			{object}	store.Product
//	@Failure		404			{object}	error
//	@Failure		500			{object}	error
//	@Router			/toko/{slug_toko} [get]
func (app *application) getProductTokoHandler(w http.ResponseWriter, r *http.Request) {
	fq := store.PaginatedFeedQuery{
		Limit:    24,
		Offset:   0,
		Sort:     "desc",
		Category: "",
		Search:   "",
	}
	slugToko := chi.URLParam(r, "slug_toko")

	ctx := r.Context()
	product, err := app.store.Products.GetProductByTokos(ctx, slugToko, fq)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, product); err != nil {
		app.internalServerError(w, r, err)
	}
}

// GetToko gdoc
//
//	@Summary		fetch a toko
//	@Description	fetch a toko by slug toko and slug product
//	@Tags			catalogue
//	@Accept			json
//	@Produce		json
//	@Param			slug_toko	path		string	true	"slug toko"
//	@Success		200			{object}	store.Product
//	@Failure		404			{object}	error
//	@Failure		500			{object}	error
//	@Router			/catalogue/{slug_toko} [get]
func (app *application) getTokoHandler(w http.ResponseWriter, r *http.Request) {
	slugToko := chi.URLParam(r, "slug_toko")

	ctx := r.Context()
	t, err := app.store.Tokos.GetBySlug(ctx, slugToko)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	products, err := app.store.Products.GetByTokoID(ctx, t.ID)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	tokoToProduct := TokoToProduct{
		UserID:       &t.UserID,
		Slug:         &t.Slug,
		Name:         &t.Name,
		ImageProfile: &t.ImageProfile,
		Country:      &t.Country,
		CreatedAt:    &t.CreatedAt,
		User:         &t.User,
		Products:     products,
	}

	if err := app.jsonResponse(w, http.StatusOK, tokoToProduct); err != nil {
		app.internalServerError(w, r, err)
	}
}

// CreateToko gdoc
//
//	@Summary		create toko
//	@Description	create toko
//	@Tags			toko
//	@Accept			json
//	@Produce		json
//	@Param			payload	body		CreateTokoRequest	true	"toko creation payload"
//	@Success		201		{object}	store.Toko
//	@Failure		400		{object}	error
//	@Failure		401		{object}	error
//	@Failure		500		{object}	error
//	@Security		ApiKeyAuth
//	@Router			/toko [post]
func (app *application) createTokoHandler(w http.ResponseWriter, r *http.Request) {
	var payload CreateTokoRequest

	err := readJSON(w, r, &payload)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	err = Validate.Struct(payload)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	t := &store.Toko{
		Slug:         payload.Slug,
		Name:         payload.Name,
		ImageProfile: payload.ImageProfile,
		Country:      payload.Country,
		UserID:       1,
	}

	if err := app.store.Tokos.Create(r.Context(), t); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusCreated, t); err != nil {
		app.internalServerError(w, r, err)
	}

}
