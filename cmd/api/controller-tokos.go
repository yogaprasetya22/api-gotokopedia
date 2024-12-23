package main

import (
	"net/http"

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
