package auth

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestAuthStore tests the SaveCookies and LoadCookies methods.
func TestAuthStore(t *testing.T) {
	// Read test data from file.
	bytes, err := os.ReadFile("../testing_data/json/cookies.json")
	require.NoError(t, err)

	var data serialisedStore

	require.NoError(t, json.Unmarshal(bytes, &data))

	// Create a new Auth object using the Simple cookie store.
	auth, err := NewAuth("test_user", "", "", Simple)
	require.NoError(t, err)

	// Remove testing file when the test is complete.
	t.Cleanup(
		func() {
			//nolint:errcheck
			os.Remove(configFilePath(&auth, Simple))
		},
	)

	// Save the cookies to the file.
	err = auth.Save(data.Account, data.Cookies)
	require.NoError(t, err)

	// Load the cookies from the file.
	account, cookies, err := auth.Load()
	require.NoError(t, err)

	// Verify the account id was loaded correctly.
	assert.Equal(t, data.Account, account)

	// Verify only first cookie was loaded after filtering.
	assert.Equal(t, data.Cookies[:1], cookies)
}
