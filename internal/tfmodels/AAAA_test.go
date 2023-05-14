package tfmodels

import (
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestModelAAAA(t *testing.T) {
	id := uint(1)

	expected := models.AAAA{
		ID:      &id,
		ZoneID:  1,
		Domain:  "example.com",
		TTL:     300,
		Data:    "::",
		Dynamic: true,
	}

	aaaa := AAAA{}

	require.NoError(t, aaaa.SetRecord(expected))

	actual, err := aaaa.GetRecord()
	require.NoError(t, err)

	assert.Equal(t, expected, actual)
}
