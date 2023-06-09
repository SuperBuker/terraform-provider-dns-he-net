package tfmodels

import (
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestModelTXT(t *testing.T) {
	id := uint(1)

	expected := models.TXT{
		ID:      &id,
		ZoneID:  1,
		Domain:  "example.com",
		TTL:     300,
		Data:    `"Just for the record"`,
		Dynamic: true,
	}

	txt := TXT{}

	require.NoError(t, txt.SetRecord(expected))

	actual, err := txt.GetRecord()
	require.NoError(t, err)

	assert.Equal(t, expected, actual)
}
