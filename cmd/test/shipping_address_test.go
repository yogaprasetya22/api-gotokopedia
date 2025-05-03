package test

import (
    "context"
    "database/sql"
    "testing"
    "time"

    "github.com/google/uuid"
    "github.com/stretchr/testify/require"
    "github.com/yogaprasetya22/api-gotokopedia/internal/store"
)

func TestShippingAddressStore_CRUD(t *testing.T) {
    ctx := context.Background()
    storeTest, db := NewTestStorage(t)
    require.NotNil(t, storeTest)
    require.NotNil(t, db)

    shippingStore := storeTest.ShippingAddresses

    // Dummy user id, pastikan user ini ada di database test
    var userID int64 = 10

    // CREATE
    addr := &store.ShippingAddresses{
        UserID:         userID,
        Label:          "Rumah",
        RecipientName:  "Budi",
        RecipientPhone: "08123456789",
        AddressLine1:   "Jl. Mawar No. 1",
        NoteForCourier: sql.NullString{String: "Tolong ketuk pintu", Valid: true},
    }
    err := shippingStore.Create(ctx, addr)
    require.NoError(t, err)
    require.NotEqual(t, uuid.Nil, addr.ID)
    require.WithinDuration(t, time.Now(), addr.CreatedAt, time.Minute)

    // GET BY ID
    got, err := shippingStore.GetByID(ctx, addr.ID, userID)
    require.NoError(t, err)
    require.Equal(t, addr.Label, got.Label)
    require.Equal(t, addr.RecipientName, got.RecipientName)
    require.Equal(t, addr.RecipientPhone, got.RecipientPhone)
    require.Equal(t, addr.AddressLine1, got.AddressLine1)
    require.Equal(t, addr.NoteForCourier.String, got.NoteForCourier.String)

    // UPDATE
    addr.Label = "Kantor"
    addr.NoteForCourier = sql.NullString{String: "Lantai 2", Valid: true}
    err = shippingStore.Update(ctx, addr)
    require.NoError(t, err)

    updated, err := shippingStore.GetByID(ctx, addr.ID, userID)
    require.NoError(t, err)
    require.Equal(t, "Kantor", updated.Label)
    require.Equal(t, "Lantai 2", updated.NoteForCourier.String)

    // DELETE
    err = shippingStore.Delete(ctx, addr.ID, userID)
    require.NoError(t, err)

    _, err = shippingStore.GetByID(ctx, addr.ID, userID)
    require.Error(t, err)
}