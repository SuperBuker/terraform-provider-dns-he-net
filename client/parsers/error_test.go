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

func TestError(t *testing.T) {
	t.Run("missing error", func(t *testing.T) {
		data, err := os.ReadFile("../testing_data/html/main.html")
		require.NoError(t, err)

		doc, err := htmlquery.Parse(bytes.NewReader(data))
		require.NoError(t, err)

		errorSlice := parsers.ParseError(doc)
		assert.Nil(t, errorSlice)
	})

	t.Run("error present", func(t *testing.T) {
		data, err := os.ReadFile("../testing_data/html/login_totp_err.html")
		require.NoError(t, err)

		doc, err := htmlquery.Parse(bytes.NewReader(data))
		require.NoError(t, err)

		errorSlice := parsers.ParseError(doc)
		assert.Equal(t, 1, len(errorSlice))
		assert.Equal(t, "The token supplied is invalid.", errorSlice[0])
	})
}
