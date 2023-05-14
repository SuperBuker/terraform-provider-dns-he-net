package tfmodels

import (
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestModelSRV(t *testing.T) {
	id := uint(1)

	expected := models.SRV{
		ID:       &id,
		ZoneID:   1,
		Domain:   "_bofher._tcp.example.com",
		TTL:      28800,
		Priority: 0,
		Weight:   0,
		Port:     22,
		Target:   "example.com",
	}

	srv := SRV{}

	require.NoError(t, srv.SetRecord(expected))

	actual, err := srv.GetRecord()
	require.NoError(t, err)

	assert.Equal(t, expected, actual)
}
