package main

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/yogaprasetya22/api-gotokopedia/internal/store"
)

type productKey string

const productCtx productKey = "product"

type CreateProductRequest struct {
	Name          string   `json:"name" validate:"required"`
	Slug          string   `json:"slug" validate:"required,max=100"`
	Description   string   `json:"description,omitempty" validate:"required,max=100"`
	Price         float64  `json:"price" validate:"required"`
	DiscountPrice float64  `json:"discount_price" validate:"required"`
	Discount      float64  `json:"discount" validate:"required"`
	Rating        float64  `json:"rating" validate:"required"`
	Estimation    string   `json:"estimation" validate:"required"`
	Stock         int      `json:"stock" validate:"required"`
	Sold          int      `json:"sold" validate:"required"`
	IsForSale     bool     `json:"is_for_sale" validate:"required"`
	IsApproved    bool     `json:"is_approved" validate:"required"`
	ImageUrls     []string `json:"image_urls" validate:"required"`
	CategoryID    int64    `json:"category_id,omitempty" `
	TokoID        int64    `json:"toko_id,omitempty" `
}

func (app *application) createProductHandler(w http.ResponseWriter, r *http.Request) {
	var payload CreateProductRequest

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

	p := &store.Product{
		Name:          payload.Name,
		Slug:          payload.Slug,
		Description:   payload.Description,
		Price:         payload.Price,
		DiscountPrice: payload.DiscountPrice,
		Discount:      payload.Discount,
		Rating:        payload.Rating,
		Estimation:    payload.Estimation,
		Stock:         payload.Stock,
		Sold:          payload.Sold,
		IsForSale:     payload.IsForSale,
		IsApproved:    payload.IsApproved,
		ImageUrls:     payload.ImageUrls,
	}

	if err := app.store.Products.Create(r.Context(), p); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusCreated, p); err != nil {
		app.internalServerError(w, r, err)
	}
}

// getProductHandler
func (app *application) getProductHandler(w http.ResponseWriter, r *http.Request) {
	product := getPostFromContext(r)

	ctx := r.Context()

	toko, _ := app.store.Tokos.GetByID(ctx, product.Toko.ID)
	category, _ := app.store.Categoris.GetByID(ctx, product.Category.ID)

	product.Toko = *toko
	product.Category = *category

if err := app.jsonResponse(w, http.StatusOK, product); err != nil {
		app.internalServerError(w, r, err)
	}
}

func (app *application) productContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		idParam := chi.URLParam(r, "productID")
		id, err := strconv.ParseInt(idParam, 10, 64)
		if err != nil {
			app.badRequestResponse(w, r, err)
			return
		}

		ctx := r.Context()
		product, err := app.store.Products.GetByID(ctx, id)
		if err != nil {
			switch {
			case errors.Is(err, store.ErrNotFound):
				app.notFoundError(w, r, err)
			default:
				app.internalServerError(w, r, err)
			}
			return
		}

		ctx = context.WithValue(ctx, productCtx, product)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getPostFromContext(r *http.Request) *store.Product {
	return r.Context().Value(productCtx).(*store.Product)
}
