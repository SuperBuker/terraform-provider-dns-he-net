package tfmodels

import (
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestModelHINFO(t *testing.T) {
	id := uint(1)

	expected := models.HINFO{
		ID:     &id,
		ZoneID: 1,
		Domain: "example.com",
		TTL:    86400,
		Data:   `"armv7 Linux"`,
	}

	hinfo := HINFO{}

	require.NoError(t, hinfo.SetRecord(expected))

	actual, err := hinfo.GetRecord()
	require.NoError(t, err)

	assert.Equal(t, expected, actual)
}
