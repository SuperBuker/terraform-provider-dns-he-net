package auth_test

import (
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client/auth"

	"github.com/pquerna/otp/totp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestAuth
func TestAuth(t *testing.T) {
	// Generate a new TOTP key.
	key, err := totp.Generate(
		totp.GenerateOpts{
			Issuer:      "issuer",
			AccountName: "account_name",
		})
	require.NoError(t, err)

	// Create a new Auth object with disabled cookie store.
	auth_, err := auth.NewAuth("user", "pass", key.Secret(), auth.Dummy)
	require.NoError(t, err)

	// Generate a TOTP code.
	passcode, err := auth_.GetCode()
	require.NoError(t, err)

	// Validate the TOTP code.
	assert.True(t, totp.Validate(passcode, key.Secret()))

	// Validate the generated auth form.
	assert.Equal(t, map[string]string{
		"email": auth_.User,
		"pass":  auth_.Password,
	}, auth_.GetAuthForm())
}

func TestAuthSimple(t *testing.T) {
	// Create a new Auth object with disabled otp.
	auth_, err := auth.NewAuth("user", "pass", "", auth.Dummy)
	require.NoError(t, err)

	// Generate a TOTP code.
	passcode, err := auth_.GetCode()
	require.ErrorIs(t, err, &auth.ErrOTPDisabled{})
	assert.Equal(t, "", passcode)

	// Validate the generated auth form.
	assert.Equal(t, map[string]string{
		"email": auth_.User,
		"pass":  auth_.Password,
	}, auth_.GetAuthForm())
}
