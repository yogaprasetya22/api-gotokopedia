package test

import (
    "context"
    "testing"
    "time"

    "github.com/stretchr/testify/require"
    "github.com/yogaprasetya22/api-gotokopedia/internal/store"
)

func TestTokoStore_CRUD(t *testing.T) {
    ctx := context.Background()
    storeTest, db := NewTestStorage(t)
    require.NotNil(t, storeTest)
    require.NotNil(t, db)

    tokoStore := storeTest.Tokos

    // Ambil user id valid dari database
    var userID int64
    err := db.QueryRow("SELECT id FROM users LIMIT 1").Scan(&userID)
    require.NoError(t, err)
    require.NotZero(t, userID)

    // CREATE
    toko := &store.Toko{
        UserID:       userID,
        Slug:         "tokotest-" + time.Now().Format("150405"),
        Name:         "Toko Test",
        ImageProfile: "https://example.com/image.jpg",
        Country:      "Indonesia",
    }
    err = tokoStore.Create(ctx, toko)
    require.NoError(t, err)
    require.NotZero(t, toko.ID)
    require.WithinDuration(t, time.Now(), toko.CreatedAt, time.Minute)

    // GET BY ID
    got, err := tokoStore.GetByID(ctx, toko.ID)
    require.NoError(t, err)
    require.Equal(t, toko.Name, got.Name)
    require.Equal(t, toko.Slug, got.Slug)
    require.Equal(t, toko.UserID, got.UserID)
    require.NotNil(t, got.User)
    require.Equal(t, userID, got.User.ID)

    // GET BY SLUG
    gotBySlug, err := tokoStore.GetBySlug(ctx, toko.Slug)
    require.NoError(t, err)
    require.Equal(t, toko.ID, gotBySlug.ID)
    require.Equal(t, toko.Slug, gotBySlug.Slug)
    require.NotNil(t, gotBySlug.User)
    require.Equal(t, userID, gotBySlug.User.ID)
}