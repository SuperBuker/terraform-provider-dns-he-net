package client_test

import (
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client"
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/auth"
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/client/filters"
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/logging"
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/models"
	"github.com/SuperBuker/terraform-provider-dns-he-net/internal/utils"
	"github.com/SuperBuker/terraform-provider-dns-he-net/test_cfg"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDDNSKey(t *testing.T) {
	_resources := test_cfg.Config.Resources
	account := _resources.Account
	arpaZone := _resources.ArpaZone
	domainZone := _resources.DomainZone

	authObj, err := account.Auth(auth.Simple)
	require.NoError(t, err)

	client, err := client.NewClient(t.Context(), authObj, logging.NewZerolog(zerolog.DebugLevel, false))
	require.NoError(t, err)

	ddnsClient := client.DDNS()

	deleteRecord := func(t *testing.T, record models.RecordX) {
		err := client.DeleteRecord(t.Context(), record)
		assert.NoError(t, err)
	}

	resetDDNSKey := func(t *testing.T, ddns models.DDNSKey) {
		ddns.Key = utils.GenerateRandomString(32) // The API doesn't support deletion, so we just set a random key
		_, err := client.SetDDNSKey(t.Context(), ddns)
		require.NoError(t, err)
	}

	allZones, err := client.GetAllZones(t.Context())
	require.NoError(t, err)

	t.Run("Create DomainZone Record + DDNS Key", func(t *testing.T) {
		_domainZone, ok := filters.ZoneById(allZones, domainZone.ID)
		require.True(t, ok)

		domains := domainZone.RandSubs("ddns-example-%04d", 10000, 1)
		require.Len(t, domains, 1) // Just in case...

		for _, domain := range domains {
			recordX, err := client.SetRecord(t.Context(), models.TXT{
				ID:      nil,
				ZoneID:  _domainZone.ID,
				Domain:  domain,
				TTL:     300,
				Data:    `"Random TXT record for testing"`,
				Dynamic: true,
			})

			require.NotNil(t, recordX)
			require.NoError(t, err)

			recordID, ok := recordX.GetID()
			require.True(t, ok)

			defer deleteRecord(t, recordX)

			ddns := models.DDNSKey{
				ZoneID: _domainZone.ID, // To test
				Domain: domain,
				Key:    utils.GenerateRandomString(32),
			}

			resp, err := client.SetDDNSKey(t.Context(), ddns)
			require.NoError(t, err)
			assert.NotNil(t, resp)

			defer resetDDNSKey(t, ddns)

			ok, err = ddnsClient.CheckAuth(t.Context(), domain, ddns.Key)
			require.NoError(t, err)
			require.True(t, ok)

			ok, err = ddnsClient.UpdateTXT(t.Context(), domain, ddns.Key, `"New random TXT record for testing"`)
			require.NoError(t, err)
			require.True(t, ok)

			records, err := client.GetRecords(t.Context(), _domainZone.ID)
			require.NoError(t, err)

			record, ok := filters.RecordById(records, recordID)
			require.True(t, ok)
			assert.Equal(t, `"New random TXT record for testing"`, record.Data)
		}
	})

	t.Run("Create ArpaZone Record + DDNS Key", func(t *testing.T) {
		_arpaZone, ok := filters.ZoneById(allZones, arpaZone.ID)
		require.True(t, ok)

		arpas := arpaZone.RandArpaSubs(16, 1)
		require.Len(t, arpas, 1) // Just in case...

		for _, arpa := range arpas {
			recordX, err := client.SetRecord(t.Context(), models.TXT{
				ID:      nil,
				ZoneID:  _arpaZone.ID,
				Domain:  arpa,
				TTL:     300,
				Data:    `"Random TXT record for testing"`,
				Dynamic: true,
			})

			require.NotNil(t, recordX)
			require.NoError(t, err)

			recordID, ok := recordX.GetID()
			require.True(t, ok)

			defer deleteRecord(t, recordX)

			ddns := models.DDNSKey{
				ZoneID: _arpaZone.ID, // To test
				Domain: arpa,
				Key:    utils.GenerateRandomString(32),
			}

			resp, err := client.SetDDNSKey(t.Context(), ddns)
			require.NoError(t, err)
			assert.NotNil(t, resp)

			defer resetDDNSKey(t, ddns)

			ok, err = ddnsClient.CheckAuth(t.Context(), arpa, ddns.Key)
			require.NoError(t, err)
			require.True(t, ok)

			ok, err = ddnsClient.UpdateTXT(t.Context(), arpa, ddns.Key, `New random TXT record for testing`)
			require.NoError(t, err)
			require.True(t, ok)

			records, err := client.GetRecords(t.Context(), _arpaZone.ID)
			require.NoError(t, err)

			record, ok := filters.RecordById(records, recordID)
			require.True(t, ok)
			assert.Equal(t, `"New random TXT record for testing"`, record.Data)
		}
	})
}
