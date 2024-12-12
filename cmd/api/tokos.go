package main

import (
	"github.com/yogaprasetya22/api-gotokopedia/internal/store"
	"net/http"
)

// type TokoKey string

// const TokoCtx TokoKey = "toko"

type CreateTokoRequest struct {
	Slug         string `json:"slug"`
	Name         string `json:"name"`
	ImageProfile string `json:"image_profile,omitempty"`
	Country      string `json:"country"`
}

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
