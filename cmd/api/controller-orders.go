package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/yogaprasetya22/api-gotokopedia/internal/store"
)

// Struct untuk request membuat order
type CreateOrderRequest struct {
	CartID              uuid.UUID `json:"cart_id"`
	PaymentMethodID     int64     `json:"payment_method_id"`
	ShippingMethodID    int64     `json:"shipping_method_id"`
	ShippingAddressesID uuid.UUID `json:"shipping_addresses_id"`
	Notes               string    `json:"notes,omitempty"`
}

// Struct untuk request update status order
type updateOrderStatusRequest struct {
	StatusID int64  `json:"status_id"`
	Notes    string `json:"notes,omitempty"`
}


// getOrderHandler godoc
//
//	@Summary		Get order by ID
//	@Description	Get order details by ID
//	@Tags			order
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Order ID"
//	@Success		200	{object}	store.Order
//	@Failure		400	{object}	error
//	@Failure		404	{object}	error
//	@Failure		403	{object}	error
//	@Failure		500	{object}	error
//	@Router			/order/{id} [get]
//	@Security		ApiKeyAuth
func (app *application) getOrderHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		app.badRequestResponse(w, r, errors.New("invalid order ID"))
		return
	}

	user := getUserFromContext(r)

	order, err := app.store.Orders.GetByID(r.Context(), id)

	// Pastikan order milik user yang terautentikasi
	if order.UserID != user.ID {
		app.forbiddenResponse(w, r)
		return
	}

	if err != nil {
		switch {
		case errors.Is(err, store.ErrNotFound):
			app.notFoundResponse(w, r, err)
		default:
			app.internalServerError(w, r, err)
		}
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, order); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

// listOrdersHandler godoc
//
//	@Summary		List orders for authenticated user
//	@Description	List orders with pagination and sorting
//	@Tags			order
//	@Accept			json
//	@Produce		json
//	@Param			limit	query		int		false	"Number of orders to return (default 5)"
//	@Param			offset	query		int		false	"Offset for pagination (default 0)"
//	@Param			sort	query		string	false	"Sort order (asc or desc, default desc)"
//	@Param			search	query		string	false	"Search term for order details"
//	@Success		200		{object}	[]store.Order
//	@Failure		400		{object}	error
//	@Failure		404		{object}	error
//	@Failure		500		{object}	error
//	@Router			/order [get]
//	@Security		ApiKeyAuth
func (app *application) listOrdersHandler(w http.ResponseWriter, r *http.Request) {
	fq := store.PaginatedFeedQuery{
		Limit:  5,
		Offset: 0,
		Sort:   "desc",
		Search: "",
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

	orders, err := app.store.Orders.GetByUserID(r.Context(), user.ID, fq)
	if err != nil {
		switch {
		case errors.Is(err, store.ErrNotFound):
			app.notFoundResponse(w, r, err)
		default:
			app.internalServerError(w, r, err)
		}
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, orders); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

// Handler untuk update status order
// func (app *application) updateOrderStatusHandler(w http.ResponseWriter, r *http.Request) {
// 	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
// 	if err != nil || id < 1 {
// 		app.notFoundResponse(w, r, err)
// 		return
// 	}

// 	user := getUserFromContext(r)

// 	// Cek apakah order ada dan milik user
// 	order, err := app.store.Orders.GetByID(r.Context(), id)
// 	if err != nil {
// 		switch {
// 		case errors.Is(err, store.ErrNotFound):
// 			app.notFoundResponse(w, r, err)
// 		default:
// 			app.internalServerError(w, r, err)
// 		}
// 		return
// 	}

// 	// Jika bukan admin, pastikan order milik user yang terautentikasi
// 	isAdmin := app.isAdmin(r)
// 	if !isAdmin && order.UserID != userID {
// 		app.notPermittedResponse(w, r)
// 		return
// 	}

// 	var req updateOrderStatusRequest
// 	err = app.readJSON(w, r, &req)
// 	if err != nil {
// 		app.badRequestResponse(w, r, err)
// 		return
// 	}

// 	// Validasi status ID
// 	v := validator.New()
// 	v.Check(req.StatusID > 0, "status_id", "tidak valid")
// 	if !v.Valid() {
// 		app.failedValidationResponse(w, r, v.Errors)
// 		return
// 	}

// 	// Update status order
// 	err = app.store.Orders.UpdateStatus(r.Context(), id, req.StatusID, req.Notes)
// 	if err != nil {
// 		app.serverErrorResponse(w, r, err)
// 		return
// 	}

// 	// Ambil data order yang sudah diupdate
// 	updatedOrder, err := app.store.Orders.GetByID(r.Context(), id)
// 	if err != nil {
// 		app.serverErrorResponse(w, r, err)
// 		return
// 	}

// 	app.writeJSON(w, http.StatusOK, envelope{"order": updatedOrder}, nil)
// }

// // Handler untuk mendapatkan daftar metode pengiriman
// func (app *application) listShippingMethodsHandler(w http.ResponseWriter, r *http.Request) {
// 	methods, err := app.store.Orders.GetShippingMethods(r.Context())
// 	if err != nil {
// 		app.serverErrorResponse(w, r, err)
// 		return
// 	}

// 	app.writeJSON(w, http.StatusOK, envelope{"shipping_methods": methods}, nil)
// }

// // Handler untuk mendapatkan daftar metode pembayaran
// func (app *application) listPaymentMethodsHandler(w http.ResponseWriter, r *http.Request) {
// 	methods, err := app.store.Orders.GetPaymentMethods(r.Context())
// 	if err != nil {
// 		app.serverErrorResponse(w, r, err)
// 		return
// 	}

// 	app.writeJSON(w, http.StatusOK, envelope{"payment_methods": methods}, nil)
// }
