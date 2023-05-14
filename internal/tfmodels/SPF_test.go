package tfmodels

import (
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestModelSPF(t *testing.T) {
	id := uint(1)

	expected := models.SPF{
		ID:     &id,
		ZoneID: 1,
		Domain: "example.com",
		TTL:    86400,
		Data:   `"v=spf1 include:_spf.example.com ~all"`,
	}

	spf := SPF{}

	require.NoError(t, spf.SetRecord(expected))

	actual, err := spf.GetRecord()
	require.NoError(t, err)

	assert.Equal(t, expected, actual)
}
