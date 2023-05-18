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

var zones = []models.Zone{
	{ID: 1234567, Name: "example.com"},
}

func TestZones(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		data, err := os.ReadFile("../testing_data/html/main.html")
		require.NoError(t, err)

		doc, err := htmlquery.Parse(bytes.NewReader(data))
		require.NoError(t, err)

		zones_, err := parsers.GetZones(doc)
		require.NoError(t, err)

		for i, zone := range zones_ {
			assert.Equal(t, zones[i], zone)
		}
	})

	t.Run("missing data", func(t *testing.T) {
		data := []byte("<html></html>")
		doc, err := htmlquery.Parse(bytes.NewReader(data))
		require.NoError(t, err)

		zones_, err := parsers.GetZones(doc)
		require.Error(t, err)
		targetErr := &parsers.ErrNotFound{}
		assert.ErrorAs(t, err, &targetErr)

		assert.Nil(t, zones_)
		assert.Equal(t, `element "//table[@id=\"domains_table\"]" not found in document`, err.Error())
	})

	t.Run("empty table", func(t *testing.T) {
		data, err := os.ReadFile("../testing_data/html/main_empty.html")
		require.NoError(t, err)

		doc, err := htmlquery.Parse(bytes.NewReader(data))
		require.NoError(t, err)

		zones_, err := parsers.GetZones(doc)
		require.NoError(t, err)

		assert.Equal(t, []models.Zone{}, zones_)
	})
}
