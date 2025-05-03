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

type CreateProductPayload struct {
	Name          string   `json:"name" validate:"required"`
	Slug          string   `json:"slug" validate:"required,max=100"`
	Description   string   `json:"description,omitempty" validate:"required,max=100"`
	Price         float64  `json:"price" validate:"required"`
	DiscountPrice float64  `json:"discount_price" validate:"omitempty"`
	Discount      float64  `json:"discount" validate:"omitempty"`
	Rating        float64  `json:"rating" validate:"required"`
	Estimation    string   `json:"estimation" validate:"required"`
	Stock         int      `json:"stock" validate:"required"`
	Sold          int      `json:"sold" validate:"required"`
	IsForSale     bool     `json:"is_for_sale" validate:"omitempty"`
	IsApproved    bool     `json:"is_approved" validate:"required"`
	ImageUrls     []string `json:"image_urls" validate:"required"`
	TokoID        int64    `json:"toko_id" validate:"required"`
	CategoryID    int64    `json:"category_id" validate:"required"`
}

// CreateProduct gdoc
//
//	@Summary		create product
//	@Description	create product
//	@Tags			product
//	@Accept			json
//	@Produce		json
//	@Param			payload	body		CreateProductPayload	true	"product creation payload"
//	@Success		201		{object}	store.Product
//	@Failure		400		{object}	error
//	@Failure		401		{object}	error
//	@Failure		500		{object}	error
//	@Security		ApiKeyAuth
//	@Router			/product [post]
func (app *application) createProductHandler(w http.ResponseWriter, r *http.Request) {
	var payload CreateProductPayload

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
		Category: &store.Category{
			ID: payload.CategoryID,
		},
		Toko: &store.Toko{
			ID: payload.TokoID,
		},
	}

	if err := app.store.Products.Create(r.Context(), p); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusCreated, p); err != nil {
		app.internalServerError(w, r, err)
	}
}

type UpdateProductPayload struct {
	Name          *string   `json:"name" validate:"omitempty"`
	Slug          *string   `json:"slug" validate:"omitempty,max=100"`
	Description   *string   `json:"description" validate:"omitempty,max=100"`
	Price         *float64  `json:"price" validate:"omitempty"`
	DiscountPrice *float64  `json:"discount_price" validate:"omitempty"`
	Discount      *float64  `json:"discount" validate:"omitempty"`
	Rating        *float64  `json:"rating" validate:"omitempty"`
	Estimation    *string   `json:"estimation" validate:"omitempty"`
	Stock         *int      `json:"stock" validate:"omitempty"`
	Sold          *int      `json:"sold" validate:"omitempty"`
	IsForSale     *bool     `json:"is_for_sale" validate:"omitempty"`
	IsApproved    *bool     `json:"is_approved" validate:"omitempty"`
	ImageUrls     *[]string `json:"image_urls" validate:"omitempty"`
	Version       *int      `json:"version"`
}

// UpdateProduct godoc
//
//	@Summary		Updates a product
//	@Description	Updates a product by ID
//	@Tags			product
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int						true	"product ID"
//	@Param			payload	body		UpdateProductPayload	true	"product payload"
//	@Success		200		{object}	store.Product
//	@Failure		400		{object}	error
//	@Failure		401		{object}	error
//	@Failure		404		{object}	error
//	@Failure		500		{object}	error
//	@Security		ApiKeyAuth
//	@Router			/product/{id} [patch]
func (app *application) updateProductHandler(w http.ResponseWriter, r *http.Request) {
	product := getProductFromContext(r)

	var payload UpdateProductPayload

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

	if payload.Name != nil {
		product.Name = *payload.Name
	}
	if payload.Slug != nil {
		product.Slug = *payload.Slug
	}
	if payload.Description != nil {
		product.Description = *payload.Description
	}
	if payload.Price != nil {
		product.Price = *payload.Price
	}
	if payload.DiscountPrice != nil {
		product.DiscountPrice = *payload.DiscountPrice
	}
	if payload.Discount != nil {
		product.Discount = *payload.Discount
	}
	if payload.Rating != nil {
		product.Rating = *payload.Rating
	}
	if payload.Estimation != nil {
		product.Estimation = *payload.Estimation
	}
	if payload.Stock != nil {
		product.Stock = *payload.Stock
	}
	if payload.Sold != nil {
		product.Sold = *payload.Sold
	}
	if payload.IsForSale != nil {
		product.IsForSale = *payload.IsForSale
	}
	if payload.IsApproved != nil {
		product.IsApproved = *payload.IsApproved
	}
	if payload.ImageUrls != nil {
		product.ImageUrls = *payload.ImageUrls
	}

	if err := app.updateProduct(r.Context(), product); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, product); err != nil {
		app.internalServerError(w, r, err)
	}
}

// DeleteProduct godoc
//
//	@Summary		Deletes a product
//	@Description	Delete a product by ID
//	@Tags			product
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"product ID"
//	@Success		204	{object}	string
//	@Failure		404	{object}	error
//	@Failure		500	{object}	error
//	@Security		ApiKeyAuth
//	@Router			/product/{id} [delete]
func (app *application) deleteProductHandler(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "productID")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	ctx := r.Context()
	if err := app.store.Products.Delete(ctx, id); err != nil {
		switch {
		case errors.Is(err, store.ErrNotFound):
			app.notFoundResponse(w, r, err)
		default:
			app.internalServerError(w, r, err)

		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
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
				app.notFoundResponse(w, r, err)
			default:
				app.internalServerError(w, r, err)
			}
			return
		}

		toko, _ := app.store.Tokos.GetByID(ctx, product.Toko.ID)

		product.Toko = toko

		ctx = context.WithValue(ctx, productCtx, product)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getProductFromContext(r *http.Request) *store.Product {
	return r.Context().Value(productCtx).(*store.Product)
}

func (app *application) updateProduct(ctx context.Context, product *store.Product) error {
	if err := app.store.Products.Update(ctx, product); err != nil {
		return err
	}

	app.cacheStorage.Users.Delete(ctx, product.ID)
	return nil
}
