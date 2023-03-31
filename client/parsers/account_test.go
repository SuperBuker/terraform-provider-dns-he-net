package parsers_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client/parsers"
	"github.com/antchfx/htmlquery"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccount(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		data, err := os.ReadFile("../testing_data/main.html")
		require.NoError(t, err)

		doc, err := htmlquery.Parse(bytes.NewReader(data))
		require.NoError(t, err)

		account, err := parsers.GetAccount(doc)
		require.NoError(t, err)

		assert.Equal(t, "tb12d34de5678901.23456789", account)
	})

	t.Run("missing data", func(t *testing.T) {
		data := []byte("<html></html>")
		doc, err := htmlquery.Parse(bytes.NewReader(data))
		require.NoError(t, err)

		account, err := parsers.GetAccount(doc)
		require.Error(t, err)

		assert.Equal(t, "", account)
		assert.Equal(t, "element \"//form[@name=\"remove_domain\"]/input[@name=\"account\"]\" not found in document", err.Error())
	})
}
