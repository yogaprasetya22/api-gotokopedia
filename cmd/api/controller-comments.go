package main

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/yogaprasetya22/api-gotokopedia/internal/store"
)

type commentKey string

const CommentCtx commentKey = "comment"

type CreateCommentsPayload struct {
	Content string `json:"content" validate:"required,max=255"`
	Rating  int    `json:"rating" validate:"required,min=1,max=5"`
}

// GetComments godoc
//
//	@Summary		Get Comments
//	@Description	Get all comments for a product
//	@Tags			comment
//	@Accept			json
//	@Produce		json
//	@Param			slug	path		string	true	"Product Slug"
//	@Param			limit	query		int		false	"limit"
//	@Param			offset	query		int		false	"offset"
//	@Param			sort	query		string	false	"sort"
//	@Success		200		{array}		store.Comment
//	@Failure		404		{object}	error
//	@Failure		500		{object}	error
//	@Router			/comment/{slug} [get]
func (app *application) getCommentsHandler(w http.ResponseWriter, r *http.Request) {
	slugProduct := chi.URLParam(r, "slug")

	fq := store.PaginatedFeedQuery{
		Limit:  5,
		Offset: 0,
		Sort:   "desc",
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

	ctx := r.Context()

	comments, err := app.store.Comments.GetComments(ctx, slugProduct, fq)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, comments); err != nil {
		app.internalServerError(w, r, err)
	}

}

// CreateComment godoc
//
//	@Summary		Create Comment
//	@Description	Create a new comment for a product
//	@Tags			comment
//	@Accept			json
//	@Produce		json
//	@Param			productID	path		int						true	"Product ID"
//	@Param			payload		body		CreateCommentsPayload	true	"Comment payload"
//	@Success		201			{object}	store.Comment
//	@Failure		400			{object}	error
//	@Failure		401			{object}	error
//	@Failure		500			{object}	error
//	@Security		ApiKeyAuth
//	@Router			/product/{productID}/comment [post]
func (app *application) createCommentHandler(w http.ResponseWriter, r *http.Request) {
	user := getUserFromContext(r)
	productIDParam := chi.URLParam(r, "productID")
	productID, err := strconv.ParseInt(productIDParam, 10, 64)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	var payload CreateCommentsPayload
	err = readJSON(w, r, &payload)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	err = Validate.Struct(payload)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	c := &store.Comment{
		Content:   payload.Content,
		UserID:    user.ID,
		ProductID: productID,
		Rating:    payload.Rating,
	}

	if err := app.store.Comments.Create(r.Context(), c); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusCreated, c); err != nil {
		app.internalServerError(w, r, err)
	}
}

type UpdateCommentsPayload struct {
	Content *string `json:"content" validate:"omitempty,max=255"`
	Rating  *int    `json:"rating" validate:"omitempty,min=1,max=5"`
}

// UpdateComment godoc
//
//	@Summary		Updates a Comment
//	@Description	Updates a Comment by ID
//	@Tags			comment
//	@Accept			json
//	@Produce		json
//	@Param			productID	path		int						true	"Product ID"
//	@Param			id			path		int						true	"Comment ID"
//	@Param			payload		body		UpdateCommentsPayload	true	"Comment payload"
//	@Success		200			{object}	store.Comment
//	@Failure		400			{object}	error
//	@Failure		401			{object}	error
//	@Failure		403			{object}	error
//	@Failure		404			{object}	error
//	@Failure		500			{object}	error
//	@Security		ApiKeyAuth
//	@Router			/product/{productID}/comment/{id} [patch]
func (app *application) updateCommentHandler(w http.ResponseWriter, r *http.Request) {
	user := getUserFromContext(r)
	comment := getCommentFromContext(r)
	product := getProductFromContext(r)

	var payload UpdateCommentsPayload

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

	if payload.Content != nil {
		comment.Content = *payload.Content
	}
	comment.ProductID = product.ID
	comment.UserID = user.ID

	if err := app.store.Comments.Update(r.Context(), comment); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, comment); err != nil {
		app.internalServerError(w, r, err)
	}
}

// DeleteComment godoc
//
//	@Summary		Deletes a Comment
//	@Description	Delete a Comment by ID
//	@Tags			comment
//	@Accept			json
//	@Produce		json
//	@Param			productID	path		int	true	"Product ID"
//	@Param			id			path		int	true	"Comment ID"
//	@Success		204			{object}	string
//	@Failure		404			{object}	error
//	@Failure		500			{object}	error
//	@Security		ApiKeyAuth
//	@Router			/product/{productID}/comment/{id} [delete]
func (app *application) deleteCommentHandler(w http.ResponseWriter, r *http.Request) {
	product := getProductFromContext(r)
	idParam := chi.URLParam(r, "commentID")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	ctx := r.Context()
	if err := app.store.Comments.Delete(ctx, id, product.ID); err != nil {
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

func (app *application) commentContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		idParam := chi.URLParam(r, "commentID")
		id, err := strconv.ParseInt(idParam, 10, 64)
		if err != nil {
			app.badRequestResponse(w, r, err)
			return
		}

		ctx := r.Context()
		c, err := app.store.Comments.GetByID(ctx, id)
		if err != nil {
			switch {
			case errors.Is(err, store.ErrNotFound):
				app.notFoundResponse(w, r, err)
			default:
				app.internalServerError(w, r, err)
			}
			return
		}

		ctx = context.WithValue(ctx, CommentCtx, c)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getCommentFromContext(r *http.Request) *store.Comment {
	return r.Context().Value(CommentCtx).(*store.Comment)
}
