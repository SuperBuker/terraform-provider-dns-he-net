package auth

import (
	"encoding/json"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAuthStore(t *testing.T) {
	data, err := os.ReadFile("../testing_data/json/cookies.json")
	require.NoError(t, err)

	var cookies []*http.Cookie

	require.NoError(t, json.Unmarshal(data, &cookies))

	auth, err := NewAuth("test_user", "", "", Simple)
	require.NoError(t, err)

	// Remove testing file
	t.Cleanup(
		func() {
			os.Remove(configFilePath(&auth, Simple))
		},
	)

	err = auth.SaveCookies(cookies)
	require.NoError(t, err)

	cookies2, err := auth.LoadCookies()
	require.NoError(t, err)
	assert.Equal(t, cookies[:1], cookies2)
}
