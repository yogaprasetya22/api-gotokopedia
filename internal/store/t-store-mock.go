package store

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

func NewMockStore() Storage {
	return Storage{
		Users: &MockUserStore{},
		// Products: &MockProductStore{},
		Carts: &MockCartStore{},
	}
}

type MockUserStore struct{ mock.Mock }

func (m *MockUserStore) CreateWithActive(ctx context.Context, tx *sql.Tx, u *User) (*User, error) {
	return &User{}, nil
}

func (m *MockUserStore) Create(ctx context.Context, tx *sql.Tx, u *User) (*User, error) {
	return &User{}, nil
}

func (m *MockUserStore) CreateByOAuth(ctx context.Context, u *User) error {
	return nil
}

func (m *MockUserStore) GetByID(ctx context.Context, userID int64) (*User, error) {
	return &User{ID: userID}, nil
}

func (m *MockUserStore) GetByEmail(context.Context, string) (*User, error) {
	return &User{}, nil
}

func (m *MockUserStore) GetByGoogleID(context.Context, string) (*User, error) {
	return &User{}, nil
}

func (m *MockUserStore) CreateAndInvite(ctx context.Context, user *User, token string, exp time.Duration) error {
	return nil
}

func (m *MockUserStore) Activate(ctx context.Context, t string) error {
	return nil
}

func (m *MockUserStore) Delete(ctx context.Context, id int64) error {
	return nil
}

type MockCartStore struct{ mock.Mock }

func (m *MockCartStore) AddToCartTransaction(ctx context.Context, userID, productID, quantity int64) (*Cart, error) {
	args := m.Called(ctx, userID, productID, quantity)
	return args.Get(0).(*Cart), args.Error(1)
}

func (m *MockCartStore) AddingQuantityCartStoreItemTransaction(ctx context.Context, cartStoreItemID uuid.UUID, userID int64) error {
	return nil
}

func (m *MockCartStore) GetCartByUserID(ctx context.Context, userID int64, query PaginatedFeedQuery) (*Cart, error) {
	return &Cart{}, nil
}

func (m *MockCartStore) GetDetailCartByCartStoreID(ctx context.Context, cartStoreID uuid.UUID, userID int64) (*CartDetailResponse, error) {
	return &CartDetailResponse{}, nil
}

func (m *MockCartStore) RemovingQuantityCartStoreItemTransaction(ctx context.Context, cartStoreItemID uuid.UUID, userID int64) error {
	return nil
}
