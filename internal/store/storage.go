package store

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
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
		CreateWithActive(ctx context.Context, tx *sql.Tx, user *User) (*User, error)
		Create(context.Context, *sql.Tx, *User) (*User, error)
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
	Carts interface {
		GetCartByUserID(ctx context.Context, userID int64, query PaginatedFeedQuery) (*Cart, error)
		AddToCartTransaction(ctx context.Context, userID, productID, quantity int64) (*Cart, error)
		AddingQuantityCartStoreItemTransaction(ctx context.Context, cartStoreItemID uuid.UUID, userID int64) error
		RemovingQuantityCartStoreItemTransaction(ctx context.Context, cartStoreItemID uuid.UUID, userID int64) error
		GetDetailCartByCartStoreID(ctx context.Context, cartStoreID uuid.UUID, userID int64) (*CartDetailResponse, error)
	}
	Orders interface {
		CreateFromCart(ctx context.Context, cartStoreID uuid.UUID, userID, paymentMethodID, shippingMethodID int64, shippingAddress, notes string) error
		UpdateStatus(ctx context.Context, orderID, statusID int64, notes string) error
		GetByID(ctx context.Context, id int64) (*Order, error)
		GetByUserID(ctx context.Context, userID int64, fq PaginatedFeedQuery) ([]*Order, error)
		GetShippingMethods(ctx context.Context) ([]*ShippingMethod, error)
		GetPaymentMethods(ctx context.Context) ([]*PaymentMethod, error)
	}

	ShippingAddresses interface {
		GetByID(context.Context, uuid.UUID, int64) (*ShippingAddresses, error)
		Create(context.Context, *ShippingAddresses) error
		Update(context.Context, *ShippingAddresses) error
		Delete(context.Context, uuid.UUID, int64) error
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Users:             &UserStore{db},
		Roles:             &RoleStore{db},
		Follow:            &FollowerStore{db},
		Categoris:         &CategoryStore{db},
		Products:          &ProductStore{db},
		Tokos:             &TokoStore{db},
		Comments:          &CommentStore{db},
		Carts:             &CartStore{db},
		Orders:            &OrderStore{db},
		ShippingAddresses: &ShippingAddresStore{db},
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
