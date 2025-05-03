package test

import (
    "context"
    "testing"

    "github.com/stretchr/testify/require"
    "github.com/yogaprasetya22/api-gotokopedia/internal/store"
)

func TestCommentStore_CRUD(t *testing.T) {
    ctx := context.Background()
    storeTest, db := NewTestStorage(t)
    require.NotNil(t, storeTest)
    require.NotNil(t, db)

    commentStore := storeTest.Comments

    // Ambil user dan product id valid dari database
    var userID, productID int64
    err := db.QueryRow("SELECT id FROM users LIMIT 1").Scan(&userID)
    require.NoError(t, err)
    err = db.QueryRow("SELECT id FROM products LIMIT 1").Scan(&productID)
    require.NoError(t, err)

    // CREATE
    comment := &store.Comment{
        Content:   "Komentar test",
        UserID:    userID,
        ProductID: productID,
        Rating:    5,
    }
    err = commentStore.Create(ctx, comment)
    require.NoError(t, err)
    require.NotZero(t, comment.ID)

    // GET BY ID
    got, err := commentStore.GetByID(ctx, comment.ID)
    require.NoError(t, err)
    require.Equal(t, comment.Content, got.Content)
    require.Equal(t, comment.UserID, got.UserID)
    require.Equal(t, comment.ProductID, got.ProductID)
    require.Equal(t, comment.Rating, got.Rating)

    // UPDATE
    comment.Content = "Komentar test update"
    err = commentStore.Update(ctx, comment)
    require.NoError(t, err)

    updated, err := commentStore.GetByID(ctx, comment.ID)
    require.NoError(t, err)
    require.Equal(t, "Komentar test update", updated.Content)

    // GET BY PRODUCT ID
    comments, err := commentStore.GetByProductID(ctx, productID)
    require.NoError(t, err)
    require.NotEmpty(t, comments)

    // DELETE
    err = commentStore.Delete(ctx, comment.ID, productID)
    require.NoError(t, err)

    _, err = commentStore.GetByID(ctx, comment.ID)
    require.Error(t, err)
}