package main

import (
	"net/http"
	"time"

	"github.com/yogaprasetya22/api-gotokopedia/internal/store"
)

type postKey string

const postCtx postKey = "product"

type CreateProductRequest struct {
	Name          string    `json:"name" validate:"required"`
	Slug          string    `json:"slug" validate:"required,max=100"`
	Description   string    `json:"description,omitempty" validate:"required,max=100"`
	Price         float64   `json:"price" validate:"required"`
	DiscountPrice float64   `json:"discount_price" validate:"required"`
	Discount      float64   `json:"discount" validate:"required"`
	Rating        float64   `json:"rating" validate:"required"`
	Estimation    string    `json:"estimation" validate:"required"`
	Stock         int       `json:"stock" validate:"required"`
	Sold          int       `json:"sold" validate:"required"`
	IsForSale     bool      `json:"is_for_sale" validate:"required"`
	IsApproved    bool      `json:"is_approved" validate:"required"`
	CreatedAt     time.Time `json:"created_at" validate:"required"`
	UpdatedAt     time.Time `json:"updated_at" validate:"required"`
	ImageUrls     []string  `json:"image_urls" validate:"required"`
	CategoryID    *string   `json:"category_id,omitempty" validate:"required"`
	TokoID        *string   `json:"toko_id,omitempty" validate:"required"`
}

func (app *application) createProductHandler(w http.ResponseWriter, r *http.Request) {
	var input CreateProductRequest

	err := readJSON(w, r, &input)
	if err != nil {
		app.jsonResponse(w, http.StatusBadRequest, err)
		return
	}

	err = Validate.Struct(input)
	if err != nil {
		app.jsonResponse(w, http.StatusUnprocessableEntity, err)
		return
	}

	p := &store.Product{
		Name:          input.Name,
		Slug:          input.Slug,
		Description:   input.Description,
		Price:         input.Price,
		DiscountPrice: input.DiscountPrice,
		Discount:      input.Discount,
		Rating:        input.Rating,
		Estimation:    input.Estimation,
		Stock:         input.Stock,
		Sold:          input.Sold,
		IsForSale:     input.IsForSale,
		IsApproved:    input.IsApproved,
		CreatedAt:     input.CreatedAt,
		UpdatedAt:     input.UpdatedAt,
		ImageUrls:     input.ImageUrls,
		CategoryID:    input.CategoryID,
		TokoID:        input.TokoID,
	}

	err = app.store.Products.Create(r.Context(), p)
	if err != nil {
		app.jsonResponse(w, http.StatusInternalServerError, err)
		return
	}

	app.jsonResponse(w, http.StatusCreated, p)
}
