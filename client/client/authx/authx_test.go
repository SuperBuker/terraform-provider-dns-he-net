package authx_test

import (
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client/auth"
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/client/authx"

	"github.com/pquerna/otp/totp"
	"github.com/stretchr/testify/require"
)

func TestAuthx(t *testing.T) {
	key, err := totp.Generate(
		totp.GenerateOpts{
			Issuer:      "issuer",
			AccountName: "account_name",
		})

	require.NoError(t, err)

	_auth, err := auth.NewAuth("Superbuker", "pass", key.Secret())
	require.NoError(t, err)

	require.Equal(t, map[string]string{
		"email":  _auth.User,
		"pass":   _auth.Password,
		"submit": "Login!",
	}, authx.Creds(_auth))

	passcode, err := _auth.GetCode()
	require.NoError(t, err)

	m, err := authx.Totp(_auth)
	require.NoError(t, err)

	if m["tfacode"] != passcode {
		passcode, err = _auth.GetCode()
		require.NoError(t, err)
	}

	require.Equal(t, map[string]string{
		"tfacode": passcode,
		"submit":  "Submit",
	}, m)

	_auth, err = auth.NewAuth("Superbuker", "pass", "this is not a valid secret")
	require.NoError(t, err)

	_, err = authx.Totp(_auth)
	require.Error(t, err)
}
