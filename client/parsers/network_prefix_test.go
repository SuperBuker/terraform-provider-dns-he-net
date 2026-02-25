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

var networkPrefixes = []models.NetworkPrefix{
	{ID: 1234567, Value: "2001:470:1f13:1::/64", Enabled: true},
	{ID: 1234568, Value: "2001:470:1::/48", Enabled: true},
	{Value: "2001:470:2::/48", Enabled: false}, // Not enabled, lacks ID
}

func TestNetworkPrefixes(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		data, err := os.ReadFile("../testing_data/html/main.html")
		require.NoError(t, err)

		doc, err := htmlquery.Parse(bytes.NewReader(data))
		require.NoError(t, err)

		prefixes_, err := parsers.GetNetworkPrefixes(doc)
		require.NoError(t, err)

		for i, prefix := range prefixes_ {
			assert.Equal(t, networkPrefixes[i], prefix)
		}
	})

	t.Run("missing data", func(t *testing.T) {
		data := []byte("<html></html>")
		doc, err := htmlquery.Parse(bytes.NewReader(data))
		require.NoError(t, err)

		prefixes_, err := parsers.GetNetworkPrefixes(doc)
		require.Error(t, err)
		targetErr := &parsers.ErrNotFound{}
		assert.ErrorAs(t, err, &targetErr)

		assert.Nil(t, prefixes_)
		// TODO: fix ErrNotFound error message
		//assert.Equal(t, `element "//table[@id='domains_table']" not found in document`, err.Error())
	})

	t.Run("empty table", func(t *testing.T) {
		data, err := os.ReadFile("../testing_data/html/main_empty.html")
		require.NoError(t, err)

		doc, err := htmlquery.Parse(bytes.NewReader(data))
		require.NoError(t, err)

		prefixes_, err := parsers.GetNetworkPrefixes(doc)
		require.NoError(t, err)

		assert.Equal(t, []models.NetworkPrefix{}, prefixes_)
	})
}
