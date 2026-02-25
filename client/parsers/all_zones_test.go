package parsers_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client/models"
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/parsers"
	"github.com/antchfx/htmlquery"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAllZones(t *testing.T) {
	var zones = []models.Zone{}
	zones = append(zones, domainZones...)

	// Initialize arpaZones from prefixes
	for _, prefix := range networkPrefixes {
		if !prefix.Enabled {
			continue
		}

		arpaZone, err := parsers.NetworkPrefixToArpaZone(prefix)

		if err != nil {
			panic(err)
		}

		zones = append(zones, arpaZone)
	}

	t.Run("ok", func(t *testing.T) {
		data, err := os.ReadFile("../testing_data/html/main.html")
		require.NoError(t, err)

		doc, err := htmlquery.Parse(bytes.NewReader(data))
		require.NoError(t, err)

		allZones, err := parsers.GetAllZones(doc)
		require.NoError(t, err)

		for i, zone := range allZones {
			assert.Equal(t, zones[i], zone)
		}
	})

	t.Run("missing data", func(t *testing.T) {
		data := []byte("<html></html>")
		doc, err := htmlquery.Parse(bytes.NewReader(data))
		require.NoError(t, err)

		allZones, err := parsers.GetAllZones(doc)
		require.Error(t, err)
		targetErr := &parsers.ErrNotFound{}
		assert.ErrorAs(t, err, &targetErr)

		assert.Nil(t, allZones)
		// TODO: fix ErrNotFound error message
		//assert.Equal(t, `element "//table[@id='domains_table']" not found in document`, err.Error())
	})

	t.Run("empty table", func(t *testing.T) {
		data, err := os.ReadFile("../testing_data/html/main_empty.html")
		require.NoError(t, err)

		doc, err := htmlquery.Parse(bytes.NewReader(data))
		require.NoError(t, err)

		allZones, err := parsers.GetAllZones(doc)
		require.NoError(t, err)

		assert.Equal(t, []models.Zone{}, allZones)
	})
}
