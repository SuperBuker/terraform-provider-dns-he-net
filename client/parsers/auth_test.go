package parsers_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client/auth"
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/parsers"
	"github.com/antchfx/htmlquery"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAuth(t *testing.T) {
	t.Run("login", func(t *testing.T) {
		files := []string{
			"../testing_data/html/login.html",
			"../testing_data/html/login_err.html",
		}

		for _, file := range files {
			data, err := os.ReadFile(file)
			require.NoError(t, err)

			doc, err := htmlquery.Parse(bytes.NewReader(data))
			require.NoError(t, err)

			status := parsers.LoginStatus(doc)
			assert.Equal(t, auth.NoAuth, status)
		}
	})

	t.Run("login_otp", func(t *testing.T) {
		data, err := os.ReadFile("../testing_data/html/login_totp.html")
		require.NoError(t, err)

		doc, err := htmlquery.Parse(bytes.NewReader(data))
		require.NoError(t, err)

		status := parsers.LoginStatus(doc)
		assert.Equal(t, auth.OTP, status)
	})

	t.Run("ok", func(t *testing.T) {
		files := []string{
			"../testing_data/html/main.html",
			"../testing_data/html/records.html",
		}

		for _, file := range files {
			data, err := os.ReadFile(file)
			require.NoError(t, err)

			doc, err := htmlquery.Parse(bytes.NewReader(data))
			require.NoError(t, err)

			status := parsers.LoginStatus(doc)
			assert.Equal(t, auth.Ok, status)
		}
	})

	t.Run("unknown", func(t *testing.T) {
		status := parsers.LoginStatus(nil)
		assert.Equal(t, auth.Unknown, status)
	})
}
