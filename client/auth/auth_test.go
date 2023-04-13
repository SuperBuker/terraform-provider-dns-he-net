package auth_test

import (
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client/auth"

	"github.com/pquerna/otp/totp"
	"github.com/stretchr/testify/require"
)

// TestAuth
func TestAuth(t *testing.T) {
	key, err := totp.Generate(
		totp.GenerateOpts{
			Issuer:      "issuer",
			AccountName: "account_name",
		})

	require.NoError(t, err)

	auth, err := auth.NewAuth("user", "pass", key.Secret(), -1)
	require.NoError(t, err)

	passcode, err := auth.GetCode()
	require.NoError(t, err)

	require.True(t, totp.Validate(passcode, key.Secret()))

	require.Equal(t, map[string]string{
		"email": auth.User,
		"pass":  auth.Password,
	}, auth.GetAuthForm())
}
