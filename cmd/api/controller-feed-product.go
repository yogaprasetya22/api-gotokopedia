package main

import (
	"net/http"

	"github.com/yogaprasetya22/api-gotokopedia/internal/store"
)

// GetCatalogueFeed gdoc
//
//	@Summary		fetch catalogue feed
//	@Description	fetch catalogue feed with pagination
//	@Tags			catalogue
//	@Accept			json
//	@Produce		json
//	@Param			limit	query		int		false	"limit"
//	@Param			offset	query		int		false	"offset"
//	@Param			sort	query		string	false	"sort"
//	@Param			search	query		string	false	"search"
//	@Success		200		{object}	[]store.Product
//	@Failure		400		{object}	error
//	@Failure		500		{object}	error
//	@Security		ApiKeyAuth
//	@Router			/catalogue/feed [get]
func (app *application) getProductFeedHandler(w http.ResponseWriter, r *http.Request) {
	fq := store.PaginatedFeedQuery{
		Limit:    12,
		Offset:   0,
		Sort:     "desc",
		Category: "",
		Search:   "",
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

// GetCatalogueCategoryFeed gdoc
//
//	@Summary		fetch catalogue category feed
//	@Description	fetch catalogue category feed with pagination
//	@Tags			catalogue
//	@Accept			json
//	@Produce		json
//	@Param			limit		query		int		false	"limit"
//	@Param			offset		query		int		false	"offset"
//	@Param			sort		query		string	false	"sort"
//	@Param			category	query		string	false	"category"
//	@Param			search		query		string	false	"search"
//	@Success		200			{object}	[]store.Product
//	@Failure		400			{object}	error
//	@Failure		500			{object}	error
//	@Security		ApiKeyAuth
//	@Router			/catalogue [get]
func (app *application) getProductCategoryFeed(w http.ResponseWriter, r *http.Request) {
	fq := store.PaginatedFeedQuery{
		Limit:    12,
		Offset:   0,
		Sort:     "desc",
		Category: "",
		Search:   "",
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

	ctx := r.Context()

	var products []*store.Product
	if fq.Category == "" || fq.Search == "" || fq.Offset == 0 {
		products, err = app.store.Products.GetAllProduct(ctx, fq)
	} else {
		products, err = app.store.Products.GetProductCategoryFeed(ctx, fq)
	}

	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, products); err != nil {
		app.internalServerError(w, r, err)
	}
}
