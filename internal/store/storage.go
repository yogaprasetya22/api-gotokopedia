package store

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

var (
	ErrAvtivated         = errors.New("akun belum diaktivasi")
	ErrNotFound          = errors.New("sumber daya tidak ditemukan")
	ErrConflict          = errors.New("sumber daya sudah ada")
	QueryTimeoutDuration = time.Second * 5
)

type Storage struct {
	Products interface {
		GetByID(context.Context, int64) (*Product, error)
		Create(context.Context, *Product) error
		Update(context.Context, *Product) error
		Delete(context.Context, int64) error
	}
	Categoris interface {
		GetByID(context.Context, int64) (*Category, error)
		GetBySlug(context.Context, string) (*Category, error)
		Create(context.Context, *Category) error
	}
	Tokos interface {
		GetByID(context.Context, int64) (*Toko, error)
		GetBySlug(context.Context, string) (*Toko, error)
		Create(context.Context, *Toko) error
	}
	Users interface {
		GetByID(context.Context, int64) (*User, error)
		Create(context.Context, *User) error
	}
	Comments interface {
		GetByProductID(context.Context, int64) ([]Comment, error)
		Create(context.Context, *Comment) error
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Products:  &ProductStore{db},
		Categoris: &CategoryStore{db},
		Tokos:     &TokoStore{db},
		Users:     &UserStore{db},
		Comments:  &CommentStore{db},
	}
}
