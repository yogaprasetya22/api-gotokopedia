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
type createOrderRequest struct {
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

// Handler untuk membuat order dari cart
func (app *application) createOrderHandler(w http.ResponseWriter, r *http.Request) {
	var payload createOrderRequest

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

	user := getUserFromContext(r)

	err = app.store.Orders.CreateFromCart(r.Context(), payload.CartID, user.ID, payload.PaymentMethodID, payload.ShippingMethodID, payload.ShippingAddressesID, payload.Notes)
	if err != nil {
		switch {
		case errors.Is(err, store.ErrNotFound):
			app.notFoundResponse(w, r, err)
		case errors.Is(err, store.ErrConflict):
			app.conflictResponse(w, r, err)
		default:
			app.internalServerError(w, r, err)
		}
		return
	}

	if err := app.jsonResponse(w, http.StatusCreated, map[string]any{"status": "order created"}); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

// Handler untuk mengambil detail order
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

// Handler untuk mengambil daftar order pengguna
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
