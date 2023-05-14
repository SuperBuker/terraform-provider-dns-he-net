package client

import (
	"context"
	"os"
	"strings"
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client/auth"
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/logging"
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/status"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClientAuth(t *testing.T) {
	user := os.Getenv("DNSHENET_USER")
	password := os.Getenv("DNSHENET_PASSWD")
	otp := os.Getenv("DNSHENET_OTP")

	t.Run("Client auth.Simple", func(t *testing.T) {
		authObj, err := auth.NewAuth(user, password, otp, auth.Simple)
		require.NoError(t, err)

		cli, err := NewClient(context.TODO(), authObj, logging.NewZerolog(zerolog.DebugLevel, false))
		require.NoError(t, err)

		assert.Equal(t, "v6643873d8c41428.97783691", cli.GetAccount())

		for _, cookie := range cli.client.Cookies {
			cookie.Value = "" // clear cookie value
			parts := strings.Split(cookie.Raw, "; ")
			parts[0] = cookie.Name + "="
			cookie.Raw = strings.Join(parts, "; ")
		}

		// Force auth failure and re-authentication before retrial
		domains, err := cli.GetDomains(context.TODO())
		require.NoError(t, err)

		assert.Equal(t, 2, len(domains))

		assert.Equal(t, "v6643873d8c41428.97783691", cli.GetAccount())

		// Not onboarded domain
		records, err := cli.GetRecords(context.TODO(), 1091256)
		require.Error(t, err)
		assert.Nil(t, records)

		records, err = cli.GetRecords(context.TODO(), 1093397)
		require.NoError(t, err)
		assert.Equal(t, 24, len(records))
	})

	t.Run("Client auth.Dummy", func(t *testing.T) {

		authObj, err := auth.NewAuth(user, password, otp, auth.Dummy)
		require.NoError(t, err)

		// Forces regular authentication with totp retrials
		cli, err := NewClient(context.TODO(), authObj, logging.NewZerolog(zerolog.DebugLevel, false), WithUserAgent("user-agent test"))
		require.NoError(t, err)

		assert.Equal(t, "user-agent test", cli.client.Header.Get("User-Agent"))

		assert.Equal(t, "v6643873d8c41428.97783691", cli.GetAccount())

		domains, err := cli.GetDomains(context.TODO())
		require.NoError(t, err)

		assert.Equal(t, 2, len(domains))

		records, err := cli.GetRecords(context.TODO(), 1091256)
		require.Error(t, err)
		assert.Nil(t, records)
	})

	t.Run("Client auth.Dummy failed", func(t *testing.T) {

		authObj, err := auth.NewAuth("user", "password", "", auth.Dummy)
		require.NoError(t, err)

		// Forces regular authentication with totp retrials
		_, err = NewClient(context.TODO(), authObj, logging.NewZerolog(zerolog.DebugLevel, false), WithDebug())
		require.ErrorIs(t, err, &status.ErrNoAuth{}) // Current status is not autheticated
	})

	t.Run("Client auth.Dummy otp failed", func(t *testing.T) {

		authObj, err := auth.NewAuth(user, password, "", auth.Dummy)
		require.NoError(t, err)

		// Forces regular authentication with totp retrials
		_, err = NewClient(context.TODO(), authObj, logging.NewZerolog(zerolog.DebugLevel, false))
		require.ErrorIs(t, err, &status.ErrOTPAuth{}) // Current status is missing OTP auth
	})
}
