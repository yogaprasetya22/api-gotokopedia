package main

import (
    "encoding/json"
    "errors"
    "net/http"
    "strconv"

    "github.com/go-chi/chi/v5"
    "github.com/yogaprasetya22/api-gotokopedia/internal/store"
)

// listPaymentMethodsHandler godoc
//	@Summary		List payment methods
//	@Description	Get all available payment methods
//	@Tags			payment-method
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		store.PaymentMethod
//	@Failure		500	{object}	error
//	@Router			/payment-methods [get]
func (app *application) listPaymentMethodsHandler(w http.ResponseWriter, r *http.Request) {
    methods, err := app.store.PaymentMethods.GetAll(r.Context())
    if err != nil {
        app.internalServerError(w, r, err)
        return
    }
    app.jsonResponse(w, http.StatusOK, methods)
}

// getPaymentMethodHandler godoc
//	@Summary		Get payment method
//	@Description	Get payment method by ID
//	@Tags			payment-method
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Payment Method ID"
//	@Success		200	{object}	store.PaymentMethod
//	@Failure		400	{object}	error
//	@Failure		404	{object}	error
//	@Failure		500	{object}	error
//	@Router			/payment-methods/{id} [get]
func (app *application) getPaymentMethodHandler(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
    if err != nil {
        app.badRequestResponse(w, r, errors.New("invalid id"))
        return
    }
    method, err := app.store.PaymentMethods.GetByID(r.Context(), id)
    if err != nil {
        app.notFoundResponse(w, r, err)
        return
    }
    app.jsonResponse(w, http.StatusOK, method)
}

// createPaymentMethodHandler godoc
//	@Summary		Create payment method
//	@Description	Create a new payment method
//	@Tags			payment-method
//	@Accept			json
//	@Produce		json
//	@Param			payment_method	body		store.PaymentMethod	true	"Payment Method"
//	@Success		201				{object}	store.PaymentMethod
//	@Failure		400				{object}	error
//	@Failure		500				{object}	error
//	@Router			/payment-methods [post]
func (app *application) createPaymentMethodHandler(w http.ResponseWriter, r *http.Request) {
    var payload store.PaymentMethod
    if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
        app.badRequestResponse(w, r, err)
        return
    }
    if payload.Name == "" {
        app.badRequestResponse(w, r, errors.New("name is required"))
        return
    }
    if err := app.store.PaymentMethods.Create(r.Context(), &payload); err != nil {
        app.internalServerError(w, r, err)
        return
    }
    app.jsonResponse(w, http.StatusCreated, payload)
}

// updatePaymentMethodHandler godoc
//	@Summary		Update payment method
//	@Description	Update an existing payment method
//	@Tags			payment-method
//	@Accept			json
//	@Produce		json
//	@Param			id				path		int					true	"Payment Method ID"
//	@Param			payment_method	body		store.PaymentMethod	true	"Payment Method"
//	@Success		200				{object}	store.PaymentMethod
//	@Failure		400				{object}	error
//	@Failure		500				{object}	error
//	@Router			/payment-methods/{id} [put]
func (app *application) updatePaymentMethodHandler(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
    if err != nil {
        app.badRequestResponse(w, r, errors.New("invalid id"))
        return
    }
    var payload store.PaymentMethod
    if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
        app.badRequestResponse(w, r, err)
        return
    }
    payload.ID = id
    if err := app.store.PaymentMethods.Update(r.Context(), &payload); err != nil {
        app.internalServerError(w, r, err)
        return
    }
    app.jsonResponse(w, http.StatusOK, payload)
}

// deletePaymentMethodHandler godoc
//	@Summary		Delete payment method
//	@Description	Delete a payment method by ID
//	@Tags			payment-method
//	@Accept			json
//	@Produce		json
//	@Param			id	path	int	true	"Payment Method ID"
//	@Success		204
//	@Failure		400	{object}	error
//	@Failure		404	{object}	error
//	@Failure		500	{object}	error
//	@Router			/payment-methods/{id} [delete]
func (app *application) deletePaymentMethodHandler(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
    if err != nil {
        app.badRequestResponse(w, r, errors.New("invalid id"))
        return
    }
    if err := app.store.PaymentMethods.Delete(r.Context(), id); err != nil {
        app.notFoundResponse(w, r, err)
        return
    }
    w.WriteHeader(http.StatusNoContent)
}