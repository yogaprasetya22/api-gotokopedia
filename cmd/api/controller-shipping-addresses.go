package main

import (
	"errors"
	"net/http"

	"github.com/yogaprasetya22/api-gotokopedia/internal/store"
)

type shipping_addres string

const ShippingAddresContext shipping_addres = "shipping_addres"

// GetShippingAddressHandler godoc
//
//	@Summary		Get shipping address
//	@Description	Get shipping address by userID
//	@Tags			shipping-address
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	store.ShippingAddresses
//	@Failure		400	{object}	error
//	@Failure		404	{object}	error
//	@Failure		500	{object}	error
//
//	@Router			/shipping-addresses [get]
//	@Security		ApiKeyAuth
func (app *application) getShippingAddressHandler(w http.ResponseWriter, r *http.Request) {
	user := getUserFromContext(r)

	ctx := r.Context()
	shippingAddress, err := app.store.ShippingAddresses.ListByUser(ctx, user.ID)
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

// GetDefaultShippingAddressHandler gdoc
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
//	@Router			/shipping-addresses/default [get]
//
//	@Security		ApiKeyAuth
func (app *application) getDefaultShippingAddressHandler(w http.ResponseWriter, r *http.Request) {
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

type shippingAddressRequest struct {
	Label          string `json:"label"`
	RecipientName  string `json:"recipient_name"`
	RecipientPhone string `json:"recipient_phone"`
	AddressLine1   string `json:"address_line1"`
	NoteForCourier string `json:"note_for_courier"`
}

// CreateShippingAddressHandler godoc
//
//	@Summary		Create shipping address
//	@Description	Create a new shipping address
//	@Tags			shipping-address
//	@Accept			json
//	@Produce		json
//	@Param			shipping_address	body		shippingAddressRequest	true	"Shipping Address"
//	@Success		201					{object}	store.ShippingAddresses
//	@Failure		400					{object}	error
//	@Failure		500					{object}	error
//	@Router			/shipping-addresses [post]
//
//	@Security		ApiKeyAuth
func (app *application) createShippingAddressHandler(w http.ResponseWriter, r *http.Request) {
	user := getUserFromContext(r)

	ctx := r.Context()
	var payload shippingAddressRequest
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

	address := &store.ShippingAddresses{
		UserID:         user.ID,
		Label:          payload.Label,
		RecipientName:  payload.RecipientName,
		RecipientPhone: payload.RecipientPhone,
		AddressLine1:   payload.AddressLine1,
		IsDefault:      false,
	}

	if err := app.store.ShippingAddresses.Create(ctx, address); err != nil {
		switch {
		case errors.Is(err, store.ErrNotFound):
			app.notFoundResponse(w, r, err)
			return
		default:
			app.internalServerError(w, r, err)
			return
		}
	}

	if err := app.jsonResponse(w, http.StatusCreated, address); err != nil {
		app.internalServerError(w, r, err)
	}
}
