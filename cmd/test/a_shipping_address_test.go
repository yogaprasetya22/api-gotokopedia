package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/yogaprasetya22/api-gotokopedia/internal/store"
)

func TestShippingAddressStore_DefaultAddress(t *testing.T) {
	ctx := context.Background()
	storeTest, db := NewTestStorage(t)
	require.NotNil(t, storeTest)
	require.NotNil(t, db)

	shippingStore := storeTest.ShippingAddresses

	//TestUntukUser110
	for userID := int64(1); userID <= 10; userID++ {
		t.Run(fmt.Sprintf("User_%d", userID), func(t *testing.T) {
			//CreateDefaultAddress (alamatPertamaOtomatisDefault)
			addr1 := &store.ShippingAddresses{
				UserID:         userID,
				Label:          "Rumah",
				RecipientName:  fmt.Sprintf("User %d", userID),
				RecipientPhone: fmt.Sprintf("0812345678%d", userID),
				AddressLine1:   fmt.Sprintf("Jl. Alamat Utama No.%d", userID),
			}

			err := shippingStore.Create(ctx, addr1)
			require.NoError(t, err)
			require.NotEqual(t, uuid.Nil, addr1.ID)
			require.True(t, addr1.IsDefault, "First address should be default")

			// Buat alamat non-default
			addr2 := &store.ShippingAddresses{
				UserID:         userID,
				Label:          "Kantor",
				RecipientName:  fmt.Sprintf("User %d", userID),
				RecipientPhone: fmt.Sprintf("0822345678%d", userID),
				AddressLine1:   fmt.Sprintf("Jl. Alamat Kantor No.%d", userID),
			}

			err = shippingStore.Create(ctx, addr2)
			require.NoError(t, err)
			require.NotEqual(t, uuid.Nil, addr2.ID)
			require.False(t, addr2.IsDefault, "Second address should not be default")

			// Verifikasi alamat default masih addr1
			defaultAddr, err := shippingStore.GetDefaultAddress(ctx, userID)
			require.NoError(t, err)
			require.Equal(t, addr1.ID, defaultAddr.ID, "Default address should be addr1")

			// Alamat daftar harus menampilkan addr1 terlebih dahulu (default)
			addresses, err := shippingStore.ListByUser(ctx, userID)
			require.NoError(t, err)
			require.Len(t, addresses, 2)
			require.Equal(t, addr1.ID, addresses[0].ID, "First address in list should be default")
			require.True(t, addresses[0].IsDefault, "First address should be default")
			require.False(t, addresses[1].IsDefault, "Second address should not be default")

			// beralih default ke addr2
			err = shippingStore.SetDefaultAddress(ctx, addr2.ID, userID)
			require.NoError(t, err)

			// Verifikasi alamat default sekarang addr2
			defaultAddr, err = shippingStore.GetDefaultAddress(ctx, userID)
			require.NoError(t, err)
			require.Equal(t, addr2.ID, defaultAddr.ID, "Default address should now be addr2")

			// Verifikasi status default kedua alamat dalam database
			var isDefault1, isDefault2 bool
			err = db.QueryRowContext(ctx,
				"SELECT is_default FROM shipping_addresses WHERE id = $1",
				addr1.ID).Scan(&isDefault1)
			require.NoError(t, err)
			require.False(t, isDefault1, "addr1 should no longer be default")

			err = db.QueryRowContext(ctx,
				"SELECT is_default FROM shipping_addresses WHERE id = $1",
				addr2.ID).Scan(&isDefault2)
			require.NoError(t, err)
			require.True(t, isDefault2, "addr2 should now be default")

			// Verifikasi Default_shipping_address_id diperbarui
			var userDefaultAddrID uuid.NullUUID
			err = db.QueryRowContext(ctx,
				"SELECT default_shipping_address_id FROM users WHERE id = $1",
				userID).Scan(&userDefaultAddrID)
			require.NoError(t, err)
			require.True(t, userDefaultAddrID.Valid)
			require.Equal(t, addr2.ID, userDefaultAddrID.UUID)

			// hapus alamat default (addr2) harus membuat addr1 default lagi
			err = shippingStore.Delete(ctx, addr2.ID, userID)
			require.NoError(t, err)

			// Verifikasi addr1 sekarang default lagi
			defaultAddr, err = shippingStore.GetDefaultAddress(ctx, userID)
			require.NoError(t, err)
			require.Equal(t, addr1.ID, defaultAddr.ID, "After deleting addr2, addr1 should be default again")

			// Verifikasi Default_shipping_address_id diperbarui
			err = db.QueryRowContext(ctx,
				"SELECT default_shipping_address_id FROM users WHERE id = $1",
				userID).Scan(&userDefaultAddrID)
			require.NoError(t, err)
			require.True(t, userDefaultAddrID.Valid)
			require.Equal(t, addr1.ID, userDefaultAddrID.UUID)

			// hapus semua alamat
			err = shippingStore.Delete(ctx, addr1.ID, userID)
			require.NoError(t, err)

			// Verifikasi tidak ada alamat default
			_, err = shippingStore.GetDefaultAddress(ctx, userID)
			require.Error(t, err, "Should return error when no default address exists")

			// Verifikasi Default_shipping_address_id adalah nol
			err = db.QueryRowContext(ctx,
				"SELECT default_shipping_address_id FROM users WHERE id = $1",
				userID).Scan(&userDefaultAddrID)
			require.NoError(t, err)
			require.False(t, userDefaultAddrID.Valid)
		})
	}
}
