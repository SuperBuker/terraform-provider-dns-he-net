package client

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client/auth"
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/logging"
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/status"
	"github.com/SuperBuker/terraform-provider-dns-he-net/test_cfg"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClientAuth(t *testing.T) {
	_datasources := test_cfg.Config.DataSouces
	account := _datasources.Account
	zoneID := _datasources.Zone.ID

	t.Run("Client auth.Simple", func(t *testing.T) {
		authObj, err := auth.NewAuth(account.User, account.Password, account.OTP, auth.Simple)
		require.NoError(t, err)

		cli, err := NewClient(context.Background(), authObj, logging.NewZerolog(zerolog.DebugLevel, false))
		require.NoError(t, err)

		assert.Equal(t, account.ID, cli.GetAccount())

		for _, cookie := range cli.client.Cookies {
			cookie.Value = "" // clear cookie value
			parts := strings.Split(cookie.Raw, "; ")
			parts[0] = cookie.Name + "="
			cookie.Raw = strings.Join(parts, "; ")
		}

		// Force auth failure and re-authentication before retrial
		zones, err := cli.GetZones(context.Background())
		require.NoError(t, err)

		assert.Equal(t, 3, len(zones))

		assert.Equal(t, account.ID, cli.GetAccount())

		// Not onboarded record
		records, err := cli.GetRecords(context.Background(), 0)
		require.Error(t, err)
		assert.Nil(t, records)

		records, err = cli.GetRecords(context.Background(), zoneID)
		require.NoError(t, err)
		assert.Equal(t, 24, len(records))
	})

	t.Run("Client auth.Dummy", func(t *testing.T) {
		authObj, err := auth.NewAuth(account.User, account.Password, account.OTP, auth.Dummy)
		require.NoError(t, err)

		// Forces regular authentication with totp retrials
		cli, err := NewClient(context.Background(), authObj, logging.NewZerolog(zerolog.DebugLevel, false), WithUserAgent("user-agent test"))
		require.NoError(t, err)
		fmt.Println(err)

		assert.Equal(t, "user-agent test", cli.client.Header.Get("User-Agent"))

		assert.Equal(t, account.ID, cli.GetAccount())

		zones, err := cli.GetZones(context.Background())
		require.NoError(t, err)

		assert.Equal(t, 3, len(zones))

		records, err := cli.GetRecords(context.Background(), 1096291)
		require.Error(t, err)
		assert.Nil(t, records)
	})

	t.Run("Client auth.Dummy failed", func(t *testing.T) {
		authObj, err := auth.NewAuth("user", "password", "", auth.Dummy)
		require.NoError(t, err)

		// Forces regular authentication with totp retrials
		_, err = NewClient(context.Background(), authObj, logging.NewZerolog(zerolog.DebugLevel, false), WithDebug())
		require.ErrorIs(t, err, &status.ErrNoAuth{}) // Current status is not autheticated
	})

	t.Run("Client auth.Dummy otp failed", func(t *testing.T) {
		authObj, err := auth.NewAuth(account.User, account.Password, "", auth.Dummy)
		require.NoError(t, err)

		// Forces regular authentication with totp retrials
		_, err = NewClient(context.Background(), authObj, logging.NewZerolog(zerolog.DebugLevel, false))
		require.ErrorIs(t, err, &status.ErrOTPAuth{}) // Current status is missing OTP auth
	})
}
