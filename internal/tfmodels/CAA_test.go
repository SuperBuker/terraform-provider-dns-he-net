package tfmodels

import (
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestModelCAA(t *testing.T) {
	id := uint(1)

	expected := models.CAA{
		Id:       &id,
		ParentId: 1,
		Domain:   "example.com",
		TTL:      86400,
		Data:     `0 issuewild ";"`,
	}

	caa := CAA{}

	require.NoError(t, caa.SetRecord(expected))

	actual, err := caa.GetRecord()
	require.NoError(t, err)

	assert.Equal(t, expected, actual)
}
