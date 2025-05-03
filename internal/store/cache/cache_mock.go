package cache

import (
	"context"

	"github.com/stretchr/testify/mock"
	"github.com/yogaprasetya22/api-gotokopedia/internal/store"
)

func NewMockStore() Storage {
	return Storage{
		Users:    &MockUserStore{},
		Products: &MockProductStore{},
		Carts:    &MockCartStore{},
	}
}

type MockUserStore struct {
	mock.Mock
}

func (m *MockUserStore) Get(ctx context.Context, userID int64) (*store.User, error) {
	args := m.Called(userID)
	return nil, args.Error(1)
}

func (m *MockUserStore) Set(ctx context.Context, user *store.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserStore) Delete(ctx context.Context, userID int64) {
	m.Called(userID)
}

type MockProductStore struct {
	mock.Mock
}

func (m *MockProductStore) Get(ctx context.Context, ProductID int64) (*store.Product, error) {
	args := m.Called(ProductID)
	return nil, args.Error(1)
}

func (m *MockProductStore) Set(ctx context.Context, product *store.Product) error {
	args := m.Called(product)
	return args.Error(0)
}

func (m *MockProductStore) Delete(ctx context.Context, ProductID int64) {
	m.Called(ProductID)
}

type MockCartStore struct {
	mock.Mock
}

func (m *MockCartStore) Get(ctx context.Context, cartID int64) (*store.Cart, error) {
	args := m.Called(cartID)
	return nil, args.Error(1)
}

func (m *MockCartStore) Set(ctx context.Context, cart *store.Cart) error {
	args := m.Called(cart)
	return args.Error(0)
}

func (m *MockCartStore) Delete(ctx context.Context, cartID int64) {
	m.Called(cartID)
}
