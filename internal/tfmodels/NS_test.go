package tfmodels

import (
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestModelNS(t *testing.T) {
	id := uint(1)

	expected := models.NS{
		Id:     &id,
		ZoneID: 1,
		Domain: "example.com",
		TTL:    172800,
		Data:   "ns1.he.net",
	}

	ns := NS{}

	require.NoError(t, ns.SetRecord(expected))

	actual, err := ns.GetRecord()
	require.NoError(t, err)

	assert.Equal(t, expected, actual)
}
