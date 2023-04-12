package auth

import (
	"net/http"
	"os"
	"testing"

	"github.com/pquerna/otp/totp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDummyStore(t *testing.T) {
	auth, err := NewAuth("test_user", "", "", Dummy)
	require.NoError(t, err)

	store := auth.store // Dummy Store

	err = store.Save(&auth, nil)
	require.NoError(t, err)

	cookies, err := store.Load(&auth)
	require.Error(t, err)
	assert.Nil(t, cookies)
}

func TestSimpleStore(t *testing.T) {
	auth, err := NewAuth("test_user", "", "", Simple)
	require.NoError(t, err)

	store := auth.store // Simple Store

	// Remove testing file
	t.Cleanup(
		func() {
			os.Remove(configFilePath(&auth))
		},
	)

	cookies := []*http.Cookie{}
	err = store.Save(&auth, cookies)
	require.NoError(t, err)

	cookies2, err := store.Load(&auth)
	require.NoError(t, err)
	assert.Equal(t, cookies, cookies2)
}

func TestEncryptStore(t *testing.T) {
	key, err := totp.Generate(
		totp.GenerateOpts{
			Issuer:      "issuer",
			AccountName: "account_name",
		})

	require.NoError(t, err)

	auth, err := NewAuth("test_user", "password", key.Secret(), Encrypted)
	require.NoError(t, err)

	store := auth.store // Encrypted Store

	// Remove testing file
	t.Cleanup(
		func() {
			os.Remove(configFilePath(&auth))
		},
	)

	cookies := []*http.Cookie{}
	err = store.Save(&auth, cookies)
	require.NoError(t, err)

	cookies2, err := store.Load(&auth)
	require.NoError(t, err)
	assert.Equal(t, cookies, cookies2)
}
