package parsers_test

import (
	"io/ioutil"
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client/auth"
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/parsers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAuth(t *testing.T) {
	t.Run("login", func(t *testing.T) {
		data, err := ioutil.ReadFile("../testing_data/login.html")
		require.NoError(t, err)

		status, err := parsers.LoginStatus(data)
		require.NoError(t, err)
		assert.Equal(t, auth.NoAuth, status)
	})

	t.Run("login_otp", func(t *testing.T) {
		data, err := ioutil.ReadFile("../testing_data/login_totp.html")
		require.NoError(t, err)

		status, err := parsers.LoginStatus(data)
		require.NoError(t, err)
		assert.Equal(t, auth.OTP, status)
	})

	t.Run("ok", func(t *testing.T) {
		files := []string{
			"../testing_data/main.html",
		}

		for _, file := range files {
			data, err := ioutil.ReadFile(file)
			require.NoError(t, err)

			status, err := parsers.LoginStatus(data)
			require.NoError(t, err)
			assert.Equal(t, auth.Ok, status)
		}
	})

	t.Run("unkown", func(t *testing.T) {
		data, err := ioutil.ReadFile("../testing_data/empty.html")
		require.NoError(t, err)

		status, err := parsers.LoginStatus(data)
		require.NoError(t, err)
		assert.Equal(t, auth.Unknown, status)
	})
}
