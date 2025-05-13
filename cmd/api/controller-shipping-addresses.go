package main

import (
	"errors"
	"net/http"

	"github.com/yogaprasetya22/api-gotokopedia/internal/store"
)

type shipping_addres string

const ShippingAddresContext shipping_addres = "shipping_addres"

// GetShippingAddressHandler gdoc
//
//	@Summary		Get default shipping address
//	@Description	Get the default shipping address for the authenticated user
//	@Tags			shipping-address
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	store.ShippingAddresses
//	@Failure		404	{object}	error
//	@Failure		500	{object}	error
//	@Security		ApiKeyAuth
//	@Router			/shipping-addresses [get]
// @Security		ApiKeyAuth
func (app *application) getShippingAddressHandler(w http.ResponseWriter, r *http.Request) {
	user := getUserFromContext(r)

	ctx := r.Context()
	shippingAddress, err := app.store.ShippingAddresses.GetDefaultAddress(ctx, user.ID)
	if err != nil {
		switch {
		case errors.Is(err, store.ErrNotFound):
			app.notFoundResponse(w, r, err)
		default:
			app.internalServerError(w, r, err)
		}
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, shippingAddress); err != nil {
		app.internalServerError(w, r, err)
	}
}
