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
		GetProduct(ctx context.Context, slug_toko, slug_product string) (*DetailProduct, error)
		GetProductByTokos(ctx context.Context, slug_toko string, query PaginatedFeedQuery) ([]*Product, error)
		GetByID(context.Context, int64) (*Product, error)
		GetByTokoID(context.Context, int64) ([]*Product, error)
		GetProductFeed(context.Context, []int64, PaginatedFeedQuery) ([]*Product, error)
		GetProductCategoryFeed(context.Context, PaginatedFeedQuery) ([]*Product, error)
		GetAllProduct(context.Context, PaginatedFeedQuery) ([]*Product, error)
		Create(context.Context, *Product) error
		Update(context.Context, *Product) error
		Delete(context.Context, int64) error
	}
	Categoris interface {
		GetAll(context.Context) ([]*Category, error)
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
		GetByGoogleID(context.Context, string) (*User, error)
		GetByID(context.Context, int64) (*User, error)
		GetByEmail(context.Context, string) (*User, error)
		Create(context.Context, *sql.Tx, *User) error
		CreateByOAuth(context.Context, *User) error
		CreateAndInvite(ctx context.Context, user *User, token string, exp time.Duration) error
		Activate(context.Context, string) error
		Delete(context.Context, int64) error
	}
	Follow interface {
		Follow(ctx context.Context, followerID, userID int64) error
		Unfollow(ctx context.Context, followerID, userID int64) error
	}
	Comments interface {
		GetComments(context.Context, string, PaginatedFeedQuery) (MetaCommentPaginated, error)
		GetByProductID(context.Context, int64) ([]Comment, error)
		GetByID(context.Context, int64) (*Comment, error)
		Create(context.Context, *Comment) error
		Update(context.Context, *Comment) error
		Delete(ctx context.Context, commentID, productID int64) error
	}
	Roles interface {
		GetByName(context.Context, string) (*Role, error)
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Products:  &ProductStore{db},
		Categoris: &CategoryStore{db},
		Tokos:     &TokoStore{db},
		Users:     &UserStore{db},
		Follow:    &FollowerStore{db},
		Comments:  &CommentStore{db},
	}
}

func withTx(db *sql.DB, ctx context.Context, f func(tx *sql.Tx) error) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	if err := f(tx); err != nil {
		_ = tx.Rollback()
		return err
	}

	return tx.Commit()
}
