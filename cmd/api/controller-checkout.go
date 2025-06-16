package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/yogaprasetya22/api-gotokopedia/internal/store"
)

type StartCheckoutPayload struct {
	CartStoreID []uuid.UUID `json:"cart_store_id" validate:"required"`
}

type CompleteCheckoutPayload struct {
	SessionID         string `json:"session_id" validate:"required"`
	ShippingMethodID  int64  `json:"shipping_method_id" validate:"required"`
	PaymentMethodID   int64  `json:"payment_method_id" validate:"required"`
	ShippingAddressID string `json:"shipping_address_id" validate:"required"`
	Notes             string `json:"notes"`
}

func (p *StartCheckoutPayload) UnmarshalJSON(data []byte) error {
	var temp struct {
		CartStoreID []string `json:"cart_store_id"`
	}

	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	for _, idStr := range temp.CartStoreID {
		id, err := uuid.Parse(idStr)
		if err != nil {
			return fmt.Errorf("invalid UUID format: %v", idStr)
		}
		p.CartStoreID = append(p.CartStoreID, id)
	}

	return nil
}

// GetCheckoutBySession godoc
//
//	@Summary		Get checkout session details
//	@Description	Get detailed information about a checkout session
//	@Tags			checkout
//	@Accept			json
//	@Produce		json
//	@Param			session_id	path		string	true	"Checkout Session ID"
//	@Success		200			{object}	store.CheckoutSession
//	@Failure		400			{object}	error
//	@Failure		401			{object}	error
//	@Failure		404			{object}	error
//	@Failure		500			{object}	error
//	@Security		ApiKeyAuth
//	@Router			/checkout/{session_id} [get]
func (app *application) getCheckoutBySessionHandler(w http.ResponseWriter, r *http.Request) {
	user := getUserFromContext(r)

	// Ambil sessionID dari path parameter
	sessionID := chi.URLParam(r, "session_id")
	if sessionID == "" {
		app.badRequestResponse(w, r, errors.New("session_id is required"))
		return
	}

	// Dapatkan session dari Redis
	checkoutSession, err := app.cacheStorage.Checkout.GetCheckoutSession(r.Context(), sessionID)
	if err != nil {
		switch {
		case errors.Is(err, store.ErrNotFound):
			app.notFoundResponse(w, r, err)
		default:
			app.internalServerError(w, r, err)
		}
		return
	}

	// Verifikasi session milik user yang benar
	if checkoutSession.UserID != user.ID {
		app.unauthorizedErrorResponse(w, r, errors.New("unauthorized to access this session"))
		return
	}

	// Response dengan data session
	if err := app.jsonResponse(w, http.StatusOK, checkoutSession); err != nil {
		app.internalServerError(w, r, err)
	}
}

// StartCheckout godoc
//
//	@Summary		Start checkout process
//	@Description	Start checkout process and create session
//	@Tags			checkout
//	@Accept			json
//	@Produce		json
//	@Param			payload	body		StartCheckoutPayload	true	"Payload"
//	@Success		200		{object}	store.CheckoutSession
//	@Failure		400		{object}	error
//	@Failure		500		{object}	error
//	@Security		ApiKeyAuth
//	@Router			/checkout/start [post]
func (app *application) startCheckoutHandler(w http.ResponseWriter, r *http.Request) {
	user := getUserFromContext(r)
	var payload StartCheckoutPayload

	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := Validate.RegisterValidation("uuid_array", func(fl validator.FieldLevel) bool {
		ids, ok := fl.Field().Interface().([]uuid.UUID)
		if !ok {
			return false
		}
		return len(ids) > 0 // validasi "required"
	}); err != nil {
		panic(err)
	}

	// Dapatkan cart items dari cart store
	cartStore, err := app.store.Carts.GetCartStoresByID(r.Context(), payload.CartStoreID)
	if err != nil {
		switch {
		case errors.Is(err, store.ErrNotFound):
			app.notFoundResponse(w, r, err)
		default:
			app.internalServerError(w, r, err)
		}
		return
	}

	// Mulai checkout session di Redis
	checkoutSession, err := app.cacheStorage.Checkout.StartCheckoutSession(r.Context(), user.ID, cartStore)
	if err != nil {
		app.badRequestResponse(w, r, fmt.Errorf("failed to start checkout: %w", err))
	}

	if err := app.jsonResponse(w, http.StatusOK, checkoutSession); err != nil {
		app.internalServerError(w, r, err)
	}
}

// CompleteCheckout godoc
//
//	@Summary		Complete checkout process
//	@Description	Complete checkout process and create order
//	@Tags			checkout
//	@Accept			json
//	@Produce		json
//	@Param			payload	body		CompleteCheckoutPayload	true	"Payload"
//	@Success		200		{object}	store.Order
//	@Failure		400		{object}	error
//	@Failure		500		{object}	error
//	@Security		ApiKeyAuth
//	@Router			/checkout/complete [post]
func (app *application) completeCheckoutHandler(w http.ResponseWriter, r *http.Request) {
	user := getUserFromContext(r)
	var payload CompleteCheckoutPayload

	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	// Dapatkan session dari Redis
	checkoutSession, err := app.cacheStorage.Checkout.GetCheckoutSession(r.Context(), payload.SessionID)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	// Verifikasi session milik user yang benar
	if checkoutSession.UserID != user.ID {
		app.unauthorizedErrorResponse(w, r, errors.New("invalid session for user"))
		return
	}

	// Dapatkan shipping address
	shippingAddress, err := app.store.ShippingAddresses.GetByID(r.Context(), uuid.MustParse(payload.ShippingAddressID), user.ID)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	// Dapatkan shipping method
	shippingMethod, err := app.store.Orders.GetShippingMethodByID(r.Context(), payload.ShippingMethodID)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	// Dapatkan payment method
	paymentMethod, err := app.store.Orders.GetPaymentMethodByID(r.Context(), payload.PaymentMethodID)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	// Update checkout session dengan data lengkap
	checkoutSession.ShippingAddress = shippingAddress
	checkoutSession.ShippingMethod = shippingMethod
	checkoutSession.PaymentMethod = paymentMethod

	// Buat order permanen di database
	err = app.store.Checkout.CreateOrderFromCheckout(r.Context(), checkoutSession)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	// Hapus session dari Redis
	err = app.cacheStorage.Checkout.CompleteCheckout(r.Context(), payload.SessionID)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	// Response dengan order yang dibuat
	// Anda mungkin perlu menyesuaikan response sesuai kebutuhan
	if err := app.jsonResponse(w, http.StatusOK, map[string]string{"status": "success"}); err != nil {
		app.internalServerError(w, r, err)
	}
}


