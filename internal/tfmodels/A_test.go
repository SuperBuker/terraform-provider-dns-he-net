package tfmodels

import (
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestModelA(t *testing.T) {
	id := uint(1)

	expected := models.A{
		Id:      &id,
		ZoneID:  1,
		Domain:  "example.com",
		TTL:     300,
		Data:    "0.0.0.0",
		Dynamic: true,
	}

	a := A{}

	require.NoError(t, a.SetRecord(expected))

	actual, err := a.GetRecord()
	require.NoError(t, err)

	assert.Equal(t, expected, actual)
}
