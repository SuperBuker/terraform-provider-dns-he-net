package status_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client/status"
	"github.com/antchfx/htmlquery"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	errAuthFailed  = &status.ErrAuthFailed{}
	errNoAuth      = &status.ErrNoAuth{}
	errOTPAuth     = &status.ErrOTPAuthFailed{}
	errPartialAuth = &status.ErrPartialAuth{}
)

func TestError(t *testing.T) {
	t.Run("missing error", func(t *testing.T) {
		data, err := os.ReadFile("../testing_data/html/main.html")
		require.NoError(t, err)

		doc, err := htmlquery.Parse(bytes.NewReader(data))
		require.NoError(t, err)

		statusMsg, errorSlice, errs := status.Check(doc)

		assert.Equal(t, "", statusMsg)
		assert.Nil(t, errorSlice)
		assert.Nil(t, errs)
	})

	t.Run("failed login error present", func(t *testing.T) {
		data, err := os.ReadFile("../testing_data/html/login_err.html")
		require.NoError(t, err)

		doc, err := htmlquery.Parse(bytes.NewReader(data))
		require.NoError(t, err)

		statusMsg, errorSlice, errs := status.Check(doc)

		assert.Equal(t, "", statusMsg)
		assert.Equal(t, []string{"Incorrect"}, errorSlice)
		assert.ErrorAs(t, errs, &errAuthFailed)
		assert.ErrorAs(t, errs, &errNoAuth)
	})

	t.Run("totp error present", func(t *testing.T) {
		data, err := os.ReadFile("../testing_data/html/login_totp_err.html")
		require.NoError(t, err)

		doc, err := htmlquery.Parse(bytes.NewReader(data))
		require.NoError(t, err)

		statusMsg, errorSlice, errs := status.Check(doc)

		assert.Equal(t, "", statusMsg)
		assert.Equal(t, []string{"The token supplied is invalid."}, errorSlice)
		assert.ErrorAs(t, errs, &errOTPAuth)
		assert.ErrorAs(t, errs, &errPartialAuth)
		assert.ErrorAs(t, errs, &errAuthFailed)
		assert.ErrorAs(t, errs, &errNoAuth)
	})
}
