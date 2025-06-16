package main

import (
    "encoding/json"
    "errors"
    "net/http"
    "strconv"

    "github.com/go-chi/chi/v5"
    "github.com/yogaprasetya22/api-gotokopedia/internal/store"
)

// listShippingMethodsHandler godoc
//	@Summary		List shipping methods
//	@Description	Get all available shipping methods
//	@Tags			shipping-method
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		store.ShippingMethod
//	@Failure		500	{object}	error
//	@Router			/shipping-methods [get]
func (app *application) listShippingMethodsHandler(w http.ResponseWriter, r *http.Request) {
    methods, err := app.store.ShippingMethods.GetAll(r.Context())
    if err != nil {
        app.internalServerError(w, r, err)
        return
    }
    app.jsonResponse(w, http.StatusOK, methods)
}

// getShippingMethodHandler godoc
//	@Summary		Get shipping method
//	@Description	Get shipping method by ID
//	@Tags			shipping-method
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Shipping Method ID"
//	@Success		200	{object}	store.ShippingMethod
//	@Failure		400	{object}	error
//	@Failure		404	{object}	error
//	@Failure		500	{object}	error
//	@Router			/shipping-methods/{id} [get]
func (app *application) getShippingMethodHandler(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
    if err != nil {
        app.badRequestResponse(w, r, errors.New("invalid id"))
        return
    }
    method, err := app.store.ShippingMethods.GetByID(r.Context(), id)
    if err != nil {
        app.notFoundResponse(w, r, err)
        return
    }
    app.jsonResponse(w, http.StatusOK, method)
}

// createShippingMethodHandler godoc
//	@Summary		Create shipping method
//	@Description	Create a new shipping method
//	@Tags			shipping-method
//	@Accept			json
//	@Produce		json
//	@Param			shipping_method	body		store.ShippingMethod	true	"Shipping Method"
//	@Success		201				{object}	store.ShippingMethod
//	@Failure		400				{object}	error
//	@Failure		500				{object}	error
//	@Router			/shipping-methods [post]
func (app *application) createShippingMethodHandler(w http.ResponseWriter, r *http.Request) {
    var payload store.ShippingMethod
    if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
        app.badRequestResponse(w, r, err)
        return
    }
    if payload.Name == "" {
        app.badRequestResponse(w, r, errors.New("name is required"))
        return
    }
    if err := app.store.ShippingMethods.Create(r.Context(), &payload); err != nil {
        app.internalServerError(w, r, err)
        return
    }
    app.jsonResponse(w, http.StatusCreated, payload)
}

// updateShippingMethodHandler godoc
//	@Summary		Update shipping method
//	@Description	Update an existing shipping method
//	@Tags			shipping-method
//	@Accept			json
//	@Produce		json
//	@Param			id				path		int						true	"Shipping Method ID"
//	@Param			shipping_method	body		store.ShippingMethod	true	"Shipping Method"
//	@Success		200				{object}	store.ShippingMethod
//	@Failure		400				{object}	error
//	@Failure		500				{object}	error
//	@Router			/shipping-methods/{id} [put]
func (app *application) updateShippingMethodHandler(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
    if err != nil {
        app.badRequestResponse(w, r, errors.New("invalid id"))
        return
    }
    var payload store.ShippingMethod
    if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
        app.badRequestResponse(w, r, err)
        return
    }
    payload.ID = id
    if err := app.store.ShippingMethods.Update(r.Context(), &payload); err != nil {
        app.internalServerError(w, r, err)
        return
    }
    app.jsonResponse(w, http.StatusOK, payload)
}

// deleteShippingMethodHandler godoc
//	@Summary		Delete shipping method
//	@Description	Delete a shipping method by ID
//	@Tags			shipping-method
//	@Accept			json
//	@Produce		json
//	@Param			id	path	int	true	"Shipping Method ID"
//	@Success		204
//	@Failure		400	{object}	error
//	@Failure		404	{object}	error
//	@Failure		500	{object}	error
//	@Router			/shipping-methods/{id} [delete]
func (app *application) deleteShippingMethodHandler(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
    if err != nil {
        app.badRequestResponse(w, r, errors.New("invalid id"))
        return
    }
    if err := app.store.ShippingMethods.Delete(r.Context(), id); err != nil {
        app.notFoundResponse(w, r, err)
        return
    }
    w.WriteHeader(http.StatusNoContent)
}