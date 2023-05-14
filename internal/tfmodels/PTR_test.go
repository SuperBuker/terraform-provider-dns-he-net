package tfmodels

import (
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestModelPTR(t *testing.T) {
	id := uint(1)

	expected := models.PTR{
		Id:     &id,
		ZoneID: 1,
		Domain: "sub.example.com",
		TTL:    300,
		Data:   "example.com",
	}

	ptr := PTR{}

	require.NoError(t, ptr.SetRecord(expected))

	actual, err := ptr.GetRecord()
	require.NoError(t, err)

	assert.Equal(t, expected, actual)
}
