package client_test

import (
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client"
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/auth"
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/client/filters"
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/logging"
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/models"
	"github.com/SuperBuker/terraform-provider-dns-he-net/test_cfg"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRecords(t *testing.T) {
	_resources := test_cfg.Config.Resources
	account := _resources.Account
	arpaZone := _resources.ArpaZone
	domainZone := _resources.DomainZone

	authObj, err := account.Auth(auth.Simple)
	require.NoError(t, err)

	client, err := client.NewClient(t.Context(), authObj, logging.NewZerolog(zerolog.DebugLevel, false))
	require.NoError(t, err)

	deleteRecord := func(t *testing.T, record models.RecordX) {
		err := client.DeleteRecord(t.Context(), record)
		assert.NoError(t, err)
	}

	allZones, err := client.GetAllZones(t.Context())
	require.NoError(t, err)

	t.Run("Domain Zone Records", func(t *testing.T) {
		_domainZone, ok := filters.ZoneById(allZones, domainZone.ID)
		require.True(t, ok)
		assert.Equal(t, domainZone.Name, _domainZone.Name)
		assert.Equal(t, domainZone.ID, _domainZone.ID)

		_, err = client.GetRecords(t.Context(), _domainZone.ID)
		require.NoError(t, err)

		domains := domainZone.RandSubs("client-example-%04d", 10000, 3)

		for _, domain := range domains {
			recordX, err := client.SetRecord(t.Context(), models.TXT{
				ID:      nil,
				ZoneID:  _domainZone.ID,
				Domain:  domain,
				TTL:     300,
				Data:    `"Random TXT record for testing"`,
				Dynamic: false,
			})

			require.NotNil(t, recordX)
			require.NoError(t, err)

			_, ok := recordX.(models.TXT)
			require.True(t, ok)

			defer deleteRecord(t, recordX)
		}
	})

	t.Run("Arpa Zone Records", func(t *testing.T) {
		_arpaZone, ok := filters.ZoneById(allZones, arpaZone.ID)
		require.True(t, ok)
		assert.Equal(t, arpaZone.Name, _arpaZone.Name)
		assert.Equal(t, arpaZone.ID, _arpaZone.ID)

		_, err = client.GetRecords(t.Context(), _arpaZone.ID)
		require.NoError(t, err)

		arpas := arpaZone.RandArpaSubs(16, 3)

		for _, arpa := range arpas {
			recordX, err := client.SetRecord(t.Context(), models.TXT{
				ID:      nil,
				ZoneID:  _arpaZone.ID,
				Domain:  arpa,
				TTL:     300,
				Data:    `"Random TXT record for testing"`,
				Dynamic: false,
			})

			require.NotNil(t, recordX)
			require.NoError(t, err)

			_, ok := recordX.(models.TXT)
			require.True(t, ok)

			defer deleteRecord(t, recordX)
		}
	})
}
