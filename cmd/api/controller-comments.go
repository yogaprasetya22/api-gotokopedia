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
	UserID  int64  `json:"user_id" validate:"required"`
}

// CreateComment godoc
//
//	@Summary		Create Comment
//	@Description	Create a new comment for a product
//	@Tags			comment
//	@Accept			json
//	@Produce		json
//	@Param			productID	path		int						true	"Product ID"
//	@Param			payload		body		CreateCommentsPayload	true	"Comment creation payload"
//	@Success		201			{object}	store.Comment
//	@Failure		400			{object}	error
//	@Failure		401			{object}	error
//	@Failure		500			{object}	error
//	@Security		ApiKeyAuth
//	@Router			/catalogue/{productID}/comment [post]
func (app *application) createCommentHandler(w http.ResponseWriter, r *http.Request) {
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
		UserID:    payload.UserID,
		ProductID: productID,
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
	UserID  *int64  `json:"user_id" validate:"omitempty"`
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
//	@Failure		404			{object}	error
//	@Failure		500			{object}	error
//	@Security		ApiKeyAuth
//	@Router			/catalogue/{productID}/comment/{id} [patch]
func (app *application) updateCommentHandler(w http.ResponseWriter, r *http.Request) {
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
	if payload.UserID != nil {
		comment.UserID = *payload.UserID
	}

	// Always use product ID from context
	comment.ProductID = product.ID

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
//	@Router			/catalogue/{productID}/comment/{id} [delete]
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
			app.notFoundError(w, r, err)
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
				app.notFoundError(w, r, err)
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
