package main

import (
	"github.com/yogaprasetya22/api-gotokopedia/internal/store"
	"net/http"
)

type CategoriedKey string

const CategoriedCtx CategoriedKey = "category"

type CreateCategoryRequest struct {
	Name        string `json:"name" validate:"required"`
	Slug        string `json:"slug" validate:"required"`
	Description string `json:"description,omitempty" validate:"required"`
}

// GetAllCategory godoc
//
//	@Summary		get all category
//	@Description	get all category
//	@Tags			category
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		store.Category
//	@Failure		400	{object}	error
//	@Failure		401	{object}	error
//	@Failure		500	{object}	error
//	@Router			/category [get]
func (app *application) getAllCategoryHandler(w http.ResponseWriter, r *http.Request) {
	categories, err := app.store.Categoris.GetAll(r.Context())
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, categories); err != nil {
		app.internalServerError(w, r, err)
	}

}

// CreateCategory gdoc
//
//	@Summary		create category
//	@Description	create category
//	@Tags			category
//	@Accept			json
//	@Produce		json
//	@Param			payload	body		CreateCategoryRequest	true	"category creation payload"
//	@Success		201		{object}	store.Category
//	@Failure		400		{object}	error
//	@Failure		401		{object}	error
//	@Failure		500		{object}	error
//	@Security		ApiKeyAuth
//	@Router			/category [post]
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
