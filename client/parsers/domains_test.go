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

var domains = []models.Domain{
	{Id: 1234567, Domain: "example.com"},
}

func TestDomains(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		data, err := os.ReadFile("../testing_data/main.html")
		require.NoError(t, err)

		doc, err := htmlquery.Parse(bytes.NewReader(data))
		require.NoError(t, err)

		_domains, err := parsers.GetDomains(doc)
		require.NoError(t, err)

		for i, domain := range _domains {
			assert.Equal(t, domains[i], domain)
		}
	})

	t.Run("missing data", func(t *testing.T) {
		data := []byte("<html></html>")
		doc, err := htmlquery.Parse(bytes.NewReader(data))
		require.NoError(t, err)

		_domains, err := parsers.GetDomains(doc)
		require.Error(t, err)

		assert.Nil(t, _domains)
		assert.Equal(t, "element \"//table[@id=\"domains_table\"]\" not found in document", err.Error())
	})

	t.Run("empty table", func(t *testing.T) {
		data, err := os.ReadFile("../testing_data/main_empty.html")
		require.NoError(t, err)

		doc, err := htmlquery.Parse(bytes.NewReader(data))
		require.NoError(t, err)

		_domains, err := parsers.GetDomains(doc)
		require.NoError(t, err)

		assert.Equal(t, []models.Domain{}, _domains)
	})
}
