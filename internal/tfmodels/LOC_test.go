package tfmodels

import (
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestModelLOC(t *testing.T) {
	id := uint(1)

	expected := models.LOC{
		ID:     &id,
		ZoneID: 1,
		Domain: "example.com",
		TTL:    86400,
		Data:   "40 27 53.86104 N 3 39 2.59092 W 712.8m 0.00m 0.00m 0.00m",
	}

	loc := LOC{}

	require.NoError(t, loc.SetRecord(expected))

	actual, err := loc.GetRecord()
	require.NoError(t, err)

	assert.Equal(t, expected, actual)
}
