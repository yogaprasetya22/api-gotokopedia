package main

import (
	"github.com/yogaprasetya22/api-gotokopedia/internal/store"
	"net/http"
)

// type TokoKey string

// const TokoCtx TokoKey = "category"

type CreateCategoryRequest struct {
	Name        string `json:"name" validate:"required"`
	Slug        string `json:"slug" validate:"required"`
	Description string `json:"description,omitempty" validate:"required"`
}

func (app *application) createCategoryHandler(w http.ResponseWriter, r *http.Request) {
	var payload CreateCategoryRequest

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

	c := &store.Category{
		Name:        payload.Name,
		Slug:        payload.Slug,
		Description: payload.Description,
	}

	if err := app.store.Categoris.Create(r.Context(), c); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusCreated, c); err != nil {
		app.internalServerError(w, r, err)
	}

}
