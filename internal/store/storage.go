package store

import (
	"context"
	"database/sql"
)

type Storage struct {
	Products interface {
		Create(context.Context, *Product) error
		GetByID(context.Context, string) (*Product, error)
	}
	Categoris interface {
		Create(context.Context, *Category) error
	}
	Tokos interface {
		Create(context.Context, *Toko) error
	}
	Users interface {
		Create(*User) error
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Products:  &ProductStore{db},
		Categoris: &CategoryStore{db},
		Tokos:     &TokoStore{db},
		Users:     &UserStore{db},
	}
}
