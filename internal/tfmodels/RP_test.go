package tfmodels

import (
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestModelRP(t *testing.T) {
	id := uint(1)

	expected := models.RP{
		Id:     &id,
		ZoneID: 1,
		Domain: "example.com",
		TTL:    86400,
		Data:   "bofher.example.com bofher.example.com",
	}

	rp := RP{}

	require.NoError(t, rp.SetRecord(expected))

	actual, err := rp.GetRecord()
	require.NoError(t, err)

	assert.Equal(t, expected, actual)
}
