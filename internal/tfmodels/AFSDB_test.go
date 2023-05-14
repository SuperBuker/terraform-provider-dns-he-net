package tfmodels

import (
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestModelAFSDB(t *testing.T) {
	id := uint(1)

	expected := models.AFSDB{
		Id:      &id,
		ZoneID:  1,
		Domain:  "example.com",
		TTL:     300,
		Data:    "2 green.example.com",
		Dynamic: true,
	}

	afsdb := AFSDB{}

	require.NoError(t, afsdb.SetRecord(expected))

	actual, err := afsdb.GetRecord()
	require.NoError(t, err)

	assert.Equal(t, expected, actual)
}
