package parsers_test

import (
	"os"
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client/parsers"
	"github.com/stretchr/testify/require"
)

func TestError(t *testing.T) {
	t.Run("missing error", func(t *testing.T) {
		data, err := os.ReadFile("../testing_data/main.html")
		require.NoError(t, err)

		errorString, err := parsers.ParseError(data)
		require.Equal(t, "", errorString)
		require.NoError(t, err)
	})

	t.Run("error present", func(t *testing.T) {
		data, err := os.ReadFile("../testing_data/login_totp_err.html")
		require.NoError(t, err)

		errorString, err := parsers.ParseError(data)
		require.Equal(t, "The token supplied is invalid.", errorString)
		require.NoError(t, err)
	})
}
