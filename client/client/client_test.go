package client

import (
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
	_datasources := test_cfg.Config.DataSources
	account := _datasources.Account
	arpaZonesCount := _datasources.ArpaZonesCount
	//ArpaZoneOk := _datasources.ArpaZones.Ok
	domainZoneCount := _datasources.DomainZonesCount
	domainOk := _datasources.DomainZones.Ok
	domainPendingDelegation := _datasources.DomainZones.PendingDelegation

	t.Run("Client auth.Simple", func(t *testing.T) {
		authObj, err := auth.NewAuth(account.User, account.Password, account.OTP, auth.Simple)
		require.NoError(t, err)

		cli, err := NewClient(t.Context(), authObj, logging.NewZerolog(zerolog.DebugLevel, false))
		require.NoError(t, err)

		assert.Equal(t, account.ID, cli.GetAccount())

		for _, cookie := range cli.client.Cookies {
			cookie.Value = "" // clear cookie value
			parts := strings.Split(cookie.Raw, "; ")
			parts[0] = cookie.Name + "="
			cookie.Raw = strings.Join(parts, "; ")
		}

		// Force auth failure and re-authentication before retrial
		assert.Equal(t, account.ID, cli.GetAccount())

		domainZones, err := cli.GetDomainZones(t.Context())
		require.NoError(t, err)

		assert.Equal(t, int(domainZoneCount), len(domainZones))

		ArpaZones, err := cli.GetArpaZones(t.Context())
		require.NoError(t, err)

		assert.Equal(t, int(arpaZonesCount), len(ArpaZones))

		allZones, err := cli.GetAllZones(t.Context())
		require.NoError(t, err)

		assert.Equal(t, int(domainZoneCount+arpaZonesCount), len(allZones))

		// Not onboarded record
		records, err := cli.GetRecords(t.Context(), 0)
		require.Error(t, err)
		assert.Nil(t, records)

		records, err = cli.GetRecords(t.Context(), domainZoneOk.ID)
		require.NoError(t, err)
		assert.Equal(t, int(domainOk.RecordCount), len(records))
	})

	t.Run("Client auth.Dummy", func(t *testing.T) {
		// Using Dummy store to force auth failure and re-authentication before retrial
		authObj, err := auth.NewAuth(account.User, account.Password, account.OTP, auth.Dummy)
		require.NoError(t, err)

		// Forces regular authentication with totp retrials
		cli, err := NewClient(t.Context(), authObj, logging.NewZerolog(zerolog.DebugLevel, false), WithUserAgent("user-agent test"))
		require.NoError(t, err)
		fmt.Println(err)

		assert.Equal(t, "user-agent test", cli.client.Header.Get("User-Agent"))

		assert.Equal(t, account.ID, cli.GetAccount())

		allZones, err := cli.GetAllZones(t.Context())
		require.NoError(t, err)

		assert.Equal(t, int(domainZoneCount), len(domains))

		// Retrieving records from a ZoneID not yet delegated to the provider.
		// Historically, the client returned an error, but now this error is ignored
		// so devs can setup the records prior to NS delegation.
		records, err := cli.GetRecords(t.Context(), domainZonePendingDelegation.ID)
		require.NoError(t, err)
		assert.Equal(t, int(domainPendingDelegation.RecordCount), len(records))
	})

	t.Run("Client auth.Dummy failed", func(t *testing.T) {
		authObj, err := auth.NewAuth("user", "password", "", auth.Dummy)
		require.NoError(t, err)

		// Forces regular authentication with totp retrials
		_, err = NewClient(t.Context(), authObj, logging.NewZerolog(zerolog.DebugLevel, false), WithDebug())
		require.ErrorIs(t, err, &status.ErrNoAuth{}) // Current status is not authenticated
	})

	t.Run("Client auth.Dummy otp failed", func(t *testing.T) {
		authObj, err := auth.NewAuth(account.User, account.Password, "", auth.Dummy)
		require.NoError(t, err)

		// Forces regular authentication with totp retrials
		_, err = NewClient(t.Context(), authObj, logging.NewZerolog(zerolog.DebugLevel, false))
		require.ErrorIs(t, err, &status.ErrMissingOTPAuth{})
	})
}
