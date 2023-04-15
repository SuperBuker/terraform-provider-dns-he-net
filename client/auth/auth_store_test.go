package auth

import (
	"encoding/json"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestAuthStore tests the SaveCookies and LoadCookies methods.
func TestAuthStore(t *testing.T) {
	// Read test data from file.
	data, err := os.ReadFile("../testing_data/json/cookies.json")
	require.NoError(t, err)

	var cookies []*http.Cookie

	require.NoError(t, json.Unmarshal(data, &cookies))

	// Create a new Auth object using the Simple cookie store.
	auth, err := NewAuth("test_user", "", "", Simple)
	require.NoError(t, err)

	// Remove testing file when the test is complete.
	t.Cleanup(
		func() {
			os.Remove(configFilePath(&auth, Simple))
		},
	)

	// Save the cookies to the file.
	err = auth.SaveCookies(cookies)
	require.NoError(t, err)

	// Load the cookies from the file.
	cookies2, err := auth.LoadCookies()
	require.NoError(t, err)

	// Verify only first cookie was loaded after filtering.
	assert.Equal(t, cookies[:1], cookies2)
}
