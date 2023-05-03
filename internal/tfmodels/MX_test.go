package tfmodels

import (
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestModelMX(t *testing.T) {
	id := uint(1)

	expected := models.MX{
		Id:       &id,
		ParentId: 1,
		Domain:   "example.com",
		TTL:      3600,
		Priority: 1,
		Data:     "mx.example.com",
	}

	mx := MX{}

	require.NoError(t, mx.SetRecord(expected))

	actual, err := mx.GetRecord()
	require.NoError(t, err)

	assert.Equal(t, expected, actual)
}
