package main

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/yogaprasetya22/api-gotokopedia/internal/store"
)

type cartKey string

const CartContext cartKey = "cart"

type AddToCartPayload struct {
	ProductID int64 `json:"product_id" validate:"required"`
	Quantity  int64 `json:"quantity" validate:"required,min=1"`
}

type UpdateCartItemPayload struct {
	ProductID int64 `json:"product_id" validate:"required"`
	Quantity  int64 `json:"quantity" validate:"required,min=0"`
}

// GetCart gdoc
//
//	@Summary		fetch a cart
//	@Description	fetch a cart
//	@Tags			cart
//	@Accept			json
//	@Produce		json
//	@Param			limit	query		int		false	"limit"
//	@Param			offset	query		int		false	"offset"
//	@Param			sort	query		string	false	"sort"
//	@Success		200		{object}	store.Cart
//	@Failure		400		{object}	error
//	@Failure		404		{object}	error
//	@Failure		500		{object}	error
//	@Router			/cart [get]
func (app *application) getCartsHandler(w http.ResponseWriter, r *http.Request) {
	fq := store.PaginatedFeedQuery{
		Limit:    24,
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

	user := getUserFromContext(r)

	ctx := r.Context()
	cart, err := app.store.Carts.GetCartByUserIDPQ(ctx, user.ID, fq)
	if err != nil {
		switch {
		case errors.Is(err, store.ErrNotFound):
			app.notFoundResponse(w, r, err)
		default:
			app.internalServerError(w, r, err)
		}
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, cart); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

// DetailCartStoreItem gdoc
//
//	@Summary		fetch a cart store item
//	@Description	fetch a cart store item by ID
//	@Tags			cart
//	@Accept			json
//	@Produce		json
//	@Param			cartStoreItemID	path		string	true	"cart store item ID"
//	@Success		200				{object}	store.Cart
//	@Failure		400				{object}	error
//	@Failure		404				{object}	error
//	@Failure		500				{object}	error
//	@Security		ApiKeyAuth
//	@Router			/cart/item/{cartStoreItemID} [get]
func (app *application) GetDetailCartStoreHandler(w http.ResponseWriter, r *http.Request) {
	cartStoreID, err := uuid.Parse(chi.URLParam(r, "cartStoreID"))
	if err != nil {
		app.badRequestResponse(w, r, errors.New("invalid cartStoreItemID format"))
		return
	}
	user := getUserFromContext(r)

	cart, err := app.store.Carts.GetDetailCartByCartStoreID(r.Context(), cartStoreID, user.ID)
	if err != nil {
		switch {
		case errors.Is(err, store.ErrNotFound):
			app.notFoundResponse(w, r, err)
		default:
			app.internalServerError(w, r, err)
		}
	}

	if err := app.jsonResponse(w, http.StatusOK, cart); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

// CreateCart gdoc
//
//	@Summary		create a cart
//	@Description	create a cart
//	@Tags			cart
//	@Accept			json
//	@Produce		json
//	@Param			payload	body		AddToCartPayload	true	"payload"
//	@Success		201		{object}	store.Cart
//	@Failure		400		{object}	error
//	@Failure		404		{object}	error
//	@Failure		500		{object}	error
//	@Security		ApiKeyAuth
//	@Router			/cart [post]
func (app *application) createCartHandler(w http.ResponseWriter, r *http.Request) {
	user := getUserFromContext(r)
	var payload AddToCartPayload

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

	cart, err := app.store.Carts.AddToCartTransaction(r.Context(), user.ID, payload.ProductID, payload.Quantity)
	if err != nil {
		switch err {
		case store.ErrNotFound:
			app.notFoundResponse(w, r, err)
		default:
			app.internalServerError(w, r, err)
		}
		return
	}

	if err := app.jsonResponse(w, http.StatusCreated, cart); err != nil {
		app.internalServerError(w, r, err)
		return
	}

}

// AddingQuantityCartStoreItem gdoc
//
//	@Summary		add quantity to cart store item
//	@Description	add quantity to cart store item
//	@Tags			cart
//	@Accept			json
//	@Produce		json
//	@Param			cartStoreItemID	path		string	true	"cart store item ID"
//	@Success		201				{object}	store.Cart
//	@Failure		400				{object}	error
//	@Failure		404				{object}	error
//	@Failure		500				{object}	error
//	@Security		ApiKeyAuth
//	@Router			/cart/item/{cartStoreItemID}/increase [patch]
func (app *application) AddingQuantityCartStoreItemHandler(w http.ResponseWriter, r *http.Request) {
	cartStoreItemID, err := uuid.Parse(chi.URLParam(r, "cartStoreItemID"))
	if err != nil {
		app.badRequestResponse(w, r, errors.New("invalid cartStoreItemID format"))
		return
	}
	user := getUserFromContext(r)

	cart := app.store.Carts.IncreaseQuantityCartStoreItemTransaction(r.Context(), cartStoreItemID, user.ID)

	if err := app.jsonResponse(w, http.StatusCreated, cart); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

// RemovingQuantityCartStoreItem gdoc
//
//	@Summary		remove quantity to cart store item
//	@Description	remove quantity to cart store item
//	@Tags			cart
//	@Accept			json
//	@Produce		json
//	@Param			cartStoreItemID	path		string	true	"cart store item ID"
//	@Success		201				{object}	store.Cart
//	@Failure		400				{object}	error
//	@Failure		404				{object}	error
//	@Failure		500				{object}	error
//	@Security		ApiKeyAuth
//	@Router			/cart/item/{cartStoreItemID}/decrease [patch]
func (app *application) RemovingQuantityCartStoreItemHandler(w http.ResponseWriter, r *http.Request) {
	cartStoreItemID, err := uuid.Parse(chi.URLParam(r, "cartStoreItemID"))
	if err != nil {
		app.badRequestResponse(w, r, errors.New("invalid cartStoreItemID format"))
		return
	}
	user := getUserFromContext(r)

	cart := app.store.Carts.DecreaseQuantityCartStoreItemTransaction(r.Context(), cartStoreItemID, user.ID)

	if err := app.jsonResponse(w, http.StatusCreated, cart); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}
