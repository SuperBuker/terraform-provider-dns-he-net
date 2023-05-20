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

func TestStatus(t *testing.T) {
	t.Run("missing status", func(t *testing.T) {
		data, err := os.ReadFile("../testing_data/html/main.html")
		require.NoError(t, err)

		doc, err := htmlquery.Parse(bytes.NewReader(data))
		require.NoError(t, err)

		status := parsers.ParseStatus(doc)
		assert.Equal(t, "", status)
	})

	t.Run("status present", func(t *testing.T) {
		data, err := os.ReadFile("../testing_data/html/records_status_msg.html")
		require.NoError(t, err)

		doc, err := htmlquery.Parse(bytes.NewReader(data))
		require.NoError(t, err)

		status := parsers.ParseStatus(doc)
		assert.Equal(t, "Successfully updated record.", status)
	})
}

func TestStatusMessage(t *testing.T) {
	t.Run("missing status", func(t *testing.T) {
		data, err := os.ReadFile("../testing_data/html/main.html")
		require.NoError(t, err)

		doc, err := htmlquery.Parse(bytes.NewReader(data))
		require.NoError(t, err)

		statusMessage, err := parsers.GetStatusMessage(doc)
		require.Error(t, err)
		targetErr := &parsers.ErrNotFound{}
		assert.ErrorAs(t, err, &targetErr)

		assert.Equal(t, models.StatusMessage{}, statusMessage)
		assert.Equal(t, errNotFoundString(statusQ), err.Error())
	})

	t.Run("status present", func(t *testing.T) {
		data, err := os.ReadFile("../testing_data/html/records_status_msg.html")
		require.NoError(t, err)

		doc, err := htmlquery.Parse(bytes.NewReader(data))
		require.NoError(t, err)

		statusMessage, err := parsers.GetStatusMessage(doc)
		require.NoError(t, err)
		assert.Equal(t, models.StatusMessage{Data: "Successfully updated record."}, statusMessage)
	})
}
