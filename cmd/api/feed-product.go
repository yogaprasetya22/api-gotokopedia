package main

import (
	"github.com/yogaprasetya22/api-gotokopedia/internal/store"
	"net/http"
)

func (app *application) getProductFeedHandler(w http.ResponseWriter, r *http.Request) {
	fq := store.PaginatedFeedQuery{
		Limit:      12,
		Offset:     0,
		Sort:       "desc",
		Category:   "",
		Search:     "",
	}

	fq, err := fq.Parse(r)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := Validate.Struct(fq); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	users, err := app.store.Products.GetProductFeed(r.Context(), []int64{3, 4}, fq)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, users); err != nil {
		app.internalServerError(w, r, err)
	}
}
